// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package events

import (
	"github.com/onosproject/onos-lib-go/pkg/sctp/connection"
	"github.com/onosproject/onos-lib-go/pkg/sctp/types"
)

// IsType checks whether the given notification is of the given type
func IsType(notification *types.Notification, notificationType types.NotificationType) bool {
	return notification.Type() == notificationType
}

// IsSCTPAssocChange checks whether the given notfication is SCTP Association change event
func IsSCTPAssocChange(notification *types.Notification) bool {
	return IsType(notification, types.SctpAssocChange)
}

// IsSCTPPeerAddrChange checks whether the given notfication is SCTP peer address change event
func IsSCTPPeerAddrChange(notification *types.Notification) bool {
	return IsType(notification, types.SctpPeerAddrChange)
}

// IsSCTPSendFailed checks whether the given notfication is SCTP send failed event
func IsSCTPSendFailed(notification *types.Notification) bool {
	return IsType(notification, types.SctpSendFailed)
}

// IsSCTPRemoteError checks whether the given notfication is SCTP remote error event
func IsSCTPRemoteError(notification *types.Notification) bool {
	return IsType(notification, types.SctpRemoteError)
}

// IsSCTPPartialDeliveryEvent checks whether the given notfication is SCTP partial delivery event
func IsSCTPPartialDeliveryEvent(notification *types.Notification) bool {
	return IsType(notification, types.SctpPartialDeliveryEvent)
}

// IsSCTPShutdownEvent checks whether the given notfication is SCTP shutdown
func IsSCTPShutdownEvent(notification *types.Notification) bool {
	return IsType(notification, types.SctpShutdownEvent)
}

// IsSCTPAdaptationIndication checks whether the given notfication is SCTP adaptation indication event
func IsSCTPAdaptationIndication(notification *types.Notification) bool {
	return IsType(notification, types.SctpAdaptationIndication)
}

// IsSCTPAuthenticationIndication checks whether the given notfication is SCTP authentication indication event
func IsSCTPAuthenticationIndication(notification *types.Notification) bool {
	return IsType(notification, types.SctpAuthenticationIndication)
}

// IsSCTPSenderDryEvent checks whether the given notfication is SCTP sender dry event
func IsSCTPSenderDryEvent(notification *types.Notification) bool {
	return IsType(notification, types.SctpSenderDryEvent)
}

// IsNotification checks whether the given message is notification
func IsNotification(flags int) bool {
	return flags&types.MsgNotification > 0
}

// IsMsgEORSet checks if MsgEOR is set or not. If all data in a single message has been delivered, MSG_EOR will be
// set in the msg_flags field of the msghdr structure
func IsMsgEORSet(flags int) bool {
	return flags&types.MsgEOR > 0
}

// GetNotfication extracts notfication from a given message
func GetNotfication(buf []byte, _ int) (*types.Notification, error) {
	notif, err := connection.SCTPParseNotification(buf)
	if err != nil {
		return nil, err
	}
	return notif, nil
}
