// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package auth

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/onosproject/onos-lib-go/pkg/auth"
	"google.golang.org/grpc/metadata"
	"gotest.tools/assert"
	"os"
	"strings"
	"testing"
	"time"
)

type RealmAccess struct {
	Roles []string
}

type Account struct {
	Roles []string `json:"roles"`
}

type ResourceAccess struct {
	Account Account `json:"account"`
}
type TestCustomClaims struct {
	jwt.RegisteredClaims
	Name              string         `json:"name"`
	Email             string         `json:"email"`
	EmailVerified     bool           `json:"email_verified"`
	PreferredUsername string         `json:"preferred_username"`
	Groups            []string       `json:"groups"`
	Roles             []string       `json:"roles"`
	Foo               int            `json:"foo"`
	Foo32             int32          `json:"foo32"`
	RealmAccess       RealmAccess    `json:"realm-access"`
	ResourceAccess    ResourceAccess `json:"resource-access"`
}

func (c TestCustomClaims) Validate() error {
	if c.Name == "" || c.Email == "" {
		return fmt.Errorf("name or email cannot be empty")
	}
	return nil
}

func Test_AuthenticationInterceptor(t *testing.T) {
	now := time.Now()

	claims := TestCustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "http://dex:32000",
			Subject:   "Test_AuthenticationInterceptor",
			Audience:  []string{"testaudience"},
			ExpiresAt: &jwt.NumericDate{Time: now.Add(24 * time.Hour)},
			NotBefore: &jwt.NumericDate{Time: now},
			IssuedAt:  &jwt.NumericDate{Time: now},
			ID:        "",
		},
		Name:              "testname",
		Email:             "test1@opennetworking.org",
		EmailVerified:     true,
		PreferredUsername: "a user Name",
		Groups:            []string{"testGroup1", "testGroup2"},
		Roles:             []string{"testRole1", "testRole2"},
		Foo:               21,
		Foo32:             22,
		RealmAccess: RealmAccess{
			Roles: []string{
				"testRole1",
				"testRole2",
			},
		},
		ResourceAccess: ResourceAccess{
			Account: Account{
				Roles: []string{
					"testRole1",
					"testRole2",
				},
			},
		},
	}

	assert.NilError(t, claims.Validate())

	signingKey := "testkey"
	err := os.Setenv(auth.SharedSecretKey, signingKey)
	assert.NilError(t, err)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := token.SignedString([]byte(signingKey))
	assert.NilError(t, err)

	mdIn := metadata.Pairs("authorization", fmt.Sprintf("bearer %s", s))
	ctx := metadata.NewIncomingContext(context.Background(), mdIn)
	intercepted, err := AuthenticationInterceptor(ctx)
	assert.NilError(t, err)
	md := metautils.ExtractIncoming(intercepted)
	assert.Assert(t, md != nil, "expected a value for Metadata")
	assert.Equal(t, "testname", md.Get("Name"))
	assert.Equal(t, "test1@opennetworking.org", md.Get("Email"))
	assert.Equal(t, "testGroup1;testGroup2", md.Get("Groups"))
	assert.Equal(t, "testRole1;testRole2", md.Get("Roles"))
	assert.Equal(t, "a user Name", md.Get("preferred_username"))
	assert.Assert(t, strings.HasPrefix(md.Get("authorization"), ContextMetadataTokenKey))
	assert.Equal(t, "testRole1;testRole2", md.Get("realm-access/roles"))
	assert.Equal(t, "testRole1;testRole2", md.Get("resource-access/account/roles"))
}
func Test_AuthenticationInterceptor_InvalidExpiry(t *testing.T) {
	now := time.Now()

	claims := TestCustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "http://dex:32000",
			Audience:  []string{"testaudience"},
			ExpiresAt: &jwt.NumericDate{Time: now.Add(-24 * time.Hour)}, // Before issued at
			NotBefore: &jwt.NumericDate{Time: now},
			IssuedAt:  &jwt.NumericDate{Time: now},
		},
		Name:  "testname",
		Email: "test1@opennetworking.org",
	}
	assert.NilError(t, claims.Validate())

	signingKey := "testkey"
	err := os.Setenv(auth.SharedSecretKey, signingKey)
	assert.NilError(t, err)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := token.SignedString([]byte(signingKey))
	assert.NilError(t, err)

	mdIn := metadata.Pairs("authorization", fmt.Sprintf("bearer %s", s))
	ctx := metadata.NewIncomingContext(context.Background(), mdIn)
	_, err = AuthenticationInterceptor(ctx)
	assert.ErrorContains(t, err, "token has invalid claims: token is expired")
}

func Test_AuthenticationInterceptor_NoAuth_NotAllowed(t *testing.T) {
	oldValue := os.Getenv(allowMissingAuthClients)
	assert.NilError(t, os.Setenv(allowMissingAuthClients, "some-other-client"))
	defer func() {
		assert.NilError(t, os.Setenv(allowMissingAuthClients, oldValue))
	}()

	signingKey := "testkey"
	err := os.Setenv(auth.SharedSecretKey, signingKey)
	assert.NilError(t, err)

	mdIn := metadata.Pairs("client", "test-client")
	ctx := metadata.NewIncomingContext(context.Background(), mdIn)
	_, err = AuthenticationInterceptor(ctx)
	assert.ErrorContains(t, err, "Request unauthenticated with bearer")
}

func Test_AuthenticationInterceptor_NoAuth_Allowed(t *testing.T) {
	oldValue := os.Getenv(allowMissingAuthClients)
	assert.NilError(t, os.Setenv(allowMissingAuthClients, "test-client,some-other-client"))
	defer func() {
		assert.NilError(t, os.Setenv(allowMissingAuthClients, oldValue))
	}()

	signingKey := "testkey"
	err := os.Setenv(auth.SharedSecretKey, signingKey)
	assert.NilError(t, err)

	mdIn := metadata.Pairs("client", "test-client")
	ctx := metadata.NewIncomingContext(context.Background(), mdIn)
	intercepted, err := AuthenticationInterceptor(ctx)
	assert.NilError(t, err)
	md := metautils.ExtractIncoming(intercepted)
	assert.Equal(t, "", md.Get("iat"))
}
