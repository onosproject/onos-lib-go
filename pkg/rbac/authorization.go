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

package rbac

import (
	"fmt"

	"strings"

	"github.com/dgrijalva/jwt-go"
	api "github.com/onosproject/onos-lib-go/api/rbac"
	"google.golang.org/grpc"
)

const (
	// GroupsKey groups key in the token claims
	GroupsKey = "groups"
)

func verifyRules(rules []*api.Rule, fullMethod string) error {
	// Retrieve service information and rpc method name
	reqService, reqVerb := getMethodInformation(fullMethod)
	for _, rule := range rules {
		for _, service := range rule.Services {
			if match(strings.ToLower(service), strings.ToLower(reqService)) {
				for _, verb := range rule.Verbs {
					if match(strings.ToLower(verb), reqVerb) {
						return nil
					}
				}
			}
		}
	}

	return fmt.Errorf("no rule found to authorize the user")
}

func extractClaimedGroups(claims jwt.MapClaims) ([]string, error) {
	var claimedGroups []interface{}
	for key := range claims {
		// extract claimed groups from the token
		if key == GroupsKey {
			claimedGroups = claims[GroupsKey].([]interface{})
		}
	}

	// If the user does not claim any groups then we cannot authorize the user
	if claimedGroups == nil {
		return nil, fmt.Errorf("groups claim cannot be empty")
	}

	var claimedGroupsList []string
	for _, group := range claimedGroups {
		claimedGroupsList = append(claimedGroupsList, group.(string))
	}

	return claimedGroupsList, nil
}

func findCandidateRules(roles map[string]*api.Role, claimedGroupsList []string) ([]*api.Rule, error) {
	var candidateRules []*api.Rule
	for _, role := range defaultRoles {
		rules := role.Rules
		for _, rule := range rules {
			matched := matchGroups(rule.Groups, claimedGroupsList)
			if matched {
				candidateRules = append(candidateRules, rule)
			}
		}
	}

	return candidateRules, nil

}

// Authorize authorize a user based on given claims
func Authorize(claims jwt.MapClaims, info *grpc.UnaryServerInfo) error {

	// Extracts claimed groups
	claimedGroupsList, err := extractClaimedGroups(claims)
	if err != nil {
		return err
	}

	// Checks the default roles first to authorize the users
	roles := GetRoles()

	// Finds list of candidate rules that should be should be checked for verification
	candidateRules, err := findCandidateRules(roles, claimedGroupsList)

	// verifies list of candidate rules
	err = verifyRules(candidateRules, info.FullMethod)
	if err != nil {
		return err
	}

	return nil
}
