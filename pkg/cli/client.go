// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package cli

import (
	"context"
	"crypto/tls"
	"github.com/onosproject/onos-lib-go/pkg/certs"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

const (
	// ServiceAddress command option
	ServiceAddress = "service-address"
	// TLSCertPathFlag command option
	TLSCertPathFlag = "tls-cert-path"
	// TLSKeyPathFlag command option
	TLSKeyPathFlag = "tls-key-path"
	// NoTLSFlag command option
	NoTLSFlag = "no-tls"
)

// GetConnection returns a gRPC client connection to the onos service
func GetConnection(cmd *cobra.Command) (*grpc.ClientConn, error) {
	address := getAddress(cmd)
	certPath := getCertPath(cmd)
	keyPath := getKeyPath(cmd)
	var opts []grpc.DialOption

	if noTLS(cmd) {
		opts = []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		}
	} else {
		if certPath != "" && keyPath != "" {
			cert, err := tls.LoadX509KeyPair(certPath, keyPath)
			if err != nil {
				return nil, err
			}
			opts = []grpc.DialOption{
				grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
					Certificates:       []tls.Certificate{cert},
					InsecureSkipVerify: true,
				})),
			}
		} else {
			// Load default Certificates
			cert, err := tls.X509KeyPair([]byte(certs.DefaultClientCrt), []byte(certs.DefaultClientKey))
			if err != nil {
				return nil, err
			}
			opts = []grpc.DialOption{
				grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
					Certificates:       []tls.Certificate{cert},
					InsecureSkipVerify: true,
				})),
			}
		}
	}

	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// NewContextWithAuthHeaderFromFlag - use from the CLI with --auth-header flag
func NewContextWithAuthHeaderFromFlag(ctx context.Context, authHeaderFlag *pflag.Flag) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	if authHeaderFlag != nil && authHeaderFlag.Value != nil && authHeaderFlag.Value.String() != "" {
		md := make(metadata.MD)
		md.Set("Authorization", authHeaderFlag.Value.String())
		ctx = metadata.NewOutgoingContext(ctx, md)
	}
	return ctx
}

// AddEndpointFlags adds service address, TLS cert path and TLS key path option to the command.
func AddEndpointFlags(cmd *cobra.Command, defaultAddress string) {
	cmd.Flags().String(ServiceAddress, defaultAddress, "service address; defaults to "+defaultAddress)
	cmd.Flags().String(TLSKeyPathFlag, "", "path to client private key")
	cmd.Flags().String(TLSCertPathFlag, "", "path to client certificate")
	cmd.Flags().Bool(NoTLSFlag, false, "if present, do not use TLS")
}

// GetServiceAddress returns the service address option value
func GetServiceAddress(cmd *cobra.Command) string {
	address, _ := cmd.Flags().GetString(ServiceAddress)
	return address
}

// GetTLSCertPath returns the TLS certificate path option value
func GetTLSCertPath(cmd *cobra.Command) string {
	certPath, _ := cmd.Flags().GetString(TLSCertPathFlag)
	return certPath
}

// GetTLSKeyPath returns the TLS key path option value
func GetTLSKeyPath(cmd *cobra.Command) string {
	keyPath, _ := cmd.Flags().GetString(TLSKeyPathFlag)
	return keyPath
}

// NoTLS returns true if the no-TLS flag is set
func NoTLS(cmd *cobra.Command) bool {
	tls, _ := cmd.Flags().GetBool(NoTLSFlag)
	return tls
}
