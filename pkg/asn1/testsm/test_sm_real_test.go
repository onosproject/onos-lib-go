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
