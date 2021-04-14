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

package connection

import (
	"github.com/onosproject/onos-lib-go/pkg/sctp/types"
)

// Config sctp connection config
type Config struct {
	addressFamily types.AddressFamily
	mode          types.SocketMode
	initMsg       types.InitMsg
	nonblocking   bool
}

// NewConfig creates a connection config
func NewConfig(options ...func(cfg *Config)) *Config {
	cfg := &Config{}
	for _, option := range options {
		option(cfg)
	}

	return cfg
}

// WithAddressFamily sets address family
func WithAddressFamily(family types.AddressFamily) func(config *Config) {
	return func(config *Config) {
		config.addressFamily = family

	}
}

// WithMode sets SCTP mode
func WithMode(mode types.SocketMode) func(config *Config) {
	return func(config *Config) {
		config.mode = mode

	}
}

// WithOptions sets options
func WithOptions(initMsg types.InitMsg) func(config *Config) {
	return func(config *Config) {
		config.initMsg = initMsg

	}
}

// WithNonBlocking sets nonblocking
func WithNonBlocking(nonblocking bool) func(config *Config) {
	return func(config *Config) {
		config.nonblocking = nonblocking
	}
}
