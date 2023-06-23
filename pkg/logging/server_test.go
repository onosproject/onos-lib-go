// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package logging

import (
	"testing"

	"google.golang.org/grpc"
)

func TestService_Register(_ *testing.T) {
	service := Service{}
	server := grpc.NewServer()
	service.Register(server)
}
