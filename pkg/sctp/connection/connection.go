// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package connection

import (
	"net"
	"sync/atomic"
	"time"

	syscall "golang.org/x/sys/unix"

	"github.com/onosproject/onos-lib-go/pkg/sctp/addressing"
	"github.com/onosproject/onos-lib-go/pkg/sctp/types"
)

// SCTPConn SCTP connection data structure
type SCTPConn struct {
	fd int32
}

// NewSCTPConnection creates an SCTP connection
func NewSCTPConnection(cfg *Config) (*SCTPConn, error) {
	fd, err := NewSocket(cfg.addressFamily.ToSyscall(), cfg.mode)
	if err != nil {
		return nil, err
	}

	// close socket on error
	defer func(f int) {
		if err != nil {
			syscall.Close(f)
		}
	}(fd)

	if err = setDefaultSockopts(fd, cfg.addressFamily.ToSyscall(), cfg.addressFamily == types.Sctp6Only); err != nil {
		return nil, err
	}

	if err = setInitOpts(fd, cfg.initMsg); err != nil {
		return nil, err
	}

	if err := syscall.SetNonblock(fd, cfg.nonblocking); err != nil {
		return nil, err
	}

	return &SCTPConn{
		fd: int32(fd),
	}, nil
}

// SetFD sets socket file descriptor
func (c *SCTPConn) SetFD(fd int32) {
	c.fd = fd
}

// GetSocketMode gets SCTP socket mode
func (c *SCTPConn) GetSocketMode() (types.SocketMode, error) {
	return getSocketMode(c.FD())
}

// GetNonblocking get nonblocking status
func (c *SCTPConn) GetNonblocking() (bool, error) {
	return getNonblocking(c.FD())
}

// SetNonblocking sets nonblocking connection
func (c *SCTPConn) SetNonblocking(val bool) error {
	return setNonblocking(c.FD(), val)
}

// Listen listens on a SCTP socket
func (c *SCTPConn) Listen() error {
	return listen(c.FD())
}

// Bind binds SCTP socket to an address
func (c *SCTPConn) Bind(laddr *addressing.Address) error {
	return bind(c.FD(), laddr, types.SctpBindxAddAddr)
}

// Connect connects to an SCTP endpoint
func (c *SCTPConn) Connect(raddr *addressing.Address) error {
	_, err := connect(c.FD(), raddr)
	return err
}

// FD returns socket file descriptor
func (c *SCTPConn) FD() int {
	return int(atomic.LoadInt32(&c.fd))
}

// Write writes on an SCTP connection
func (c *SCTPConn) Write(b []byte) (int, error) {
	//E2 interface should use ppid 70(70=0x46).["The Payload Protocol Identifier assigned by IANA to be used by SCTP for the application layer protocol E2AP is 70" defined in chapter 6.1 of <E2 GAP> spec.]
	return c.SCTPWrite(b, &types.SndRcvInfo{PPID: 0x46000000})
}

// Read reads on an SCTP connection
func (c *SCTPConn) Read(b []byte) (int, error) {
	n, _, _, err := c.SCTPRead(b)
	if n < 0 {
		n = 0
	}
	return n, err
}

// SetEvents set SCTP connection events
func (c *SCTPConn) SetEvents(events ...types.Event) error {
	eventSubscribe := &types.EventSubscribe{}
	for _, option := range events {
		option(eventSubscribe)
	}
	return setEvents(c.FD(), *eventSubscribe)
}

// GetEvents gets SCTP connection events
func (c *SCTPConn) GetEvents() (int, error) {
	return getEvents(c.FD())
}

// SetDefaultSentParam sets default sending parameters
func (c *SCTPConn) SetDefaultSentParam(info *types.SndRcvInfo) error {
	return setDefaultSentParam(c.FD(), info)
}

// GetDefaultSentParam gets default sending parameters
func (c *SCTPConn) GetDefaultSentParam() (*types.SndRcvInfo, error) {
	return getDefaultSentParam(c.FD())
}

// SCTPGetPrimaryPeerAddr returns SCTP primary peer address
func (c *SCTPConn) SCTPGetPrimaryPeerAddr() (*addressing.Address, error) {
	return getAddrs(c.FD(), 0, types.SctpPrimaryAddr)
}

// SCTPLocalAddr returns SCTP local address
func (c *SCTPConn) SCTPLocalAddr(id uint16) (*addressing.Address, error) {
	return getLocalAddr(c.FD(), id)
}

// LocalAddr returns local address
func (c *SCTPConn) LocalAddr() net.Addr {
	addr, err := c.SCTPLocalAddr(0)
	if err != nil {
		return nil
	}
	return addr
}

// SCTPRemoteAddr returns SCTP remote address
func (c *SCTPConn) SCTPRemoteAddr(id uint16) (*addressing.Address, error) {
	return getRemoteAddr(c.FD(), id)
}

// RemoteAddr gets remote address
func (c *SCTPConn) RemoteAddr() net.Addr {
	addr, err := c.SCTPRemoteAddr(0)
	if err != nil {
		return nil
	}
	return addr
}

// PeelOff peels off SCTP connection
func (c *SCTPConn) PeelOff(id int32) (*SCTPConn, error) {
	fd, err := peelOff(c.FD(), id)
	if err != nil {
		return nil, err
	}

	conn := &SCTPConn{
		fd: int32(fd),
	}

	blocking, err := c.GetNonblocking()
	if err != nil {
		return nil, err
	}

	if err := conn.SetNonblocking(blocking); err != nil {
		return nil, err
	}

	return conn, nil

}

// SetDeadline sets deadline for SCTP connection
func (c *SCTPConn) SetDeadline(_ time.Time) error {
	return syscall.EOPNOTSUPP
}

// SetReadDeadline sets read deadline
func (c *SCTPConn) SetReadDeadline(_ time.Time) error {
	return syscall.EOPNOTSUPP
}

// SetWriteDeadline sets write deadline
func (c *SCTPConn) SetWriteDeadline(_ time.Time) error {
	return syscall.EOPNOTSUPP
}

// SCTPWrite writes on SCTP connection
func (c *SCTPConn) SCTPWrite(b []byte, info *types.SndRcvInfo) (int, error) {
	return write(c.FD(), b, info)
}

// SCTPRead reads from a SCTP connection
func (c *SCTPConn) SCTPRead(b []byte) (int, *types.OOBMessage, int, error) {
	return read(c.FD(), b)
}

// Close closes an SCTP connection
func (c *SCTPConn) Close() error {
	return close(c.FD())
}
