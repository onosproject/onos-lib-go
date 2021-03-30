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
	"bytes"
	"net"
	"strconv"
	"strings"

	"github.com/onosproject/onos-lib-go/pkg/sctp/defs"
	"github.com/onosproject/onos-lib-go/pkg/sctp/utils"

	"github.com/onosproject/onos-lib-go/pkg/errors"

	syscall "golang.org/x/sys/unix"
)

// Address sctp address
type Address struct {
	IPAddrs       []net.IPAddr
	Port          int
	AddressFamily defs.AddressFamily
}

// ResolveAddress resolves sctp address
func ResolveAddress(addressFamily defs.AddressFamily, addrs string) (*Address, error) {
	elems := strings.Split(addrs, "/")
	if len(elems) == 0 {
		return nil, errors.NewInvalid("invalid address:")
	}

	lastE := elems[len(elems)-1]
	ipaddrs := make([]net.IPAddr, 0, len(elems))
	addr, port, err := net.SplitHostPort(lastE)
	if err != nil {
		return nil, errors.NewInvalid("invalid port:", err.Error())

	}

	if port == "" {
		return nil, errors.NewInvalid("port cannot be empty")
	}

	iPort, err := strconv.Atoi(port)
	if err != nil {
		return nil, errors.NewInvalid("invalid input: Non-integer port: %s", addrs)
	}

	elems[len(elems)-1] = addr
	for _, e := range elems {
		family := addressFamily.String()
		if !strings.Contains(e, ":") && addressFamily == defs.Sctp6 {
			family = defs.Sctp4.String()
		}
		ipa, err := net.ResolveIPAddr(family, e)
		if err != nil {
			return nil, err
		}

		if ipa.IP != nil {
			if ipa.IP.To4() == nil {
				if addressFamily == defs.Sctp4 {
					return nil, errors.New(errors.NotFound, "IPv6 address detected but addressFamily is IPv4")
				}
			}
			ipaddrs = append(ipaddrs, net.IPAddr{IP: ipa.IP, Zone: ipa.Zone})
		} else {
			var ip net.IPAddr
			switch addressFamily {
			case defs.Sctp4:
				ip = net.IPAddr{IP: net.IPv4zero, Zone: ""}
			case defs.Sctp6:
				ip = net.IPAddr{IP: net.IPv6zero, Zone: ""}
			default:
				return nil, errors.NewUnknown("Unknown addressFamily: %s", addressFamily)
			}
			ipaddrs = append(ipaddrs, ip)
		}
	}

	return &Address{
		IPAddrs:       ipaddrs,
		Port:          iPort,
		AddressFamily: addressFamily,
	}, nil
}

// Network returns network
func (a *Address) Network() string { return "sctp" }

// ToRawSockAddrBuf ...
func (a *Address) ToRawSockAddrBuf() []byte {
	p := utils.Htons(uint16(a.Port))
	if len(a.IPAddrs) == 0 { // if a.IPAddrs list is empty - fall back to IPv4 zero addr
		s := syscall.RawSockaddrInet4{
			Family: syscall.AF_INET,
			Port:   p,
		}
		copy(s.Addr[:], net.IPv4zero)
		return utils.ToBuf(s)
	}
	buf := []byte{}
	for _, ip := range a.IPAddrs {
		ipBytes := ip.IP
		if len(ipBytes) == 0 {
			ipBytes = net.IPv4zero
		}
		if ip4 := ipBytes.To4(); ip4 != nil {
			s := syscall.RawSockaddrInet4{
				Family: syscall.AF_INET,
				Port:   p,
			}
			copy(s.Addr[:], ip4)
			buf = append(buf, utils.ToBuf(s)...)
		} else {
			var scopeid uint32
			ifi, err := net.InterfaceByName(ip.Zone)
			if err == nil {
				scopeid = uint32(ifi.Index)
			}
			s := syscall.RawSockaddrInet6{
				Family:   syscall.AF_INET6,
				Port:     p,
				Scope_id: scopeid,
			}
			copy(s.Addr[:], ipBytes)
			buf = append(buf, utils.ToBuf(s)...)
		}
	}
	return buf
}

// String converts to string representation
func (a *Address) String() string {
	var b bytes.Buffer

	for n, i := range a.IPAddrs {
		if i.IP.To4() != nil {
			b.WriteString(i.String())
		} else if i.IP.To16() != nil {
			if n == len(a.IPAddrs)-1 {
				b.WriteRune('[')
				b.WriteString(i.String())
				b.WriteRune(']')
			} else {
				b.WriteString(i.String())
			}
		}
		if n < len(a.IPAddrs)-1 {
			b.WriteRune('/')
		}
	}
	b.WriteRune(':')
	b.WriteString(strconv.Itoa(a.Port))
	return b.String()
}

/*func (a *Address) isWildcard() bool {
	if a == nil {
		return true
	}
	if 0 == len(a.IPAddrs) {
		return true
	}

	return a.IPAddrs[0].IP.IsUnspecified()
}*/

// GetAddressFamily returns family address
func GetAddressFamily(laddr *Address, raddr *Address) (family int, ipv6only bool) {

	if laddr != nil && raddr != nil {
		if laddr.AddressFamily == raddr.AddressFamily {
			return laddr.AddressFamily.ToSyscall(), (laddr.AddressFamily == defs.Sctp6)
		}

		if utils.SupportsIPv4map() || !utils.SupportsIPv4() {
			return defs.Sctp6.ToSyscall(), false
		}
	}
	return defs.Sctp4.ToSyscall(), false
}
