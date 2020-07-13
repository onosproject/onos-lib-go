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

package grpcinterceptors

import (
	"context"
	"strings"

	"github.com/dgrijalva/jwt-go"

	"google.golang.org/grpc"

	"github.com/onosproject/onos-lib-go/pkg/auth"

	"github.com/onosproject/onos-lib-go/pkg/logging"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
)

const (
	// ContextMetadataTokenKey metadata token key
	ContextMetadataTokenKey = "bearer"
	GroupsKey               = "groups"
)

var log = logging.GetLogger("interceptors")

func getMethodInformation(fullMethod string) (service, verb string) {
	parts := strings.Split(fullMethod, "/")

	if len(parts) > 1 {
		splitedParts := strings.Split(parts[1], ".")
		index := len(splitedParts) - 1
		service = splitedParts[index]
	}

	if len(parts) > 2 {
		verb = strings.ToLower(parts[2])
	}
	return service, verb
}

func authorize(claims jwt.MapClaims, info *grpc.UnaryServerInfo) error {

	// Retrieve service information and rpc method name
	reqService, reqVerb := getMethodInformation(info.FullMethod)
	log.Info(reqService, reqVerb)
	claimedGroups := make([]interface{}, len(claims))
	//defaultRoles := rbac.GetDefaultRoles()
	for key, _ := range claims {
		// extract claimed groups from the token
		if key == GroupsKey {
			claimedGroups = claims[GroupsKey].([]interface{})
			log.Info(claimedGroups)
		}

	}

	return nil
}

// AuthorizationUnaryInterceptor an unary interceptor for authorization
func AuthorizationUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		log.Info("Authorizing the user")
		// Extract token from metadata in the context
		tokenString, err := grpc_auth.AuthFromMD(ctx, ContextMetadataTokenKey)
		if err != nil {
			return nil, err
		}

		jwtAuth := new(auth.JwtAuthenticator)
		claims, err := jwtAuth.ParseAndValidate(tokenString)
		if err != nil {
			return ctx, err
		}

		err = authorize(claims, info)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)

	}
}

// AuthenticationInterceptor an interceptor for authentication
func AuthenticationInterceptor(ctx context.Context) (context.Context, error) {
	log.Info("Authenticating the user")

	// Extract token from metadata in the context
	tokenString, err := grpc_auth.AuthFromMD(ctx, ContextMetadataTokenKey)
	if err != nil {
		return nil, err
	}

	// Authenticate the jwt token
	jwtAuth := new(auth.JwtAuthenticator)
	_, err = jwtAuth.ParseAndValidate(tokenString)
	if err != nil {
		return ctx, err
	}

	return ctx, nil

}
