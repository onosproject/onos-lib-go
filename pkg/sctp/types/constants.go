// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package types

import (
	syscall "golang.org/x/sys/unix"
)

const (
	// SolSctp ...
	SolSctp = 132
	// SctpBindxAddAddr  SCTP_BINDX_ADD_ADDR directs SCTP to add the given addresses to the
	//   socket (i.e., endpoint)
	SctpBindxAddAddr = 0x01
	// SctpBindxRemAddr SCTP_BINDX_REM_ADDR directs SCTP to
	//   remove the given addresses from the socket.
	SctpBindxRemAddr = 0x02

	// MsgNotification  If a notification has arrived, recvmsg() will return the notification
	//   in the msg_iov field and set the MSG_NOTIFICATION flag in msg_flags.
	//   If the MSG_NOTIFICATION flag is not set, recvmsg() will return data.
	MsgNotification = 0x8000

	// MsgEOR  If all portions of a data frame or notification have been read,
	//   recvmsg() will return with MSG_EOR set in msg_flags. If the application
	//   does not provide enough buffer space to completely
	//   receive a data message, MSG_EOR will not be set in msg_flags.
	//   Successive reads will consume more of the same message until the
	//   entire message has been delivered, and MSG_EOR will be set.
	MsgEOR = 0x80
)

const (
	// SctpRtoinfoIota ...
	SctpRtoinfoIota = iota
	// SctpAssocinfo ...
	SctpAssocinfo
	// SctpInitmsg applications can specify protocol parameters for the default
	//   association initialization. Setting initialization parameters is effective only on an unconnected
	//   socket (for one-to-many style sockets, only future associations are
	//   affected by the change)
	SctpInitmsg

	// SctpNodelay this option turns on/off any Nagle-like algorithm.  This means that
	//   packets are generally sent as soon as possible, and no unnecessary
	//   delays are introduced, at the cost of more packets in the network.
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
	// SctpUnordered this flag is present when the message was sent unordered.
	SctpUnordered = 1 << iota

	// SctpAddrOver  for a one-to-many style socket requests that the SCTP stack override
	//   the primary destination address with the address found with the sendto sendmsg call.
	SctpAddrOver

	// SctpAbort Setting this flag causes the specified association to abort by sending an ABORT message to the peer.
	SctpAbort

	// SctpSackImmediately ...
	SctpSackImmediately

	// SctpEOF Setting this flag invokes the SCTP graceful shutdown
	//  procedures on the specified association.
	SctpEOF
)

const (
	// SctpMaxStream ...
	SctpMaxStream = 0xffff
)

// State SCTP association event value.
type State uint16

const (
	// SctpCommUp this state means a new association is now ready, and data may be
	// exchanged with this peer. When an association has been
	// established successfully, this notification should be the
	// first one.
	SctpCommUp = State(iota)

	// SctpCommLost this state means the association has failed. The association is
	// now in the closed state. If SEND_FAILED notifications are turned on, an
	// SCTP_COMM_LOST is accompanied by a series of SCTP_SEND_FAILED_EVENT events,
	// one for each outstanding message.
	SctpCommLost

	// SctpRestart SCTP has detected that the peer has restarted.
	SctpRestart

	// SctpShutdownComp the association has gracefully closed.
	SctpShutdownComp

	// SctpCantStrAssoc the association setup failed.  If non-blocking mode is set and data was sent (on a one-to-many
	//  style socket), an SCTP_CANT_STR_ASSOC is accompanied by a
	//  series of SCTP_SEND_FAILED_EVENT events, one for each
	//  outstanding message.
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
		return "UNKNOWN_STATE"
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

// ToSyscall converts to syscall address family constant values
func (af AddressFamily) ToSyscall() int {

	switch af {
	case Sctp4:
		return syscall.AF_INET
	case Sctp6:
		return syscall.AF_INET6
	case Sctp6Only:
		return syscall.AF_INET6
	default:
		return -1
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
	// OneToOne one-to-one mode. the goal of this style is to follow as closely as possible the
	//   current practice of using the sockets interface for a connection-
	//   oriented protocol such as TCP.  This style enables existing
	//   applications using connection-oriented protocols to be ported to SCTP
	//   with very little effort.
	OneToOne = SocketMode(iota)
	// OneToMany one to many mode. This set of semantics is
	// similar to that defined for connectionless protocols, such as
	//  UDP.  A one-to-many style SCTP socket should be able to control
	//  multiple SCTP associations.  This is similar to a UDP socket,
	//  which can communicate with many peer endpoints.
	OneToMany
)

// PeerChangeState ...
type PeerChangeState uint32

const (
	// SctpAddrAvailable   address reachable notification.
	SctpAddrAvailable = iota

	// SctpAddrUnreachable address unreachable notification
	SctpAddrUnreachable

	// SctpAddrRemoved the address is no longer part of the association.
	SctpAddrRemoved

	// SctpAddrAdded the address is now part of the association.
	SctpAddrAdded

	// SctpAddrMadePrim  the address has now been made the primary
	// destination address.
	SctpAddrMadePrim
)
