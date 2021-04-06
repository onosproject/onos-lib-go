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

package sctp

import (
	"net"

	"github.com/onosproject/onos-lib-go/pkg/sctp/addressing"

	"github.com/onosproject/onos-lib-go/pkg/sctp/connection"

	"github.com/onosproject/onos-lib-go/pkg/sctp/defs"
)

/*var nativeEndian binary.ByteOrder
var sndRcvInfoSize uintptr

func init() {
	i := uint16(1)
	if *(*byte)(unsafe.Pointer(&i)) == 0 {
		nativeEndian = binary.BigEndian
	} else {
		nativeEndian = binary.LittleEndian
	}
	sndRcvInfoSize = unsafe.Sizeof(defs.SndRcvInfo{})
}*/

// DialOptions is SCTP options
type DialOptions struct {
	addressFamily defs.AddressFamily
	mode          defs.SocketMode
	options       defs.InitMsg
	nonblocking   bool
}

// NewDialOptions creates dial options
func NewDialOptions(options ...func(options *DialOptions)) *DialOptions {
	dialOptions := &DialOptions{}
	for _, option := range options {
		option(dialOptions)
	}

	return dialOptions
}

// WithAddressFamily sets address family
func WithAddressFamily(family defs.AddressFamily) func(options *DialOptions) {
	return func(options *DialOptions) {
		options.addressFamily = family

	}
}

// WithMode sets SCTP mode
func WithMode(mode defs.SocketMode) func(options *DialOptions) {
	return func(options *DialOptions) {
		options.mode = mode

	}
}

// WithOptions sets options
func WithOptions(initOptions defs.InitMsg) func(options *DialOptions) {
	return func(options *DialOptions) {
		options.options = initOptions

	}
}

// WithNonBlocking sets nonblocking
func WithNonBlocking(nonblocking bool) func(options *DialOptions) {
	return func(options *DialOptions) {
		options.nonblocking = nonblocking
	}
}

// DialSCTP creates a new SCTP connection
func DialSCTP(addr net.Addr, dialOptions *DialOptions) (*connection.SCTPConn, error) {
	cfg := connection.NewConfig(
		connection.WithAddressFamily(dialOptions.addressFamily),
		connection.WithOptions(dialOptions.options),
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
