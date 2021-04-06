// SPDX-FileCopyrightText: ${year}-present Open Networking Foundation <info@opennetworking.org>
// SPDX-License-Identifier: Apache-2.0

package listener

import (
	"net"
	"testing"

	"github.com/onosproject/onos-lib-go/pkg/sctp/addressing"
	"github.com/onosproject/onos-lib-go/pkg/sctp/defs"
)

var sctpListenerNameTests = []*addressing.Address{
	&addressing.Address{IPAddrs: []net.IPAddr{net.IPAddr{IP: net.IPv4(127, 0, 0, 1)}}},
	&addressing.Address{},
	nil,
	&addressing.Address{Port: 7777},
}

func TestSCTPListenerName(t *testing.T) {
	for _, tt := range sctpListenerNameTests {
		ln, err := NewListener(tt, defs.InitMsg{}, defs.OneToOne, false)
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
