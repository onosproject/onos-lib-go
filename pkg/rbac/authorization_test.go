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
	"testing"

	"github.com/stretchr/testify/assert"

	api "github.com/onosproject/onos-lib-go/api/rbac"
)

func TestAuthorize(t *testing.T) {

}

func TestVerifyRules(t *testing.T) {
	tests := []struct {
		denied     bool
		fullMethod string
		rules      []*api.Rule
	}{
		{
			denied:     false,
			fullMethod: "/onos.config.admin.ConfigAdminService/UploadRegisterModel",
			rules: []*api.Rule{
				{
					Groups: []string{
						"*",
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
		{
			denied:     true,
			fullMethod: "/onos.config.admin.ConfigAdminService/UploadRegisterModel",
			rules: []*api.Rule{
				{
					Groups: []string{
						"*",
					},

					Services: []string{
						"*",
					},
					Verbs: []string{
						"ListRegisteredModels",
						"RollbackNetworkChange",
						"ListSnapshots",
						"CompactChanges",
					},
				},
			},
		},
		{
			denied:     false,
			fullMethod: "/onos.config.admin.ConfigAdminService/UploadRegisterModel",
			rules: []*api.Rule{
				{
					Groups: []string{
						"*",
					},

					Services: []string{
						"*",
					},
					Verbs: []string{
						"UploadRegisterModel",
					},
				},
			},
		},
		{
			denied:     true,
			fullMethod: "/onos.config.admin.ConfigAdminService/UploadRegisterModel",
			rules: []*api.Rule{
				{
					Groups: []string{
						"*",
					},

					Services: []string{
						"TopoAdminService",
					},
					Verbs: []string{
						"*",
					},
				},
			},
		},
		{
			denied:     false,
			fullMethod: "/onos.config.admin.ConfigAdminService/UploadRegisterModel",
			rules: []*api.Rule{
				{
					Groups: []string{
						"*",
					},

					Services: []string{
						"ConfigAdminService",
					},
					Verbs: []string{
						"*",
					},
				},
			},
		},
		{
			denied:     true,
			fullMethod: "/onos.config.admin.ConfigAdminService/UploadRegisterModel",
			rules: []*api.Rule{
				{
					Groups: []string{
						"*",
					},

					Services: []string{
						"ConfigAdminService",
					},
					Verbs: []string{
						"CompactChanges",
					},
				},
			},
		},
		{
			denied:     false,
			fullMethod: "/onos.config.admin.ConfigAdminService/UploadRegisterModel",
			rules: []*api.Rule{
				{
					Groups: []string{
						"*",
					},

					Services: []string{
						"ConfigAdminService",
					},
					Verbs: []string{
						"UploadRegisterModel",
					},
				},
			},
		},
		{
			denied:     false,
			fullMethod: "/onos.config.admin.ConfigAdminService/UploadRegisterModel",
			rules: []*api.Rule{
				{
					Groups: []string{
						"*",
					},

					Services: []string{
						"configAdminService",
					},
					Verbs: []string{
						"uploadRegisterModel",
					},
				},
			},
		},
	}

	for _, test := range tests {
		err := verifyRules(test.rules, test.fullMethod)
		if test.denied {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}

	}

}
