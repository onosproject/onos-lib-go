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

package sctp

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"net"
	"runtime"
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

func TestSCTPConcurrentAccept(t *testing.T) {
	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(4))
	addr, _ := addressing.ResolveAddress(defs.Sctp4, address)
	ln, err := listener.NewListener(addr, defs.InitMsg{}, defs.OneToMany, false)
	if err != nil {
		t.Fatal(err)
	}

	raddr, err := ln.SCTPLocalAddr(0)
	if err != nil {
		t.Fatal(err)
	}

	const N = 10
	var wg sync.WaitGroup
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			for {
				c, err := ln.Accept()
				if err != nil {
					break
				}
				c.Close()
			}
			wg.Done()
		}()
	}
	attempts := 10 * N
	fails := 0
	for i := 0; i < attempts; i++ {
		cfg := connection.NewConfig(
			connection.WithAddressFamily(raddr.AddressFamily),
			connection.WithOptions(defs.InitMsg{}),
			connection.WithMode(defs.OneToOne),
			connection.WithNonBlocking(false))
		c, err := connection.NewSCTPConnection(cfg)
		if err != nil {
			fails++
		} else {
			c.Close()
		}
	}
	ln.Close()
	if fails > 0 {
		t.Fatalf("# of failed Dials: %v", fails)
	}
}

func TestSCTPCloseRecv(t *testing.T) {
	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(4))
	addr, _ := addressing.ResolveAddress(defs.Sctp4, address)
	ln, err := listener.NewListener(addr, defs.InitMsg{}, defs.OneToOne, false)
	if err != nil {
		t.Fatal(err)
	}

	raddr, err := ln.SCTPLocalAddr(0)
	if err != nil {
		t.Fatal(err)
	}

	var wg sync.WaitGroup
	connReady := make(chan struct{}, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		conn, err := ln.Accept()
		if err != nil {
			t.Fatal(err)
		}

		connReady <- struct{}{}
		buf := make([]byte, 256)
		_, err = conn.Read(buf)
		if err != io.EOF && err != syscall.EBADF {
			t.Fatalf("read failed: %v", err)
		}
	}()

	cfg := connection.NewConfig(
		connection.WithAddressFamily(raddr.AddressFamily),
		connection.WithOptions(defs.InitMsg{}),
		connection.WithMode(defs.OneToOne),
		connection.WithNonBlocking(false))
	c, err := connection.NewSCTPConnection(cfg)
	if err != nil {
		t.Fatalf("failed to dial: %s", err)
	}

	if err := c.Connect(raddr); err != nil {
		t.Fatalf("failed to dial: %s", err)
	}

	<-connReady
	err = c.Close()
	if err != nil {
		t.Fatalf("close failed: %v", err)
	}
	wg.Wait()
}

func TestSCTPConcurrentOneToMany(t *testing.T) {
	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(4))
	addr, _ := addressing.ResolveAddress(defs.Sctp4, address)
	ln, err := listener.NewListener(addr, defs.InitMsg{}, defs.OneToMany, false)
	assert.NoError(t, err)

	raddr, err := ln.SCTPLocalAddr(0)
	assert.NoError(t, err)

	ln.SetEvents(defs.SctpEventDataIo | defs.SctpEventAssociation)

	const N = 10
	for i := 0; i < N; i++ {
		go func() {
			for {
				buf := make([]byte, 512)
				n, _, flags, err := ln.SCTPRead(buf)
				assert.NoError(t, err)

				if flags&defs.MsgNotification > 0 {
					notif, _ := connection.SCTPParseNotification(buf[:n])
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
	}
	attempts := 10 * N
	for i := 0; i < attempts; i++ {
		cfg := connection.NewConfig(
			connection.WithAddressFamily(raddr.AddressFamily),
			connection.WithOptions(defs.InitMsg{}),
			connection.WithMode(defs.OneToOne),
			connection.WithNonBlocking(false))
		c, err := connection.NewSCTPConnection(cfg)
		assert.NoError(t, err)
		err = c.Connect(raddr)

	}
	ln.Close()

}

func TestOneToManyPeelOff(t *testing.T) {

	var wg sync.WaitGroup
	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(4))
	addr, _ := addressing.ResolveAddress(defs.Sctp4, address)
	ln, err := listener.NewListener(addr, defs.InitMsg{NumOstreams: StreamTestStreams, MaxInstreams: StreamTestStreams}, defs.OneToMany, false)
	if err != nil {
		t.Fatal(err)
	}

	laddr, _ := ln.LocalAddr().(*addressing.Address)

	ln.SetEvents(defs.SctpEventAssociation)

	go func() {
		test := 999
		count := 0
		for {
			t.Logf("[%d]Reading from server socket...\n", test)
			buf := make([]byte, 512)
			n, oob, flags, err := ln.SCTPRead(buf)
			if err != nil {
				if err == io.EOF {
					break
				}
				t.Fatalf("[%d]Got an error reading from main socket", test)
			}

			if flags&defs.MsgNotification > 0 {
				t.Logf("[%d]Got a notification. Bytes read: %v\n", test, n)
				notif, _ := connection.SCTPParseNotification(buf[:n])
				switch notif.Type() {
				case defs.SctpAssocChange:
					t.Logf("[%d]Got an association change notification\n", test)
					assocChange := notif.GetAssociationChange()
					if assocChange.State == defs.SctpCommUp {
						t.Logf("[%d]SCTP_COMM_UP. Creating socket for association: %v\n", test, assocChange.AssocID)
						newSocket, err := ln.PeelOff(assocChange.AssocID)
						if err != nil {
							t.Fatalf("Failed to peel off socket: %v", err)
						}
						t.Logf("[%d]Peeled off socket: %#+v\n", test, newSocket)
						if err := newSocket.SetEvents(defs.SctpEventDataIo); err != nil {
							t.Logf("[%d]Failed to subscribe to data io for peeled off socket: %v -> %#+v\n", test, err, newSocket)
						}
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
			cfg := connection.NewConfig(
				connection.WithAddressFamily(l.AddressFamily),
				connection.WithOptions(defaultOptions),
				connection.WithMode(defs.OneToOne),
				connection.WithNonBlocking(false))
			c, err := connection.NewSCTPConnection(cfg)
			if err != nil {
				t.Fatalf("[%d]Failed to connect to SCTP server: %v", client, err)
			}
			if err := c.Connect(l); err != nil {
				t.Fatalf("[%d]Failed to connect to SCTP server: %v", client, err)
			}

			c.SetEvents(defs.SctpEventDataIo)
			for q := range []int{0, 1} {
				rstring := randomString(10)
				_, err = c.SCTPWrite(
					[]byte(rstring),
					&defs.SndRcvInfo{
						Stream: uint16(StreamTestStreams),
						PPID:   uint32(q),
					},
				)
				if err != nil {
					t.Fatalf("Failed to send data to SCTP server: %v", err)
				}

				t.Logf("[%d]Reading from client socket...\n", client)
				buf := make([]byte, 512)
				n, oob, _, err := c.SCTPRead(buf)
				if err != nil {
					t.Fatalf("Failed to read from client socket: %v", err)
				}
				if oob == nil {
					t.Fatal("WTF. OOB is nil?!")
				}
				t.Logf("[%d]***Read from client socket\n", client)
				if oob.GetSndRcvInfo().Stream != uint16(StreamTestStreams) {
					t.Fatalf("Data received on a stream(%v) we didn't send(%v) on",
						oob.GetSndRcvInfo().Stream,
						StreamTestStreams)
				}
				if string(buf[:n]) != rstring {
					t.Fatalf("Data from server doesn't match what client sent\nSent: %v\nReceived: %v",
						rstring,
						string(buf[:n]),
					)
				}
				t.Logf("[%d]Client read success! MsgCount: %v\n", client, q)
			}
			c.Close()

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
		if err != nil {
			if err == io.EOF || err == io.ErrUnexpectedEOF || err == syscall.ENOTCONN {
				t.Logf("[%d]Got EOF...\n", goroutine)
				sock.Close()
				break
			}
			t.Fatalf("[%d]Failed to read from socket: %#+v", goroutine, err)
		}

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
	addr, _ := addressing.ResolveAddress(defs.Sctp4, address)
	ln, err := listener.NewListener(addr, defs.InitMsg{NumOstreams: StreamTestStreams, MaxInstreams: StreamTestStreams}, defs.OneToMany, true)
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}
	raddr := ln.LocalAddr().(*addressing.Address)
	t.Logf("Listening on: %v\n", raddr)

	ln.SetEvents(defs.SctpEventDataIo)

	t.Logf("Starting main server loop...\n")
	go func() {
		type ready struct {
			SndRcvInfo *defs.SndRcvInfo
			Data       []byte
		}
		b := make(map[int32]map[uint16]bytes.Buffer)
		c := make([]*ready, 0)
		for {
			buf := make([]byte, 64)
			t.Logf("Server read\n")
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
					t.Fatalf("Server socket error: %v", err)
				}
			}

			t.Logf("DATA: %v, N: %d, OOB: %#+v, FLAGS: %d, ERR: %v\n", buf[:n], n, oob, flags, err)

			if flags&defs.MsgEOR > 0 {
				info := oob.GetSndRcvInfo()
				assocId := info.AssocID
				if _, ok := b[assocId]; !ok {
					b[assocId] = make(map[uint16]bytes.Buffer)
				}
				bucket := b[assocId]

				stream := bucket[info.Stream]
				stream.Write(buf[:n])

				data := stream.Bytes()
				dataCopy := make([]byte, stream.Len())
				copy(dataCopy, data)

				stream.Reset()

				sndrcv := &defs.SndRcvInfo{Stream: info.Stream, AssocID: info.AssocID}
				c = append(c, &ready{SndRcvInfo: sndrcv, Data: dataCopy})
				t.Logf("Write data queued: %#+v\n", c)

			} else {
				info := oob.GetSndRcvInfo()
				assocId := info.AssocID
				if _, ok := b[assocId]; !ok {
					b[assocId] = make(map[uint16]bytes.Buffer)
				}
				bucket := b[assocId]

				stream := bucket[info.Stream]
				stream.Write(buf[:n])

				t.Logf("No EOR\n")
			}
		WRITE:
			for {
				if len(c) > 0 {
					var r *ready
					r = c[0]
					c = c[1:]
					t.Logf("Writing: %v, %#+v\n", r.Data, r.SndRcvInfo)
					_, err := ln.SCTPWrite(r.Data, r.SndRcvInfo)
					if err != nil {
						if err == syscall.EWOULDBLOCK {
							t.Logf("WRITE EWOULDBLOCK\n")
							c = append(c, r)
							break
						}
						t.Logf("Something went wrong?: %v", err)
					}
				} else {
					t.Logf("No queued writes\n")
					break
				}
			}

			<-time.Tick(time.Millisecond * 10)
			t.Logf("tick!\n")
		}
	}()

	t.Logf("Starting client connections...\n")
	var wg sync.WaitGroup
	for i := 0; i < StreamTestClients; i++ {
		wg.Add(1)
		go func(test int) {
			defer wg.Done()
			options := defs.InitMsg{NumOstreams: StreamTestStreams, MaxInstreams: StreamTestStreams}
			cfg := connection.NewConfig(
				connection.WithAddressFamily(defs.Sctp6),
				connection.WithOptions(options),
				connection.WithMode(defs.OneToOne),
				connection.WithNonBlocking(false))
			conn, err := connection.NewSCTPConnection(cfg)
			if err != nil {
				t.Errorf("failed to dial address %s, test #%d: %v", raddr.String(), test, err)
				return
			}
			t.Logf("Connecting to: %v...", raddr)
			if err := conn.Connect(raddr); err != nil {
				t.Fatalf("Failed to connect to server: %v", err)
			}
			t.Logf("Success!\n")
			defer conn.Close()
			conn.SetEvents(defs.SctpEventDataIo)
			for ppid := uint16(0); ppid < StreamTestStreams; ppid++ {
				info := &defs.SndRcvInfo{
					Stream: uint16(ppid),
					PPID:   uint32(ppid),
				}
				text := fmt.Sprintf("[%s,%d,%d]", randomString(10), test, ppid)
				t.Logf("Sending data to server: %v\n", text)
				n, err := conn.SCTPWrite([]byte(text), info)
				if err != nil {
					t.Errorf("failed to write %s, len: %d, err: %v, bytes written: %d, info: %+v", text, len(text), err, n, info)
					return
				}
				var b bytes.Buffer
				for {
					buf := make([]byte, 64)
					cn, oob, flags, err := conn.SCTPRead(buf)
					t.Logf("Client read data count: %d", cn)
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

					if flags&defs.MsgEOR > 0 {
						if oob.GetSndRcvInfo().Stream != ppid {
							t.Errorf("Mismatched PPIDs: %d != %d", oob.GetSndRcvInfo().Stream, ppid)
							return
						}
						rtext := string(b.Bytes())
						b.Reset()
						if rtext != text {
							t.Fatalf("Mismatched payload: %s != %s", []byte(rtext), []byte(text))
						}
						t.Logf("Data read from server matched what we sent")

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
		n, err = conn.Write(buf[:n])
		if err != nil {
			return err
		}
	}
}

func TestStreamsOneToOneNew(t *testing.T) {
	addr, _ := addressing.ResolveAddress(defs.Sctp4, address)
	ln, err := listener.NewListener(addr, defs.InitMsg{NumOstreams: StreamTestStreams, MaxInstreams: StreamTestStreams}, defs.OneToOne, false)
	assert.NoError(t, err)
	addr = ln.LocalAddr().(*addressing.Address)

	go func() {
		for {
			t.Log("Accept")
			c, err := ln.Accept()
			sconn := c.(*connection.SCTPConn)
			assert.NoError(t, err)
			defer sconn.Close()
			go serveClient(t, sconn, 64)
		}
	}()

	wait := make(chan struct{})
	i := 0
	for ; i < StreamTestClients; i++ {
		go func(test int) {
			defer func() { wait <- struct{}{} }()
			cfg := connection.NewConfig(
				connection.WithAddressFamily(addr.AddressFamily),
				connection.WithOptions(defaultOptions),
				connection.WithMode(defs.OneToOne),
				connection.WithNonBlocking(false))
			conn, err := connection.NewSCTPConnection(cfg)
			assert.NoError(t, err)
			conn.Connect(addr)
			defer conn.Close()
			for ppid := uint16(0); ppid < StreamTestStreams; ppid++ {
				info := &defs.SndRcvInfo{
					Stream: uint16(ppid),
					PPID:   uint32(ppid),
				}
				//randomLen := r.Intn(5) + 1
				text := fmt.Sprintf("[%s,%d,%d]", "test", test, ppid)
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
						t.Logf("EOF on server connection. Total bytes received: %d, bytes received: %d", len(b.Bytes()), cn)
					}

					b.Write(buf[:cn])
					rtext := string(b.Bytes())
					t.Log("Recevie:", rtext)
					b.Reset()
					if rtext != text {
						t.Fatalf("Mismatched payload: %s != %s", []byte(rtext), []byte(text))
					}
					break

				}
			}
		}(i)
	}

	for ; i > 0; i-- {
		select {
		case <-wait:
		case <-time.After(time.Second * 30):
			close(wait)
			t.Fatal("timed out")
		}
	}
}

func TestStreamsOneToOne(t *testing.T) {
	addr, _ := addressing.ResolveAddress(defs.Sctp4, address)
	ln, err := listener.NewListener(addr, defs.InitMsg{NumOstreams: StreamTestStreams, MaxInstreams: StreamTestStreams}, defs.OneToOne, false)
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}
	addr = ln.LocalAddr().(*addressing.Address)

	go func() {
		for {
			c, err := ln.Accept()
			sconn := c.(*connection.SCTPConn)
			if err != nil {
				t.Errorf("failed to accept: %v", err)
				return
			}
			defer sconn.Close()

			sconn.SetEvents(defs.SctpEventDataIo | defs.SctpEventAssociation)

			go func() {
				totalrcvd := 0
				var b bytes.Buffer
				for {
					buf := make([]byte, 64)
					n, oob, flags, err := sconn.SCTPRead(buf)
					if err != nil {
						if err == io.EOF || err == io.ErrUnexpectedEOF {
							if n == 0 {
								break
							}
							t.Logf("EOF on server connection. Total bytes received: %d, bytes received: %d", totalrcvd, n)
						} else {
							t.Errorf("Server connection read err: %v. Total bytes received: %d, bytes received: %d", err, totalrcvd, n)
							return
						}
					}

					b.Write(buf[:n])

					if flags&defs.MsgNotification > 0 {
						if !(flags&defs.MsgEOR > 0) {
							t.Log("buffer not large enough for notification")
							continue
						}
					} else if flags&defs.MsgEOR > 0 {
						info := oob.GetSndRcvInfo()
						data := b.Bytes()
						n, err = sconn.SCTPWrite(data, &defs.SndRcvInfo{
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

	wait := make(chan struct{})
	i := 0
	for ; i < StreamTestClients; i++ {
		go func(test int) {
			defer func() { wait <- struct{}{} }()
			cfg := connection.NewConfig(
				connection.WithAddressFamily(addr.AddressFamily),
				connection.WithOptions(defaultOptions),
				connection.WithMode(defs.OneToOne),
				connection.WithNonBlocking(false))
			conn, err := connection.NewSCTPConnection(cfg)
			if err != nil {
				t.Errorf("failed to dial address %s, test #%d: %v", addr.String(), test, err)
				return
			}
			conn.Connect(addr)
			defer conn.Close()
			conn.SetEvents(defs.SctpEventDataIo)
			for ppid := uint16(0); ppid < StreamTestStreams; ppid++ {
				info := &defs.SndRcvInfo{
					Stream: uint16(ppid),
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
						rtext := string(b.Bytes())
						b.Reset()
						if rtext != text {
							t.Fatalf("Mismatched payload: %s != %s", []byte(rtext), []byte(text))
						}

						break
					}
				}
			}
		}(i)
	}
	for ; i > 0; i-- {
		select {
		case <-wait:
		case <-time.After(time.Second * 30):
			close(wait)
			t.Fatal("timed out")
		}
	}
}

func TestStreamsOneToMany(t *testing.T) {
	addr, _ := addressing.ResolveAddress(defs.Sctp4, address)
	ln, err := listener.NewListener(addr, defs.InitMsg{NumOstreams: StreamTestStreams, MaxInstreams: StreamTestStreams}, defs.OneToMany, false)
	assert.NoError(t, err)
	addr = ln.LocalAddr().(*addressing.Address)
	ln.SetEvents(defs.SctpEventDataIo)

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
				n, err = ln.SCTPWrite(data, &defs.SndRcvInfo{
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

	wait := make(chan struct{})
	i := 0
	t.Log("Spinning up clients")
	for ; i < StreamTestClients; i++ {
		go func(test int) {
			defer func() { wait <- struct{}{} }()
			t.Log("Creating client connection")
			cfg := connection.NewConfig(
				connection.WithAddressFamily(addr.AddressFamily),
				connection.WithOptions(defaultOptions),
				connection.WithMode(defs.OneToOne),
				connection.WithNonBlocking(false))
			conn, err := connection.NewSCTPConnection(cfg)
			assert.NoError(t, err)
			conn.Connect(addr)
			defer conn.Close()
			conn.SetEvents(defs.SctpEventDataIo)
			for ppid := uint16(0); ppid < StreamTestStreams; ppid++ {
				info := &defs.SndRcvInfo{
					Stream: uint16(ppid),
					PPID:   uint32(ppid),
				}
				text := randomString(10)
				t.Logf("Sending data to server: %v", text)
				_, err := conn.SCTPWrite([]byte(text), info)
				assert.NoError(t, err)
				var b bytes.Buffer
				for {
					buf := make([]byte, 64)
					cn, oob, flags, err := conn.SCTPRead(buf)
					//t.Logf("Client read data count: %d", cn)
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

					if flags&defs.MsgEOR > 0 {
						if oob.GetSndRcvInfo().Stream != ppid {
							t.Errorf("Mismatched PPIDs: %d != %d", oob.GetSndRcvInfo().Stream, ppid)
							return
						}
						rtext := string(b.Bytes())
						b.Reset()
						if rtext != text {
							t.Fatalf("Mismatched payload: %s != %s", []byte(rtext), []byte(text))
						}
						t.Log("Data read from server matched what we sent")

						break
					}
				}
			}
		}(i)
	}
	for ; i > 0; i-- {
		select {
		case <-wait:
		case <-time.After(time.Second * 10):
			close(wait)
			t.Fatal("timed out")
		}
	}
	ln.Close()
}
