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

package main

import (
	"flag"
	"log"
	"net"
	"strconv"

	"github.com/onosproject/onos-lib-go/pkg/sctp/listener"

	"github.com/onosproject/onos-lib-go/pkg/sctp/addressing"
	"github.com/onosproject/onos-lib-go/pkg/sctp/types"
)

func serveClient(conn net.Conn, bufsize int) error {
	for {
		buf := make([]byte, bufsize+128)
		n, err := conn.Read(buf)
		log.Printf("Recevied %s from client", string(buf))
		if err != nil {
			return err
		}
		_, err = conn.Write(buf[:n])
		log.Printf("Sending %s to the client", string(buf))
		if err != nil {
			return err
		}
	}
}

func main() {
	var ip = flag.String("ip", "127.0.0.1", "")
	var port = flag.Int("port", 36421, "")

	flag.Parse()

	address := *ip + ":" + strconv.Itoa(*port)
	addr, err := addressing.ResolveAddress(types.Sctp4, address)
	if err != nil {
		log.Println(err)
		return
	}
	ln, err := listener.NewListener(addr,
		listener.WithMode(types.OneToOne),
		listener.WithNonBlocking(false))
	if err != nil {
		log.Println(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			break
		}
		go func() {
			_ = serveClient(conn, 64)
		}()
	}
}
