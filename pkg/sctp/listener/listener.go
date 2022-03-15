// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package listener

import (
	"fmt"
	"net"

	"github.com/onosproject/onos-lib-go/pkg/errors"

	"github.com/onosproject/onos-lib-go/pkg/sctp/connection"

	"github.com/onosproject/onos-lib-go/pkg/sctp/addressing"
	"github.com/onosproject/onos-lib-go/pkg/sctp/types"
)

// Listener SCTP listener
type Listener struct {
	connection.SCTPConn
	socketMode types.SocketMode
}

// Options SCTP listener options
type Options struct {
	mode        types.SocketMode
	initMsg     types.InitMsg
	nonblocking bool
}

// Option listener options
type Option func(options *Options)

// WithMode sets SCTP mode
func WithMode(mode types.SocketMode) func(options *Options) {
	return func(options *Options) {
		options.mode = mode

	}
}

// WithInitMsg sets options
func WithInitMsg(initMsg types.InitMsg) func(options *Options) {
	return func(options *Options) {
		options.initMsg = initMsg

	}
}

// WithNonBlocking sets nonblocking
func WithNonBlocking(nonblocking bool) func(options *Options) {
	return func(options *Options) {
		options.nonblocking = nonblocking
	}
}

// NewListener creates a new SCTP listener instance
func NewListener(laddr *addressing.Address, opts ...Option) (*Listener, error) {
	if laddr == nil {
		return nil, errors.NewInvalid("Local SCTPAddr is required")
	}
	listenerOptions := &Options{}
	for _, option := range opts {
		option(listenerOptions)
	}

	cfg := connection.NewConfig(
		connection.WithAddressFamily(laddr.AddressFamily),
		connection.WithOptions(listenerOptions.initMsg),
		connection.WithMode(listenerOptions.mode),
		connection.WithNonBlocking(listenerOptions.nonblocking))
	conn, err := connection.NewSCTPConnection(cfg)
	if err != nil {
		return nil, err
	}
	ln := &Listener{SCTPConn: *conn, socketMode: listenerOptions.mode}
	ln.socketMode = listenerOptions.mode

	if err := ln.Bind(laddr); err != nil {
		return nil, err
	}

	if err := ln.Listen(); err != nil {
		return nil, err
	}
	return ln, nil
}

// accept waits for and returns the next SCTP connection to the listener.
func (ln *Listener) accept() (*connection.SCTPConn, error) {
	if ln.socketMode == types.OneToMany {
		return nil, fmt.Errorf("Calling Accept on OneToMany socket is invalid")
	}

	fd, err := connection.Accept(ln.FD())
	if err != nil {
		return nil, err
	}
	blocking, err := ln.GetNonblocking()
	if err != nil {
		return nil, err
	}
	conn := &connection.SCTPConn{}
	conn.SetFD(int32(fd))

	err = conn.SetNonblocking(blocking)
	if err != nil {
		return nil, err
	}

	return conn, nil

}

// Accept waits for and returns the next connection connection to the listener.
func (ln *Listener) Accept() (net.Conn, error) {
	return ln.accept()
}

// SCTPRead reads from an SCTP connection
func (ln *Listener) SCTPRead(b []byte) (int, *types.OOBMessage, int, error) {
	if ln.socketMode == types.OneToOne {
		return -1, nil, -1, errors.NewInvalid("Invalid state: SCTPRead on OneToOne socket not allowed")
	}

	return ln.SCTPConn.SCTPRead(b)
}

// SCTPWrite writes on an SCTP connection
func (ln *Listener) SCTPWrite(b []byte, info *types.SndRcvInfo) (int, error) {
	if ln.socketMode == types.OneToOne {
		return -1, errors.NewInvalid("Invalid state: SCTPWrite on OneToOne socket not allowed")
	}

	return ln.SCTPConn.SCTPWrite(b, info)
}
