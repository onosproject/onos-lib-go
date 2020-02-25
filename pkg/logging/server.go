// Copyright 2020-present Open Networking Foundation.
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

package logging

import (
	"github.com/onosproject/onos-lib-go/api/logging"
	"github.com/onosproject/onos-lib-go/pkg/logging/service"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// NewService returns a new device Service
func NewService() (service.Service, error) {
	return &Service{}, nil
}

// Service is an implementation of C1 service.
type Service struct {
	service.Service
}

// Register registers the logging Service with the gRPC server.
func (s Service) Register(r *grpc.Server) {
	server := &Server{}
	logging.RegisterLoggerServer(r, server)
}

// Server implements the logging gRPC service
type Server struct {
}

// SetLevel implements SetLevel rpc function to set a logger level
func (s *Server) SetLevel(ctx context.Context, req *logging.SetLevelRequest) (*logging.SetLevelResponse, error) {

	return &logging.SetLevelResponse{}, nil
}

// SetSink implements SetSink rpc function to set a sink for a logger
func (s *Server) SetSink(ctx context.Context, req *logging.SetSinkRequest) (*logging.SetSinkResponse, error) {

	return &logging.SetSinkResponse{}, nil
}
