// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package auth

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/onosproject/onos-lib-go/pkg/auth"
	"strings"
)

const (
	// ContextMetadataTokenKey metadata token key
	ContextMetadataTokenKey = "bearer"
)

// AuthenticationInterceptor an interceptor for authentication
func AuthenticationInterceptor(ctx context.Context) (context.Context, error) {
	// Extract token from metadata in the context
	tokenString, err := grpc_auth.AuthFromMD(ctx, ContextMetadataTokenKey)
	if err != nil {
		return nil, err
	}

	// Authenticate the jwt token
	jwtAuth := new(auth.JwtAuthenticator)
	authClaimsIf, err := jwtAuth.ParseAndValidate(tokenString)
	if err != nil {
		return ctx, err
	}

	niceMd := metautils.ExtractIncoming(ctx)

	authClaims, isMap := authClaimsIf.(jwt.MapClaims)
	if !isMap {
		return nil, fmt.Errorf("error converting claims to a map")
	}
	for k, v := range authClaims {
		switch vt := v.(type) {
		case string:
			niceMd.Set(k, vt)
		case float64:
			niceMd.Set(k, fmt.Sprintf("%v", vt))
		case []interface{}:
			items := make([]string, 0)
			for _, item := range vt {
				items = append(items, fmt.Sprintf("%v", item))
			}
			niceMd.Set(k, strings.Join(items, ";"))
		default:
			return nil, fmt.Errorf("unhandled type %T", vt)
		}
	}

	return niceMd.ToIncoming(ctx), nil
}
