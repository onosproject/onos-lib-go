// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package controller

// Filter filters individual events for a node
// Each time an event is received from a watcher, the filter has the option to discard the request by
// returning false.
type Filter interface {
	// Accept indicates whether to accept the given object
	Accept(id ID) bool
}
