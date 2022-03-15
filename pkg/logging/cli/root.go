// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package cli

import (
	"github.com/onosproject/onos-lib-go/pkg/cli"
	"github.com/spf13/cobra"
)

//var viper = viperapi.New()

// init initializes the command line
func init() {
	cli.InitConfig("logging")
}

// Init is a hook called after cobra initialization
func Init() {
	// noop for now
}

// GetCommand returns the root command for the logging service
func GetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "log {set/get} level [args]",
		Short: "logging api commands",
	}

	cmd.AddCommand(getSetCommand())
	cmd.AddCommand(getGetCommand())

	return cmd
}
