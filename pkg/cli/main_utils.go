// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package cli

import (
	"fmt"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

const (
	// BindPortFlag command option
	BindPortFlag = "bind-port"
	// TLSCACertPathFlag command option
	TLSCACertPathFlag = "tls-ca-cert-path"
	// DefaultBindPort command options
	DefaultBindPort = 5150
)

// Run runs the body of the given root command; exists with status 1 if an error is encountered
func Run(rootCommand *cobra.Command) {
	if err := rootCommand.Execute(); err != nil {
		println(err)
		os.Exit(1)
	}
}

// Daemon is a simple abstraction of a process that can be started and stopped
type Daemon interface {
	// Start starts the daemon in the background
	Start() error
	// Stop stops the daemon
	Stop()
}

// RunDaemon starts the given deamon, waits until it receives SIGTERM and then stops the daemon.
func RunDaemon(daemon Daemon) error {
	if err := daemon.Start(); err != nil {
		return err
	}
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh
	daemon.Stop()
	return nil
}

// AddServiceEndpointFlags injects standard server service endpoint flags to the given command
func AddServiceEndpointFlags(cmd *cobra.Command, name string) {
	cmd.Flags().Int(BindPortFlag, DefaultBindPort, fmt.Sprintf("listen TCP port of the %s service", name))
	cmd.Flags().Bool(NoTLSFlag, false, fmt.Sprintf("if set, do not use TLS for the %s service", name))
	cmd.Flags().String(TLSCACertPathFlag, "", "path to CA certificate")
	cmd.Flags().String(TLSKeyPathFlag, "", "path to client private key")
	cmd.Flags().String(TLSCertPathFlag, "", "path to client certificate")
}

// ServiceEndpointFlags carries values extracted from the command-line options
type ServiceEndpointFlags struct {
	BindPort int
	NoTLS    bool
	CAPath   string
	KeyPath  string
	CertPath string
}

// ExtractServiceEndpointFlags extracts the standard server service endpont flags from the given command
func ExtractServiceEndpointFlags(cmd *cobra.Command) (*ServiceEndpointFlags, error) {
	var err error
	flags := &ServiceEndpointFlags{}

	if flags.BindPort, err = cmd.Flags().GetInt(BindPortFlag); err != nil {
		return nil, err
	}
	if flags.NoTLS, err = cmd.Flags().GetBool(NoTLSFlag); err != nil {
		return nil, err
	}
	flags.CAPath, _ = cmd.Flags().GetString(TLSCACertPathFlag)
	flags.KeyPath, _ = cmd.Flags().GetString(TLSKeyPathFlag)
	flags.CertPath, _ = cmd.Flags().GetString(TLSCertPathFlag)

	// If no certs and key were specified, revert to no-TLS mode
	if len(flags.CAPath) == 0 && len(flags.KeyPath) == 0 && len(flags.CertPath) == 0 {
		flags.NoTLS = true
	}
	return flags, nil
}

// ServerConfigFromFlags creates a new server config based on the specified flags and security config descriptor
func ServerConfigFromFlags(flags *ServiceEndpointFlags, secCfg northbound.SecurityConfig) *northbound.ServerConfig {
	cfg := northbound.NewInsecureServerConfig(int16(flags.BindPort))
	if !flags.NoTLS {
		cfg = northbound.NewServerCfg(flags.CAPath, flags.KeyPath, flags.CertPath, int16(flags.BindPort), false, secCfg)
	}
	return cfg
}
