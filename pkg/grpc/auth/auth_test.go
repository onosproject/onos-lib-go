// Copyright 2021-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
