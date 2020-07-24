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
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const configDir = ".onos"

// SinkType is the type of a sink
type SinkType string

func (t SinkType) String() string {
	return string(t)
}

const (
	StdoutSinkType SinkType = "stdout"
	KafkaSinkType  SinkType = "kafka"
)

// SinkEncoding is the encoding for a sink
type SinkEncoding string

func (e SinkEncoding) String() string {
	return string(e)
}

const (
	ConsoleEncoding SinkEncoding = "console"
	JSONEncoding    SinkEncoding = "json"
)

const (
	rootLoggerName  = "root"
	defaultSinkName = "default"
)

// Config logging configuration
type Config struct {
	Loggers map[string]LoggerConfig `yaml:"loggers"`
	Sinks   map[string]SinkConfig   `yaml:"sinks"`
}

// GetRootLogger returns the root logger configuration
func (c Config) GetRootLogger() LoggerConfig {
	root, ok := c.Loggers[rootLoggerName]
	if ok {
		return root
	}
	return LoggerConfig{Output: map[string]OutputConfig{}}
}

// GetLoggers returns the configured loggers
func (c Config) GetLoggers() []LoggerConfig {
	loggers := make([]LoggerConfig, 0, len(c.Loggers))
	for name, logger := range c.Loggers {
		if name != rootLoggerName {
			logger.Name = name
			loggers = append(loggers, logger)
		}
	}
	return loggers
}

// GetLogger returns a logger by name
func (c Config) GetLogger(name string) (LoggerConfig, bool) {
	if name == rootLoggerName {
		return LoggerConfig{}, false
	}

	logger, ok := c.Loggers[name]
	if ok {
		logger.Name = name
		return logger, true
	}
	return LoggerConfig{}, false
}

// GetDefaultSink returns the default sink
func (c Config) GetDefaultSink() SinkConfig {
	sink, ok := c.Sinks[defaultSinkName]
	if !ok {
		sinkType := StdoutSinkType
		sinkEncoding := ConsoleEncoding
		sink = SinkConfig{
			Name:     StdoutSinkType.String(),
			Type:     &sinkType,
			Encoding: &sinkEncoding,
		}
	}
	return sink
}

// GetSinks returns the configured sinks
func (c Config) GetSinks() []SinkConfig {
	sinks := make([]SinkConfig, 0, len(c.Sinks))
	for name, sink := range c.Sinks {
		if name != defaultSinkName {
			sink.Name = name
			sinks = append(sinks, sink)
		}
	}
	return sinks
}

// GetSink returns a sink by name
func (c Config) GetSink(name string) (SinkConfig, bool) {
	if name == defaultSinkName {
		return SinkConfig{}, false
	}

	sink, ok := c.Sinks[name]
	if ok {
		sink.Name = name
		return sink, true
	}
	return SinkConfig{}, false
}

// LoggerConfig is the configuration for a logger
type LoggerConfig struct {
	Name   string                  `yaml:"name"`
	Level  *string                 `yaml:"level,omitempty"`
	Output map[string]OutputConfig `yaml:"output,omitempty"`
}

// GetLevel returns the logger level
func (c LoggerConfig) GetLevel() Level {
	level := c.Level
	if level != nil {
		return levelStringToLevel(*level)
	}
	return ErrorLevel
}

// GetOutputs returns the logger outputs
func (c LoggerConfig) GetOutputs() []OutputConfig {
	outputs := c.Output
	if outputs == nil {
		return []OutputConfig{}
	}

	outputsList := make([]OutputConfig, 0, len(outputs))
	for _, output := range outputs {
		outputsList = append(outputsList, output)
	}
	return outputsList
}

// SinkRefConfig is the configuration for a sink instance
type OutputConfig struct {
	Sink  string  `yaml:"sink"`
	Level *string `yaml:"level,omitempty"`
}

// GetLevel returns the output level
func (c OutputConfig) GetLevel() Level {
	level := c.Level
	if level != nil {
		return levelStringToLevel(*level)
	}
	return DebugLevel
}

// SinkConfig is the configuration for a sink
type SinkConfig struct {
	Name     string            `yaml:"name"`
	Type     *SinkType         `yaml:"type,omitempty"`
	Encoding *SinkEncoding     `yaml:"encoding,omitempty"`
	Stdout   *StdoutSinkConfig `yaml:"stdout,omitempty"`
	Kafka    *KafkaSinkConfig  `yaml:"kafka,omitempty"`
}

// GetType returns the sink type
func (c SinkConfig) GetType() SinkType {
	sinkType := c.Type
	if sinkType != nil {
		return *sinkType
	}
	return StdoutSinkType
}

// GetEncoding returns the sink encoding
func (c SinkConfig) GetEncoding() SinkEncoding {
	encoding := c.Encoding
	if encoding != nil {
		return *encoding
	}
	return ConsoleEncoding
}

// GetStdoutSinkConfig returns the stdout sink configuration
func (c SinkConfig) GetStdoutSinkConfig() StdoutSinkConfig {
	config := c.Stdout
	if config != nil {
		return *config
	}
	return StdoutSinkConfig{}
}

// GetKafkaSinkConfig returns the Kafka sink configuration
func (c SinkConfig) GetKafkaSinkConfig() KafkaSinkConfig {
	config := c.Kafka
	if config != nil {
		return *config
	}
	return KafkaSinkConfig{}
}

// StdoutSinkConfig is the configuration for an stdout sink
type StdoutSinkConfig struct {
}

// KafkaSinkConfig is the configuration for a Kafka sink
type KafkaSinkConfig struct {
	Topic   string   `yaml:"topic"`
	Key     string   `yaml:"key"`
	Brokers []string `yaml:"brokers"`
}

// load loads the configuration
func load(config *Config) error {
	home, err := homedir.Dir()
	if err != nil {
		return err
	}

	// Set the file name of the configurations file
	viper.SetConfigName("logging")

	// Set the path to look for the configurations file
	viper.AddConfigPath("./" + configDir + "/config")
	viper.AddConfigPath(home + "/" + configDir + "/config")
	viper.AddConfigPath("/etc/onos/config")
	viper.AddConfigPath(".")

	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	err = viper.Unmarshal(config)
	if err != nil {
		return err
	}
	return nil
}
