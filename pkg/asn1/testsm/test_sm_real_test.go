// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package testsm

import (
	"encoding/hex"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testreal struct {
	values   []float64
	expected []byte
}

func Test_TestUnconstrainedRealEncode(t *testing.T) {
	testCases := []testreal{
		{
			[]float64{1.234, 1.234},
			[]byte{0x09, 0x80, 0xcf, 0x02, 0x77, 0xce, 0xd9, 0x16, 0x87, 0x2b, 0x09, 0x80, 0xcf, 0x02, 0x77, 0xce, 0xd9, 0x16, 0x87, 0x2b},
		},
		//{
		//	[]float64{12345.6789, 98765.4321},
		//	[]byte{0x09, 0x80, 0xd9, 0x18, 0x1c, 0xd6, 0xe6, 0x31, 0xf8, 0xa1, 0x09, 0x80, 0xdd, 0x0c, 0x0e, 0x6b, 0x74, 0xf0, 0xf8, 0x45},
		//},
		{
			[]float64{-2.5, -2.6},
			[]byte{0x03, 0xc0, 0xff, 0x05, 0x09, 0xc0, 0xcd, 0x14, 0xcc, 0xcc, 0xcc, 0xcc, 0xcc, 0xcd},
		},
		{
			[]float64{-65535.0, 65534.0},
			[]byte{0x05, 0xc0, 0x00, 0x00, 0xff, 0xff, 0x05, 0x80, 0x01, 0x00, 0x7f, 0xff},
		},
		{
			[]float64{2.1, -16.375},
			[]byte{0x09, 0x80, 0xcd, 0x10, 0xcc, 0xcc, 0xcc, 0xcc, 0xcc, 0xcd, 0x04, 0xc0, 0xfd, 0x00, 0x83},
		},
		{
			[]float64{-101.0, -100.0},
			[]byte{0x04, 0xc0, 0x00, 0x00, 0x65, 0x03, 0xc0, 0x02, 0x19},
		},
		{
			[]float64{64.0, -2.0},
			[]byte{0x03, 0x80, 0x06, 0x01, 0x03, 0xc0, 0x01, 0x01},
		},
		{
			[]float64{-3.0, 65.0},
			[]byte{0x03, 0xc0, 0x00, 0x03, 0x04, 0x80, 0x00, 0x00, 0x41},
		},
		{
			[]float64{10.0, -16777215.0},
			[]byte{0x03, 0x80, 0x01, 0x05, 0x06, 0xc0, 0x00, 0x00, 0xff, 0xff, 0xff},
		},
		{
			[]float64{21.7, -653.43},
			[]byte{0x09, 0x80, 0xd0, 0x15, 0xb3, 0x33, 0x33, 0x33, 0x33, 0x33, 0x09, 0xc0, 0xd5, 0x14, 0x6b, 0x70, 0xa3, 0xd7, 0x0a, 0x3d},
		},
	}

	for _, tc := range testCases {
		a := tc.values[0]
		b := tc.values[1]
		test1 := &TestUnconstrainedReal{
			AttrUcrA: a,
			AttrUcrB: b,
		}
		assert.NotNil(t, tc.expected)

		per, err := aper.Marshal(test1, Choicemap, CanonicalChoicemap)
		assert.Nil(t, err)
		t.Logf("APER bytes are\n%v", hex.Dump(per))
		assert.EqualValues(t, per, tc.expected)
	}

	for _, tc := range testCases {
		res := &TestUnconstrainedReal{}
		err := aper.Unmarshal(tc.expected, res, Choicemap, CanonicalChoicemap)
		assert.Nil(t, err)
		t.Logf("Decoded struct is\n%v", res)
		assert.EqualValues(t, tc.values[0], res.AttrUcrA)
		assert.EqualValues(t, tc.values[1], res.AttrUcrB)
	}

	valuesWithDistortion := []byte{0x09, 0x80, 0xd9, 0x18, 0x1c, 0xd6, 0xe6, 0x31, 0xf8, 0xa1, 0x09, 0x80, 0xdd, 0x0c, 0x0e, 0x6b, 0x74, 0xf0, 0xf8, 0x45}
	res1 := &TestUnconstrainedReal{}
	err := aper.Unmarshal(valuesWithDistortion, res1, Choicemap, CanonicalChoicemap)
	assert.Nil(t, err)
	t.Logf("Decoded struct is\n%v", res1)

}

func Test_RealErrors(t *testing.T) {
	test1 := &TestUnconstrainedReal{
		AttrUcrA: 0.0,
		AttrUcrB: 21.546,
	}

	_, err := aper.Marshal(test1, Choicemap, CanonicalChoicemap)
	assert.EqualError(t, err, "Error encoding REAL - numerical argument is out of domain")

	test2 := &TestConstrainedReal{
		AttrCrA: 21.4587,
		AttrCrB: 654654.651564,
		AttrCrC: -5464.98421,
		AttrCrD: 554421.0,
		AttrCrE: 10.0,
		AttrCrF: 11.354,
	}

	_, err = aper.Marshal(test2, Choicemap, CanonicalChoicemap)
	assert.EqualError(t, err, "Error encoding REAL - value (554421) is higher than upperbound (20)")

	test3 := &TestConstrainedReal{
		AttrCrA: 21.4587,
		AttrCrB: -1.275,
		AttrCrC: -5464.98421,
		AttrCrD: 554421.0,
		AttrCrE: 10.0,
		AttrCrF: 11.354,
	}

	_, err = aper.Marshal(test3, Choicemap, CanonicalChoicemap)
	assert.EqualError(t, err, "Error encoding REAL - value (-1.275) is lower than lowerbound (10)")
}
