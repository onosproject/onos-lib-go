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
	"errors"
	"strings"

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

// SetDebug enable/disable debug mode for logging package
func (s *Server) SetDebug(ctx context.Context, req *logging.SetDebugModeRequest) (*logging.SetDebugModeResponse, error) {
	return &logging.SetDebugModeResponse{
		ResponseStatus: logging.ResponseStatus_OK,
	}, nil

}

// SetLevel implements SetLevel rpc function to set a logger level
func (s *Server) SetLevel(ctx context.Context, req *logging.SetLevelRequest) (*logging.SetLevelResponse, error) {
	name := req.GetLoggerName()
	level := req.GetLevel()
	if name == "" {
		return &logging.SetLevelResponse{
			ResponseStatus: logging.ResponseStatus_PRECONDITION_FAILED,
		}, errors.New("precondition for set level request is failed")
	}

	names := strings.Split(name, "/")
	logger := GetLogger(names...)

	switch level {
	case logging.Level_INFO:
		logger.SetLevel(InfoLevel)
	case logging.Level_DEBUG:
		logger.SetLevel(DebugLevel)
	case logging.Level_WARN:
		logger.SetLevel(WarnLevel)
	case logging.Level_ERROR:
		logger.SetLevel(ErrorLevel)
	case logging.Level_PANIC:
		logger.SetLevel(PanicLevel)
	case logging.Level_DPANIC:
		logger.SetLevel(DPanicLevel)
	case logging.Level_FATAL:
		logger.SetLevel(FatalLevel)

	default:
		return &logging.SetLevelResponse{
			ResponseStatus: logging.ResponseStatus_PRECONDITION_FAILED,
		}, errors.New("the requested level is not supported")

	}

	return &logging.SetLevelResponse{
		ResponseStatus: logging.ResponseStatus_OK,
	}, nil
}
