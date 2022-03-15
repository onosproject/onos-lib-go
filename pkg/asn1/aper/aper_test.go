// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

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
	assert.Nil(t, err1)

	pd2 := perBitData{
		bytes:      []byte{0x40, 0x0F},
		byteOffset: 0,
		bitsOffset: 2,
	}

	err2 := pd2.getCanonicalChoiceIndex()
	assert.Nil(t, err2)

	pd3 := perBitData{
		bytes:      []byte{0x40, 0x7F, 0x00, 0x00, 0x00},
		byteOffset: 0,
		bitsOffset: 2,
	}

	err3 := pd3.getCanonicalChoiceIndex()
	assert.Nil(t, err3)
}
