// SPDX-FileCopyrightText: ${year}-present Open Networking Foundation <info@opennetworking.org>
// SPDX-License-Identifier: Apache-2.0

package listener

import (
	"fmt"
	"net"

	"github.com/onosproject/onos-lib-go/pkg/sctp/primitives"

	"github.com/onosproject/onos-lib-go/pkg/sctp/addressing"
	"github.com/onosproject/onos-lib-go/pkg/sctp/connection"
	"github.com/onosproject/onos-lib-go/pkg/sctp/defs"
)

// Listener sctp listener
type Listener struct {
	connection.Connection
	socketMode defs.SocketMode
}

// NewListener creates a new sctp listener
func NewListener(laddr *addressing.Address, init defs.InitMsg, mode defs.SocketMode, nonblocking bool) (*Listener, error) {
	if laddr == nil {
		return nil, fmt.Errorf("Local SCTPAddr is required")
	}

	conn, err := connection.NewConnection(laddr.AddressFamily, init, mode, nonblocking)
	if err != nil {
		return nil, err
	}
	ln := &Listener{Connection: *conn, socketMode: mode}
	ln.socketMode = mode

	if err := ln.Bind(laddr); err != nil {
		return nil, err
	}

	if err := ln.Listen(); err != nil {
		return nil, err
	}
	return ln, nil
}

// acceptSCTP waits for and returns the next SCTP connection to the listener.
func (ln *Listener) acceptSCTP() (*connection.Connection, error) {
	if ln.socketMode == defs.OneToMany {
		return nil, fmt.Errorf("Calling Accept on OneToMany socket is invalid")
	}

	fd, err := primitives.Accept(ln.FD())
	if err != nil {
		return nil, err
	}
	blocking, err := ln.GetNonBlocking()
	if err != nil {
		return nil, err
	}
	conn := &connection.Connection{}
	conn.SetFd(int32(fd))

	err = conn.SetNonBlocking(blocking)
	if err != nil {
		return nil, err
	}

	return conn, nil

}

// Accept waits for and returns the next connection connection to the listener.
func (ln *Listener) Accept() (net.Conn, error) {
	return ln.acceptSCTP()
}

// SCTPRead ...
func (ln *Listener) SCTPRead(b []byte) (int, *defs.OOBMessage, int, error) {
	if ln.socketMode == defs.OneToOne {
		return -1, nil, -1, fmt.Errorf("Invalid state: SCTPRead on OneToOne socket not allowed")
	}

	return ln.Connection.SCTPRead(b)
}

// SCTPWrite ...
func (ln *Listener) SCTPWrite(b []byte, info *defs.SndRcvInfo) (int, error) {
	if ln.socketMode == defs.OneToOne {
		return -1, fmt.Errorf("Invalid state: SCTPWrite on OneToOne socket not allowed")
	}

	return ln.Connection.SCTPWrite(b, info)
}
