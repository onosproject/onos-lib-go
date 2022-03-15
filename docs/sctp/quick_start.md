<!--
SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
SPDX-License-Identifier: Apache-2.0
-->

# Quick Start 

## Installation 
The SCTP library is managed using Go modules. First you need to include the onos-lib-go in your Go application by adding the `github.com/onosproject/onos-lib-go` module to your go.mod: 

```bash
go get github.com/onosproject/onos-lib-go
```

Then you need to import the required packages under [SCTP folder](../../pkg/sctp) in your source code files. For example:

```go
import "github.com/onosproject/onos-lib-go/pkg/sctp"
```

## A Simple SCTP Client and Server Example

A simple client and a server are written as an example to show the usage of the library 
that is available under [examples](../../pkg/sctp/examples) folder. To run the 
example, first you need to build client and server code:

To build the server code:
```bash
cd examples/server
go build server.go
```

To build the client code:
```bash
cd examples/client
go build client.go
```

Then, if you need to open two terminals to run the server and the client
executables:

In one terminal run:
```bash
cd examples/server
./server
```

In another terminal run:
```bash
cd examples/client
./client
```

In the above example, the server reads whatever the client is sent and 
send it back to the client.

The output in the server side will be like this:

```bash
2021/04/13 18:28:33 Recevied touched-walleye from client
2021/04/13 18:28:33 Sending touched-walleye to the client
2021/04/13 18:28:34 Recevied generous-poodle from client
2021/04/13 18:28:34 Sending generous-poodle to the client
2021/04/13 18:28:35 Recevied feasible-caiman from client
2021/04/13 18:28:35 Sending feasible-caiman to the client
2021/04/13 18:28:36 Recevied precise-airedale from client
2021/04/13 18:28:36 Sending precise-airedale to the client
```

The output in the client side will be like this:

```bash
2021/04/13 18:28:33 Sending touched-walleye to the server:
2021/04/13 18:28:33 Recevied touched-walleye from server:
2021/04/13 18:28:34 Sending generous-poodle to the server:
2021/04/13 18:28:34 Recevied generous-poodle from server:
2021/04/13 18:28:35 Sending feasible-caiman to the server:
2021/04/13 18:28:35 Recevied feasible-caiman from server:
2021/04/13 18:28:36 Sending precise-airedale to the server:
2021/04/13 18:28:36 Recevied precise-airedale from server
```


