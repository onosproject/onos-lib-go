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
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"os"
	"strings"
)

var log = logging.GetLogger("grpc", "auth")

const (
	// ContextMetadataTokenKey metadata token key
	ContextMetadataTokenKey = "bearer"
	allowMissingAuthClients = "ALLOW_MISSING_AUTH_CLIENTS"
)

// AuthenticationInterceptor an interceptor for authentication
// If there is an environment variable ALLOW_MISSING_AUTH_CLIENTS (and there is no
// Authentication Token attached to request) then look through its values and see
// the current requests client matches anything.
// ALLOW_MISSING_AUTH_CLIENTS is acomma separated list of client names
func AuthenticationInterceptor(ctx context.Context) (context.Context, error) {
	niceMd := metautils.ExtractIncoming(ctx)

	// Extract token from metadata in the context
	tokenString, err := grpc_auth.AuthFromMD(ctx, ContextMetadataTokenKey)
	if err != nil {
		acceptNoAuth := os.Getenv(allowMissingAuthClients)
		allowedMissingClients := strings.Split(acceptNoAuth, ",")
		requestClient := niceMd.Get("client")
		var foundMissingAuthClient bool
		for _, amc := range allowedMissingClients {
			if requestClient == strings.TrimSpace(strings.ToLower(amc)) {
				foundMissingAuthClient = true
				break
			}
		}
		if foundMissingAuthClient && err.Error() == `rpc error: code = Unauthenticated desc = Request unauthenticated with bearer` {
			log.Warnf("Allowing unauthenticated gRPC request from client: %s", niceMd.Get("client"))
			return ctx, nil
		}
		return nil, err
	}

	// Authenticate the jwt token
	jwtAuth := new(auth.JwtAuthenticator)
	authClaimsIf, err := jwtAuth.ParseAndValidate(tokenString)
	if err != nil {
		return ctx, err
	}

	authClaims, isMap := authClaimsIf.(jwt.MapClaims)
	if !isMap {
		return nil, fmt.Errorf("error converting claims to a map")
	}
	for k, v := range authClaims {
		err = HandleClaim(&niceMd, []string{k}, v)
		if err != nil {
			return nil, err
		}
	}

	log.Debugf("JWT token is valid, proceeding with processing")

	return niceMd.ToIncoming(ctx), nil
}

// HandleClaim function converts claims extracted from JWT to the string and appends them to the context
func HandleClaim(niceMd *metautils.NiceMD, key []string, value interface{}) error {
	k := strings.Join(key, "/")
	switch vt := value.(type) {
	case string:
		niceMd.Set(k, vt)
	case float64:
		niceMd.Set(k, fmt.Sprintf("%v", vt))
	case bool:
		if vt {
			niceMd.Set(k, "true")
		} else {
			niceMd.Set(k, "false")
		}
	case []interface{}:
		for _, item := range vt {
			niceMd.Add(k, fmt.Sprintf("%v", item))
		}
	case map[string]interface{}:
		for k, v := range vt {
			err := HandleClaim(niceMd, append(key, k), v)
			if err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("gRPC metadata unhandled type %T", vt)
	}
	return nil
}
