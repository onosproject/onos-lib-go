// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package testsm

import (
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
			[]byte{0x01, 0x01},
		},
		{
			[]float64{12345.6789, 98765.4321},
			[]byte{0x01, 0x01},
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
		aper, err := aper.Marshal(test1, Choicemap, CanonicalChoicemap)
		assert.Nil(t, aper)
		assert.EqualError(t, err, "unsupported: Type:float64 Kind:float64")
	}
}
