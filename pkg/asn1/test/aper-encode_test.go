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

package test

import (
	"encoding/hex"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log := logging.GetLogger("asn1")
	log.SetLevel(logging.DebugLevel)
	os.Exit(m.Run())
}

type testint struct {
	values   []int
	expected []byte
}

type testBitString struct {
	value    []*asn1.BitString
	expected []byte
}

func Test_TestUnconstrainedIntEncode(t *testing.T) {
	testCases := []testint{
		{
			[]int{0, -1}, []byte{0x01, 0x00, 0x01, 0xff},
		},
		{
			[]int{1, 200}, []byte{0x01, 0x01, 0x02, 0x00, 0xc8},
		},
		{
			[]int{123456789, 1234567890987654321}, []byte{0x04, 0x07, 0x5b, 0xcd, 0x15, 0x08, 0x11, 0x22, 0x10, 0xf4, 0xb1, 0x6c, 0x1c, 0xb1},
		},
		{
			[]int{-123456789, -1234567890987654321}, []byte{0x04, 0xf8, 0xa4, 0x32, 0xeb, 0x08, 0xee, 0xdd, 0xef, 0x0b, 0x4e, 0x93, 0xe3, 0x4f},
		},
	}

	for _, tc := range testCases {
		a := int32(tc.values[0])
		b := int64(tc.values[1])
		test1 := &TestUnconstrainedInt{
			AttrUciA: a,
			AttrUciB: b,
		}

		aper, err := aper.Marshal(test1)
		assert.NoError(t, err)
		assert.NotNil(t, aper)
		t.Logf("%d %d gives APER %s", a, b, hex.Dump(aper))
		assert.EqualValues(t, tc.expected, aper)
		lenA := aper[0]
		lenB := aper[lenA+1]
		// Check that the lengths stated reflect the true length
		assert.Equal(t, int(1+lenA+1+lenB), len(aper))
	}

}

func Test_TestConstrainedIntEncode(t *testing.T) {
	testCases := []testint{
		{
			[]int{10, 10, 10, 10, 10},
			[]byte{
				0x00,       // A value - no len, value = 0 (over 10) shifted << by 1 ?
				0x01, 0x00, // B value - len = 1, value = 0 (over 10)
				0x01, 0x0a, // C value - len = 1, value = 10
				0x00, // D values - length = 0 and no values because they equal lower bound
				// E value missing because it always has to be 10
			},
		},
		{
			[]int{20, 20, 20, 20, 10},
			[]byte{
				0x14,      // A value - no len, value = 10 (over 10) shifted << by 1 ?
				0x01, 0xa, // B value - len = 1, value = 10 (over 10)
				0x01, 0x14, // C value - len = 1, value = 20
				0xa0, // D value - no len, no value - equals to upper bound
				// E value missing because it always has to be 10
			},
		},
		{
			[]int{30, 30, 30, 15, 10},
			[]byte{
				0x28,       // A value - no len, value = 20 (over 10) shifted << by 1 ?
				0x01, 0x14, // B value - len = 1, value = 20 (over 10)
				0x01, 0x1e, // C value - len = 1, value = 30
				0x50, // D value - value = 5 (over 10) shifted << by 4 ?
				// E value missing because it always has to be 10
			},
		},
		{
			[]int{100, 100, 100, 20, 10},
			[]byte{
				0xb4,       // A value - no len, value = 90 (over 10) shifted << by 1 ?
				0x01, 0x5a, // B value - len = 1, value = 90 (over 10)
				0x01, 0x64, // C value - len = 1, value = 100
				0xa0, // D value - no len, no value - equals to upper bound
				// E value missing because it always has to be 10
			},
		},
	}

	for _, tc := range testCases {
		a := int32(tc.values[0])
		b := int64(tc.values[1])
		c := int64(tc.values[2])
		d := int64(tc.values[3])
		e := int64(tc.values[4])
		test1 := &TestConstrainedInt{
			AttrCiA: a,
			AttrCiB: b,
			AttrCiC: c,
			AttrCiD: d,
			AttrCiE: e,
		}

		aper, err := aper.Marshal(test1)
		assert.NoError(t, err)
		assert.NotNil(t, aper)
		t.Logf("%d %d %d %d %d gives APER %s", a, b, c, d, e, hex.Dump(aper))
		assert.EqualValues(t, tc.expected, aper)
	}
}

func Test_BitString(t *testing.T) {
	testCases := []testBitString{
		{
			[]*asn1.BitString{
				{Value: []byte{0xff, 0xee, 0xdf}, Len: 22},
				{Value: []byte{0xff, 0xee, 0xdf}, Len: 20},
				{Value: []byte{0xff, 0xee, 0xdf}, Len: 23},
				{Value: []byte{0xff, 0xee, 0xdd, 0xcf}, Len: 28},
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
				{Value: []byte{0xff, 0xee, 0xdd, 0xcf}, Len: 20},
				{Value: []byte{0xff, 0xee, 0xdd, 0xcf}, Len: 32},
				{Value: []byte{0xff, 0xee, 0xdd, 0xcf}, Len: 32},
				{Value: []byte{0xff, 0xee, 0xdd, 0xcc, 0xbb}, Len: 36},
			},
			[]byte{
				0x80,                   // flag to show optional 5th item is present - there's only 1 optional field so it's first bit
				0x20,                   // length = 32
				0xff, 0xee, 0xdd, 0xcf, // no bit zeroed - fits byte boundary
				// no length as it's fixed
				0xff, 0xee, 0xd0, 0xcf, // last 12 bits are zeroed, so dd cf becomes d0 -- we should not have given 4th byte
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
		aper, err := aper.Marshal(tbs)
		assert.NoError(t, err)
		assert.NotNil(t, aper)
		t.Logf("%v %v %v gives APER %s", attrBs1, attrBs2, attrBs3, hex.Dump(aper))
		assert.EqualValues(t, tc.expected, aper)
	}

}
