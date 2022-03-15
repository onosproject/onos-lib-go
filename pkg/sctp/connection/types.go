// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package connection

// GetAddrs ...
type GetAddrs struct {
	assocID int32
	addrNum uint32
	addrs   [4096]byte
}

// PeeloffArg ...
type PeeloffArg struct {
	assocID int32
	sd      int
}
