// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"encoding/hex"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testBitString struct {
	value    []*asn1.BitString
	expected []byte
}

func Test_BitString(t *testing.T) {
	testCases := []testBitString{
		{
			[]*asn1.BitString{
				{Value: []byte{0xff, 0xee, 0xdc}, Len: 22},
				{Value: []byte{0xff, 0xee, 0xd0}, Len: 20},
				{Value: []byte{0xff, 0xee, 0xde}, Len: 23},
				{Value: []byte{0xff, 0xee, 0xdd, 0xc0}, Len: 28},
				{Value: []byte{0xff, 0xee, 0xdd, 0xcf}, Len: 32},
			},
			[]byte{
				0x80,             // flag to show optional 5th item is present - there's only 1 optional field so it's first bit
				0x16,             // length = 22
				0xff, 0xee, 0xdc, // last 2 bits are zeroed, so df becomes dc
				// no length as it's fixed
				0xff, 0xee, 0xd1, // last 4 bits are zeroed, so df should be d0 - is d1 (maybe this belongs to next part)
				0xff, 0xee, 0xde, // length is 1 (over 22). last 1 bit are zeroed, so df becomes de
				0x00,                   // for is extensible
				0xff, 0xee, 0xdd, 0xc0, // length is 0 (over 28). last 1 bit are zeroed, so cf becomes c0
				0xff, 0xee, 0xdd, 0xcf, // length is 0 (over 32)
			},
		},
		{
			[]*asn1.BitString{
				{Value: []byte{0xff, 0xee, 0xdd, 0xcf}, Len: 32},
				{Value: []byte{0xff, 0xee, 0xc0}, Len: 20},
				{Value: []byte{0xff, 0xee, 0xdd, 0xcf}, Len: 32},
				{Value: []byte{0xff, 0xee, 0xdd, 0xcf}, Len: 32},
				{Value: []byte{0xff, 0xee, 0xdd, 0xcc, 0xb0}, Len: 36},
			},
			[]byte{
				0x80,                   // flag to show optional 5th item is present - there's only 1 optional field so it's first bit
				0x20,                   // length = 32
				0xff, 0xee, 0xdd, 0xcf, // no bit zeroed - fits byte boundary
				// no length as it's fixed
				0xff, 0xee, 0xca, // last 4 bits are zeroed, so cf becomes c0
				// no length - was expecting 12 (over 20)
				0xff, 0xee, 0xdd, 0xcf, // no bits are zeroed - fits byte boundary
				0x80,                   // first bit for extensible - then
				0xff, 0xee, 0xdd, 0xcf, // length is 0 (over 28). last 1 bit are zeroed, so cf becomes c0
				0x80,                         // length = 4 (over 32)
				0xff, 0xee, 0xdd, 0xcc, 0xb0, // last 4 bits are zeroed do bb became b0
			},
		},
	}

	for _, tc := range testCases {
		attrBs1 := tc.value[0]
		attrBs2 := tc.value[1]
		attrBs3 := tc.value[2]
		attrBs4 := tc.value[3]
		attrBs5 := tc.value[4]
		tbs := &TestBitString{
			AttrBs1: attrBs1,
			AttrBs2: attrBs2,
			AttrBs3: attrBs3,
			AttrBs4: attrBs4,
			AttrBs5: attrBs5,
		}
		aper, err := aper.Marshal(tbs, nil, nil)
		assert.NoError(t, err)
		assert.NotNil(t, aper)
		t.Logf("%v %v %v gives APER %s", attrBs1, attrBs2, attrBs3, hex.Dump(aper))
		assert.EqualValues(t, tc.expected, aper)
	}

}
