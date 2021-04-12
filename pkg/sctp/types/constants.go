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

package types

import (
	"fmt"

	syscall "golang.org/x/sys/unix"
)

const (
	// SolSctp ...
	SolSctp = 132
	// SctpBindxAddAddr ...
	SctpBindxAddAddr = 0x01
	// SctpBindxRemAddr ...
	SctpBindxRemAddr = 0x02

	// MsgNotification ...
	MsgNotification = 0x8000
	// MsgEOR ...
	MsgEOR = 0x80
)

const (
	// SctpRtoinfoIota ...
	SctpRtoinfoIota = iota
	// SctpAssocinfo ...
	SctpAssocinfo
	// SctpInitmsg Applications can specify protocol parameters for the default
	//   association initialization. Setting initialization parameters is effective only on an unconnected
	//   socket (for one-to-many style sockets, only future associations are
	//   affected by the change)
	SctpInitmsg
	// SctpNodelay ...
	SctpNodelay
	// SctpAutoclose ...
	SctpAutoclose
	// SctpSetPeerPrimaryAddr ...
	SctpSetPeerPrimaryAddr
	// SctpPrimaryAddr ...
	SctpPrimaryAddr
	// SctpAdaptationLayer ...
	SctpAdaptationLayer
	// SctpDisableFragments ...
	SctpDisableFragments
	// SctpPeerAddrParams ...
	SctpPeerAddrParams
	// SctpDefaultSentParam ...
	SctpDefaultSentParam
	// SctpEvents ...
	SctpEvents

	// SctpIWantMappedV4Addr ...
	SctpIWantMappedV4Addr
	// SctpMaxseg ...
	SctpMaxseg
	// SctpStatus ...
	SctpStatus
	// SctpGetPeerAddrInfo ...
	SctpGetPeerAddrInfo

	// SctpDelayedAckTime ...
	SctpDelayedAckTime
	// SctpDelayedAck ...
	SctpDelayedAck = SctpDelayedAckTime
	// SctpDelayedSack ...
	SctpDelayedSack = SctpDelayedAckTime

	// SctpSockoptBindxAdd ...
	SctpSockoptBindxAdd = 100
	// SctpSockoptBindxRem ...
	SctpSockoptBindxRem = 101
	// SctpSockoptPeeloff ...
	SctpSockoptPeeloff = 102
	// SctpGetPeerAddrs  ...
	SctpGetPeerAddrs = 108
	// SctpGetLocalAddrs ...
	SctpGetLocalAddrs = 109
	// SctpSockoptConnectx ...
	SctpSockoptConnectx = 110
	// SctpSockoptConnectx3 ...
	SctpSockoptConnectx3 = 111
)

// CmsgType ...
type CmsgType int32

// I32 ...
func (t CmsgType) I32() int32 { return int32(t) }

const (
	// SctpCmsgInit ...
	SctpCmsgInit = CmsgType(iota)
	// SctpCmsgSndrcv ...
	SctpCmsgSndrcv
	// SctpCmsgSndinfo ...
	SctpCmsgSndinfo
	// SctpCmsgRcvinfo ...
	SctpCmsgRcvinfo
	// SctpCmsgNxtinfo ...
	SctpCmsgNxtinfo
)

const (
	// SctpUnordered ...
	SctpUnordered = 1 << iota
	// SctpAddrOver ...
	SctpAddrOver
	// SctpAbort ...
	SctpAbort
	// SctpSackImmediately ...
	SctpSackImmediately
	// SctpEOF ...
	SctpEOF
)

const (
	// SctpMaxStream ...
	SctpMaxStream = 0xffff
)

// State ...
type State uint16

const (
	// SctpCommUp ...
	SctpCommUp = State(iota)
	// SctpCommLost ...
	SctpCommLost
	// SctpRestart ...
	SctpRestart
	// SctpShutdownComp ...
	SctpShutdownComp
	// SctpCantStrAssoc ...
	SctpCantStrAssoc
)

func (s State) String() string {
	switch s {
	case SctpCommUp:
		return "SCTP_COMM_UP"
	case SctpCommLost:
		return "SCTP_COMM_LOST"
	case SctpRestart:
		return "SCTP_RESTART"
	case SctpShutdownComp:
		return "SCTP_SHUTDOWN_COMP"
	case SctpCantStrAssoc:
		return "SCTP_CANT_STR_ASSOC"
	default:
		panic(fmt.Sprintf("Unknown SCTPState: %d", s))
	}
}

// AddressFamily ...
type AddressFamily int

const (
	// Sctp4 sctp4
	Sctp4 = AddressFamily(iota)
	// Sctp6 sctp6
	Sctp6
	// Sctp6Only sctp6
	Sctp6Only
)

// ToSyscall ...
func (af AddressFamily) ToSyscall() int {

	switch af {
	case Sctp4:
		return syscall.AF_INET
	case Sctp6:
		return syscall.AF_INET6
	case Sctp6Only:
		return syscall.AF_INET6
	default:
		panic("Invalid SCTPAddressFamily")
	}
}

func (af AddressFamily) String() string {
	switch af {
	case Sctp4:
		return "ip4"
	case Sctp6:
		return "ip6"
	case Sctp6Only:
		return "ip6"
	default:
		panic("Invalid SCTPAddressFamily")
	}
}

// SocketMode ...
type SocketMode int

const (
	// OneToOne one to one mode
	OneToOne = SocketMode(iota)
	// OneToMany one to many mode
	OneToMany
)

// PeerChangeState ...
type PeerChangeState uint32

const (
	// SctpAddrAvailable ...
	SctpAddrAvailable = iota
	// SctpAddrUnreachable ...
	SctpAddrUnreachable
	// SctpAddrRemoved ...
	SctpAddrRemoved
	// SctpAddrAdded ...
	SctpAddrAdded
	// SctpAddrMadePrim ...
	SctpAddrMadePrim
)
