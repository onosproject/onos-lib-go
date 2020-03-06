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

package southbound

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

// Connect establishes a client-side connection to the gRPC end-point.
func Connect(ctx context.Context, address string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	conn, err := grpc.DialContext(ctx, address, opts...)
	if err != nil {
		fmt.Println("Can't connect", err)
		return nil, err
	}
	return conn, nil
}
