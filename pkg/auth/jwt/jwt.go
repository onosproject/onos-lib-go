// Copyright 2020-present Open Networking Foundation.
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

package jwt

import (
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

const (
	// HsSecretKey signing secret key for HMAC with SHA-256 algorithm
	HsSecretKey = "HS_SECRET_KEY"
)

// ParseToken parse a jwt string token and returns a jwt token
func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// HS256, HS384, or HS512
		if strings.HasPrefix(token.Method.Alg(), "HS") {
			key := os.Getenv(HsSecretKey)
			return []byte(key), nil
		}
		return nil, fmt.Errorf("unknown signining algorithm: %s", token.Method.Alg())
	})

	return token, err
}
