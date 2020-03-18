// Copyright 2019-present Open Networking Foundation.
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

package atomix

import (
	"fmt"
	"github.com/onosproject/onos-lib-go/pkg/env"
)

// Config is the Atomix configuration
type Config struct {
	// Controller is the Atomix controller address
	Controller string `yaml:"controller,omitempty"`
	// Namespace is the Atomix namespace
	Namespace string `yaml:"namespace,omitempty"`
	// Scope is the Atomix client/application scope
	Scope string `yaml:"scope,omitempty"`
	// Protocols is a mapping of protocol types to databases
	Protocols map[string]string `yaml:"protocols"`
}

// GetController gets the Atomix controller address
func (c Config) GetController() string {
	if c.Controller == "" {
		namespace := c.GetNamespace()
		if namespace != "" {
			c.Controller = fmt.Sprintf("atomix-controller.%s.svc.cluster.local:5679", namespace)
		}
	}
	return c.Controller
}

// GetNamespace gets the Atomix client namespace
func (c Config) GetNamespace() string {
	if c.Namespace == "" {
		c.Namespace = env.GetServiceNamespace()
	}
	return c.Namespace
}

// GetScope gets the Atomix client scope
func (c Config) GetScope() string {
	if c.Scope == "" {
		c.Scope = env.GetServiceName()
	}
	if c.Scope == "" {
		c.Scope = c.GetNamespace()
	}
	return c.Scope
}

// GetDatabase gets the database name for the given protocol
func (c Config) GetDatabase(protocol string) string {
	if db, ok := c.Protocols[protocol]; ok {
		return db
	}
	return protocol
}
