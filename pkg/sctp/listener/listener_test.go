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

// +build !darwin

package listener

import (
	"net"
	"testing"

	"github.com/onosproject/onos-lib-go/pkg/sctp/addressing"
	"github.com/onosproject/onos-lib-go/pkg/sctp/types"
)

var sctpListenerNameTests = []*addressing.Address{
	{IPAddrs: []net.IPAddr{{IP: net.IPv4(127, 0, 0, 1)}}},
	{},
	nil,
	{Port: 7777},
}

func TestSCTPListenerName(t *testing.T) {
	for _, tt := range sctpListenerNameTests {

		ln, err := NewListener(tt, WithInitMsg(types.InitMsg{}), WithMode(types.OneToOne), WithNonBlocking(false))
		if err != nil {
			if tt == nil {
				continue
			}
			t.Fatal(err)
		}
		defer ln.Close()
		la := ln.LocalAddr()
		if a, ok := la.(*addressing.Address); !ok || a.Port == 0 {
			t.Fatalf("got %v; expected a proper address with non-zero port number", la)
		}

	}
}
