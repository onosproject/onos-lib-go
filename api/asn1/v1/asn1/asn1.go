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

package asn1

import (
	"encoding/binary"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"math"
)

// GetValueBytes - convert the internal uint64 to []byte format
func (m *BitString) GetValueBytes() []byte {
	if m != nil {
		numBytes := uint(math.Ceil(float64(m.Len) / 8.0))
		valueBytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(valueBytes, m.Value)
		return valueBytes[:numBytes]
	}
	return nil
}

// UpdateValue - replace the uint64 value with values from a new []byte
// the size stays the same
func (m *BitString) UpdateValue(newBytes []byte) (uint64, error) {
	if m == nil {
		return m.Value, errors.NewInvalid("null")
	}
	expectedLen := int(math.Ceil(float64(m.Len) / 8.0))
	if len(newBytes) != expectedLen {
		return m.Value, errors.NewInvalid("too many bytes %d. Expecting %d", len(newBytes), expectedLen)
	}
	fullBytes := make([]byte, 8)
	copy(fullBytes, newBytes)
	m.Value = binary.LittleEndian.Uint64(fullBytes)
	return m.Value, nil
}
