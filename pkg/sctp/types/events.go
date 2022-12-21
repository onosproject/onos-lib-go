// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package types

// NotificationType sctp notification type
type NotificationType uint16

const (
	// SctpSnTypeBase ...
	SctpSnTypeBase = NotificationType(iota + (1 << 15))
	// SctpAssocChange ...
	SctpAssocChange
	// SctpPeerAddrChange ...
	SctpPeerAddrChange
	// SctpSendFailed ...
	SctpSendFailed
	// SctpRemoteError ...
	SctpRemoteError
	// SctpShutdownEvent ...
	SctpShutdownEvent
	// SctpPartialDeliveryEvent ...
	SctpPartialDeliveryEvent
	// SctpAdaptationIndication ...
	SctpAdaptationIndication
	// SctpAuthenticationIndication ...
	SctpAuthenticationIndication
	// SctpSenderDryEvent ...
	SctpSenderDryEvent
)

func (n NotificationType) String() string {
	switch n {
	case SctpAssocChange:
		return "SCTP_ASSOC_CHANGE"
	case SctpPeerAddrChange:
		return "SCTP_PEER_ADDR_CHANGE"
	case SctpSendFailed:
		return "SCTP_SEND_FAILED"
	case SctpRemoteError:
		return "SCTP_REMOTE_ERROR"
	case SctpShutdownEvent:
		return "SCTP_SHUTDOWN_EVENT"
	case SctpPartialDeliveryEvent:
		return "SCTP_PARTIAL_DELIVERY_EVENT"
	case SctpAdaptationIndication:
		return "SCTP_ADAPTATION_INDICATION"
	case SctpAuthenticationIndication:
		return "SCTP_AUTHENTICATION_INDICATION"
	case SctpSenderDryEvent:
		return "SCTP_SENDER_DRY_EVENT"
	default:
		return "UNKNOWN"

	}
}

const (
	// SctpEventDataIo ...
	SctpEventDataIo = 1 << iota
	// SctpEventAssociation communication notifications inform the application that an SCTP
	// association has either begun or ended.
	SctpEventAssociation
	// SctpEventAddress when a destination address of a multi-homed peer encounters a state
	//   change, a peer address change event is sent.
	SctpEventAddress
	// SctpEventSendFailure If SCTP cannot deliver a message, it can return back the message as a
	//   notification if the SCTP_SEND_FAILED_EVENT event is enabled.
	SctpEventSendFailure
	// SctpEventPeerError ...
	SctpEventPeerError
	// SctpEventShutdown  When a peer sends a SHUTDOWN, SCTP delivers this notification to
	// inform the application that it should cease sending data.
	SctpEventShutdown
	// SctpEventPartialDelivery  When a receiver is engaged in a partial delivery of a message, this
	//   notification will be used to indicate various events.
	SctpEventPartialDelivery
	// SctpEventAdaptationLayer When a peer sends an Adaptation Layer Indication parameter as
	//   described in [RFC5061], SCTP delivers this notification to inform the
	//   application about the peer's adaptation layer indication.
	SctpEventAdaptationLayer
	// SctpEventAuthentication defines an extension to authenticate SCTP messages.  The
	//   following notification is used to report different events relating to
	//   the use of this extension.
	SctpEventAuthentication
	// SctpEventSenderDry   When the SCTP stack has no more user data to send or retransmit, this
	//   notification is given to the user.  Also, at the time when a user app
	//   subscribes to this event, if there is no data to be sent or
	//   retransmit, the stack will immediately send up this notification.
	SctpEventSenderDry

	// SctpEventAll ...
	SctpEventAll = SctpEventDataIo | SctpEventAssociation | SctpEventAddress | SctpEventSendFailure | SctpEventPeerError | SctpEventShutdown | SctpEventPartialDelivery | SctpEventAdaptationLayer | SctpEventAuthentication | SctpEventSenderDry
)

// EventSubscribe Data structure for different types of SCTP events.
//
//	An SCTP application may need to understand and process events and
//	errors that happen on the SCTP stack.  These events include network
//	status changes, association startups, remote operational errors, and
//	undeliverable messages.  All of these can be essential for the
//	application.
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

// Event SCTP event
type Event func(*EventSubscribe)

// WithDataIO sets data io event
func WithDataIO() func(*EventSubscribe) {
	return func(subscribe *EventSubscribe) {
		subscribe.DataIO = 1

	}
}

// WithAssociation sets association event
func WithAssociation() func(*EventSubscribe) {
	return func(subscribe *EventSubscribe) {
		subscribe.Association = 1

	}
}

// WithAddress sets address event
func WithAddress() func(*EventSubscribe) {
	return func(subscribe *EventSubscribe) {
		subscribe.Address = 1

	}
}

// WithSendFailure sets send failure event
func WithSendFailure() func(*EventSubscribe) {
	return func(subscribe *EventSubscribe) {
		subscribe.SendFailure = 1

	}
}

// WithPeerError sets peer error event
func WithPeerError() func(*EventSubscribe) {
	return func(subscribe *EventSubscribe) {
		subscribe.PeerError = 1

	}
}

// WithShutdown sets shut down event
func WithShutdown() func(*EventSubscribe) {
	return func(subscribe *EventSubscribe) {
		subscribe.Shutdown = 1

	}
}

// WithPartialDelivery sets partial delivery event
func WithPartialDelivery() func(*EventSubscribe) {
	return func(subscribe *EventSubscribe) {
		subscribe.PartialDelivery = 1

	}
}

// WithAdaptationLayer sets adaptation layer event
func WithAdaptationLayer() func(*EventSubscribe) {
	return func(subscribe *EventSubscribe) {
		subscribe.AdaptationLayer = 1

	}
}

// WithAuthentication sets authentication event
func WithAuthentication() func(*EventSubscribe) {
	return func(subscribe *EventSubscribe) {
		subscribe.Authentication = 1

	}
}

// WithSenderDry sets sender dry event
func WithSenderDry() func(*EventSubscribe) {
	return func(subscribe *EventSubscribe) {
		subscribe.SenderDry = 1

	}
}
