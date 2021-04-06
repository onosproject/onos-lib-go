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

package defs

import (
	"unsafe"

	syscall "golang.org/x/sys/unix"
)

// EventSubscribe ...
type EventSubscribe struct {
	DataIO          uint8
	Association     uint8
	Address         uint8
	SendFailure     uint8
	PeerError       uint8
	Shutdown        uint8
	PartialDelivery uint8
	AdaptationLayer uint8
	Authentication  uint8
	SenderDry       uint8
}

// InitMsg ...
type InitMsg struct {
	NumOstreams    uint16
	MaxInstreams   uint16
	MaxAttempts    uint16
	MaxInitTimeout uint16
}

// NewDefaultInitMsg ...
func NewDefaultInitMsg() *InitMsg {
	return &InitMsg{
		NumOstreams: uint16(10),
	}
}

// SndRcvInfo ...
type SndRcvInfo struct {
	Stream  uint16
	SSN     uint16
	Flags   uint16
	_       uint16
	PPID    uint32
	Context uint32
	TTL     uint32
	TSN     uint32
	CumTSN  uint32
	AssocID int32
}

// NxtInfo ...
type NxtInfo struct {
	Stream  uint16
	Flags   uint16
	PPID    uint32
	Length  uint32
	AssocID int32
}

// SndInfo ...
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

// AssociationChange ...
type AssociationChange struct {
	Type            NotificationType
	Flags           uint16
	Length          uint32
	State           State
	Error           uint16
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

// RemoteError ...
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

// AdaptationIndication ...
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
