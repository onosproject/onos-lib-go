// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package sctp

import (
	"net"

	"github.com/onosproject/onos-lib-go/pkg/sctp/addressing"

	"github.com/onosproject/onos-lib-go/pkg/sctp/connection"

	"github.com/onosproject/onos-lib-go/pkg/sctp/types"
)

// DialOptions is SCTP options
type DialOptions struct {
	addressFamily types.AddressFamily
	mode          types.SocketMode
	initMsg       types.InitMsg
	nonblocking   bool
}

// DialOption dial option function
type DialOption func(*DialOptions)

// WithAddressFamily sets address family
func WithAddressFamily(family types.AddressFamily) func(options *DialOptions) {
	return func(options *DialOptions) {
		options.addressFamily = family

	}
}

// WithMode sets SCTP mode
func WithMode(mode types.SocketMode) func(options *DialOptions) {
	return func(options *DialOptions) {
		options.mode = mode

	}
}

// WithInitMsg sets options
func WithInitMsg(initMsg types.InitMsg) func(options *DialOptions) {
	return func(options *DialOptions) {
		options.initMsg = initMsg

	}
}

// WithNonBlocking sets nonblocking
func WithNonBlocking(nonblocking bool) func(options *DialOptions) {
	return func(options *DialOptions) {
		options.nonblocking = nonblocking
	}
}

// DialSCTP creates a new SCTP connection
func DialSCTP(addr net.Addr, opts ...DialOption) (*connection.SCTPConn, error) {
	dialOptions := &DialOptions{}
	for _, option := range opts {
		option(dialOptions)
	}
	cfg := connection.NewConfig(
		connection.WithAddressFamily(dialOptions.addressFamily),
		connection.WithOptions(dialOptions.initMsg),
		connection.WithMode(dialOptions.mode),
		connection.WithNonBlocking(dialOptions.nonblocking))
	conn, err := connection.NewSCTPConnection(cfg)
	if err != nil {
		return nil, err
	}
	sctpAddress := addr.(*addressing.Address)
	err = conn.Connect(sctpAddress)
	if err != nil {
		return nil, err
	}

	return conn, nil

}
