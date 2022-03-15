// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package southbound

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/onosproject/onos-lib-go/pkg/certs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Connect establishes a client-side connection to the gRPC end-point.
func Connect(ctx context.Context, address string, certPath string, keyPath string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	var tlsOpts []grpc.DialOption
	if certPath != "" && keyPath != "" {
		cert, err := tls.LoadX509KeyPair(certPath, keyPath)
		if err != nil {
			return nil, err
		}
		tlsOpts = []grpc.DialOption{
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
		tlsOpts = []grpc.DialOption{
			grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
				Certificates:       []tls.Certificate{cert},
				InsecureSkipVerify: true,
			})),
		}
	}

	opts = append(tlsOpts, opts...)
	conn, err := grpc.DialContext(ctx, address, opts...)
	if err != nil {
		fmt.Println("Can't connect", err)
		return nil, err
	}
	return conn, nil
}
