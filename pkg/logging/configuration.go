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
	zp "go.uber.org/zap"
	zc "go.uber.org/zap/zapcore"
)

// Configuration log config structure
type Configuration struct {
	zapConfig zp.Config
}

// ConfigurationBuilder log configuration builder interface
type ConfigurationBuilder interface {
	SetLevel(Level) ConfigurationBuilder
	SetEncoding(string) ConfigurationBuilder
	SetDevelopment(bool) ConfigurationBuilder
	SetOutputPaths([]string) ConfigurationBuilder
	SetErrorOutputPaths([]string) ConfigurationBuilder
	//SetInitialFields(Fields) ConfigurationBuilder
	SetECLevelKey(string) ConfigurationBuilder
	SetECMsgKey(string) ConfigurationBuilder
	SetECTimeKey(string) ConfigurationBuilder
	SetECStackTraceKey(string) ConfigurationBuilder
	SetECLineEnding(string) ConfigurationBuilder
	SetECEncodeLevel(zc.LevelEncoder) ConfigurationBuilder
	SetECTimeEncoder(zc.TimeEncoder) ConfigurationBuilder
	SetName(...string) ConfigurationBuilder
	SetECEncodeDuration(zc.DurationEncoder) ConfigurationBuilder
	SetECEncoderCaller(encoder zc.CallerEncoder) ConfigurationBuilder
	Build() Configuration
}

// GetZapConfig gets zap configuration
func (c *Configuration) GetZapConfig() zp.Config {
	return c.zapConfig
}

// SetLevel sets log level
func (c *Configuration) SetLevel(level Level) ConfigurationBuilder {
	c.zapConfig.Level = intToAtomicLevel(level)
	return c
}

// SetEncoding sets log encoding
func (c *Configuration) SetEncoding(encoding string) ConfigurationBuilder {
	c.zapConfig.Encoding = encoding
	return c
}

// SetDevelopment sets log development
func (c *Configuration) SetDevelopment(development bool) ConfigurationBuilder {
	c.zapConfig.Development = development
	return c
}

// SetOutputPaths sets log output paths
func (c *Configuration) SetOutputPaths(outputPaths []string) ConfigurationBuilder {
	c.zapConfig.OutputPaths = outputPaths
	return c
}

// SetErrorOutputPaths sets log error output paths
func (c *Configuration) SetErrorOutputPaths(errorOutputPaths []string) ConfigurationBuilder {
	c.zapConfig.ErrorOutputPaths = errorOutputPaths
	return c
}

// SetInitialFields sets log initial fields
/*func (c *Configuration) SetInitialFields(initFields Fields) ConfigurationBuilder {
	c.zapConfig.InitialFields = initFields
	return c
}*/

// SetEncoderConfigLevelKey sets log encoder config level key
func (c *Configuration) SetECLevelKey(levelKey string) ConfigurationBuilder {
	c.zapConfig.EncoderConfig.LevelKey = levelKey
	return c
}

// SetECNameKey sets encoder config name key
func (c *Configuration) SetName(names ...string) ConfigurationBuilder {
	nameKey := buildTreeName(names...)
	c.zapConfig.EncoderConfig.NameKey = nameKey
	return c
}

// SetEncoderConfigMsgKey sets log encoder config message key
func (c *Configuration) SetECMsgKey(msgKey string) ConfigurationBuilder {
	c.zapConfig.EncoderConfig.MessageKey = msgKey
	return c
}

// SetECTimeKey sets log encoder config time key
func (c *Configuration) SetECTimeKey(timeKey string) ConfigurationBuilder {
	c.zapConfig.EncoderConfig.TimeKey = timeKey
	return c
}

// SetECStackTraceKey sets log encoder config start trace key
func (c *Configuration) SetECStackTraceKey(stackTraceKey string) ConfigurationBuilder {
	c.zapConfig.EncoderConfig.StacktraceKey = stackTraceKey
	return c
}

// SetECLineEnding sets log encoder config line ending
func (c *Configuration) SetECLineEnding(lineEnding string) ConfigurationBuilder {
	c.zapConfig.EncoderConfig.LineEnding = lineEnding
	return c
}

// SetECEncodeLevel sets log encoder config encode level
func (c *Configuration) SetECEncodeLevel(encodeLevel zc.LevelEncoder) ConfigurationBuilder {
	c.zapConfig.EncoderConfig.EncodeLevel = encodeLevel
	return c
}

// SetECTimeEncoder sets log encoder config time encoder
func (c *Configuration) SetECTimeEncoder(timeEncoder zc.TimeEncoder) ConfigurationBuilder {
	c.zapConfig.EncoderConfig.EncodeTime = timeEncoder
	return c
}

// SetECEncodeDuration sets log encoder config encode duration
func (c *Configuration) SetECEncodeDuration(encodeDuration zc.DurationEncoder) ConfigurationBuilder {
	c.zapConfig.EncoderConfig.EncodeDuration = encodeDuration
	return c
}

// SetECEncoderCaller sets log encoder config encoder caller
func (c *Configuration) SetECEncoderCaller(encoderCaller zc.CallerEncoder) ConfigurationBuilder {
	c.zapConfig.EncoderConfig.EncodeCaller = encoderCaller
	return c
}

// Build builds a custom log configuration
func (c *Configuration) Build() Configuration {
	return Configuration{
		zapConfig: c.zapConfig,
	}

}
