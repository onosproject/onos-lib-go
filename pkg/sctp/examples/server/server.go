// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

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
