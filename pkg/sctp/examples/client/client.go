// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"flag"
	"io"
	"strconv"
	"time"

	"log"

	petname "github.com/dustinkirkland/golang-petname"

	"github.com/onosproject/onos-lib-go/pkg/sctp"
	"github.com/onosproject/onos-lib-go/pkg/sctp/addressing"
	"github.com/onosproject/onos-lib-go/pkg/sctp/types"
)

func newPetName(words int) string {
	return petname.Generate(words, "-")
}

func main() {

	var ip = flag.String("ip", "127.0.0.1", "")
	var port = flag.Int("port", 36421, "")

	flag.Parse()

	address := *ip + ":" + strconv.Itoa(*port)
	raddr, err := addressing.ResolveAddress(types.Sctp4, address)
	if err != nil {
		log.Println(err)
		return
	}
	conn, err := sctp.DialSCTP(raddr,
		sctp.WithAddressFamily(raddr.AddressFamily),
		sctp.WithMode(types.OneToOne),
		sctp.WithNonBlocking(false))

	if err != nil {
		log.Println(err)
		return
	}

	var b bytes.Buffer
	for {
		buf := make([]byte, 64)
		petName := newPetName(2)
		_, err = conn.Write([]byte(petName))
		if err != nil {
			log.Println(err)
		}
		log.Printf("Sending %s to the server:\n", petName)
		cn, err := conn.Read(buf)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			if cn == 0 {
				break
			}
		}

		if err != nil {
			log.Println(err)
		}
		b.Write(buf[:cn])
		log.Printf("Recevied %s from server:\n", b.String())
		b.Reset()
		time.Sleep(time.Second)
	}

}
