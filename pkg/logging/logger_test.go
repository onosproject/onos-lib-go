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
	"testing"
)

func TestLoggerConfig(t *testing.T) {
	logger := GetLogger()
	assert.Equal(t, InfoLevel, logger.GetLevel())
	logger.Debug("should not be printed")
	logger.Info("should be printed")

	logger = GetLogger("test")
	assert.Equal(t, InfoLevel, logger.GetLevel())
	logger.Debug("should not be printed")
	logger.Info("should be printed")

	logger = GetLogger("test", "1")
	assert.Equal(t, DebugLevel, logger.GetLevel())
	logger.Debug("should be printed")
	logger.Info("should be printed")

	logger = GetLogger("test", "1", "2")
	assert.Equal(t, DebugLevel, logger.GetLevel())
	logger.Debug("should be printed")
	logger.Info("should be printed")

	logger = GetLogger("test")
	assert.Equal(t, InfoLevel, logger.GetLevel())
	logger.Debug("should not be printed")
	logger.Info("should be printed")

	logger = GetLogger("test", "2")
	assert.Equal(t, WarnLevel, logger.GetLevel())
	logger.Debug("should not be printed")
	logger.Info("should not be printed")
	logger.Warn("should be printed")

	logger = GetLogger("test", "2", "3")
	assert.Equal(t, WarnLevel, logger.GetLevel())
	logger.Debug("should not be printed")
	logger.Info("should not be printed")
	logger.Warn("should be printed")

	logger = GetLogger("test/2")
	logger.SetLevel(DebugLevel)
	assert.Equal(t, DebugLevel, logger.GetLevel())
	logger.Debug("should not be printed")
	logger.Info("should be printed")
	logger.Warn("should be printed")

	logger = GetLogger("test/2/3")
	assert.Equal(t, DebugLevel, logger.GetLevel())
	logger.Debug("should not be printed")
	logger.Info("should be printed")
	logger.Warn("should be printed")

	logger = GetLogger("test/3")
	assert.Equal(t, InfoLevel, logger.GetLevel())
	logger.Debug("should not be printed")
	logger.Info("should be printed")
	logger.Warn("should be printed twice")

	//logger = GetLogger("test", "kafka")
	//assert.Equal(t, InfoLevel, logger.GetLevel())
}
