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

package addressing

import (
	"net"
	"reflect"
	"testing"

	"github.com/onosproject/onos-lib-go/pkg/sctp/defs"
	"github.com/stretchr/testify/assert"
)

type resolveSCTPAddrTest struct {
	network       defs.AddressFamily
	litAddrOrName string
	addr          *Address
	err           error
}

var ipv4loop = net.IPv4(127, 0, 0, 1)

var resolveSCTPAddrTests = []resolveSCTPAddrTest{
	{defs.Sctp4, "127.0.0.1:0", &Address{AddressFamily: defs.Sctp4, IPAddrs: []net.IPAddr{{IP: ipv4loop}}, Port: 0}, nil},
	{defs.Sctp4, "127.0.0.1:65535", &Address{AddressFamily: defs.Sctp4, IPAddrs: []net.IPAddr{{IP: ipv4loop}}, Port: 65535}, nil},

	{defs.Sctp6, "[::1]:0", &Address{AddressFamily: defs.Sctp6, IPAddrs: []net.IPAddr{{IP: net.ParseIP("::1")}}, Port: 0}, nil},
	{defs.Sctp6, "[::1]:65535", &Address{AddressFamily: defs.Sctp6, IPAddrs: []net.IPAddr{{IP: net.ParseIP("::1")}}, Port: 65535}, nil},

	{defs.Sctp6, "[::1%lo0]:0", &Address{AddressFamily: defs.Sctp6, IPAddrs: []net.IPAddr{{IP: net.ParseIP("::1"), Zone: "lo0"}}, Port: 0}, nil},
	{defs.Sctp6, "[::1%lo0]:65535", &Address{AddressFamily: defs.Sctp6, IPAddrs: []net.IPAddr{{IP: net.ParseIP("::1"), Zone: "lo0"}}, Port: 65535}, nil},
	{defs.Sctp4, "0.0.0.0:12345", &Address{AddressFamily: defs.Sctp4, IPAddrs: []net.IPAddr{{IP: net.IPv4zero, Zone: ""}}, Port: 12345}, nil},
	{defs.Sctp4, "127.0.0.1/10.0.0.1:0", &Address{IPAddrs: []net.IPAddr{{IP: net.IPv4(127, 0, 0, 1)}, {IP: net.IPv4(10, 0, 0, 1)}}, Port: 0}, nil},
	{defs.Sctp4, "127.0.0.1/10.0.0.1:65535", &Address{IPAddrs: []net.IPAddr{{IP: net.IPv4(127, 0, 0, 1)}, {IP: net.IPv4(10, 0, 0, 1)}}, Port: 65535}, nil},
	{defs.Sctp6, "::1%lo0/127.0.0.1:1234", &Address{AddressFamily: defs.Sctp6, IPAddrs: []net.IPAddr{{IP: net.ParseIP("::1"), Zone: "lo0"}, {IP: ipv4loop, Zone: ""}}, Port: 1234}, nil},
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
		if !reflect.DeepEqual(addr, tt.addr) || !reflect.DeepEqual(err, tt.err) {
			t.Errorf("ResolveSCTPAddr(%q, %q) = %#v, %v, want %#v, %v", tt.network, tt.litAddrOrName, addr, err, tt.addr, tt.err)
			continue
		}
		if err == nil {
			addr2, err := ResolveAddress(addr.AddressFamily, addr.String())
			if !reflect.DeepEqual(addr2, tt.addr) || err != tt.err {
				t.Errorf("(%q, %q): ResolveSCTPAddr(%q, %q) = %#v, %v, want %#v, %v", tt.network, tt.litAddrOrName, addr.Network(), addr.String(), addr2, err, tt.addr, tt.err)
			}
		}
	}
}
