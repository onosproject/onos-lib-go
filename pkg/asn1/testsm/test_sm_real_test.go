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
		{
			[]float64{12345.6789, 98765.4321},
			[]byte{0x09, 0x80, 0xd9, 0x18, 0x1c, 0xd6, 0xe6, 0x31, 0xf8, 0xa1, 0x09, 0x80, 0xdd, 0x0c, 0x0e, 0x6b, 0x74, 0xf0, 0xf8, 0x45},
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
		assert.Equal(t, per, tc.expected)
	}
}
