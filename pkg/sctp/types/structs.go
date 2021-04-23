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
	"unsafe"

	syscall "golang.org/x/sys/unix"
)

// InitMsg  the default association initialization.
type InitMsg struct {
	// NumOstreams specifies the number of
	// streams to which the application wishes to be able to send.
	NumOstreams uint16
	// MaxInstreams specifies the maximum number of
	// inbound streams the application is prepared to support.
	MaxInstreams uint16
	// MaxAttempts specifies how many attempts the
	//  SCTP endpoint should make at resending the INIT.
	MaxAttempts uint16
	// MaxInitTimeout specifies the largest timeout or
	// retransmission timeout (RTO) value (in milliseconds) to use in attempting an INIT.
	MaxInitTimeout uint16
}

// NewDefaultInitMsg creates new InitMsg
func NewDefaultInitMsg() *InitMsg {
	return &InitMsg{
		NumOstreams: uint16(10),
	}
}

// SndRcvInfo ...
type SndRcvInfo struct {
	// Stream stream number
	Stream uint16
	// SSN  this value contains the stream sequence
	// number that the remote endpoint placed in the DATA chunk.
	SSN uint16
	// Flags ...
	Flags uint16
	_     uint16
	// PPID This value in sendmsg() is an unsigned integer that is
	// passed to the remote end in each user message.
	PPID uint32
	// Context This value is an opaque 32-bit context datum that is
	// used in the sendmsg() function.  This value is passed back to the
	// upper layer if an error occurs on the send of a message and is
	// retrieved with each undelivered message.
	Context uint32
	// TTL For the sending side, this field contains the
	// message's time to live, in milliseconds.
	TTL uint32
	// TSN For the receiving side, this field holds a Transmission Sequence Number (TSN)
	// that was assigned to one of the SCTP DATA chunks.
	TSN uint32
	// CumTSN This field will hold the current cumulative TSN as known by the underlying SCTP layer.
	CumTSN uint32
	// AssocID  The association handle field, sinfo_assoc_id, holds the identifier for the association announced in the SCTP_COMM_UP
	//  notification.  All notifications for a given association have the
	//  same identifier.  This field is ignored for one-to-one style sockets.
	AssocID int32
}

// NxtInfo information of the next message that will be delivered  if this information
// is already available when delivering the current message.
type NxtInfo struct {
	// Stream next message's stream number
	Stream uint16
	Flags  uint16
	PPID   uint32
	// Length length of the message currently within
	// the socket buffer
	Length uint32
	// holds the identifier for the association announced
	// in the SCTP_COMM_UP notification.
	AssocID int32
}

// SndInfo SCTP options for sending a msg.
type SndInfo struct {
	Stream  uint16
	Flags   uint16
	PPID    uint32
	Context uint32
	AssocID int32
}

// GetAddrsOld ...
type GetAddrsOld struct {
	AssocID int32
	AddrNum int32
	Addrs   uintptr
}

// NotificationHeader ...
type NotificationHeader struct {
	Type   NotificationType
	Flags  uint16
	Length uint32
}

// Notification ...
type Notification struct {
	Data []byte
}

// Header gets notification header
func (n *Notification) Header() *NotificationHeader {
	return (*NotificationHeader)(unsafe.Pointer(&n.Data[0]))
}

// Type gets notification type
func (n *Notification) Type() NotificationType {
	return n.Header().Type
}

// GetAssociationChange ...
func (n *Notification) GetAssociationChange() *AssociationChange {
	return (*AssociationChange)(unsafe.Pointer(&n.Data[0]))
}

// GetPeerAddrChange  ...
func (n *Notification) GetPeerAddrChange() *PeerAddrChange {
	return (*PeerAddrChange)(unsafe.Pointer(&n.Data[0]))
}

// GetRemoteError ...
func (n *Notification) GetRemoteError() *RemoteError {
	return (*RemoteError)(unsafe.Pointer(&n.Data[0]))
}

// GetSendFailed ...
func (n *Notification) GetSendFailed() *SendFailed {
	return (*SendFailed)(unsafe.Pointer(&n.Data[0]))
}

// GetAdaptationIndication ...
func (n *Notification) GetAdaptationIndication() *AdaptationIndication {
	return (*AdaptationIndication)(unsafe.Pointer(&n.Data[0]))
}

// GetPartialDelivery ...
func (n *Notification) GetPartialDelivery() *PartialDelivery {
	return (*PartialDelivery)(unsafe.Pointer(&n.Data[0]))
}

// GetAuthentication ...
func (n *Notification) GetAuthentication() *Authentication {
	return (*Authentication)(unsafe.Pointer(&n.Data[0]))
}

// GetSenderDry ...
func (n *Notification) GetSenderDry() *SenderDry {
	return (*SenderDry)(unsafe.Pointer(&n.Data[0]))
}

// AssociationChange Communication notifications inform the application that an SCTP
//   association has either begun or ended.
type AssociationChange struct {
	Type  NotificationType
	Flags uint16
	// Length length of the notification data,
	Length uint32
	State  State
	//  If the state was reached due to an error condition (e.g.,
	//  SCTP_COMM_LOST), any relevant error information is available in this field.
	Error uint16
	// OutboundStreams  The maximum number of
	// streams allowed in each direction is available in
	// sac_outbound_streams and sac_inbound streams.
	OutboundStreams uint16
	InboundStreams  uint16
	AssocID         int32
	Info            []byte
}

// PeerAddrChange ...
type PeerAddrChange struct {
	Type   NotificationType
	Length uint32
	//Addr    C.struct_sockaddr_storage
	State   PeerChangeState
	Error   uint32
	AssocID int32
}

// RemoteError a remote peer may send an Operation Error message to its peer.
type RemoteError struct {
	Type    NotificationType
	Flags   uint16
	Length  uint32
	Error   uint16
	AssocID int32
	Info    []byte
}

// SendFailed ....
type SendFailed struct {
	Type    NotificationType
	Flags   uint16
	Length  uint32
	Error   uint16
	SndInfo SndInfo
	AssocID int32
	Data    []byte
}

// AdaptationIndication for informing the application about the peer's adaptation layer indication.
type AdaptationIndication struct {
	Type       NotificationType
	Flags      uint16
	Length     uint32
	Indication uint32
	AssocID    int32
}

// PartialDelivery ...
type PartialDelivery struct {
	Type           NotificationType
	Flags          uint16
	Length         uint32
	Indication     uint32
	StreamID       uint32
	SequenceNumber uint32
}

// Authentication ...
type Authentication struct {
	Type       NotificationType
	Flags      uint16
	Length     uint32
	KeyNumber  uint16
	Indication uint32
	AssocID    int32
}

// SenderDry ...
type SenderDry struct {
	Type    NotificationType
	Flags   uint16
	Length  uint32
	AssocID int32
}

// OOBMessage ...
type OOBMessage struct {
	syscall.SocketControlMessage
}

// IsSCTP ...
func (o *OOBMessage) IsSCTP() bool {
	return o.Header.Level == syscall.IPPROTO_SCTP
}

// Type ...
func (o *OOBMessage) Type() CmsgType {
	return CmsgType(o.Header.Type)
}

// GetSndRcvInfo  ...
func (o *OOBMessage) GetSndRcvInfo() *SndRcvInfo {
	return (*SndRcvInfo)(unsafe.Pointer(&o.Data[0]))
}

// GetSndInfo ...
func (o *OOBMessage) GetSndInfo() *SndInfo {
	return (*SndInfo)(unsafe.Pointer(&o.Data[0]))
}

// GetNxtInfo ...
func (o *OOBMessage) GetNxtInfo() *NxtInfo {
	return (*NxtInfo)(unsafe.Pointer(&o.Data[0]))
}
