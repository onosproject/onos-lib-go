// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//go:build !darwin
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
