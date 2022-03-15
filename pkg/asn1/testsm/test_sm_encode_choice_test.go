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

func Test_encodeChoice1(t *testing.T) {
	testChoice1 := &TestChoices{
		OtherAttr: "choice1only",
		Choice1: &Choice1{
			Choice1: &Choice1_Choice1A{
				Choice1A: 10,
			},
		},
		Choice2: &Choice2{
			Choice2: &Choice2_Choice2B{
				Choice2B: 20,
			},
		},
		Choice3: &Choice3{
			Choice3: &Choice3_Choice3B{
				Choice3B: 30,
			},
		},
		Choice4: &Choice4{
			Choice4: &Choice4_Choice4A{
				Choice4A: 10,
			},
		},
	}
	tcExpected := []byte{
		0x0b,                                                             // The length of the following text == 11 - not constrained so uses all 8 bits
		0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x31, 0x6f, 0x6e, 0x6c, 0x79, // the text "choice1only"
		// no bits for choice1 as there is only 1 option - choice1_a
		0x01, // num bytes for integer - 0000 0001 - not constrained so uses all 8 bits
		0x0a, // Value of Choice 1a int = 10
		0x80, // Choice 2b (2 of 2) 1 bits for choice (1), 0 for extensible
		0x01, // num bytes for integer - 0000 0001 - not constrained so uses all 8 bits
		0x14, // Value of Choice 2b int = 20
		0x40, // Choice 3b (2 of 3) 2 bits for choice (01), 0 for extensible
		0x01, // num bytes for integer - 0000 0001 - not constrained so uses all 8 bits
		0x1e, // Value of Choice 3b int = 30
		0x00, // Choice 4a (1 of 1) 0 bits for choice, 0 for extensible on the oneof
		0x01, // num bytes for integer - 0000 0001 - not constrained so uses all 8 bits
		0x0a, // Value of Choice 4a int = 10
	}

	aper, err := aper.Marshal(testChoice1, Choicemap, CanonicalChoicemap)
	assert.NoError(t, err)
	assert.NotNil(t, aper)
	t.Logf("Choice 1 APER %s", hex.Dump(aper))
	assert.EqualValues(t, tcExpected, aper)
}
