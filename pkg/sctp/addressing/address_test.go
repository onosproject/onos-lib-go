// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package addressing

import (
	"net"
	"testing"

	"github.com/onosproject/onos-lib-go/pkg/sctp/types"
	"github.com/stretchr/testify/assert"
)

type resolveSCTPAddrTest struct {
	network       types.AddressFamily
	litAddrOrName string
	addr          *Address
}

var ipv4loop = net.IPv4(127, 0, 0, 1)

var resolveSCTPAddrTests = []resolveSCTPAddrTest{
	{types.Sctp4, "127.0.0.1:0", &Address{AddressFamily: types.Sctp4, IPAddrs: []net.IPAddr{{IP: ipv4loop}}, Port: 0}},
	{types.Sctp4, "127.0.0.1:65535", &Address{AddressFamily: types.Sctp4, IPAddrs: []net.IPAddr{{IP: ipv4loop}}, Port: 65535}},

	{types.Sctp6, "[::1]:0", &Address{AddressFamily: types.Sctp6, IPAddrs: []net.IPAddr{{IP: net.ParseIP("::1")}}, Port: 0}},
	{types.Sctp6, "[::1]:65535", &Address{AddressFamily: types.Sctp6, IPAddrs: []net.IPAddr{{IP: net.ParseIP("::1")}}, Port: 65535}},

	{types.Sctp6, "[::1%lo0]:0", &Address{AddressFamily: types.Sctp6, IPAddrs: []net.IPAddr{{IP: net.ParseIP("::1"), Zone: "lo0"}}, Port: 0}},
	{types.Sctp6, "[::1%lo0]:65535", &Address{AddressFamily: types.Sctp6, IPAddrs: []net.IPAddr{{IP: net.ParseIP("::1"), Zone: "lo0"}}, Port: 65535}},
	{types.Sctp4, "0.0.0.0:12345", &Address{AddressFamily: types.Sctp4, IPAddrs: []net.IPAddr{{IP: net.IPv4zero, Zone: ""}}, Port: 12345}},
	{types.Sctp4, "127.0.0.1/10.0.0.1:0", &Address{IPAddrs: []net.IPAddr{{IP: net.IPv4(127, 0, 0, 1)}, {IP: net.IPv4(10, 0, 0, 1)}}, Port: 0}},
	{types.Sctp4, "127.0.0.1/10.0.0.1:65535", &Address{IPAddrs: []net.IPAddr{{IP: net.IPv4(127, 0, 0, 1)}, {IP: net.IPv4(10, 0, 0, 1)}}, Port: 65535}},
	{types.Sctp6, "::1%lo0/127.0.0.1:1234", &Address{AddressFamily: types.Sctp6, IPAddrs: []net.IPAddr{{IP: net.ParseIP("::1"), Zone: "lo0"}, {IP: ipv4loop, Zone: ""}}, Port: 1234}},
}

func TestSCTPAddrString(t *testing.T) {
	for _, tt := range resolveSCTPAddrTests {
		s := tt.addr.String()
		assert.Equal(t, tt.litAddrOrName, s)
	}
}

func TestResolveSCTPAddr(t *testing.T) {
	for _, tt := range resolveSCTPAddrTests {
		addr, err := ResolveAddress(tt.network, tt.litAddrOrName)
		assert.NoError(t, err)
		assert.Equal(t, tt.addr, addr)
		addr2, err := ResolveAddress(addr.AddressFamily, addr.String())
		assert.NoError(t, err)
		assert.Equal(t, tt.addr, addr2)
	}
}
