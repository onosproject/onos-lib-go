// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

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
