// Copyright 2022-present Open Networking Foundation.
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
	"testing"
)

func Test_getCanonicalChoiceIndexError(t *testing.T) {
	pd1 := perBitData{
		bytes:      []byte{0x40, 0x80, 0xFF, 0x00},
		byteOffset: 0,
		bitsOffset: 2,
	}

	err1 := pd1.getCanonicalChoiceIndex()
	assert.EqualError(t, err1, "Checksum didn't pass. Expecting 255 bytes, but have 1 bytes to decode")

	pd2 := perBitData{
		bytes:      []byte{0x40, 0x0F},
		byteOffset: 0,
		bitsOffset: 2,
	}

	err2 := pd2.getCanonicalChoiceIndex()
	assert.EqualError(t, err2, "Checksum didn't pass. Expecting 15 bytes, but have 0 bytes to decode")

	pd3 := perBitData{
		bytes:      []byte{0x40, 0x7F, 0x00, 0x00, 0x00},
		byteOffset: 0,
		bitsOffset: 2,
	}

	err3 := pd3.getCanonicalChoiceIndex()
	assert.EqualError(t, err3, "Checksum didn't pass. Expecting 127 bytes, but have 3 bytes to decode")
}
