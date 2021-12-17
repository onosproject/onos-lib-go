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

package aper

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_appendEnumerated(t *testing.T) {
	pd := perRawBitData{}
	lowerBound := int64(1)
	upperBound := int64(2)
	err := pd.appendEnumerated(1, true, &lowerBound, &upperBound)
	assert.NoError(t, err)
}

func Test_appendOpenType(t *testing.T) {
	pd := &perRawBitData{}
	v := reflect.ValueOf("abc")
	params := fieldParameters{
		sizeExtensible: true,
	}
	err := pd.appendOpenType(v, params)
	assert.NoError(t, err)
}

func Test_appendNormallySmallNonNegativeWholeNumber(t *testing.T) {

	pd := perRawBitData{}
	err := pd.appendNormallySmallNonNegativeWholeNumber(uint64(122))
	assert.NoError(t, err)
	assert.Equal(t, pd.bytes, []byte{0x7a})

	pd1 := perRawBitData{}
	err = pd1.appendNormallySmallNonNegativeWholeNumber(uint64(131))
	assert.NoError(t, err)
	assert.Equal(t, pd1.bytes, []byte{0x80, 0x83})

	pd2 := perRawBitData{}
	err = pd2.appendNormallySmallNonNegativeWholeNumber(uint64(2310))
	assert.NoError(t, err)
	assert.Equal(t, pd2.bytes, []byte{0x89, 0x06})

	pd3 := perRawBitData{}
	err = pd3.appendNormallySmallNonNegativeWholeNumber(uint64(3))
	assert.NoError(t, err)
	assert.Equal(t, pd3.bytes, []byte{0x03})

	pd4 := perRawBitData{}
	err = pd4.appendNormallySmallNonNegativeWholeNumber(uint64(77))
	assert.NoError(t, err)
	assert.Equal(t, pd4.bytes, []byte{0x4d})

	pd5 := perRawBitData{}
	err = pd5.appendNormallySmallNonNegativeWholeNumber(uint64(113))
	assert.NoError(t, err)
	assert.Equal(t, pd5.bytes, []byte{0x71})

	pd6 := perRawBitData{}
	err = pd6.appendNormallySmallNonNegativeWholeNumber(uint64(177))
	assert.NoError(t, err)
	assert.Equal(t, pd6.bytes, []byte{0x80, 0xb1})

	pd7 := perRawBitData{}
	err = pd7.appendNormallySmallNonNegativeWholeNumber(uint64(204))
	assert.NoError(t, err)
	assert.Equal(t, pd7.bytes, []byte{0x80, 0xcc})

	pd8 := perRawBitData{}
	err = pd8.appendNormallySmallNonNegativeWholeNumber(uint64(231))
	assert.NoError(t, err)
	assert.Equal(t, pd8.bytes, []byte{0x80, 0xe7})

	pd9 := perRawBitData{}
	err = pd9.appendNormallySmallNonNegativeWholeNumber(uint64(249))
	assert.NoError(t, err)
	assert.Equal(t, pd9.bytes, []byte{0x80, 0xf9})

	pd10 := perRawBitData{}
	err = pd10.appendNormallySmallNonNegativeWholeNumber(uint64(258))
	assert.NoError(t, err)
	assert.Equal(t, pd10.bytes, []byte{0x81, 0x02})

	pd11 := perRawBitData{}
	err = pd11.appendNormallySmallNonNegativeWholeNumber(uint64(285))
	assert.NoError(t, err)
	assert.Equal(t, pd11.bytes, []byte{0x81, 0x1d})

	pd12 := perRawBitData{}
	err = pd12.appendNormallySmallNonNegativeWholeNumber(uint64(465))
	assert.NoError(t, err)
	assert.Equal(t, pd12.bytes, []byte{0x81, 0xd1})

	pd13 := perRawBitData{}
	err = pd13.appendNormallySmallNonNegativeWholeNumber(uint64(915))
	assert.NoError(t, err)
	assert.Equal(t, pd13.bytes, []byte{0x83, 0x93})

	pd14 := perRawBitData{}
	err = pd14.appendNormallySmallNonNegativeWholeNumber(uint64(1365))
	assert.NoError(t, err)
	assert.Equal(t, pd14.bytes, []byte{0x85, 0x55})

	pd15 := perRawBitData{}
	err = pd15.appendNormallySmallNonNegativeWholeNumber(uint64(1815))
	assert.NoError(t, err)
	assert.Equal(t, pd15.bytes, []byte{0x87, 0x17})

	pd16 := perRawBitData{}
	err = pd16.appendNormallySmallNonNegativeWholeNumber(uint64(2319))
	assert.NoError(t, err)
	assert.Equal(t, pd16.bytes, []byte{0x89, 0x0f})
}
