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

type TestCustomClaims struct {
	jwt.RegisteredClaims
	Name              string   `json:"name"`
	Email             string   `json:"email"`
	PreferredUsername string   `json:"preferred_username"`
	Groups            []string `json:"groups"`
	Roles             []string `json:"roles"`
	Foo               int
	Foo32             int32
}

func (c *TestCustomClaims) Validate() error {
	if c.Name == "" || c.Email == "" {
		return fmt.Errorf("Name or Email cannot be empty")
	}
	return nil
}

func Test_AuthenticationInterceptor(t *testing.T) {
	now := time.Now()
	in100kSec := time.Unix(now.Unix()+1000000, 0)
	signingKey := "testkey"
	err := os.Setenv(auth.SharedSecretKey, signingKey)
	assert.NilError(t, err)

	claims := TestCustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "http://dex:32000",
			Subject:   "Test_AuthenticationInterceptor",
			Audience:  []string{"testaudience"},
			ExpiresAt: &jwt.NumericDate{Time: in100kSec},
			NotBefore: &jwt.NumericDate{Time: now},
			IssuedAt:  &jwt.NumericDate{Time: now},
			ID:        "",
		},
		Name:              "testname",
		Email:             "test1@opennetworking.org",
		PreferredUsername: "a user Name",
		Groups:            []string{"testGroup1", "testGroup2"},
		Roles:             []string{"testRole1", "testRole2"},
		Foo:               21,
		Foo32:             22,
	}
	assert.NilError(t, claims.Validate())

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Valid = true
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
}
