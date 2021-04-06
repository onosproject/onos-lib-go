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

package connection

import (
	"fmt"
	"io"
	"net"
	"os"
	"sync/atomic"
	"unsafe"

	"github.com/onosproject/onos-lib-go/pkg/sctp/defs"
	"github.com/onosproject/onos-lib-go/pkg/sctp/utils"

	"github.com/onosproject/onos-lib-go/pkg/sctp/addressing"

	syscall "golang.org/x/sys/unix"
)

func NewSocket(af int, mode defs.SocketMode) (int, error) {
	socketType := syscall.SOCK_SEQPACKET

	if mode == defs.OneToOne {
		socketType = syscall.SOCK_STREAM
	}

	fd, err := syscall.Socket(
		af,
		socketType,
		syscall.IPPROTO_SCTP,
	)
	if err != nil {
		return -1, err
	}

	return fd, nil
}

func getSocketMode(fd int) (defs.SocketMode, error) {
	optname := syscall.SO_TYPE
	optval := int(0)
	optlen := unsafe.Sizeof(optname)
	r0, _, err := syscall.Syscall6(syscall.SYS_GETSOCKOPT,
		uintptr(fd),
		syscall.SOL_SOCKET,
		uintptr(optname),
		uintptr(unsafe.Pointer(&optval)),
		uintptr(optlen),
		0)

	if err != 0 {
		return -1, err
	}

	switch r0 {
	case syscall.SOCK_STREAM:
		return defs.OneToOne, nil
	case syscall.SOCK_SEQPACKET:
		return defs.OneToMany, nil
	default:
		panic("Not an SCTP socket type!")
	}
}

func setInitOpts(fd int, options defs.InitMsg) error {
	optlen := unsafe.Sizeof(options)
	_, _, err := setsockopt(fd, defs.SctpInitmsg, uintptr(unsafe.Pointer(&options)), uintptr(optlen))
	return err
}

func getInitOpts(fd int) (defs.InitMsg, error) {
	options := defs.InitMsg{}
	optlen := unsafe.Sizeof(options)
	_, _, err := getsockopt(fd, defs.SctpInitmsg, uintptr(unsafe.Pointer(&options)), uintptr(optlen))
	return options, err
}

func connect(fd int, addr *addressing.Address) (int, error) {
	buf := addr.ToRawSockAddrBuf()
	param := defs.GetAddrsOld{
		AddrNum: int32(len(buf)),
		Addrs:   uintptr(uintptr(unsafe.Pointer(&buf[0]))),
	}
	optlen := unsafe.Sizeof(param)
	_, _, err := getsockopt(fd, defs.SctpSockoptConnectx3, uintptr(unsafe.Pointer(&param)), uintptr(unsafe.Pointer(&optlen)))
	if err == nil {
		return int(param.AssocID), nil
	} else if err != syscall.ENOPROTOOPT {
		return 0, err
	}
	r0, _, err := setsockopt(fd, defs.SctpSockoptConnectx, uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	return int(r0), err
}

func bind(fd int, addr *addressing.Address, flags int) error {
	var option uintptr
	switch flags {
	case defs.SctpBindxAddAddr:
		option = defs.SctpSockoptBindxAdd
	case defs.SctpBindxRemAddr:
		option = defs.SctpSockoptBindxRem
	default:
		return syscall.EINVAL
	}

	buf := addr.ToRawSockAddrBuf()
	_, _, err := setsockopt(fd, option, uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	return err
}

func listen(fd int) error {
	return syscall.Listen(fd, syscall.SOMAXCONN)
}

func Accept(fd int) (int, error) {
	fd, _, err := syscall.Accept(fd)
	return fd, err
}

func write(fd int, b []byte, info *defs.SndRcvInfo) (int, error) {
	var cbuf []byte
	if info != nil {
		cmsgBuf := utils.ToBuf(info)
		hdr := &syscall.Cmsghdr{
			Level: syscall.IPPROTO_SCTP,
			Type:  defs.SctpCmsgSndrcv.I32(),
		}

		// bitwidth of hdr.Len is platform-specific,
		// so we use hdr.SetLen() rather than directly setting hdr.Len
		hdr.SetLen(syscall.CmsgSpace(len(cmsgBuf)))
		cbuf = append(utils.ToBuf(hdr), cmsgBuf...)
	}
	return syscall.SendmsgN(fd, b, cbuf, nil, 0)
}

func read(fd int, b []byte) (dataCount int, oob *defs.OOBMessage, flags int, err error) {

	oobBuffer := make([]byte, 254)
	oobCount := 0

	dataCount, oobCount, flags, _, err = syscall.Recvmsg(fd, b, oobBuffer, 0)

	if err != nil {
		return
	}

	if dataCount == 0 && oobCount == 0 {
		err = io.EOF
		return
	}

	if oobCount > 0 {
		oob, err = SCTPParseOOB(oobBuffer[:oobCount])
	}

	return
}

func close(fd int) error {
	if fd > 0 {
		fdq := int32(fd)
		fd = int(atomic.SwapInt32(&fdq, -1))
		if fd > 0 {
			info := &defs.SndRcvInfo{
				Flags: defs.SctpEOF,
			}
			write(fd, nil, info)
			err := syscall.Shutdown(fd, syscall.SHUT_RDWR)
			if err != nil {
				return err
			}
			return syscall.Close(fd)
		}
	}
	return syscall.EBADF
}

func setNonblocking(fd int, nonblocking bool) error {
	return syscall.SetNonblock(fd, nonblocking)
}

func getNonblocking(fd int) (bool, error) {
	flags, err := syscall.FcntlInt(uintptr(fd), syscall.F_GETFL, 0)
	if err != nil {
		return false, err
	}
	return flags&syscall.O_NONBLOCK > 0, nil
}

func getLocalAddr(fd int, stream uint16) (*addressing.Address, error) {
	return getAddrs(fd, stream, defs.SctpGetLocalAddrs)
}

func getRemoteAddr(fd int, stream uint16) (*addressing.Address, error) {
	return getAddrs(fd, stream, defs.SctpGetPeerAddrs)
}

func getAddrs(fd int, id uint16, optname int) (*addressing.Address, error) {

	type getaddrs struct {
		assocId int32
		addrNum uint32
		addrs   [4096]byte
	}
	param := getaddrs{
		assocId: int32(id),
	}
	optlen := unsafe.Sizeof(param)
	_, _, err := getsockopt(fd, uintptr(optname), uintptr(unsafe.Pointer(&param)), uintptr(unsafe.Pointer(&optlen)))
	if err != nil {
		return nil, err
	}

	ptr, n := unsafe.Pointer(&param.addrs), int(param.addrNum)

	addr := &addressing.Address{
		IPAddrs: make([]net.IPAddr, n),
	}

	switch family := (*(*syscall.RawSockaddrAny)(ptr)).Addr.Family; family {
	case syscall.AF_INET:
		addr.Port = int(utils.Ntohs(uint16((*(*syscall.RawSockaddrInet4)(ptr)).Port)))
		tmp := syscall.RawSockaddrInet4{}
		size := unsafe.Sizeof(tmp)
		for i := 0; i < n; i++ {
			a := *(*syscall.RawSockaddrInet4)(unsafe.Pointer(
				uintptr(ptr) + size*uintptr(i)))
			addr.IPAddrs[i] = net.IPAddr{IP: a.Addr[:]}
		}
	case syscall.AF_INET6:
		addr.Port = int(utils.Ntohs(uint16((*(*syscall.RawSockaddrInet4)(ptr)).Port)))
		tmp := syscall.RawSockaddrInet6{}
		size := unsafe.Sizeof(tmp)
		for i := 0; i < n; i++ {
			a := *(*syscall.RawSockaddrInet6)(unsafe.Pointer(
				uintptr(ptr) + size*uintptr(i)))
			var zone string
			ifi, err := net.InterfaceByIndex(int(a.Scope_id))
			if err == nil {
				zone = ifi.Name
			}
			addr.IPAddrs[i] = net.IPAddr{IP: a.Addr[:], Zone: zone}
		}
	default:
		return nil, fmt.Errorf("unknown address family: %d", family)
	}
	return addr, nil
}

func getDefaultSentParam(fd int) (*defs.SndRcvInfo, error) {
	info := &defs.SndRcvInfo{}
	optlen := unsafe.Sizeof(*info)
	_, _, err := getsockopt(fd, defs.SctpDefaultSentParam, uintptr(unsafe.Pointer(info)), uintptr(unsafe.Pointer(&optlen)))
	return info, err
}

func setDefaultSentParam(fd int, info *defs.SndRcvInfo) error {
	optlen := unsafe.Sizeof(*info)
	_, _, err := setsockopt(fd, defs.SctpDefaultSentParam, uintptr(unsafe.Pointer(info)), uintptr(optlen))
	return err
}

func peelOff(fd int, associd int32) (int, error) {
	type peeloffArg struct {
		assocId int32
		sd      int
	}
	param := peeloffArg{
		assocId: associd,
	}
	optlen := unsafe.Sizeof(param)
	r0, _, err := getsockopt(fd, defs.SctpSockoptPeeloff, uintptr(unsafe.Pointer(&param)), uintptr(unsafe.Pointer(&optlen)))
	if err != nil {
		return -1, err
	}
	// Note, for some reason, the struct isn't getting populated after the syscall. But the return values are right, so we use r0 which is our fd that we want.
	if param.sd == -1 || r0 == 0 {
		return -1, fmt.Errorf("Returned fd is negative!")
	}
	return int(r0), nil

}

func setEvents(fd, flags int) error {

	var d, a, ad, sf, p, sh, pa, ada, au, se uint8
	if flags&defs.SctpEventDataIo > 0 {
		d = 1
	}
	if flags&defs.SctpEventAssociation > 0 {
		a = 1
	}
	if flags&defs.SctpEventAddress > 0 {
		ad = 1
	}
	if flags&defs.SctpEventSendFailure > 0 {
		sf = 1
	}
	if flags&defs.SctpEventPeerError > 0 {
		p = 1
	}
	if flags&defs.SctpEventShutdown > 0 {
		sh = 1
	}
	if flags&defs.SctpEventPartialDelivery > 0 {
		pa = 1
	}
	if flags&defs.SctpEventAdaptationLayer > 0 {
		ada = 1
	}
	if flags&defs.SctpEventAuthentication > 0 {
		au = 1
	}
	if flags&defs.SctpEventSenderDry > 0 {
		se = 1
	}
	param := defs.EventSubscribe{
		DataIO:          d,
		Association:     a,
		Address:         ad,
		SendFailure:     sf,
		PeerError:       p,
		Shutdown:        sh,
		PartialDelivery: pa,
		AdaptationLayer: ada,
		Authentication:  au,
		SenderDry:       se,
	}
	optlen := unsafe.Sizeof(param)
	_, _, err := setsockopt(fd, defs.SctpEvents, uintptr(unsafe.Pointer(&param)), uintptr(optlen))
	return err
}

func getEvents(fd int) (int, error) {
	param := defs.EventSubscribe{}
	optlen := unsafe.Sizeof(param)
	_, _, err := getsockopt(fd, defs.SctpEvents, uintptr(unsafe.Pointer(&param)), uintptr(unsafe.Pointer(&optlen)))
	if err != nil {
		return 0, err
	}
	var flags int
	if param.DataIO > 0 {
		flags |= defs.SctpEventDataIo
	}
	if param.Association > 0 {
		flags |= defs.SctpEventAssociation
	}
	if param.Address > 0 {
		flags |= defs.SctpEventAddress
	}
	if param.SendFailure > 0 {
		flags |= defs.SctpEventSendFailure
	}
	if param.PeerError > 0 {
		flags |= defs.SctpEventPeerError
	}
	if param.Shutdown > 0 {
		flags |= defs.SctpEventShutdown
	}
	if param.PartialDelivery > 0 {
		flags |= defs.SctpEventPartialDelivery
	}
	if param.AdaptationLayer > 0 {
		flags |= defs.SctpEventAdaptationLayer
	}
	if param.Authentication > 0 {
		flags |= defs.SctpEventAuthentication
	}
	if param.SenderDry > 0 {
		flags |= defs.SctpEventSenderDry
	}
	return flags, nil
}

func setsockopt(fd int, optname, optval, optlen uintptr) (uintptr, uintptr, error) {
	// FIXME: syscall.SYS_SETSOCKOPT is undefined on 386
	r0, r1, errno := syscall.Syscall6(syscall.SYS_SETSOCKOPT,
		uintptr(fd),
		defs.SolSctp,
		optname,
		optval,
		optlen,
		0)
	if errno != 0 {
		return r0, r1, errno
	}
	return r0, r1, nil
}

//from https://github.com/golang/go
//Changes: it is for SCTP only
func setDefaultSockopts(s int, family int, ipv6only bool) error {
	if family == syscall.AF_INET6 {
		// Allow both IP versions even if the OS default
		// is otherwise. Note that some operating systems
		// never admit this option.
		err := syscall.SetsockoptInt(s, syscall.IPPROTO_IPV6, syscall.IPV6_V6ONLY, utils.BoolToInt(ipv6only))
		if err != nil {
			return err
		}
	}
	// Allow broadcast.
	return os.NewSyscallError("setsockopt", syscall.SetsockoptInt(s, syscall.SOL_SOCKET, syscall.SO_BROADCAST, 1))
}

func getsockopt(fd int, optname, optval, optlen uintptr) (uintptr, uintptr, error) {
	// FIXME: syscall.SYS_GETSOCKOPT is undefined on 386
	r0, r1, errno := syscall.Syscall6(syscall.SYS_GETSOCKOPT,
		uintptr(fd),
		defs.SolSctp,
		optname,
		optval,
		optlen,
		0)
	if errno != 0 {
		return r0, r1, errno
	}
	return r0, r1, nil
}

func SCTPParseOOB(b []byte) (*defs.OOBMessage, error) {
	msgs, err := syscall.ParseSocketControlMessage(b)
	if err != nil {
		return nil, err
	}
	for _, msg := range msgs {
		m := &defs.OOBMessage{msg}
		if m.IsSCTP() {
			return m, nil
		}
	}
	return nil, nil
}

func SCTPParseNotification(b []byte) (*defs.Notification, error) {
	return &defs.Notification{Data: b}, nil
}
