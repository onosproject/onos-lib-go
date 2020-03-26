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

package cli

import (
	"github.com/spf13/cobra"
	"gotest.tools/assert"
	"testing"
)

// These are mainly placeholders for the moment for functions that are not
// otherwise called
func Test_InitConfig(t *testing.T) {
	InitConfig("test-module")

	assert.Equal(t, "test-module", configName)
}

func Test_GetConnection(t *testing.T) {
	conn, err := GetConnection(&cobra.Command{
		Short: "test command",
	})
	assert.NilError(t, err)
	assert.Assert(t, conn != nil)
}

func Test_AddConfigFlags(t *testing.T) {
	AddConfigFlags(&cobra.Command{
		Short: "test command",
	},
		"localhost:5150")
}

func Test_GetConfigCommand(t *testing.T) {
	GetConfigCommand()
}

func Test_RunConfigSetCommand(t *testing.T) {
	err := runConfigSetCommand(&cobra.Command{
		Short: "test command",
	}, []string{"a", "b", "c"})
	assert.ErrorContains(t, err, "Config File \"test-module\" Not Found")
}

func Test_GetCertPath(t *testing.T) {
	path := getCertPath(&cobra.Command{
		Short: "test command",
	})
	assert.Assert(t, len(path) == 0)
}
