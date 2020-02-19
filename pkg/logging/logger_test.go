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
	zc "go.uber.org/zap/zapcore"
)

func TestHierarchicalLogger(t *testing.T) {
	fooLogger := GetLogger("foo")
	assert.NotNil(t, fooLogger.stdLogger)

	// Inherits from foo logger (i.e. warn level)
	fooBarLogger := GetLogger("foo", "bar")
	assert.NotNil(t, fooBarLogger.stdLogger)

	fooBarBazLogger := GetLogger("foo", "bar", "baz")
	assert.NotNil(t, fooBarBazLogger.stdLogger)

	fooBarBadLogger := GetLogger("foo", "bar", "bad")
	assert.NotNil(t, fooBarBadLogger.stdLogger)

	cfg := Configuration{}

	cfg.SetEncoding("json").
		SetLevel(WarnLevel).
		SetOutputPaths([]string{"stdout"}).
		SetName("config", "foo", "bar").
		SetErrorOutputPaths([]string{"stderr"}).
		SetECMsgKey("Msg").
		SetECLevelKey("Level").
		SetECTimeKey("Ts").
		SetECTimeEncoder(zc.ISO8601TimeEncoder).
		SetECEncodeLevel(zc.CapitalLevelEncoder).
		Build()
	cfg.AddLogger()
	loggerWithCustomConfig := GetLogger("config", "foo", "bar")
	assert.NotNil(t, loggerWithCustomConfig.stdLogger)

}
