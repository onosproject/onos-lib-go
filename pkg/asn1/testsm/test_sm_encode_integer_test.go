// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package testsm

import (
	"encoding/hex"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log := logging.GetLogger("asn1")
	log.SetLevel(logging.DebugLevel)
	//aper.ChoiceMap = Choicemap // from choiceOptions.go - generated with protoc-gen-choice
	os.Exit(m.Run())
}

type testint struct {
	values   []int
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
			[]int{123456789, 1987654321}, []byte{0x04, 0x07, 0x5b, 0xcd, 0x15, 0x04, 0x76, 0x79, 0x32, 0xb1},
		},
		{
			[]int{-123456789, -1987654321}, []byte{0x04, 0xf8, 0xa4, 0x32, 0xeb, 0x04, 0x89, 0x86, 0xcd, 0x4f},
		},
	}

	for _, tc := range testCases {
		a := int32(tc.values[0])
		b := int32(tc.values[1])
		test1 := &TestUnconstrainedInt{
			AttrUciA: a,
			AttrUciB: b,
		}

		aper, err := aper.Marshal(test1, Choicemap, CanonicalChoicemap)
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
			[]int{10, 255, 10, 10, 10, 10, 10},
			[]byte{
				0x00,       // A value - no len, value = 0 (over 10) shifted << by 1 ?
				0x00, 0x00, // B value - no len, 16 bits need 2 bytes to be encoded, value = 0 (over 255)
				// ToDo - after introducing an upperbound of 4294967295 (maximum range of long in Clang, it started to encode length as 0. Is that correct? Yes, because length is now a known value (we know range)
				0x00, 0x00, // B value - len = 1, value = 0 (over 10)
				0x01, 0x0a, // C value - len = 1, value = 10
				0x00, // D values not octet aligned since under 16. no length. Values = 0 since equals lower bound
				// E value missing because it always has to be 10
				// F value missing because it always has to be 10 if not extended
			},
		},
		{
			[]int{20, 255, 20, 20, 20, 10, 10},
			[]byte{
				0x14,       // A value - no len, value = 10 (over 10) shifted << by 1 ?
				0x00, 0x00, // B value - no len, 16 bits need 2 bytes to be encoded, value = 0 (over 255)
				// ToDo - after introducing an upperbound of 4294967295 (maximum range of long in Clang, it started to encode length as 0. Is that correct? Yes, because length is now a known value (we know range)
				0x00, 0xa, // B value - len = 1, value = 10 (over 10)
				0x01, 0x14, // C value - len = 1, value = 20
				0xa0, // D values not octet aligned since under 16. no length. Values = 0 since equals lower bound
				// E value missing because it always has to be 10
				// F value missing because it always has to be 10 if not extended
			},
		},
		{
			[]int{30, 255, 30, 30, 15, 10, 10},
			[]byte{
				0x28,       // A value - no len, value = 20 (over 10) shifted << by 1 ?
				0x00, 0x00, // B value - no len, 16 bits need 2 bytes to be encoded, value = 0 (over 255)
				// ToDo - after introducing an upperbound of 4294967295 (maximum range of long in Clang, it started to encode length as 0. Is that correct? Yes, because length is now a known value (we know range)
				0x00, 0x14, // B value - len = 1, value = 20 (over 10)
				0x01, 0x1e, // C value - len = 1, value = 30
				0x50, // 0101 0000 D value = 5 (over 10) - then 0 since F is not extended
				// E value missing because it always has to be 10
				// F value missing because it always has to be 10 if not extended
			},
		},
		{
			[]int{100, 65534, 100, 100, 20, 10, 20},
			[]byte{
				0xb4,       // A value - no len, value = 90 (over 10) shifted << by 1 ?
				0xfe, 0xff, // B value - no len (have bound in definition), value = 65279 (over 255)
				// ToDo - after introducing an upperbound of 4294967295 (maximum range of long in Clang, it started to encode length as 0. Is that correct? Yes, because length is now a known value (we know range)
				0x00, 0x5a, // C value - len = 1, value = 90 (over 10)
				0x01, 0x64, // D value - len = 1, value = 100
				0xa8, // 1010 1000
				// E = 4 bits 1010 = a = 10 (over 10)
				// F = will never have any value
				// G 1 bit to say it's extended
				0x01, // number of bytes length of G since it is extended
				0x14, // The value byte of G = 20
			},
		},
	}

	for _, tc := range testCases {
		a := int32(tc.values[0])
		b := int32(tc.values[1])
		c := int32(tc.values[2])
		d := int32(tc.values[3])
		e := int32(tc.values[4])
		f := int32(tc.values[5])
		g := int32(tc.values[6])
		test1 := &TestConstrainedInt{
			AttrCiA: a,
			AttrCiB: b,
			AttrCiC: c,
			AttrCiD: d,
			AttrCiE: e,
			AttrCiF: f,
			AttrCiG: g,
		}

		aper, err := aper.Marshal(test1, Choicemap, CanonicalChoicemap)
		assert.NoError(t, err)
		assert.NotNil(t, aper)
		t.Logf("%d %d %d %d %d %d gives APER %s", a, b, c, d, e, f, hex.Dump(aper))
		assert.EqualValues(t, tc.expected, aper)
	}
}
