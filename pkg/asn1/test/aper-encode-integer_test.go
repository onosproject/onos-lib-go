// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package test

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

		aper, err := aper.Marshal(test1, Choicemap, nil)
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

		aper, err := aper.Marshal(test1, Choicemap, nil)
		assert.NoError(t, err)
		assert.NotNil(t, aper)
		t.Logf("%d %d %d %d %d gives APER %s", a, b, c, d, e, hex.Dump(aper))
		assert.EqualValues(t, tc.expected, aper)
	}
}
