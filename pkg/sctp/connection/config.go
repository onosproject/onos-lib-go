// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

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
