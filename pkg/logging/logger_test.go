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
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)

func TestDefaultLogger(t *testing.T) {
	logger := GetLogger()
	assert.Equal(t, "github.com/onosproject/onos-lib-go/pkg/logging", logger.Name())
	logger.Info("foo: bar")
	logger.Info("bar: baz")
	logger.Debug("baz: bar")
}

func TestLoggerConfig(t *testing.T) {
	config := Config{}
	bytes, err := ioutil.ReadFile("test.yaml")
	assert.NoError(t, err)
	err = yaml.Unmarshal(bytes, &config)
	assert.NoError(t, err)
	err = configure(config)
	assert.NoError(t, err)

	// The root logger should be configured with INFO level
	logger := GetLogger()
	assert.Equal(t, InfoLevel, logger.GetLevel())
	logger.Debug("should not be printed")
	logger.Info("should be printed")

	// The "test" logger should inherit the INFO level from the root logger
	logger = GetLogger("test")
	assert.Equal(t, InfoLevel, logger.GetLevel())
	logger.Debug("should not be printed")
	logger.Info("should be printed")

	// The "test/1" logger should be configured with DEBUG level
	logger = GetLogger("test", "1")
	assert.Equal(t, DebugLevel, logger.GetLevel())
	logger.Debug("should be printed")
	logger.Info("should be printed")

	// The "test/1/2" logger should inherit the DEBUG level from "test/1"
	logger = GetLogger("test", "1", "2")
	assert.Equal(t, DebugLevel, logger.GetLevel())
	logger.Debug("should be printed")
	logger.Info("should be printed")

	// The "test" logger should still inherit the INFO level from the root logger
	logger = GetLogger("test")
	assert.Equal(t, InfoLevel, logger.GetLevel())
	logger.Debug("should not be printed")
	logger.Info("should be printed")

	// The "test/2" logger should be configured with WARN level
	logger = GetLogger("test", "2")
	assert.Equal(t, WarnLevel, logger.GetLevel())
	logger.Debug("should not be printed")
	logger.Info("should not be printed")
	logger.Warn("should be printed twice")

	// The "test/2/3" logger should be configured with INFO level
	logger = GetLogger("test", "2", "3")
	assert.Equal(t, InfoLevel, logger.GetLevel())
	logger.Debug("should not be printed")
	logger.Info("should be printed twice")
	logger.Warn("should be printed twice")

	// The "test/2/4" logger should inherit the WARN level from "test/2"
	logger = GetLogger("test", "2", "4")
	assert.Equal(t, WarnLevel, logger.GetLevel())
	logger.Debug("should not be printed")
	logger.Info("should not be printed")
	logger.Warn("should be printed twice")

	// The "test/2" logger level should be changed to DEBUG
	logger = GetLogger("test/2")
	logger.SetLevel(DebugLevel)
	assert.Equal(t, DebugLevel, logger.GetLevel())
	logger.Debug("should be printed")
	logger.Info("should be printed twice")
	logger.Warn("should be printed twice")

	// The "test/2/3" logger should not inherit the change to the "test/2" logger since its level has been explicitly set
	logger = GetLogger("test/2/3")
	assert.Equal(t, InfoLevel, logger.GetLevel())
	logger.Debug("should not be printed")
	logger.Info("should be printed twice")
	logger.Warn("should be printed twice")

	// The "test/2/4" logger should inherit the change to the "test/2" logger since its level has not been explicitly set
	// The "test/2/4" logger should not output DEBUG messages since the output level is explicitly set to WARN
	logger = GetLogger("test/2/4")
	assert.Equal(t, DebugLevel, logger.GetLevel())
	logger.Debug("should be printed")
	logger.Info("should be printed twice")
	logger.Warn("should be printed twice")

	// The "test/3" logger should be configured with INFO level
	// The "test/3" logger should write to multiple outputs
	logger = GetLogger("test/3")
	assert.Equal(t, InfoLevel, logger.GetLevel())
	logger.Debug("should not be printed")
	logger.Info("should be printed")
	logger.Warn("should be printed twice")

	// The "test/3/4" logger should inherit INFO level from "test/3"
	// The "test/3/4" logger should inherit multiple outputs from "test/3"
	logger = GetLogger("test/3/4")
	assert.Equal(t, InfoLevel, logger.GetLevel())
	logger.Debug("should not be printed")
	logger.Info("should be printed")
	logger.Warn("should be printed twice")

	//logger = GetLogger("test", "kafka")
	//assert.Equal(t, InfoLevel, logger.GetLevel())
}
