// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//go:build !darwin
// +build !darwin

package sctp

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"net"
	"strconv"
	"sync"

	"github.com/onosproject/onos-lib-go/pkg/sctp/events"

	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-lib-go/pkg/sctp/listener"

	"github.com/onosproject/onos-lib-go/pkg/sctp/connection"

	"github.com/onosproject/onos-lib-go/pkg/sctp/addressing"
	"github.com/onosproject/onos-lib-go/pkg/sctp/types"

	"testing"
	"time"

	syscall "golang.org/x/sys/unix"
)

const (
	StreamTestClients  = 10
	StreamTestStreams  = 100
	address            = "127.0.0.1:0"
	ServerRoutineCount = 10
	ClientRoutineCount = 100
	testClients        = 10
)

var defaultOptions types.InitMsg

func init() {
	defaultOptions = types.InitMsg{NumOstreams: StreamTestStreams, MaxInstreams: StreamTestStreams}
}

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
func serveClient(t *testing.T, conn net.Conn, bufsize int) error {
	for {
		buf := make([]byte, bufsize+128)
		n, err := conn.Read(buf)
		if err != nil {
			return err
		}
		_, err = conn.Write(buf[:n])
		if err != nil {
			return err
		}
	}
}

func createOneToOneListener(t *testing.T, defaultInit bool) *listener.Listener {
	addr, err := addressing.ResolveAddress(types.Sctp4, address)
	assert.NoError(t, err)
	if !defaultInit {

		ln, err := listener.NewListener(addr,
			listener.WithMode(types.OneToOne),
			listener.WithNonBlocking(false))
		assert.NoError(t, err)
		return ln
	}
	ln, err := listener.NewListener(addr,
		listener.WithMode(types.OneToOne),
		listener.WithInitMsg(defaultOptions),
		listener.WithNonBlocking(false))
	assert.NoError(t, err)
	return ln

}

func createOneToManyListener(t *testing.T, defaultInit bool) *listener.Listener {
	addr, err := addressing.ResolveAddress(types.Sctp4, address)
	assert.NoError(t, err)
	if !defaultInit {

		ln, err := listener.NewListener(addr,
			listener.WithMode(types.OneToMany),
			listener.WithNonBlocking(false))
		assert.NoError(t, err)
		return ln
	}
	ln, err := listener.NewListener(addr,
		listener.WithMode(types.OneToMany),
		listener.WithInitMsg(defaultOptions),
		listener.WithNonBlocking(false))
	assert.NoError(t, err)
	return ln

}

func connect(t *testing.T, raddr *addressing.Address, defaultInit bool) (*connection.SCTPConn, error) {
	if !defaultInit {
		conn, err := DialSCTP(raddr,
			WithAddressFamily(raddr.AddressFamily),
			WithMode(types.OneToOne),
			WithNonBlocking(false))
		return conn, err
	}
	conn, err := DialSCTP(raddr,
		WithAddressFamily(raddr.AddressFamily),
		WithInitMsg(defaultOptions),
		WithMode(types.OneToOne),
		WithNonBlocking(false))
	return conn, err

}

// TestSCTPConcurrentAccept test multiple clients connecting to one server concurrently
func TestSCTPConcurrentAccept(t *testing.T) {
	ln := createOneToOneListener(t, false)
	raddr, err := ln.SCTPLocalAddr(0)
	assert.NoError(t, err)

	go func() {
		for {
			conn, err := ln.Accept()
			assert.NoError(t, err)
			conn.Close()
		}
	}()

	attempts := 2 * testClients
	fails := 0
	for i := 0; i < attempts; i++ {
		conn, err := connect(t, raddr, false)
		assert.NoError(t, err)

		if err != nil {
			fails++
		} else {
			conn.Close()
		}
	}
	ln.Close()
	assert.Equal(t, 0, fails)
}

// TestSCTPCloseRecv checks the server recevies EOF when the connection is closed by the client
func TestSCTPCloseRecv(t *testing.T) {
	ln := createOneToOneListener(t, false)
	raddr, err := ln.SCTPLocalAddr(0)
	assert.NoError(t, err)

	var wg sync.WaitGroup
	connReady := make(chan struct{}, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		conn, err := ln.Accept()
		assert.NoError(t, err)

		connReady <- struct{}{}
		buf := make([]byte, 256)
		_, err = conn.Read(buf)
		if err != io.EOF && err != syscall.EBADF {
			t.Fail()
		}
	}()
	conn, err := connect(t, raddr, false)
	assert.NoError(t, err)

	<-connReady
	err = conn.Close()
	assert.NoError(t, err)
	wg.Wait()
}

// TestSCTPConcurrentOneToMany tests SCTP one to many mode with multiple clients
func TestSCTPConcurrentOneToMany(t *testing.T) {
	ln := createOneToManyListener(t, false)
	raddr, err := ln.SCTPLocalAddr(0)
	assert.NoError(t, err)

	err = ln.SetEvents(types.WithDataIO(), types.WithAssociation())
	assert.NoError(t, err)

	go func() {
		for {
			buf := make([]byte, 512)
			n, _, flags, err := ln.SCTPRead(buf)
			assert.NoError(t, err)
			if events.IsNotification(flags) {
				notif, err := events.GetNotfication(buf[:n], flags)
				assert.NoError(t, err)
				if events.IsSCTPAssocChange(notif) {
					assocChange := notif.GetAssociationChange()
					if assocChange.State == types.SctpCommUp {
						ln.SCTPWrite([]byte{0}, &types.SndRcvInfo{Flags: types.SctpEOF, AssocID: assocChange.AssocID})
					}
				}
			}
		}
	}()

	attempts := 10 * testClients
	for i := 0; i < attempts; i++ {
		conn, err := connect(t, raddr, false)
		assert.NoError(t, err)
		err = conn.Close()
		assert.NoError(t, err)

	}
	ln.Close()
}

func TestOneToManyPeelOff(t *testing.T) {
	// TODO this test needs to be improved
	var wg sync.WaitGroup
	ln := createOneToManyListener(t, true)
	raddr, ok := ln.LocalAddr().(*addressing.Address)
	assert.Equal(t, true, ok)

	err := ln.SetEvents(types.WithAssociation())
	assert.NoError(t, err)

	go func() {
		test := 999
		count := 0
		for {
			buf := make([]byte, 512)
			n, oob, flags, err := ln.SCTPRead(buf)
			if err == io.EOF {
				break
			}
			assert.NoError(t, err)
			if events.IsNotification(flags) {
				notif, err := events.GetNotfication(buf[:n], flags)
				assert.NoError(t, err)
				if events.IsSCTPAssocChange(notif) {
					assocChange := notif.GetAssociationChange()
					if assocChange.State == types.SctpCommUp {
						newSocket, err := ln.PeelOff(assocChange.AssocID)
						assert.NoError(t, err)
						err = newSocket.SetEvents(types.WithDataIO())
						assert.NoError(t, err)
						count++
						go socketReaderMirror(newSocket, t, test-count)
						continue
					}
				}
			}

			if events.IsMsgEORSet(flags) {
				info := oob.GetSndRcvInfo()
				_, err := ln.SCTPWrite(buf[:n],
					&types.SndRcvInfo{
						AssocID: info.AssocID,
						Stream:  info.Stream,
						PPID:    info.PPID,
					},
				)
				if err != nil {
					return
				}
				continue
			}
		}
	}()

	for i := ClientRoutineCount; i > 0; i-- {
		wg.Add(1)
		go func(client int, raddr *addressing.Address) {
			defer wg.Done()
			t.Logf("[%d]Creating new client connection\n", client)

			conn, err := connect(t, raddr, true)
			assert.NoError(t, err)

			err = conn.SetEvents(types.WithDataIO())
			assert.NoError(t, err)
			for q := range []int{0, 1} {
				rstring := randomString(10)
				rstream := uint16(rand.Intn(StreamTestStreams))
				_, err = conn.SCTPWrite(
					[]byte(rstring),
					&types.SndRcvInfo{
						Stream: rstream,
						PPID:   uint32(q),
					},
				)
				assert.NoError(t, err)

				buf := make([]byte, 512)
				n, oob, _, err := conn.SCTPRead(buf)
				assert.NoError(t, err)
				assert.NotNil(t, oob)
				assert.Equal(t, oob.GetSndRcvInfo().Stream, rstream)
				assert.Equal(t, string(buf[:n]), rstring)
			}
			conn.Close()

		}(i, raddr)
	}
	wg.Wait()
	ln.Close()
}

func socketReaderMirror(sock *connection.SCTPConn, t *testing.T, goroutine int) {
	for {
		buf := make([]byte, 512)
		n, oob, flags, err := sock.SCTPRead(buf)
		if err == io.EOF || err == io.ErrUnexpectedEOF || err == syscall.ENOTCONN {
			sock.Close()
			break
		}
		assert.NoError(t, err)

		if events.IsNotification(flags) {
			if notif, err := connection.SCTPParseNotification(buf[:n]); err == nil {
				t.Logf("[%d]Notification type: %v\n", goroutine, notif.Type().String())
			}
		}
		info := oob.GetSndRcvInfo()
		wn, err := sock.SCTPWrite(buf[:n],
			&types.SndRcvInfo{
				AssocID: info.AssocID,
				Stream:  info.Stream,
				PPID:    info.PPID,
			},
		)
		if err != nil {
			t.Errorf("[%d]failed to write %s, len: %d, err: %v, bytes written: %d, info: %+v", goroutine, string(buf[:n]), len(buf[:n]), err, wn, info)
			return
		}
	}
}

func TestNonBlockingServerOneToMany(t *testing.T) {
	// TODO must be improved
	ln := createOneToManyListener(t, true)
	raddr := ln.LocalAddr().(*addressing.Address)

	err := ln.SetEvents(types.WithDataIO())
	assert.NoError(t, err)

	go func() {
		type ready struct {
			SndRcvInfo *types.SndRcvInfo
			Data       []byte
		}
		b := make(map[int32]map[uint16]bytes.Buffer)
		c := make([]*ready, 0)
		for {
			buf := make([]byte, 64)
			n, oob, flags, err := ln.SCTPRead(buf)
			if err != nil {
				switch err {
				case syscall.EAGAIN:
					goto WRITE
				case syscall.EBADF:
					return
				case syscall.ENOTCONN:
					return
				default:
					t.Fail()
				}
			}

			if events.IsMsgEORSet(flags) {
				info := oob.GetSndRcvInfo()
				assocID := info.AssocID
				if _, ok := b[assocID]; !ok {
					b[assocID] = make(map[uint16]bytes.Buffer)
				}
				bucket := b[assocID]

				stream := bucket[info.Stream]
				stream.Write(buf[:n])

				data := stream.Bytes()
				dataCopy := make([]byte, stream.Len())
				copy(dataCopy, data)

				stream.Reset()

				sndrcv := &types.SndRcvInfo{Stream: info.Stream, AssocID: info.AssocID}
				c = append(c, &ready{SndRcvInfo: sndrcv, Data: dataCopy})

			} else {
				info := oob.GetSndRcvInfo()
				assocID := info.AssocID
				if _, ok := b[assocID]; !ok {
					b[assocID] = make(map[uint16]bytes.Buffer)
				}
				bucket := b[assocID]

				stream := bucket[info.Stream]
				stream.Write(buf[:n])
			}
		WRITE:
			for {
				if len(c) > 0 {
					r := c[0]
					c = c[1:]
					_, err := ln.SCTPWrite(r.Data, r.SndRcvInfo)
					if err != nil {
						if err == syscall.EWOULDBLOCK {
							c = append(c, r)
							break
						}
					}
				} else {
					break
				}
			}

			<-time.Tick(time.Millisecond * 10)
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < StreamTestClients; i++ {
		wg.Add(1)
		go func(test int) {
			defer wg.Done()
			conn, err := connect(t, raddr, true)
			assert.NoError(t, err)

			defer conn.Close()
			err = conn.SetEvents(types.WithDataIO())
			assert.NoError(t, err)
			for ppid := uint16(0); ppid < StreamTestStreams; ppid++ {
				info := &types.SndRcvInfo{
					Stream: ppid,
					PPID:   uint32(ppid),
				}
				text := fmt.Sprintf("[%s,%d,%d]", randomString(10), test, ppid)
				_, err := conn.SCTPWrite([]byte(text), info)
				assert.NoError(t, err)
				var b bytes.Buffer
				for {
					buf := make([]byte, 64)
					cn, oob, flags, err := conn.SCTPRead(buf)
					if err == io.EOF || err == io.ErrUnexpectedEOF {
						if cn == 0 {
							break
						}
					}

					assert.NoError(t, err)
					b.Write(buf[:cn])

					if events.IsMsgEORSet(flags) {
						assert.Equal(t, oob.GetSndRcvInfo().Stream, ppid)
						rtext := b.String()
						b.Reset()
						assert.Equal(t, rtext, text)
						break
					}
				}
			}
		}(i)
	}

	wg.Wait()
	ln.Close()
}

func TestStreamsOneToOneWithoutEvents(t *testing.T) {
	ln := createOneToOneListener(t, true)
	raddr := ln.LocalAddr().(*addressing.Address)

	go func() {
		for {
			c, err := ln.Accept()
			sconn := c.(*connection.SCTPConn)
			assert.NoError(t, err)
			go func() {
				_ = serveClient(t, sconn, 64)
			}()
		}
	}()

	var wg sync.WaitGroup
	i := 0
	for ; i < StreamTestClients; i++ {
		wg.Add(1)
		go func(test int) {
			defer wg.Done()
			conn, err := connect(t, raddr, true)
			assert.NoError(t, err)
			defer conn.Close()
			for ppid := uint16(0); ppid < StreamTestStreams; ppid++ {
				info := &types.SndRcvInfo{
					Stream: ppid,
					PPID:   uint32(ppid),
				}
				text := fmt.Sprintf("[%s,%d,%d]", randomString(10), test, ppid)
				_, err := conn.SCTPWrite([]byte(text), info)
				assert.NoError(t, err)
				var b bytes.Buffer
				for {
					buf := make([]byte, 64)
					cn, _, _, err := conn.SCTPRead(buf)
					if err == io.EOF || err == io.ErrUnexpectedEOF {
						if cn == 0 {
							break
						}
					}

					assert.NoError(t, err)

					b.Write(buf[:cn])
					rtext := b.String()
					assert.Equal(t, rtext, text)
					b.Reset()
					break

				}
			}
		}(i)
	}
	wg.Wait()
}

func TestStreamsOneToOneWithEvents(t *testing.T) {
	ln := createOneToOneListener(t, true)
	raddr := ln.LocalAddr().(*addressing.Address)

	go func() {
		for {
			c, err := ln.Accept()
			assert.NoError(t, err)
			sconn := c.(*connection.SCTPConn)
			err = sconn.SetEvents(types.WithDataIO(), types.WithAssociation())
			assert.NoError(t, err)

			go func() {
				var b bytes.Buffer
				for {
					buf := make([]byte, 64)
					n, oob, flags, err := sconn.SCTPRead(buf)
					if err == io.EOF || err == io.ErrUnexpectedEOF {
						if n == 0 {
							break
						}

					}
					assert.NoError(t, err)
					b.Write(buf[:n])
					if events.IsNotification(flags) {
						if !(events.IsMsgEORSet(flags)) {
							continue
						}
					} else if events.IsMsgEORSet(flags) {
						info := oob.GetSndRcvInfo()
						data := b.Bytes()
						_, err = sconn.SCTPWrite(data, &types.SndRcvInfo{
							Stream: info.Stream,
							PPID:   info.PPID,
						})
						assert.NoError(t, err)
					}
					b.Reset()
				}
			}()
		}
	}()

	i := 0
	var wg sync.WaitGroup
	for ; i < StreamTestClients; i++ {
		wg.Add(1)
		go func(test int) {
			defer wg.Done()
			conn, err := connect(t, raddr, true)
			assert.NoError(t, err)
			defer conn.Close()
			err = conn.SetEvents(types.WithDataIO())
			assert.NoError(t, err)

			for ppid := uint16(0); ppid < StreamTestStreams; ppid++ {
				info := &types.SndRcvInfo{
					Stream: ppid,
					PPID:   uint32(ppid),
				}
				text := fmt.Sprintf("[%s,%d,%d]", randomString(10), test, ppid)
				_, err := conn.SCTPWrite([]byte(text), info)
				assert.NoError(t, err)
				var b bytes.Buffer
				for {
					buf := make([]byte, 64)
					cn, oob, flags, err := conn.SCTPRead(buf)
					if err != nil {
						if err == io.EOF || err == io.ErrUnexpectedEOF {
							if cn == 0 {
								break
							}
						} else {
							t.Errorf("Client connection read err: %v. Total bytes received: %d, bytes received: %d", err, len(b.Bytes()), cn)
							return
						}
					}

					b.Write(buf[:cn])

					if events.IsNotification(flags) {
						if !(events.IsMsgEORSet(flags)) {
							t.Log("buffer not large enough for notification")
							continue
						}
					} else if events.IsMsgEORSet(flags) {
						if oob.GetSndRcvInfo().Stream != ppid {
							t.Errorf("Mismatched PPIDs: %d != %d", oob.GetSndRcvInfo().Stream, ppid)
							return
						}
						rtext := b.String()
						b.Reset()
						assert.Equal(t, rtext, text)
						break
					}
				}
			}
		}(i)
	}

	wg.Wait()
}

func TestStreamsOneToMany(t *testing.T) {
	ln := createOneToManyListener(t, true)
	raddr := ln.LocalAddr().(*addressing.Address)
	err := ln.SetEvents(types.WithDataIO())
	assert.NoError(t, err)

	go func() {
		var b bytes.Buffer
		for {
			buf := make([]byte, 64)
			n, oob, flags, err := ln.SCTPRead(buf)
			assert.NoError(t, err)
			b.Write(buf[:n])

			if events.IsMsgEORSet(flags) {
				info := oob.GetSndRcvInfo()
				data := b.Bytes()
				_, err = ln.SCTPWrite(data, &types.SndRcvInfo{
					Stream:  info.Stream,
					PPID:    info.PPID,
					AssocID: info.AssocID,
				})

				b.Reset()
				assert.NoError(t, err)
			}
		}
	}()

	var wg sync.WaitGroup
	i := 0
	for ; i < StreamTestClients; i++ {
		wg.Add(1)
		go func(testClient int) {
			defer wg.Done()
			conn, err := connect(t, raddr, true)
			assert.NoError(t, err)
			defer conn.Close()
			err = conn.SetEvents(types.WithDataIO())
			assert.NoError(t, err)
			for ppid := uint16(0); ppid < StreamTestStreams; ppid++ {
				info := &types.SndRcvInfo{
					Stream: ppid,
					PPID:   uint32(ppid),
				}
				text := randomString(10) + ":" + strconv.Itoa(testClient)
				_, err := conn.SCTPWrite([]byte(text), info)
				assert.NoError(t, err)
				var b bytes.Buffer
				for {
					buf := make([]byte, 64)
					cn, oob, flags, err := conn.SCTPRead(buf)
					if err == io.EOF || err == io.ErrUnexpectedEOF {
						if cn == 0 {
							break
						}
					}
					assert.NoError(t, err)

					b.Write(buf[:cn])

					if events.IsMsgEORSet(flags) {
						assert.Equal(t, ppid, oob.GetSndRcvInfo().Stream)
						rtext := b.String()
						b.Reset()
						assert.Equal(t, rtext, text)
						break
					}
				}
			}
		}(i)
	}
	wg.Wait()
	ln.Close()
}
