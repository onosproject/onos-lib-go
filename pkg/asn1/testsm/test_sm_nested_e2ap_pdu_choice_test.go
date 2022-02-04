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

func createSampleNestedE2ApPduChoice(i int) *SampleNestedE2ApPduChoice {
	msg := &SampleNestedE2ApPduChoice{
		Criticality: 1,
	}
	switch i {
	case 1:
		msg.Id = int32(CanonicalNestedChoiceIDSampleOctetString)
		msg.Ch = &CanonicalNestedChoice{
			CanonicalNestedChoice: &CanonicalNestedChoice_Ch1{
				Ch1: &SampleOctetString{
					Value: []byte{0x23, 0x64, 0x81, 0x37, 0xFF, 0x4A, 0xD5, 0x7B, 0xDE, 0xC7},
				},
			},
		}
	case 2:
		msg.Id = int32(CanonicalNestedChoiceIDSampleConstrainedInteger)
		msg.Ch = &CanonicalNestedChoice{
			CanonicalNestedChoice: &CanonicalNestedChoice_Ch2{
				Ch2: &SampleConstrainedInteger{
					Value: 255,
				},
			},
		}
	case 3:
		msg.Id = int32(CanonicalNestedChoiceIDSampleBitString)
		msg.Ch = &CanonicalNestedChoice{
			CanonicalNestedChoice: &CanonicalNestedChoice_Ch3{
				Ch3: &SampleBitString{
					Value: &asn1.BitString{
						Value: []byte{0x23, 0x64, 0x81, 0xFC},
						Len:   30,
					},
				},
			},
		}
	case 4:
		msg.Id = int32(CanonicalNestedChoiceIDTestListExtensible1)
		listExt := &TestListExtensible1{
			Value: make([]*Item, 0),
		}
		item1 := &Item{
			Item2: &asn1.BitString{
				Value: []byte{0xDE},
				Len:   7,
			},
		}
		var ie21 int32 = -56
		item2 := &Item{
			Item1: &ie21,
			Item2: &asn1.BitString{
				Value: []byte{0xAE},
				Len:   7,
			},
		}
		listExt.Value = append(listExt.Value, item1)
		listExt.Value = append(listExt.Value, item2)

		msg.Ch = &CanonicalNestedChoice{
			CanonicalNestedChoice: &CanonicalNestedChoice_Ch4{
				Ch4: listExt,
			},
		}
	default:
		msg.Id = int32(CanonicalNestedChoiceIDSampleOctetString)
		msg.Ch = &CanonicalNestedChoice{
			CanonicalNestedChoice: &CanonicalNestedChoice_Ch1{
				Ch1: &SampleOctetString{
					Value: []byte{0x23, 0x64, 0x81, 0x37},
				},
			},
		}
	}

	return msg
}

func Test_CanonicalNestedChoice(t *testing.T) {

	// Satisfying a ChoiceMap constraint
	//aper.ChoiceMap = Choicemap
	//aper.CanonicalChoiceMap = CanonicalChoicemap

	for i := 1; i <= 5; i++ {

		msg := createSampleNestedE2ApPduChoice(i)

		aperBytes, err := aper.Marshal(msg, Choicemap, CanonicalChoicemap)
		assert.NilError(t, err)
		assert.Assert(t, aperBytes != nil)
		t.Logf("APER \n%s", hex.Dump(aperBytes))

		// Now decode the bytes and compare messages
		result := &SampleNestedE2ApPduChoice{}
		err = aper.Unmarshal(aperBytes, result, Choicemap, CanonicalChoicemap)
		assert.NilError(t, err)
		assert.Assert(t, result != nil)
		assert.Equal(t, msg.String(), result.String())
		//t.Logf("Decoded message is\n%v", result)
	}
}

func Test_CanonicalNestedChoiceIncorrectMapping(t *testing.T) {

	// Satisfying a ChoiceMap constraint
	//aper.ChoiceMap = Choicemap
	//aper.CanonicalChoiceMap = CanonicalChoicemap

	msg1 := &SampleNestedE2ApPduChoice{
		Id:          12,
		Criticality: 1,
		Ch: &CanonicalNestedChoice{
			CanonicalNestedChoice: &CanonicalNestedChoice_Ch1{
				Ch1: &SampleOctetString{
					Value: []byte{0x23, 0x64, 0x81, 0x37},
				},
			},
		},
	}

	_, err := aper.Marshal(msg1, Choicemap, CanonicalChoicemap)
	assert.ErrorContains(t, err, "Expected to have key (12) in CanonicalChoiceMap\nmap[11:testsm.CanonicalNestedChoice_Ch1 21:testsm.CanonicalNestedChoice_Ch2 31:testsm.CanonicalNestedChoice_Ch3 41:testsm.CanonicalNestedChoice_Ch4]")

	msg2 := &SampleNestedE2ApPduChoice{
		Id:          21,
		Criticality: 1,
		Ch: &CanonicalNestedChoice{
			CanonicalNestedChoice: &CanonicalNestedChoice_Ch1{
				Ch1: &SampleOctetString{
					Value: []byte{0x23, 0x64, 0x81, 0x37},
				},
			},
		},
	}

	_, err = aper.Marshal(msg2, Choicemap, CanonicalChoicemap)
	assert.ErrorContains(t, err, "UNIQUE ID (21) doesn't correspond to it's choice option (CanonicalNestedChoice_Ch2), got CanonicalNestedChoice_Ch1")
}
