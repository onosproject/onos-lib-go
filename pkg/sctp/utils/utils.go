// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"bytes"
	"encoding/binary"
	"net"
	"sync"
	"unsafe"

	syscall "golang.org/x/sys/unix"
)

// BoolToInt ...
func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func getNativeByteOrder() binary.ByteOrder {
	i := uint16(1)
	var nativeEndian binary.ByteOrder
	if *(*byte)(unsafe.Pointer(&i)) == 0 {
		nativeEndian = binary.BigEndian
	} else {
		nativeEndian = binary.LittleEndian
	}

	return nativeEndian
}

// ToBuf ...
func ToBuf(v interface{}) []byte {
	var buf bytes.Buffer
	binary.Write(&buf, getNativeByteOrder(), v)
	return buf.Bytes()
}

// Htons ...
func Htons(h uint16) uint16 {
	if getNativeByteOrder() == binary.LittleEndian {
		return (h << 8 & 0xff00) | (h >> 8 & 0xff)
	}
	return h
}

// from https://github.com/golang/go
type ipStackCapabilities struct {
	sync.Once             // guards following
	ipv4Enabled           bool
	ipv6Enabled           bool
	ipv4MappedIPv6Enabled bool
}

// from https://github.com/golang/go
var ipStackCaps ipStackCapabilities

// SupportsIPv4 supportsIPv4 reports whether the platform supports IPv4 networking
// functionality.
func SupportsIPv4() bool {
	ipStackCaps.Once.Do(ipStackCaps.probe)
	return ipStackCaps.ipv4Enabled
}

//from https://github.com/golang/go
// supportsIPv6 reports whether the platform supports IPv6 networking
// functionality.
/*func supportsIPv6() bool {
	ipStackCaps.Once.Do(ipStackCaps.probe)
	return ipStackCaps.ipv6Enabled
}*/

// SupportsIPv4map supportsIPv4map reports whether the platform supports mapping an
//
//	IPv4 address inside an IPv6 address at transport layer
//	protocols. See RFC 4291, RFC 4038 and RFC 3493.
func SupportsIPv4map() bool {
	ipStackCaps.Once.Do(ipStackCaps.probe)
	return ipStackCaps.ipv4MappedIPv6Enabled
}

// from https://github.com/golang/go
// Probe probes IPv4, IPv6 and IPv4-mapped IPv6 communication
// capabilities which are controlled by the IPV6_V6ONLY socket option
// and kernel configuration.
//
// Should we try to use the IPv4 socket interface if we're only
// dealing with IPv4 sockets? As long as the host system understands
// IPv4-mapped IPv6, it's okay to pass IPv4-mapeed IPv6 addresses to
// the IPv6 interface. That simplifies our code and is most
// general. Unfortunately, we need to run on kernels built without
// IPv6 support too. So probe the kernel to figure it out.
func (p *ipStackCapabilities) probe() {
	s, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	switch err {
	case syscall.EAFNOSUPPORT, syscall.EPROTONOSUPPORT:
	case nil:
		syscall.Close(s)
		p.ipv4Enabled = true
	}
	var probes = []struct {
		laddr net.TCPAddr
		value int
	}{
		// IPv6 communication capability
		{laddr: net.TCPAddr{IP: net.IPv6loopback}, value: 1},
		// IPv4-mapped IPv6 address communication capability
		{laddr: net.TCPAddr{IP: net.IPv4(127, 0, 0, 1)}, value: 0},
	}

	for i := range probes {
		s, err := syscall.Socket(syscall.AF_INET6, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
		if err != nil {
			continue
		}
		defer syscall.Close(s)
		err = syscall.SetsockoptInt(s, syscall.IPPROTO_IPV6, syscall.IPV6_V6ONLY, probes[i].value)
		if err != nil {
			continue
		}
		sa, err := Sockaddr(&(probes[i].laddr), syscall.AF_INET6)
		if err != nil {
			continue
		}
		if err := syscall.Bind(s, sa); err != nil {
			continue
		}
		if i == 0 {
			p.ipv6Enabled = true
		} else {
			p.ipv4MappedIPv6Enabled = true
		}
	}
}

// IPToSockaddr ...
func IPToSockaddr(family int, ip net.IP, port int, _ string) (syscall.Sockaddr, error) {
	switch family {
	case syscall.AF_INET:
		if len(ip) == 0 {
			ip = net.IPv4zero
		}
		ip4 := ip.To4()
		if ip4 == nil {
			return nil, &net.AddrError{Err: "non-IPv4 address", Addr: ip.String()}
		}
		sa := &syscall.SockaddrInet4{Port: port}
		copy(sa.Addr[:], ip4)
		return sa, nil
	case syscall.AF_INET6:
		// In general, an IP wildcard address, which is either
		// "0.0.0.0" or "::", means the entire IP addressing
		// space. For some historical reason, it is used to
		// specify "any available address" on some operations
		// of IP node.
		//
		// When the IP node supports IPv4-mapped IPv6 address,
		// we allow an listener to listen to the wildcard
		// address of both IP addressing spaces by specifying
		// IPv6 wildcard address.
		if len(ip) == 0 || ip.Equal(net.IPv4zero) {
			ip = net.IPv6zero
		}
		// We accept any IPv6 address including IPv4-mapped
		// IPv6 address.
		ip6 := ip.To16()
		if ip6 == nil {
			return nil, &net.AddrError{Err: "non-IPv6 address", Addr: ip.String()}
		}
		//we set ZoneId to 0, as currently we use this functon only to probe the IP capabilities of the host
		//if real Zone handling is required, the zone cache implementation in golang/net should be pulled here
		sa := &syscall.SockaddrInet6{Port: port, ZoneId: 0}
		copy(sa.Addr[:], ip6)
		return sa, nil
	}
	return nil, &net.AddrError{Err: "invalid address family", Addr: ip.String()}
}

// Sockaddr ...
func Sockaddr(a *net.TCPAddr, family int) (syscall.Sockaddr, error) {
	if a == nil {
		return nil, nil
	}
	return IPToSockaddr(family, a.IP, a.Port, a.Zone)
}

// Ntohs ...
var Ntohs = Htons
