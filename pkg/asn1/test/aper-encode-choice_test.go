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

package test

import (
	"encoding/hex"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_encodeChoice1(t *testing.T) {
	testChoice1 := &TestChoices{
		OtherAttr: "choice1only",
		Choice1: &TestChoices_Choice1A{
			Choice1A: 10,
		},
		Choice2: &TestChoices_Choice2B{
			Choice2B: "test2",
		},
		Choice3: &TestChoices_Choice3B{
			Choice3B: "test3",
		},
	}
	tcExpected := []byte{
		0x0b,                                                             // The length of the following text == 11 - not constrained so uses all 8 bits
		0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x31, 0x6f, 0x6e, 0x6c, 0x79, // the text "choice1only"
		0x00,                         // Choice 1a (1 of 1)
		0x01,                         // Value not extended and length of integer - [0]000 0001 - not constrained so uses all 8 bits
		0x0a,                         // Value of Choice 1a int = 10
		0x40,                         // Choice 2b (2 of 2)
		0x05,                         // Value not extended and length of integer - [0]000 0001 - not constrained so uses all 8 bits
		0x74, 0x65, 0x73, 0x74, 0x32, // Value of Choice 2b string = test2
		0x40,                         // Choice 3b (2 of 3)
		0x05,                         // Value not extended and length of integer - [0]000 0001 - not constrained so uses all 8 bits
		0x74, 0x65, 0x73, 0x74, 0x33, // Value of Choice 3b string = test3
	}

	aper, err := aper.Marshal(testChoice1, Choicemap, nil)
	assert.NoError(t, err)
	assert.NotNil(t, aper)
	t.Logf("Choice 1 APER %s", hex.Dump(aper))
	assert.EqualValues(t, tcExpected, aper)
}
