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

// LoggerConfig log config structure
type LoggerConfig struct {
	zapConfig zp.Config
	sinkURLs  []SinkURL
}

// LoggerBuilder log configuration builder interface
type LoggerBuilder interface {
	SetLevel(Level) LoggerBuilder
	SetEncoding(string) LoggerBuilder
	SetDevelopment(bool) LoggerBuilder
	SetOutputPaths([]string) LoggerBuilder
	SetErrorOutputPaths([]string) LoggerBuilder
	SetSinkURLs([]SinkURL) LoggerBuilder
	//SetInitialFields(Fields) LoggerBuilder
	SetECLevelKey(string) LoggerBuilder
	SetECMsgKey(string) LoggerBuilder
	SetECTimeKey(string) LoggerBuilder
	SetECStackTraceKey(string) LoggerBuilder
	SetECLineEnding(string) LoggerBuilder
	SetECEncodeLevel(zc.LevelEncoder) LoggerBuilder
	SetECTimeEncoder(zc.TimeEncoder) LoggerBuilder
	SetName(...string) LoggerBuilder
	SetECEncodeDuration(zc.DurationEncoder) LoggerBuilder
	SetECEncoderCaller(string, zc.CallerEncoder) LoggerBuilder
	Build() LoggerConfig
}

// GetSinkURLs gets sink urls
func (c *LoggerConfig) GetSinkURLs() []SinkURL {
	return c.sinkURLs
}

// SetSinkURLs sets sink urls
func (c *LoggerConfig) SetSinkURLs(sinkURLs []SinkURL) LoggerBuilder {
	c.sinkURLs = sinkURLs
	return c
}

// GetZapConfig gets zap configuration
func (c *LoggerConfig) GetZapConfig() zp.Config {
	return c.zapConfig
}

// SetLevel sets log level
func (c *LoggerConfig) SetLevel(level Level) LoggerBuilder {
	c.zapConfig.Level = intToAtomicLevel(level)
	return c
}

// SetEncoding sets log encoding
func (c *LoggerConfig) SetEncoding(encoding string) LoggerBuilder {
	c.zapConfig.Encoding = encoding
	return c
}

// SetDevelopment sets log development
func (c *LoggerConfig) SetDevelopment(development bool) LoggerBuilder {
	c.zapConfig.Development = development
	return c
}

// SetOutputPaths sets log output paths
func (c *LoggerConfig) SetOutputPaths(outputPaths []string) LoggerBuilder {
	c.zapConfig.OutputPaths = outputPaths
	return c
}

// SetErrorOutputPaths sets log error output paths
func (c *LoggerConfig) SetErrorOutputPaths(errorOutputPaths []string) LoggerBuilder {
	c.zapConfig.ErrorOutputPaths = errorOutputPaths
	return c
}

// SetInitialFields sets log initial fields
/*func (c *LoggerConfig) SetInitialFields(initFields Fields) LoggerBuilder {
	c.zapConfig.InitialFields = initFields
	return c
}*/

// SetECLevelKey sets log encoder config level key
func (c *LoggerConfig) SetECLevelKey(levelKey string) LoggerBuilder {
	c.zapConfig.EncoderConfig.LevelKey = levelKey
	return c
}

// SetName sets encoder config name key
func (c *LoggerConfig) SetName(names ...string) LoggerBuilder {
	nameKey := buildTreeName(names...)
	c.zapConfig.EncoderConfig.NameKey = nameKey
	return c
}

// SetECMsgKey sets log encoder config message key
func (c *LoggerConfig) SetECMsgKey(msgKey string) LoggerBuilder {
	c.zapConfig.EncoderConfig.MessageKey = msgKey
	return c
}

// SetECTimeKey sets log encoder config time key
func (c *LoggerConfig) SetECTimeKey(timeKey string) LoggerBuilder {
	c.zapConfig.EncoderConfig.TimeKey = timeKey
	return c
}

// SetECStackTraceKey sets log encoder config start trace key
func (c *LoggerConfig) SetECStackTraceKey(stackTraceKey string) LoggerBuilder {
	c.zapConfig.EncoderConfig.StacktraceKey = stackTraceKey
	return c
}

// SetECLineEnding sets log encoder config line ending
func (c *LoggerConfig) SetECLineEnding(lineEnding string) LoggerBuilder {
	c.zapConfig.EncoderConfig.LineEnding = lineEnding
	return c
}

// SetECEncodeLevel sets log encoder config encode level
func (c *LoggerConfig) SetECEncodeLevel(encodeLevel zc.LevelEncoder) LoggerBuilder {
	c.zapConfig.EncoderConfig.EncodeLevel = encodeLevel
	return c
}

// SetECTimeEncoder sets log encoder config time encoder
func (c *LoggerConfig) SetECTimeEncoder(timeEncoder zc.TimeEncoder) LoggerBuilder {
	c.zapConfig.EncoderConfig.EncodeTime = timeEncoder
	return c
}

// SetECEncodeDuration sets log encoder config encode duration
func (c *LoggerConfig) SetECEncodeDuration(encodeDuration zc.DurationEncoder) LoggerBuilder {
	c.zapConfig.EncoderConfig.EncodeDuration = encodeDuration
	return c
}

// SetECEncoderCaller sets log encoder config encoder caller
func (c *LoggerConfig) SetECEncoderCaller(callerKey string, encoderCaller zc.CallerEncoder) LoggerBuilder {
	c.zapConfig.EncoderConfig.EncodeCaller = encoderCaller
	c.zapConfig.EncoderConfig.CallerKey = callerKey
	return c
}

// Build builds a custom log configuration
func (c *LoggerConfig) Build() LoggerConfig {
	return LoggerConfig{
		zapConfig: c.zapConfig,
	}

}
