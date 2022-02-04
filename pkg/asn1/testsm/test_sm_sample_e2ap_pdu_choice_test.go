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
	"encoding/hex"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"gotest.tools/assert"
	"testing"
)

func Test_CanonicalChoice(t *testing.T) {

	// Satisfying a ChoiceMap constraint
	//aper.ChoiceMap = Choicemap
	//aper.CanonicalChoiceMap = CanonicalChoicemap

	for i := 1; i <= 4; i++ {
		msg := &SampleE2ApPduChoice{
			Id:          int32(CanonicalChoiceIDSampleNestedE2apPduChoice),
			Criticality: 0,
			Ch: &CanonicalChoice{
				CanonicalChoice: &CanonicalChoice_Ch6{
					Ch6: createSampleNestedE2ApPduChoice(i),
				},
			},
		}

		aperBytes, err := aper.Marshal(msg, Choicemap, CanonicalChoicemap)
		assert.NilError(t, err)
		assert.Assert(t, aperBytes != nil)
		t.Logf("APER \n%s", hex.Dump(aperBytes))

		// Now decode the bytes and compare messages
		result := &SampleE2ApPduChoice{}
		err = aper.Unmarshal(aperBytes, result, Choicemap, CanonicalChoicemap)
		assert.NilError(t, err)
		assert.Assert(t, result != nil)
		assert.Equal(t, msg.String(), result.String())
		//t.Logf("Decoded message is\n%v", result)
	}

	var ie21 int32 = -56
	item2 := &Item{
		Item1: &ie21,
		Item2: &asn1.BitString{
			Value: []byte{0xFE},
			Len:   7,
		},
	}

	msg := &SampleE2ApPduChoice{
		Id:          int32(CanonicalChoiceIDItem),
		Criticality: 0,
		Ch: &CanonicalChoice{
			CanonicalChoice: &CanonicalChoice_Ch5{
				Ch5: item2,
			},
		},
	}

	aperBytes, err := aper.Marshal(msg, Choicemap, CanonicalChoicemap)
	assert.NilError(t, err)
	assert.Assert(t, aperBytes != nil)
	t.Logf("APER \n%s", hex.Dump(aperBytes))

	// Now decode the bytes and compare messages
	result := &SampleE2ApPduChoice{}
	err = aper.Unmarshal(aperBytes, result, Choicemap, CanonicalChoicemap)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	assert.Equal(t, msg.String(), result.String())
	//t.Logf("Decoded message is\n%v", result)
}
