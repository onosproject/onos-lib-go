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

// +build !darwin

package sctp

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"net"
	"sync"

	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-lib-go/pkg/sctp/listener"

	"github.com/onosproject/onos-lib-go/pkg/sctp/connection"

	"github.com/onosproject/onos-lib-go/pkg/sctp/addressing"
	"github.com/onosproject/onos-lib-go/pkg/sctp/defs"

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

var defaultOptions defs.InitMsg

func init() {
	defaultOptions = defs.InitMsg{NumOstreams: StreamTestStreams, MaxInstreams: StreamTestStreams}
}

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

// TestSCTPConcurrentAccept test multiple clients connecting to one server concurrently
func TestSCTPConcurrentAccept(t *testing.T) {
	addr, err := addressing.ResolveAddress(defs.Sctp4, address)
	assert.NoError(t, err)

	ln, err := listener.NewListener(addr, defs.InitMsg{}, defs.OneToOne, false)
	assert.NoError(t, err)

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
		options := NewDialOptions(
			WithAddressFamily(raddr.AddressFamily),
			WithOptions(defs.InitMsg{}),
			WithMode(defs.OneToOne),
			WithNonBlocking(false))

		conn, err := DialSCTP(ln.LocalAddr(), options)
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
	addr, err := addressing.ResolveAddress(defs.Sctp4, address)
	assert.NoError(t, err)

	ln, err := listener.NewListener(addr, defs.InitMsg{}, defs.OneToOne, false)
	assert.NoError(t, err)

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

	options := NewDialOptions(
		WithAddressFamily(raddr.AddressFamily),
		WithOptions(defs.InitMsg{}),
		WithMode(defs.OneToOne),
		WithNonBlocking(false))

	conn, err := DialSCTP(ln.LocalAddr(), options)
	assert.NoError(t, err)

	<-connReady
	err = conn.Close()
	assert.NoError(t, err)
	wg.Wait()
}

// TestSCTPConcurrentOneToMany tests SCTP one to many mode with multiple clients
func TestSCTPConcurrentOneToMany(t *testing.T) {
	addr, err := addressing.ResolveAddress(defs.Sctp4, address)
	assert.NoError(t, err)

	ln, err := listener.NewListener(addr, defs.InitMsg{}, defs.OneToMany, false)
	assert.NoError(t, err)

	raddr, err := ln.SCTPLocalAddr(0)
	assert.NoError(t, err)

	err = ln.SetEvents(defs.SctpEventDataIo | defs.SctpEventAssociation)
	assert.NoError(t, err)

	go func() {
		for {
			buf := make([]byte, 512)
			n, _, flags, err := ln.SCTPRead(buf)
			assert.NoError(t, err)

			if flags&defs.MsgNotification > 0 {
				notif, err := connection.SCTPParseNotification(buf[:n])
				assert.NoError(t, err)
				switch notif.Type() {
				case defs.SctpAssocChange:
					assocChange := notif.GetAssociationChange()
					if assocChange.State == defs.SctpCommUp {
						ln.SCTPWrite([]byte{0}, &defs.SndRcvInfo{Flags: defs.SctpEOF, AssocID: assocChange.AssocID})
					}
				}
			}
		}
	}()

	attempts := 10 * testClients
	for i := 0; i < attempts; i++ {
		options := NewDialOptions(
			WithAddressFamily(raddr.AddressFamily),
			WithOptions(defs.InitMsg{}),
			WithMode(defs.OneToOne),
			WithNonBlocking(false))

		conn, err := DialSCTP(ln.LocalAddr(), options)
		assert.NoError(t, err)
		err = conn.Close()
		assert.NoError(t, err)

	}
	ln.Close()
}

func TestOneToManyPeelOff(t *testing.T) {
	var wg sync.WaitGroup
	addr, _ := addressing.ResolveAddress(defs.Sctp4, address)
	ln, err := listener.NewListener(addr, defs.InitMsg{NumOstreams: StreamTestStreams, MaxInstreams: StreamTestStreams}, defs.OneToMany, false)
	assert.NoError(t, err)

	laddr, _ := ln.LocalAddr().(*addressing.Address)

	err = ln.SetEvents(defs.SctpEventAssociation)
	assert.NoError(t, err)

	go func() {
		test := 999
		count := 0
		for {
			t.Logf("[%d]Reading from server socket...\n", test)
			buf := make([]byte, 512)
			n, oob, flags, err := ln.SCTPRead(buf)
			if err == io.EOF {
				break
			}
			assert.NoError(t, err)

			if flags&defs.MsgNotification > 0 {
				t.Logf("[%d]Got a notification. Bytes read: %v\n", test, n)
				notif, err := connection.SCTPParseNotification(buf[:n])
				assert.NoError(t, err)

				switch notif.Type() {
				case defs.SctpAssocChange:
					t.Logf("[%d]Got an association change notification\n", test)
					assocChange := notif.GetAssociationChange()
					if assocChange.State == defs.SctpCommUp {
						t.Logf("[%d]SCTP_COMM_UP. Creating socket for association: %v\n", test, assocChange.AssocID)
						newSocket, err := ln.PeelOff(assocChange.AssocID)
						assert.NoError(t, err)
						t.Logf("[%d]Peeled off socket: %#+v\n", test, newSocket)
						err = newSocket.SetEvents(defs.SctpEventDataIo)
						assert.NoError(t, err)
						count++
						go socketReaderMirror(newSocket, t, test-count)
						continue
					}
				}
			}

			if flags&defs.MsgEOR > 0 {
				info := oob.GetSndRcvInfo()
				t.Logf("[%d]Got data on main socket, but it wasn't a notification: %#+v \n", test, info)
				wn, werr := ln.SCTPWrite(buf[:n],
					&defs.SndRcvInfo{
						AssocID: info.AssocID,
						Stream:  info.Stream,
						PPID:    info.PPID,
					},
				)
				if werr != nil {
					t.Errorf("[%d]failed to write %s, len: %d, err: %v, bytes written: %d, info: %+v", test, string(buf[:n]), len(buf[:n]), werr, wn, info)
					return
				}
				continue
			}
			t.Logf("[%d]No clue wtf is happening", test)
		}
	}()

	for i := ClientRoutineCount; i > 0; i-- {
		wg.Add(1)
		go func(client int, l *addressing.Address) {
			defer wg.Done()
			t.Logf("[%d]Creating new client connection\n", client)
			options := NewDialOptions(
				WithAddressFamily(addr.AddressFamily),
				WithOptions(defaultOptions),
				WithMode(defs.OneToOne),
				WithNonBlocking(false))

			conn, err := DialSCTP(ln.LocalAddr(), options)
			assert.NoError(t, err)

			err = conn.SetEvents(defs.SctpEventDataIo)
			assert.NoError(t, err)
			for q := range []int{0, 1} {
				rstring := randomString(10)
				rstream := uint16(rand.Intn(StreamTestStreams))
				_, err = conn.SCTPWrite(
					[]byte(rstring),
					&defs.SndRcvInfo{
						Stream: rstream,
						PPID:   uint32(q),
					},
				)
				assert.NoError(t, err)

				t.Logf("[%d]Reading from client socket...\n", client)
				buf := make([]byte, 512)
				n, oob, _, err := conn.SCTPRead(buf)
				assert.NoError(t, err)
				assert.NotNil(t, oob)
				t.Logf("[%d]***Read from client socket\n", client)
				assert.Equal(t, oob.GetSndRcvInfo().Stream, rstream)
				assert.Equal(t, string(buf[:n]), rstring)
				t.Logf("[%d]Client read success! MsgCount: %v\n", client, q)
			}
			conn.Close()

		}(i, laddr)
	}
	wg.Wait()
	ln.Close()
}

func socketReaderMirror(sock *connection.SCTPConn, t *testing.T, goroutine int) {
	for {
		t.Logf("[%d]Reading peel off server socket...\n", goroutine)
		buf := make([]byte, 512)
		n, oob, flags, err := sock.SCTPRead(buf)
		if err == io.EOF || err == io.ErrUnexpectedEOF || err == syscall.ENOTCONN {
			sock.Close()
			break
		}
		assert.NoError(t, err)

		if flags&defs.MsgNotification > 0 {
			t.Logf("[%d]Notification received. Byte count: %v, OOB: %#+v, Flags: %v\n", goroutine, n, oob, flags)
			if notif, err := connection.SCTPParseNotification(buf[:n]); err == nil {
				t.Logf("[%d]Notification type: %v\n", goroutine, notif.Type().String())
			}
		}
		t.Logf("[%d]Writing peel off server socket...\n", goroutine)
		info := oob.GetSndRcvInfo()
		wn, werr := sock.SCTPWrite(buf[:n],
			&defs.SndRcvInfo{
				AssocID: info.AssocID,
				Stream:  info.Stream,
				PPID:    info.PPID,
			},
		)
		if werr != nil {
			t.Errorf("[%d]failed to write %s, len: %d, err: %v, bytes written: %d, info: %+v", goroutine, string(buf[:n]), len(buf[:n]), werr, wn, info)
			return
		}
	}
}

func TestNonBlockingServerOneToMany(t *testing.T) {
	// TODO must be improved
	addr, err := addressing.ResolveAddress(defs.Sctp4, address)
	assert.NoError(t, err)

	ln, err := listener.NewListener(addr, defs.InitMsg{NumOstreams: StreamTestStreams, MaxInstreams: StreamTestStreams}, defs.OneToMany, true)
	assert.NoError(t, err)

	raddr := ln.LocalAddr().(*addressing.Address)

	err = ln.SetEvents(defs.SctpEventDataIo)
	assert.NoError(t, err)

	go func() {
		type ready struct {
			SndRcvInfo *defs.SndRcvInfo
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
					t.Log("EAGAIN")
					goto WRITE
				case syscall.EBADF:
					return
				case syscall.ENOTCONN:
					return
				default:
					t.Fail()
				}
			}

			//t.Logf("DATA: %v, N: %d, OOB: %#+v, FLAGS: %d, ERR: %v\n", buf[:n], n, oob, flags, err)

			if flags&defs.MsgEOR > 0 {
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

				sndrcv := &defs.SndRcvInfo{Stream: info.Stream, AssocID: info.AssocID}
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

				t.Logf("No EOR\n")
			}
		WRITE:
			for {
				if len(c) > 0 {
					r := c[0]
					c = c[1:]
					//t.Logf("Writing: %v, %#+v\n", r.Data, r.SndRcvInfo)
					_, err := ln.SCTPWrite(r.Data, r.SndRcvInfo)
					if err != nil {
						if err == syscall.EWOULDBLOCK {
							t.Logf("WRITE EWOULDBLOCK\n")
							c = append(c, r)
							break
						}
						//t.Logf("Something went wrong?: %v", err)
					}
				} else {
					break
				}
			}

			<-time.Tick(time.Millisecond * 10)
			//t.Logf("tick!\n")
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < StreamTestClients; i++ {
		wg.Add(1)
		go func(test int) {
			defer wg.Done()
			options := NewDialOptions(
				WithAddressFamily(defs.Sctp6),
				WithOptions(defaultOptions),
				WithMode(defs.OneToOne),
				WithNonBlocking(false))

			conn, err := DialSCTP(raddr, options)
			assert.NoError(t, err)

			defer conn.Close()
			err = conn.SetEvents(defs.SctpEventDataIo)
			assert.NoError(t, err)
			for ppid := uint16(0); ppid < StreamTestStreams; ppid++ {
				info := &defs.SndRcvInfo{
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

					if flags&defs.MsgEOR > 0 {
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

func TestStreamsOneToOneWithoutEvents(t *testing.T) {
	addr, _ := addressing.ResolveAddress(defs.Sctp4, address)
	ln, err := listener.NewListener(addr, defs.InitMsg{NumOstreams: StreamTestStreams, MaxInstreams: StreamTestStreams}, defs.OneToOne, false)
	assert.NoError(t, err)
	addr = ln.LocalAddr().(*addressing.Address)

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
			options := NewDialOptions(
				WithAddressFamily(addr.AddressFamily),
				WithOptions(defaultOptions),
				WithMode(defs.OneToOne),
				WithNonBlocking(false))

			conn, err := DialSCTP(ln.LocalAddr(), options)
			assert.NoError(t, err)
			defer conn.Close()
			for ppid := uint16(0); ppid < StreamTestStreams; ppid++ {
				info := &defs.SndRcvInfo{
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
					t.Log(text)
					b.Reset()
					break

				}
			}
		}(i)
	}
	wg.Wait()
}

func TestStreamsOneToOneWithEvents(t *testing.T) {
	addr, err := addressing.ResolveAddress(defs.Sctp4, address)
	assert.NoError(t, err)
	ln, err := listener.NewListener(addr, defs.InitMsg{NumOstreams: StreamTestStreams, MaxInstreams: StreamTestStreams}, defs.OneToOne, false)
	assert.NoError(t, err)
	addr = ln.LocalAddr().(*addressing.Address)

	go func() {
		for {
			c, err := ln.Accept()
			assert.NoError(t, err)
			sconn := c.(*connection.SCTPConn)
			err = sconn.SetEvents(defs.SctpEventDataIo | defs.SctpEventAssociation)
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
					if flags&defs.MsgNotification > 0 {
						if !(flags&defs.MsgEOR > 0) {
							t.Log("buffer not large enough for notification")
							continue
						}
					} else if flags&defs.MsgEOR > 0 {
						info := oob.GetSndRcvInfo()
						data := b.Bytes()
						_, err = sconn.SCTPWrite(data, &defs.SndRcvInfo{
							Stream: info.Stream,
							PPID:   info.PPID,
						})
						if err != nil {
							t.Error(err)
							return
						}
					} else {
						t.Logf("No flags match?: %v", flags&defs.MsgEOR)
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
			options := NewDialOptions(
				WithAddressFamily(addr.AddressFamily),
				WithOptions(defaultOptions),
				WithMode(defs.OneToOne),
				WithNonBlocking(false))

			conn, err := DialSCTP(ln.LocalAddr(), options)
			assert.NoError(t, err)
			defer conn.Close()
			err = conn.SetEvents(defs.SctpEventDataIo)
			assert.NoError(t, err)

			for ppid := uint16(0); ppid < StreamTestStreams; ppid++ {
				info := &defs.SndRcvInfo{
					Stream: ppid,
					PPID:   uint32(ppid),
				}
				text := fmt.Sprintf("[%s,%d,%d]", randomString(10), test, ppid)
				n, err := conn.SCTPWrite([]byte(text), info)
				if err != nil {
					t.Errorf("failed to write %s, len: %d, err: %v, bytes written: %d, info: %+v", text, len(text), err, n, info)
					return
				}
				var b bytes.Buffer
				for {
					buf := make([]byte, 64)
					cn, oob, flags, err := conn.SCTPRead(buf)
					if err != nil {
						if err == io.EOF || err == io.ErrUnexpectedEOF {
							if cn == 0 {
								break
							}
							t.Logf("EOF on server connection. Total bytes received: %d, bytes received: %d", len(b.Bytes()), cn)
						} else {
							t.Errorf("Client connection read err: %v. Total bytes received: %d, bytes received: %d", err, len(b.Bytes()), cn)
							return
						}
					}

					b.Write(buf[:cn])

					if flags&defs.MsgNotification > 0 {
						if !(flags&defs.MsgEOR > 0) {
							t.Log("buffer not large enough for notification")
							continue
						}
					} else if flags&defs.MsgEOR > 0 {
						if oob.GetSndRcvInfo().Stream != ppid {
							t.Errorf("Mismatched PPIDs: %d != %d", oob.GetSndRcvInfo().Stream, ppid)
							return
						}
						rtext := b.String()
						b.Reset()
						assert.Equal(t, rtext, text)
						t.Log(rtext)

						break
					}
				}
			}
		}(i)
	}

	wg.Wait()
}

func TestStreamsOneToMany(t *testing.T) {
	addr, err := addressing.ResolveAddress(defs.Sctp4, address)
	assert.NoError(t, err)

	ln, err := listener.NewListener(addr, defs.InitMsg{NumOstreams: StreamTestStreams, MaxInstreams: StreamTestStreams}, defs.OneToMany, false)
	assert.NoError(t, err)

	addr = ln.LocalAddr().(*addressing.Address)

	err = ln.SetEvents(defs.SctpEventDataIo)
	assert.NoError(t, err)

	go func() {
		var b bytes.Buffer
		for {
			buf := make([]byte, 64)
			n, oob, flags, err := ln.SCTPRead(buf)
			//t.Logf("Server read data count: %d", n)
			assert.NoError(t, err)
			b.Write(buf[:n])

			if flags&defs.MsgEOR > 0 {
				info := oob.GetSndRcvInfo()
				data := b.Bytes()
				t.Logf("Server received data: %s", string(data))
				_, err = ln.SCTPWrite(data, &defs.SndRcvInfo{
					Stream:  info.Stream,
					PPID:    info.PPID,
					AssocID: info.AssocID,
				})

				b.Reset()
				assert.NoError(t, err)
			} else {
				t.Logf("No flags match?: %v", flags&defs.MsgEOR)
			}

		}
	}()

	var wg sync.WaitGroup
	i := 0
	for ; i < StreamTestClients; i++ {
		wg.Add(1)
		go func(test int) {
			defer wg.Done()
			options := NewDialOptions(
				WithAddressFamily(addr.AddressFamily),
				WithOptions(defaultOptions),
				WithMode(defs.OneToOne),
				WithNonBlocking(false))

			conn, err := DialSCTP(ln.LocalAddr(), options)
			assert.NoError(t, err)
			defer conn.Close()
			err = conn.SetEvents(defs.SctpEventDataIo)
			assert.NoError(t, err)
			for ppid := uint16(0); ppid < StreamTestStreams; ppid++ {
				info := &defs.SndRcvInfo{
					Stream: ppid,
					PPID:   uint32(ppid),
				}
				text := randomString(10)
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

					if flags&defs.MsgEOR > 0 {
						if oob.GetSndRcvInfo().Stream != ppid {
							t.Errorf("mismatched PPIDs: %d != %d", oob.GetSndRcvInfo().Stream, ppid)
							return
						}
						rtext := b.String()
						b.Reset()
						assert.Equal(t, rtext, text)
						t.Log("Data read from server matched what we sent")
						break
					}
				}
			}
		}(i)
	}
	wg.Wait()
	ln.Close()
}
