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

package controller

// NewID creates a new object identifier
func NewID(value interface{}) ID {
	return ID{value}
}

// ID is an object identifier
type ID struct {
	// Value is the raw identifier
	Value interface{}
}

// Int returns the identifier as an integer
func (i ID) Int() int {
	return i.Value.(int)
}

// Int32 returns the identifier as an integer
func (i ID) Int32() int32 {
	return i.Value.(int32)
}

// Int64 returns the identifier as an integer
func (i ID) Int64() int64 {
	return i.Value.(int64)
}

// Uint returns the identifier as an integer
func (i ID) Uint() uint {
	return i.Value.(uint)
}

// Uint32 returns the identifier as an integer
func (i ID) Uint32() uint32 {
	return i.Value.(uint32)
}

// Uint64 returns the identifier as an integer
func (i ID) Uint64() uint64 {
	return i.Value.(uint64)
}

// String returns the identifier as a string
func (i ID) String() string {
	return i.Value.(string)
}
