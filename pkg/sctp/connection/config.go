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
	"github.com/onosproject/onos-lib-go/pkg/sctp/defs"
)

// Config sctp connection config
type Config struct {
	addressFamily defs.AddressFamily
	mode          defs.SocketMode
	options       defs.InitMsg
	nonblocking   bool
}

// NewConfig creates a conection config
func NewConfig(options ...func(cfg *Config)) *Config {
	cfg := &Config{}
	for _, option := range options {
		option(cfg)
	}

	return cfg
}

// WithAddressFamily sets address family
func WithAddressFamily(family defs.AddressFamily) func(config *Config) {
	return func(config *Config) {
		config.addressFamily = family

	}
}

// WithMode sets SCTP mode
func WithMode(mode defs.SocketMode) func(config *Config) {
	return func(config *Config) {
		config.mode = mode

	}
}

// WithOptions sets options
func WithOptions(options defs.InitMsg) func(config *Config) {
	return func(config *Config) {
		config.options = options

	}
}

// WithNonBlocking sets nonblocking
func WithNonBlocking(nonblocking bool) func(config *Config) {
	return func(config *Config) {
		config.nonblocking = nonblocking
	}
}
