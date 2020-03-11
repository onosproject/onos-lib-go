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

	art "github.com/plar/go-adaptive-radix-tree"

	"github.com/onosproject/onos-lib-go/pkg/logging/config"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	zc "go.uber.org/zap/zapcore"
)

func TestPreConfiguredLogger(t *testing.T) {
	loggersConfig, err := config.GetConfig()
	assert.Nil(t, err)
	AddConfiguredLoggers(loggersConfig)

	for _, configuredLogger := range loggersConfig.Logging.Loggers {
		loggers.ForEachPrefix(art.Key(configuredLogger.Name), func(node art.Node) bool {
			return assert.NotNil(t, node.Key())
		})
	}
}

func TestCustomLogger(t *testing.T) {
	cfgFooLogger := Configuration{}
	cfgFooLogger.SetEncoding("json").
		SetLevel(ErrorLevel).
		SetOutputPaths([]string{"stdout"}).
		SetName("foo").
		SetErrorOutputPaths([]string{"stderr"}).
		SetECMsgKey("msg").
		SetECLevelKey("level").
		SetECTimeKey("ts").
		SetECTimeEncoder(zc.ISO8601TimeEncoder).
		SetECEncodeLevel(zc.CapitalLevelEncoder).
		Build()

	cfgFooBarLogger := Configuration{}

	cfgFooBarLogger.SetEncoding("json").
		SetLevel(WarnLevel).
		SetOutputPaths([]string{"stdout"}).
		SetName("foo", "bar").
		SetErrorOutputPaths([]string{"stderr"}).
		SetECMsgKey("msg").
		SetECLevelKey("level").
		SetECTimeKey("ts").
		SetECTimeEncoder(zc.ISO8601TimeEncoder).
		SetECEncodeLevel(zc.CapitalLevelEncoder).
		Build()

	fooLogger := cfgFooLogger.GetLogger()
	fooBarLogger := cfgFooBarLogger.GetLogger()

	assert.NotNil(t, fooLogger.stdLogger)
	assert.Equal(t, fooLogger.stdLogger.Core().Enabled(zap.DebugLevel), false)
	assert.Equal(t, fooLogger.stdLogger.Core().Enabled(zap.InfoLevel), false)
	assert.Equal(t, fooLogger.stdLogger.Core().Enabled(zap.WarnLevel), false)
	assert.Equal(t, fooLogger.stdLogger.Core().Enabled(zap.ErrorLevel), true)

	assert.NotNil(t, fooBarLogger.stdLogger)
	assert.Equal(t, fooBarLogger.stdLogger.Core().Enabled(zap.DebugLevel), false)
	assert.Equal(t, fooBarLogger.stdLogger.Core().Enabled(zap.InfoLevel), false)
	assert.Equal(t, fooBarLogger.stdLogger.Core().Enabled(zap.WarnLevel), true)
	assert.Equal(t, fooBarLogger.stdLogger.Core().Enabled(zap.ErrorLevel), true)

	fooLogger.SetLevel(InfoLevel)
	assert.Equal(t, fooLogger.stdLogger.Core().Enabled(zap.DebugLevel), false)
	assert.Equal(t, fooLogger.stdLogger.Core().Enabled(zap.InfoLevel), true)
	assert.Equal(t, fooLogger.stdLogger.Core().Enabled(zap.WarnLevel), true)
	assert.Equal(t, fooLogger.stdLogger.Core().Enabled(zap.ErrorLevel), true)

	assert.Equal(t, fooBarLogger.stdLogger.Core().Enabled(zap.DebugLevel), false)
	assert.Equal(t, fooBarLogger.stdLogger.Core().Enabled(zap.InfoLevel), true)
	assert.Equal(t, fooBarLogger.stdLogger.Core().Enabled(zap.WarnLevel), true)
	assert.Equal(t, fooBarLogger.stdLogger.Core().Enabled(zap.ErrorLevel), true)

}

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

	assert.Equal(t, fooLogger.stdLogger.Core().Enabled(zap.DebugLevel), false)
	assert.Equal(t, fooLogger.stdLogger.Core().Enabled(zap.InfoLevel), true)

	fooLogger.SetLevel(DebugLevel)
	assert.Equal(t, fooLogger.stdLogger.Core().Enabled(zap.DebugLevel), true)
	assert.Equal(t, fooLogger.stdLogger.Core().Enabled(zap.InfoLevel), true)

	assert.Equal(t, fooBarBazLogger.stdLogger.Core().Enabled(zap.DebugLevel), true)
	assert.Equal(t, fooBarBazLogger.stdLogger.Core().Enabled(zap.InfoLevel), true)

	assert.Equal(t, fooBarBadLogger.stdLogger.Core().Enabled(zap.DebugLevel), true)
	assert.Equal(t, fooBarBazLogger.stdLogger.Core().Enabled(zap.InfoLevel), true)

	fooBarLogger.SetLevel(InfoLevel)
	assert.Equal(t, fooLogger.stdLogger.Core().Enabled(zap.DebugLevel), true)
	assert.Equal(t, fooLogger.stdLogger.Core().Enabled(zap.InfoLevel), true)

	assert.Equal(t, fooBarBazLogger.stdLogger.Core().Enabled(zap.InfoLevel), true)
	assert.Equal(t, fooBarBazLogger.stdLogger.Core().Enabled(zap.DebugLevel), false)

	assert.Equal(t, fooBarBazLogger.stdLogger.Core().Enabled(zap.InfoLevel), true)
	assert.Equal(t, fooBarBadLogger.stdLogger.Core().Enabled(zap.DebugLevel), false)

	fooLogger.SetLevel(ErrorLevel)
	assert.Equal(t, fooLogger.stdLogger.Core().Enabled(zap.DebugLevel), false)
	assert.Equal(t, fooLogger.stdLogger.Core().Enabled(zap.InfoLevel), false)
	assert.Equal(t, fooLogger.stdLogger.Core().Enabled(zap.WarnLevel), false)
	assert.Equal(t, fooLogger.stdLogger.Core().Enabled(zap.ErrorLevel), true)

	assert.Equal(t, fooBarLogger.stdLogger.Core().Enabled(zap.DebugLevel), false)
	assert.Equal(t, fooBarLogger.stdLogger.Core().Enabled(zap.InfoLevel), false)
	assert.Equal(t, fooBarLogger.stdLogger.Core().Enabled(zap.WarnLevel), false)
	assert.Equal(t, fooBarLogger.stdLogger.Core().Enabled(zap.ErrorLevel), true)

	assert.Equal(t, fooBarBazLogger.stdLogger.Core().Enabled(zap.DebugLevel), false)
	assert.Equal(t, fooBarBazLogger.stdLogger.Core().Enabled(zap.InfoLevel), false)
	assert.Equal(t, fooBarBazLogger.stdLogger.Core().Enabled(zap.WarnLevel), false)
	assert.Equal(t, fooBarBazLogger.stdLogger.Core().Enabled(zap.ErrorLevel), true)

	assert.Equal(t, fooBarBadLogger.stdLogger.Core().Enabled(zap.DebugLevel), false)
	assert.Equal(t, fooBarBadLogger.stdLogger.Core().Enabled(zap.InfoLevel), false)
	assert.Equal(t, fooBarBadLogger.stdLogger.Core().Enabled(zap.WarnLevel), false)
	assert.Equal(t, fooBarBadLogger.stdLogger.Core().Enabled(zap.ErrorLevel), true)

}
