// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package testsm

import (
	"encoding/hex"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"gotest.tools/assert"
	"testing"
)

func Test_MixedChoice(t *testing.T) {

	// Satisfying a ChoiceMap constraint
	//aper.ChoiceMap = Choicemap
	//aper.CanonicalChoiceMap = CanonicalChoicemap

	msg1 := &MixedChoice{
		MixedChoice: &MixedChoice_Ch1{
			Ch1: &SampleE2ApPduChoice{
				Id:          int32(CanonicalChoiceIDSampleNestedE2apPduChoice),
				Criticality: 0,
				Ch: &CanonicalChoice{
					CanonicalChoice: &CanonicalChoice_Ch6{
						Ch6: createSampleNestedE2ApPduChoice(1),
					},
				},
			},
		},
	}

	aperBytes1, err := aper.Marshal(msg1, Choicemap, CanonicalChoicemap)
	assert.NilError(t, err)
	assert.Assert(t, aperBytes1 != nil)
	t.Logf("APER \n%s", hex.Dump(aperBytes1))

	// Now decode the bytes and compare messages
	result1 := &MixedChoice{}
	err = aper.Unmarshal(aperBytes1, result1, Choicemap, CanonicalChoicemap)
	assert.NilError(t, err)
	assert.Assert(t, result1 != nil)
	assert.Equal(t, msg1.String(), result1.String())

	msg2 := &MixedChoice{
		MixedChoice: &MixedChoice_Ch1{
			Ch1: &SampleE2ApPduChoice{
				Id:          int32(CanonicalChoiceIDSampleNestedE2apPduChoice),
				Criticality: 0,
				Ch: &CanonicalChoice{
					CanonicalChoice: &CanonicalChoice_Ch6{
						Ch6: createSampleNestedE2ApPduChoice(4),
					},
				},
			},
		},
	}

	aperBytes2, err := aper.Marshal(msg2, Choicemap, CanonicalChoicemap)
	assert.NilError(t, err)
	assert.Assert(t, aperBytes2 != nil)
	t.Logf("APER \n%s", hex.Dump(aperBytes2))

	// Now decode the bytes and compare messages
	result2 := &MixedChoice{}
	err = aper.Unmarshal(aperBytes2, result2, Choicemap, CanonicalChoicemap)
	assert.NilError(t, err)
	assert.Assert(t, result2 != nil)
	assert.Equal(t, msg2.String(), result2.String())

	msg3 := &MixedChoice{
		MixedChoice: &MixedChoice_Ch2{
			Ch2: &SampleConstrainedInteger{
				Value: 254,
			},
		},
	}

	aperBytes3, err := aper.Marshal(msg3, Choicemap, CanonicalChoicemap)
	assert.NilError(t, err)
	assert.Assert(t, aperBytes3 != nil)
	t.Logf("APER \n%s", hex.Dump(aperBytes3))

	// Now decode the bytes and compare messages
	result3 := &MixedChoice{}
	err = aper.Unmarshal(aperBytes3, result3, Choicemap, CanonicalChoicemap)
	assert.NilError(t, err)
	assert.Assert(t, result3 != nil)
	assert.Equal(t, msg3.String(), result3.String())
}
