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

package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHierarchicalLogger(t *testing.T) {
	AddLogger("warn", "foo")
	AddLogger("", "foo", "bar")
	AddLogger("error", "foo", "bar", "baz")

	fooLogger, found := GetLogger("foo")
	assert.Equal(t, found, true)
	assert.NotNil(t, fooLogger.stdLogger)

	// Inherits from foo logger (i.e. warn level)
	fooBarLogger, found := GetLogger("foo", "bar")
	assert.Equal(t, found, true)
	assert.NotNil(t, fooBarLogger.stdLogger)

	fooBarBazLogger, found := GetLogger("foo", "bar", "baz")
	assert.Equal(t, found, true)
	assert.NotNil(t, fooBarBazLogger.stdLogger)

	fooBarBadLogger, found := GetLogger("foo", "bar", "bad")
	assert.Equal(t, found, false)
	assert.Nil(t, fooBarBadLogger.stdLogger)
}
