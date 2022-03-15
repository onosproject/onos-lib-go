// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package auth

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/onosproject/onos-lib-go/pkg/auth"
	"google.golang.org/grpc/metadata"
	"gotest.tools/assert"
	"os"
	"testing"
	"time"
)

func Test_AuthenticationInterceptor(t *testing.T) {
	now := time.Now().Unix()
	signingKey := "testkey"
	err := os.Setenv(auth.SharedSecretKey, signingKey)
	assert.NilError(t, err)

	claims := jwt.MapClaims{}
	claims["name"] = "testname"
	claims["email"] = "test1@opennetworking.org"
	claims["aud"] = "testaudience"
	claims["iat"] = float64(now)
	claims["exp"] = float64(now + 100000)
	claims["iss"] = "http://dex:32000"
	claims["sub"] = "Test_AuthenticationInterceptor"
	claims["preferred_username"] = "a user name"
	claims["groups"] = []string{"testGroup1", "testGroup2"}
	claims["roles"] = []string{"testRole1", "testRole2"}
	claims["nbf"] = 0.0
	assert.NilError(t, claims.Valid())

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
	assert.Equal(t, "testname", md.Get("name"))
	assert.Equal(t, "test1@opennetworking.org", md.Get("email"))
	assert.Equal(t, "testGroup1;testGroup2", md.Get("groups"))
	assert.Equal(t, "testRole1;testRole2", md.Get("roles"))
	assert.Equal(t, "a user name", md.Get("preferred_username"))
}
