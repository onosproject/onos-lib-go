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
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"math"
)

// UpdateValue - replace the bytes value with values from a new []byte
// the size stays the same
func (m *BitString) UpdateValue(newBytes []byte) ([]byte, error) {
	if m == nil {
		return m.Value, errors.NewInvalid("null")
	}
	expectedLen := int(math.Ceil(float64(m.Len) / 8.0))
	if len(newBytes) != expectedLen {
		return m.Value, errors.NewInvalid("too many bytes %d. Expecting %d", len(newBytes), expectedLen)
	}
	m.Value = newBytes
	return m.Value, nil
}

// TruncateValue - truncates value of trailing bits in the BitString the size stays the same
// Assuming that BitString has a non-empty length
func (m *BitString) TruncateValue() ([]byte, error) {
	if m == nil {
		return m.Value, errors.NewInvalid("null")
	}
	if m.Len == 0 {
		return nil, errors.NewInvalid("Length should not be 0")
	}
	// Computing the number of bytes
	expectedLen := int(math.Ceil(float64(m.Len) / 8.0))
	if len(m.Value) != expectedLen {
		return m.Value, errors.NewInvalid("too many bytes %d. Expecting %d", len(m.Value), expectedLen)
	}
	// Creating set of truncated bytes, with trailing zeroes
	// Since we've got there, value in expectedLen is correct
	truncBytes := make([]byte, expectedLen)
	for i := 0; i < expectedLen; i++ {
		truncBytes[i] = m.Value[i]
	}

	bitsFull := expectedLen * 8
	trailingBits := uint32(bitsFull) - m.Len

	mask := ^((1 << trailingBits) - 1)
	truncBytes[len(truncBytes)-1] = truncBytes[len(truncBytes)-1] & byte(mask)
	//fmt.Printf("Last byte after truncation is %x\n", truncBytes[len(truncBytes)-1])
	m.Value = truncBytes
	return m.Value, nil
}
