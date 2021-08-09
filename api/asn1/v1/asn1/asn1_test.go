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

package asn1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_UpdateValue(t *testing.T) {
	bs := BitString{
		Value: []byte{0x3f, 0xff, 0xfd},
		Len:   22,
	}
	t.Logf("%x", bs.GetValue())
	newValue, err := bs.UpdateValue([]byte{0xfd, 0xee, 0x3f})
	assert.NoError(t, err)
	assert.EqualValues(t, []byte{0xfd, 0xee, 0x3f}, newValue)
	t.Logf("%x", newValue)
	assert.EqualValues(t, []byte{0xfd, 0xee, 0x3f}, bs.GetValue())
	t.Logf("%x", bs.GetValue())

	// Testing case of an empty value
	bs1 := BitString{
		//Value: []byte{0x3f, 0xff, 0xfd},
		Len: 31,
	}
	t.Logf("%x", bs1.GetValue())
	newValue1, err := bs1.UpdateValue([]byte{0xfd, 0xe4, 0xff, 0x1c})
	assert.NoError(t, err)
	assert.EqualValues(t, []byte{0xfd, 0xe4, 0xff, 0x1c}, newValue1)
	t.Logf("%x", newValue1)
	assert.EqualValues(t, []byte{0xfd, 0xe4, 0xff, 0x1c}, bs1.GetValue())
	t.Logf("%x", bs1.GetValue())
}
