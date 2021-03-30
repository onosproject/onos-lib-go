// SPDX-FileCopyrightText: ${year}-present Open Networking Foundation <info@opennetworking.org>
// SPDX-License-Identifier: Apache-2.0

package connection

import (
	"net"
	"sync/atomic"
	"time"

	"github.com/onosproject/onos-lib-go/pkg/sctp/primitives"

	"github.com/onosproject/onos-lib-go/pkg/sctp/addressing"
	"github.com/onosproject/onos-lib-go/pkg/sctp/defs"

	syscall "golang.org/x/sys/unix"
)

// Connection SCTP connection data structure
type Connection struct {
	fd int32
}

// Config sctp connection config
type Config struct {
	//adressFamily defs.AddressFamily
}

// NewConnection creates a new SCTP connection
func NewConnection(af defs.AddressFamily, options defs.InitMsg, mode defs.SocketMode, nonblocking bool) (*Connection, error) {
	fd, err := primitives.NewSocket(af.ToSyscall(), mode)
	if err != nil {
		return nil, err
	}

	// close socket on error
	defer func(f int) {
		if err != nil {
			syscall.Close(f)
		}
	}(fd)

	if err = primitives.SetDefaultSockOpts(fd, af.ToSyscall(), af == defs.Sctp6Only); err != nil {
		return nil, err
	}

	if err = primitives.SetInitOpts(fd, options); err != nil {
		return nil, err
	}

	if err := syscall.SetNonblock(fd, nonblocking); err != nil {
		return nil, err
	}

	return &Connection{
		fd: int32(fd),
	}, nil
}

// GetFd gets fd
func (c *Connection) GetFd() int32 {
	return c.fd
}

// SetWriteBuffer ...
func (c *Connection) SetWriteBuffer(bytes int) error {
	return syscall.SetsockoptInt(c.FD(), syscall.SOL_SOCKET, syscall.SO_SNDBUF, bytes)
}

// GetSocketMode gets socket mode
func (c *Connection) GetSocketMode() (defs.SocketMode, error) {
	return primitives.GetSocketMode(c.FD())
}

// GetNonBlocking checks non blocking
func (c *Connection) GetNonBlocking() (bool, error) {
	return primitives.GetNonblocking(c.FD())
}

// SetNonBlocking sets non blocking
func (c *Connection) SetNonBlocking(val bool) error {
	return primitives.SetNonblocking(c.FD(), val)
}

// Listen ....
func (c *Connection) Listen() error {
	return primitives.Listen(c.FD())
}

// Bind ...
func (c *Connection) Bind(laddr *addressing.Address) error {
	return primitives.Bind(c.FD(), laddr, defs.SctpBindxAddAddr)
}

// Connect ...
func (c *Connection) Connect(raddr *addressing.Address) error {
	_, err := primitives.Connect(c.FD(), raddr)
	return err
}

// FD ...
func (c *Connection) FD() int {
	return int(atomic.LoadInt32(&c.fd))
}

// SetFd ...
func (c *Connection) SetFd(fd int32) {
	c.fd = fd
}

// Write ...
func (c *Connection) Write(b []byte) (int, error) {
	return c.SCTPWrite(b, nil)
}

// Read ...
func (c *Connection) Read(b []byte) (int, error) {
	n, _, _, err := c.SCTPRead(b)
	if n < 0 {
		n = 0
	}
	return n, err
}

// SetEvents sets events
func (c *Connection) SetEvents(flags int) error {
	return primitives.SetEvents(c.FD(), flags)
}

// GetEvents gets events
func (c *Connection) GetEvents() (int, error) {
	return primitives.GetEvents(c.FD())
}

// SetDefaultSentParam sets default parameters
func (c *Connection) SetDefaultSentParam(info *defs.SndRcvInfo) error {
	return primitives.SetDefaultSentParam(c.FD(), info)
}

// GetDefaultSentParam gets default parameters
func (c *Connection) GetDefaultSentParam() (*defs.SndRcvInfo, error) {
	return primitives.GetDefaultSentParam(c.FD())
}

// SCTPGetPrimaryPeerAddr ...
func (c *Connection) SCTPGetPrimaryPeerAddr() (*addressing.Address, error) {
	return primitives.GetAddrs(c.FD(), 0, defs.SctpPrimaryAddr)
}

// SCTPLocalAddr ...
func (c *Connection) SCTPLocalAddr(id uint16) (*addressing.Address, error) {
	return primitives.GetLocalAddr(c.FD(), id)
}

// LocalAddr ...
func (c *Connection) LocalAddr() net.Addr {
	addr, err := c.SCTPLocalAddr(0)
	if err != nil {
		return nil
	}
	return addr
}

// SCTPRemoteAddr ...
func (c *Connection) SCTPRemoteAddr(id uint16) (*addressing.Address, error) {
	return primitives.GetRemoteAddr(c.FD(), id)
}

// RemoteAddr ...
func (c *Connection) RemoteAddr() net.Addr {
	addr, err := c.SCTPRemoteAddr(0)
	if err != nil {
		return nil
	}
	return addr
}

// PeelOff ...
func (c *Connection) PeelOff(id int32) (*Connection, error) {
	fd, err := primitives.PeelOff(c.FD(), id)
	if err != nil {
		return nil, err
	}

	conn := &Connection{
		fd: int32(fd),
	}

	blocking, err := c.GetNonBlocking()
	if err != nil {
		return nil, err
	}

	if err := conn.SetNonBlocking(blocking); err != nil {
		return nil, err
	}

	return conn, nil

}

// SetDeadline ...
func (c *Connection) SetDeadline(t time.Time) error {
	return syscall.EOPNOTSUPP
}

// SetReadDeadline ...
func (c *Connection) SetReadDeadline(t time.Time) error {
	return syscall.EOPNOTSUPP
}

// SetWriteDeadline ...
func (c *Connection) SetWriteDeadline(t time.Time) error {
	return syscall.EOPNOTSUPP
}

// SCTPWrite ....
func (c *Connection) SCTPWrite(b []byte, info *defs.SndRcvInfo) (int, error) {
	return primitives.Write(c.FD(), b, info)
}

// SCTPRead ...
func (c *Connection) SCTPRead(b []byte) (int, *defs.OOBMessage, int, error) {
	return primitives.Read(c.FD(), b)
}

// Close closes connection
func (c *Connection) Close() error {
	return primitives.Close(c.FD())
}
