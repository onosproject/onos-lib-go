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
	"github.com/onosproject/onos-lib-go/api/rbac"
)

// DefaultRole
type DefaultRole struct {
	role *rbac.Role
}

const (
	SystemAdminRoleName = "system.admin"
)

var (
	defaultRoles = map[string]*rbac.Role{
		SystemAdminRoleName: &rbac.Role{
			Metadata: &rbac.Metadata{
				Name: SystemAdminRoleName,
			},
			Rules: []*rbac.Rule{
				{
					Groups: []string{
						"admin",
					},
					Services: []string{
						"*",
					},
					Verbs: []string{
						"*",
					},
				},
			},
		},
	}
)

// GetDefaultRoles returns the list of default roles in the system
func GetDefaultRoles() map[string]*rbac.Role {
	return defaultRoles
}
