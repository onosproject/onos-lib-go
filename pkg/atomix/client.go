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
	"context"
	"fmt"
	"github.com/atomix/api/proto/atomix/controller"
	"github.com/atomix/go-client/pkg/client"
	netutil "github.com/atomix/go-client/pkg/client/util/net"
	"github.com/atomix/go-framework/pkg/atomix"
	"github.com/atomix/go-framework/pkg/atomix/registry"
	"github.com/atomix/go-local/pkg/atomix/local"
	"net"
	"time"
)

const basePort = 45000

// StartLocalNode starts a single local Atomix node
func StartLocalNode() (*atomix.Node, netutil.Address) {
	for port := basePort; port < basePort+100; port++ {
		address := netutil.Address(fmt.Sprintf("localhost:%d", port))
		lis, err := net.Listen("tcp", string(address))
		if err != nil {
			continue
		}
		node := local.NewNode(lis, registry.Registry, []*controller.PartitionId{
			{
				Partition: 1,
			},
		})
		_ = node.Start()
		return node, address
	}
	panic("cannot find open port")
}

// GetClient returns the Atomix client
func GetClient(config Config) (*client.Client, error) {
	opts := []client.Option{
		client.WithNamespace(config.GetNamespace()),
		client.WithScope(config.GetScope()),
	}
	return client.NewClient(config.GetController(), opts...)
}

// GetDatabase returns the Atomix database
func GetDatabase(config Config, database string) (*client.Database, error) {
	client, err := GetClient(config)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	return client.GetDatabase(ctx, database)
}
