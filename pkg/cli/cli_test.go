// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package cli

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	viper.Set(addressKey, "http://localhost")
	conn, err := GetConnection(&cobra.Command{
		Short: "test command",
	})
	assert.NilError(t, err)
	assert.Assert(t, conn != nil)
}

func Test_AddConfigFlags(_ *testing.T) {
	AddConfigFlags(&cobra.Command{
		Short: "test command",
	}, "localhost:5150")
}

func Test_GetConfigCommand(_ *testing.T) {
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
