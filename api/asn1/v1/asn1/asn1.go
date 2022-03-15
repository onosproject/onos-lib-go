// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

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
	expectedBytesLen := int(math.Ceil(float64(m.Len) / 8.0))
	if len(m.Value) != expectedBytesLen {
		return m.Value, errors.NewInvalid("too many bytes %d. Expecting %d", len(m.Value), expectedBytesLen)
	}
	// Creating set of truncated bytes, with trailing zeroes
	// Since we've got there, value in expectedLen is correct
	truncBytes := make([]byte, expectedBytesLen)
	for i := 0; i < expectedBytesLen; i++ {
		truncBytes[i] = m.Value[i]
	}

	bitsFull := expectedBytesLen * 8
	trailingBits := uint32(bitsFull) - m.Len

	mask := ^((1 << trailingBits) - 1)
	truncBytes[len(truncBytes)-1] = truncBytes[len(truncBytes)-1] & byte(mask)
	//fmt.Printf("Last byte after truncation is %x\n", truncBytes[len(truncBytes)-1])
	m.Value = truncBytes
	return m.Value, nil
}
