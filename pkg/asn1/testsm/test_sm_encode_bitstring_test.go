// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package testsm

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
				{Value: []byte{0xff, 0xc0}, Len: 10},
				{Value: []byte{0xff, 0xee, 0xd0}, Len: 20},
				{Value: []byte{0xff, 0xee, 0xd0}, Len: 20},
				{Value: []byte{}, Len: 0},
				{Value: []byte{0xff, 0xee, 0xfc}, Len: 22},
				{Value: []byte{0xff, 0xee, 0xdd, 0xc0}, Len: 28},
				{Value: []byte{0xff, 0xee, 0xfc}, Len: 22},
			},
			[]byte{
				0x80,       // flag to show optional Bs7 is present - there's only 1 optional field so it's first bit
				0x0a,       // length of bBs1 (unconstrained) 10 - it's given the full byte because it could be anything
				0xff, 0xc0, // value of Bs1
				0xff, 0xee, 0xd0, // value of Bs2 - no length needed, since constrained
				0xff, 0xee, 0xd0, // value of Bs3 - no length needed, since constrained
				0x00,             // length of Bs4 in first 5 bits = 0, and no value. Then length of Bs5 in next 3 bits = 0 (over 22)
				0xff, 0xee, 0xfc, // value of Bs5
				0x00,                   // for is size extensible and 0 for length (over 28)
				0xff, 0xee, 0xdd, 0xc0, // value of Bs6
				// Does length of Bs7 get rolled in to previous, or is omitted when 0?
				0xff, 0xee, 0xfc, // value of Bs7 over
			},
		},
		{
			[]*asn1.BitString{
				{Value: []byte{0xff, 0xee, 0xdd, 0xcf}, Len: 32},
				{Value: []byte{0xff, 0xee, 0xc0}, Len: 20},
				{Value: []byte{0xff, 0xee, 0xc0}, Len: 20},
				{Value: []byte{0xff, 0xee, 0xc0}, Len: 18},
				{Value: []byte{0xff, 0xee, 0xdd, 0xcc}, Len: 32},
				{Value: []byte{0xff, 0xee, 0xdd, 0xcc, 0xbb}, Len: 40},
				{Value: []byte{0xff, 0xee, 0xdd, 0xcc, 0xb0}, Len: 36},
			},
			[]byte{
				0x80,                   // flag to show optional Bs7 is present - there's only 1 optional field so it's first bit
				0x20,                   // length of Bs1 = 32 (unconstrained)
				0xff, 0xee, 0xdd, 0xcf, // value of Bs1 - no bit zeroed - fits byte boundary
				// no length for Bs2 as it's fixed
				0xff, 0xee, 0xc0, // value of Bs2
				// no length for Bs3 as it's fixed
				0xff, 0xee, 0xc9, // value of Bs3. Last 4 bits and 1 bit from next byte give 10010 = 18 (length of Bs4)
				0x00,             // Bs4 length
				0xff, 0xee, 0xe8, // Bs4 value with last byte containing 2 bits = 11 from Bs4 value, then 1010 (10 over 22 = 32) for the length of Bs5
				0xff, 0xee, 0xdd, 0xcc, // Value of Bs5
				0x80,                         // Bs6 size extended
				0x28,                         // Length of Bs6 = 40
				0xff, 0xee, 0xdd, 0xcc, 0xbb, // Value of Bs6
				0xe0, // Length of Bs7 1110 = 14 (over 22)
				0xff, 0xee, 0xdd, 0xcc, 0xb0,
			},
		},
	}

	for _, tc := range testCases {
		attrBs1 := tc.value[0]
		attrBs2 := tc.value[1]
		attrBs3 := tc.value[2]
		attrBs4 := tc.value[3]
		attrBs5 := tc.value[4]
		attrBs6 := tc.value[5]
		attrBs7 := tc.value[6]
		tbs := &TestBitString{
			AttrBs1: attrBs1,
			AttrBs2: attrBs2,
			AttrBs3: attrBs3,
			AttrBs4: attrBs4,
			AttrBs5: attrBs5,
			AttrBs6: attrBs6,
			AttrBs7: attrBs7,
		}
		aper, err := aper.Marshal(tbs, Choicemap, CanonicalChoicemap)
		assert.NoError(t, err)
		assert.NotNil(t, aper)
		t.Logf("APER \n%s", hex.Dump(aper))
		assert.EqualValues(t, tc.expected, aper)
	}

}
