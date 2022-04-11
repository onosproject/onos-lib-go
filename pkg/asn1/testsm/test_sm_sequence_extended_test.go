// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package testsm

import (
	"encoding/hex"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"gotest.tools/assert"
	"testing"
)

func Test_SequenceExtended(t *testing.T) {

	list := &TestList2{
		Value: make([]*ItemExtensible, 0),
	}
	item := &ItemExtensible{
		Item1: 1234,
		Item2: []byte{0xaa, 0xbb, 0xcc},
	}
	list.Value = append(list.Value, item)

	msg1 := &SequenceExtended{
		Se1: &SampleConstrainedInteger{
			Value: 256,
		},
		Se2: &TestOctetString{
			AttrOs1: []byte{0xff, 0xac, 0xbd, 0xef, 0x3d},
			AttrOs2: []byte{0xff, 0xac},
			AttrOs3: []byte{0xff, 0xac, 0xbd},
			AttrOs4: []byte{0xff, 0xac, 0xbd},
			AttrOs5: []byte{0xff, 0xac, 0xbd},
			AttrOs6: []byte{0xff, 0xac, 0xbd, 0xef, 0x3d},
		},
		Se3: list,
	}

	aperBytes1, err := aper.MarshalWithParams(msg1, "valueExt", Choicemap, CanonicalChoicemap)
	assert.NilError(t, err)
	assert.Assert(t, aperBytes1 != nil)
	t.Logf("APER \n%s", hex.Dump(aperBytes1))

	// Now decode the bytes and compare messages
	result1 := &SequenceExtended{}
	err = aper.UnmarshalWithParams(aperBytes1, result1, "valueExt", Choicemap, CanonicalChoicemap)
	assert.NilError(t, err)
	assert.Assert(t, result1 != nil)
	t.Logf("Decoded message is\n%v", result1)
	assert.Equal(t, msg1.String(), result1.String())

	msg2 := &SequenceExtended{
		Se1: &SampleConstrainedInteger{
			Value: 256,
		},
		Se2: &TestOctetString{
			AttrOs1: []byte{0xff, 0xac, 0xbd, 0xef, 0x3d},
			AttrOs2: []byte{0xff, 0xac},
			AttrOs3: []byte{0xff, 0xac, 0xbd},
			AttrOs4: []byte{0xff, 0xac, 0xbd},
			AttrOs5: []byte{0xff, 0xac, 0xbd},
			AttrOs6: []byte{0xff, 0xac, 0xbd, 0xef, 0x3d},
		},
		Se3: list,
		Se4: &TestConstrainedInt{
			AttrCiA: 11,
			AttrCiB: 256,
			AttrCiC: 99,
			AttrCiD: -21,
			AttrCiE: 20,
			AttrCiF: 10,
			AttrCiG: 11,
		},
	}

	aperBytes2, err := aper.MarshalWithParams(msg2, "valueExt", Choicemap, CanonicalChoicemap)
	assert.NilError(t, err)
	assert.Assert(t, aperBytes2 != nil)
	t.Logf("APER \n%s", hex.Dump(aperBytes2))

	// Now decode the bytes and compare messages
	result2 := &SequenceExtended{}
	err = aper.UnmarshalWithParams(aperBytes2, result2, "valueExt", Choicemap, CanonicalChoicemap)
	assert.NilError(t, err)
	assert.Assert(t, result2 != nil)
	t.Logf("Decoded message is\n%v", result2)
	assert.Equal(t, msg2.String(), result2.String())

	var se5 = "onfForever"
	msg3 := &SequenceExtended{
		Se1: &SampleConstrainedInteger{
			Value: 256,
		},
		Se2: &TestOctetString{
			AttrOs1: []byte{0xff, 0xac, 0xbd, 0xef, 0x3d},
			AttrOs2: []byte{0xff, 0xac},
			AttrOs3: []byte{0xff, 0xac, 0xbd},
			AttrOs4: []byte{0xff, 0xac, 0xbd},
			AttrOs5: []byte{0xff, 0xac, 0xbd},
			AttrOs6: []byte{0xff, 0xac, 0xbd, 0xef, 0x3d},
		},
		Se3: list,
		Se4: &TestConstrainedInt{
			AttrCiA: 11,
			AttrCiB: 256,
			AttrCiC: 99,
			AttrCiD: -21,
			AttrCiE: 20,
			AttrCiF: 10,
			AttrCiG: 11,
		},
		Se5: &se5,
	}

	aperBytes3, err := aper.MarshalWithParams(msg3, "valueExt", Choicemap, CanonicalChoicemap)
	assert.NilError(t, err)
	assert.Assert(t, aperBytes3 != nil)
	t.Logf("APER \n%s", hex.Dump(aperBytes3))

	// Now decode the bytes and compare messages
	result3 := &SequenceExtended{}
	err = aper.UnmarshalWithParams(aperBytes3, result3, "valueExt", Choicemap, CanonicalChoicemap)
	assert.NilError(t, err)
	assert.Assert(t, result3 != nil)
	t.Logf("Decoded message is\n%v", result3)
	assert.Equal(t, msg3.String(), result3.String())
}
