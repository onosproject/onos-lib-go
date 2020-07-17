// Copyright 2020-present Open Networking Foundation.
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

package rbac

import (
	"testing"

	"gotest.tools/assert"
)

func TestMatch(t *testing.T) {

	tests := []struct {
		matchFound bool
		s1         string
		s2         string
	}{
		{
			matchFound: false,
			s1:         "",
			s2:         "",
		},
		{
			matchFound: true,
			s1:         "*",
			s2:         "foobar",
		},
		{
			matchFound: true,
			s1:         "********",
			s2:         "foobar",
		},
		{
			matchFound: false,
			s1:         "foo",
			s2:         "bar",
		},
		{
			matchFound: false,
			s1:         "*foo",
			s2:         "barbaz",
		},
		{
			matchFound: false,
			s1:         "foo*",
			s2:         "barbaz",
		},
		{
			matchFound: false,
			s1:         "*foo*",
			s2:         "barbaz",
		},
		{
			matchFound: true,
			s1:         "*foo",
			s2:         "barbazfoo",
		},
		{
			matchFound: true,
			s1:         "bar*",
			s2:         "barfoobaz",
		},
		{
			matchFound: true,
			s1:         "*foo*",
			s2:         "barfoobaz",
		},
		{
			matchFound: false,
			s1:         "foo",
			s2:         "foobarbaz",
		},
	}

	for index, test := range tests {
		assert.Equal(t, test.matchFound, match(test.s1, test.s2), index)

	}
}
