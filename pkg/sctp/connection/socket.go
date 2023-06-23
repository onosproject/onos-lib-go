// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package connection

import (
	"fmt"
	"io"
	"net"
	"os"
	"sync/atomic"
	"unsafe"

	"github.com/onosproject/onos-lib-go/pkg/errors"

	"github.com/onosproject/onos-lib-go/pkg/sctp/types"
	"github.com/onosproject/onos-lib-go/pkg/sctp/utils"

	"github.com/onosproject/onos-lib-go/pkg/sctp/addressing"

	syscall "golang.org/x/sys/unix"
)

// NewSocket creates a new SCTP socket based on a given mode
func NewSocket(af int, mode types.SocketMode) (int, error) {
	socketType := syscall.SOCK_SEQPACKET

	if mode == types.OneToOne {
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

func getSocketMode(fd int) (types.SocketMode, error) {
	optname := syscall.SO_TYPE
	optval := int(0)
	optlen := unsafe.Sizeof(optname)
	r0, _, err := syscall.Syscall6(syscall.SYS_GETSOCKOPT,
		uintptr(fd),
		syscall.SOL_SOCKET,
		uintptr(optname),
		uintptr(unsafe.Pointer(&optval)),
		optlen,
		0)

	if err != 0 {
		return -1, err
	}

	switch r0 {
	case syscall.SOCK_STREAM:
		return types.OneToOne, nil
	case syscall.SOCK_SEQPACKET:
		return types.OneToMany, nil
	default:
		panic("Not an SCTP socket type!")
	}
}

func setInitOpts(fd int, options types.InitMsg) error {
	optlen := unsafe.Sizeof(options)
	_, _, err := setsockopt(fd, types.SctpInitmsg, uintptr(unsafe.Pointer(&options)), optlen)
	return err
}

/*func getInitOpts(fd int) (defs.InitMsg, error) {
	options := defs.InitMsg{}
	optlen := unsafe.Sizeof(options)
	_, _, err := getsockopt(fd, defs.SctpInitmsg, uintptr(unsafe.Pointer(&options)), optlen)
	return options, err
}*/

func connect(fd int, addr *addressing.Address) (int, error) {
	buf := addr.ToRawSockAddrBuf()
	param := types.GetAddrsOld{
		AddrNum: int32(len(buf)),
		Addrs:   uintptr(unsafe.Pointer(&buf[0])),
	}
	optlen := unsafe.Sizeof(param)
	_, _, err := getsockopt(fd, types.SctpSockoptConnectx3, uintptr(unsafe.Pointer(&param)), uintptr(unsafe.Pointer(&optlen)))
	if err == nil {
		return int(param.AssocID), nil
	} else if err != syscall.ENOPROTOOPT {
		return 0, err
	}
	r0, _, err := setsockopt(fd, types.SctpSockoptConnectx, uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	return int(r0), err
}

func bind(fd int, addr *addressing.Address, flags int) error {
	var option uintptr
	switch flags {
	case types.SctpBindxAddAddr:
		option = types.SctpSockoptBindxAdd
	case types.SctpBindxRemAddr:
		option = types.SctpSockoptBindxRem
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

// Accept accepts an SCTP connection request
func Accept(fd int) (int, error) {
	fd, _, err := syscall.Accept(fd)
	return fd, err
}

func write(fd int, b []byte, info *types.SndRcvInfo) (int, error) {
	var cbuf []byte
	if info != nil {
		cmsgBuf := utils.ToBuf(info)
		hdr := &syscall.Cmsghdr{
			Level: syscall.IPPROTO_SCTP,
			Type:  types.SctpCmsgSndrcv.I32(),
		}

		// bitwidth of hdr.Len is platform-specific,
		// so we use hdr.SetLen() rather than directly setting hdr.Len
		hdr.SetLen(syscall.CmsgSpace(len(cmsgBuf)))
		cbuf = append(utils.ToBuf(hdr), cmsgBuf...)
	}
	return syscall.SendmsgN(fd, b, cbuf, nil, 0)
}

func read(fd int, b []byte) (dataCount int, oob *types.OOBMessage, flags int, err error) {

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

// nolint: all
func close(fd int) error {
	if fd > 0 {
		fdq := int32(fd)
		fd = int(atomic.SwapInt32(&fdq, -1))
		if fd > 0 {
			info := &types.SndRcvInfo{
				Flags: types.SctpEOF,
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
	return getAddrs(fd, stream, types.SctpGetLocalAddrs)
}

func getRemoteAddr(fd int, stream uint16) (*addressing.Address, error) {
	return getAddrs(fd, stream, types.SctpGetPeerAddrs)
}

func getAddrs(fd int, id uint16, optname int) (*addressing.Address, error) {

	param := GetAddrs{
		assocID: int32(id),
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
		addr.Port = int(utils.Ntohs((*(*syscall.RawSockaddrInet4)(ptr)).Port))
		size := unsafe.Sizeof(syscall.RawSockaddrInet4{})
		for i := 0; i < n; i++ {
			a := *(*syscall.RawSockaddrInet4)(unsafe.Pointer(
				uintptr(ptr) + size*uintptr(i)))
			addr.IPAddrs[i] = net.IPAddr{IP: a.Addr[:]}
		}
	case syscall.AF_INET6:
		addr.Port = int(utils.Ntohs((*(*syscall.RawSockaddrInet4)(ptr)).Port))
		size := unsafe.Sizeof(syscall.RawSockaddrInet6{})
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

func getDefaultSentParam(fd int) (*types.SndRcvInfo, error) {
	info := &types.SndRcvInfo{}
	optlen := unsafe.Sizeof(*info)
	_, _, err := getsockopt(fd, types.SctpDefaultSentParam, uintptr(unsafe.Pointer(info)), uintptr(unsafe.Pointer(&optlen)))
	return info, err
}

func setDefaultSentParam(fd int, info *types.SndRcvInfo) error {
	optlen := unsafe.Sizeof(*info)
	_, _, err := setsockopt(fd, types.SctpDefaultSentParam, uintptr(unsafe.Pointer(info)), optlen)
	return err
}

func peelOff(fd int, assocID int32) (int, error) {

	param := PeeloffArg{
		assocID: assocID,
	}
	optlen := unsafe.Sizeof(param)
	r0, _, err := getsockopt(fd, types.SctpSockoptPeeloff, uintptr(unsafe.Pointer(&param)), uintptr(unsafe.Pointer(&optlen)))
	if err != nil {
		return -1, err
	}
	// Note, for some reason, the struct isn't getting populated after the syscall. But the return values are right, so we use r0 which is our fd that we want.
	if param.sd == -1 || r0 == 0 {
		return -1, errors.New(errors.Unknown, "returned fd is negative")
	}
	return int(r0), nil

}

func setEvents(fd int, param types.EventSubscribe) error {

	/*var d, a, ad, sf, p, sh, pa, ada, au, se uint8
	if flags&types.SctpEventDataIo > 0 {
		d = 1
	}
	if flags&types.SctpEventAssociation > 0 {
		a = 1
	}
	if flags&types.SctpEventAddress > 0 {
		ad = 1
	}
	if flags&types.SctpEventSendFailure > 0 {
		sf = 1
	}
	if flags&types.SctpEventPeerError > 0 {
		p = 1
	}
	if flags&types.SctpEventShutdown > 0 {
		sh = 1
	}
	if flags&types.SctpEventPartialDelivery > 0 {
		pa = 1
	}
	if flags&types.SctpEventAdaptationLayer > 0 {
		ada = 1
	}
	if flags&types.SctpEventAuthentication > 0 {
		au = 1
	}
	if flags&types.SctpEventSenderDry > 0 {
		se = 1
	}*/
	/*param := types.EventSubscribe{
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
	}*/
	optlen := unsafe.Sizeof(param)
	_, _, err := setsockopt(fd, types.SctpEvents, uintptr(unsafe.Pointer(&param)), optlen)
	return err
}

func getEvents(fd int) (int, error) {
	param := types.EventSubscribe{}
	optlen := unsafe.Sizeof(param)
	_, _, err := getsockopt(fd, types.SctpEvents, uintptr(unsafe.Pointer(&param)), uintptr(unsafe.Pointer(&optlen)))
	if err != nil {
		return 0, err
	}
	var flags int
	if param.DataIO > 0 {
		flags |= types.SctpEventDataIo
	}
	if param.Association > 0 {
		flags |= types.SctpEventAssociation
	}
	if param.Address > 0 {
		flags |= types.SctpEventAddress
	}
	if param.SendFailure > 0 {
		flags |= types.SctpEventSendFailure
	}
	if param.PeerError > 0 {
		flags |= types.SctpEventPeerError
	}
	if param.Shutdown > 0 {
		flags |= types.SctpEventShutdown
	}
	if param.PartialDelivery > 0 {
		flags |= types.SctpEventPartialDelivery
	}
	if param.AdaptationLayer > 0 {
		flags |= types.SctpEventAdaptationLayer
	}
	if param.Authentication > 0 {
		flags |= types.SctpEventAuthentication
	}
	if param.SenderDry > 0 {
		flags |= types.SctpEventSenderDry
	}
	return flags, nil
}

func setsockopt(fd int, optname, optval, optlen uintptr) (uintptr, uintptr, error) {
	// FIXME: syscall.SYS_SETSOCKOPT is undefined on 386
	r0, r1, errno := syscall.Syscall6(syscall.SYS_SETSOCKOPT,
		uintptr(fd),
		types.SolSctp,
		optname,
		optval,
		optlen,
		0)
	if errno != 0 {
		return r0, r1, errno
	}
	return r0, r1, nil
}

// from https://github.com/golang/go
// Changes: it is for SCTP only
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
		types.SolSctp,
		optname,
		optval,
		optlen,
		0)
	if errno != 0 {
		return r0, r1, errno
	}
	return r0, r1, nil
}

// SCTPParseOOB ...
func SCTPParseOOB(b []byte) (*types.OOBMessage, error) {
	msgs, err := syscall.ParseSocketControlMessage(b)
	if err != nil {
		return nil, err
	}
	for _, msg := range msgs {
		m := &types.OOBMessage{SocketControlMessage: msg}
		if m.IsSCTP() {
			return m, nil
		}
	}
	return nil, nil
}

// SCTPParseNotification ...
func SCTPParseNotification(b []byte) (*types.Notification, error) {
	return &types.Notification{Data: b}, nil
}
