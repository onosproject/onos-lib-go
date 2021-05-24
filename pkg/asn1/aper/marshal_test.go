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

package aper

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_appendEnumerated(t *testing.T) {
	pd := perRawBitData{}
	lowerBound := int64(1)
	upperBound := int64(2)
	err := pd.appendEnumerated(1, true, &lowerBound, &upperBound)
	assert.NoError(t, err)
}

func Test_appendOpenType(t *testing.T) {
	pd := &perRawBitData{}
	v := reflect.ValueOf("abc")
	params := fieldParameters{
		sizeExtensible: true,
	}
	err := pd.appendOpenType(v, params)
	assert.NoError(t, err)
}
