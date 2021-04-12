// Copyright 2021-present Open Networking Foundation.
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

// NewListener creates a new SCTP listener instance
func NewListener(laddr *addressing.Address, options types.InitMsg, mode types.SocketMode, nonblocking bool) (*Listener, error) {
	if laddr == nil {
		return nil, errors.NewInvalid("Local SCTPAddr is required")
	}

	cfg := connection.NewConfig(
		connection.WithAddressFamily(laddr.AddressFamily),
		connection.WithOptions(options),
		connection.WithMode(mode),
		connection.WithNonBlocking(nonblocking))
	conn, err := connection.NewSCTPConnection(cfg)
	if err != nil {
		return nil, err
	}
	ln := &Listener{SCTPConn: *conn, socketMode: mode}
	ln.socketMode = mode

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
