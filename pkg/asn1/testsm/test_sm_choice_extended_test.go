// Copyright 2022-present Open Networking Foundation.
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
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"gotest.tools/assert"
	"testing"
)

func Test_ChoiceExtended(t *testing.T) {

	// Satisfying a ChoiceMap constraint
	aper.ChoiceMap = Choicemap
	aper.CanonicalChoiceMap = CanonicalChoicemap

	msg1 := &ChoiceExtended{
		ChoiceExtended: &ChoiceExtended_ChoiceExtendedC{
			ChoiceExtendedC: 1,
		},
	}

	aperBytes1, err := aper.MarshalWithParams(msg1, "choiceExt")
	assert.NilError(t, err)
	assert.Assert(t, aperBytes1 != nil)
	t.Logf("APER \n%s", hex.Dump(aperBytes1))

	//Now decode the bytes and compare messages
	result1 := &ChoiceExtended{}
	err1 := aper.UnmarshalWithParams(aperBytes1, result1, "choiceExt")
	assert.NilError(t, err1)
	assert.Assert(t, result1 != nil)
	assert.Equal(t, msg1.String(), result1.String())
	t.Logf("Decoded message is\n%v", result1)

	msg2 := &ChoiceExtended{
		ChoiceExtended: &ChoiceExtended_ChoiceExtendedD{
			ChoiceExtendedD: 1,
		},
	}

	aperBytes2, err := aper.Marshal(msg2)
	assert.NilError(t, err)
	assert.Assert(t, aperBytes2 != nil)
	t.Logf("APER \n%s", hex.Dump(aperBytes2))

	// Now decode the bytes and compare messages
	result2 := &ChoiceExtended{}
	err2 := aper.UnmarshalWithParams(aperBytes2, result2, "choiceExt")
	assert.NilError(t, err2)
	assert.Assert(t, result2 != nil)
	assert.Equal(t, msg2.String(), result2.String())
	t.Logf("Decoded message is\n%v", result2)

	msg3 := &ChoiceExtended{
		ChoiceExtended: &ChoiceExtended_ChoiceExtendedD{
			ChoiceExtendedD: 3,
		},
	}

	aperBytes3, err := aper.Marshal(msg3)
	assert.NilError(t, err)
	assert.Assert(t, aperBytes3 != nil)
	t.Logf("APER \n%s", hex.Dump(aperBytes3))

	// Now decode the bytes and compare messages
	result3 := &ChoiceExtended{}
	err3 := aper.UnmarshalWithParams(aperBytes3, result3, "choiceExt")
	assert.NilError(t, err3)
	assert.Assert(t, result3 != nil)
	assert.Equal(t, msg3.String(), result3.String())
	t.Logf("Decoded message is\n%v", result3)

	msg4 := &ChoiceExtended{
		ChoiceExtended: &ChoiceExtended_ChoiceExtendedB{
			ChoiceExtendedB: 16,
		},
	}

	aperBytes4, err := aper.MarshalWithParams(msg4, "choiceExt")
	assert.NilError(t, err)
	assert.Assert(t, aperBytes4 != nil)
	t.Logf("APER \n%s", hex.Dump(aperBytes4))

	// Now decode the bytes and compare messages
	result4 := &ChoiceExtended{}
	err4 := aper.UnmarshalWithParams(aperBytes4, result4, "choiceExt")
	assert.NilError(t, err4)
	assert.Assert(t, result4 != nil)
	assert.Equal(t, msg4.String(), result4.String())
	t.Logf("Decoded message is\n%v", result4)
}
