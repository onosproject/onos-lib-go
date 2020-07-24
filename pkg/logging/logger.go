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
	"fmt"
	"strings"
)

var root *zapLogger

const nameSep = "/"

func init() {
	config := Config{}
	if err := load(&config); err != nil {
		panic(err)
	} else if err := configure(config); err != nil {
		panic(err)
	}
}

// configure configures the loggers
func configure(config Config) error {
	rootLogger, err := newZapLogger(config, config.GetRootLogger())
	if err != nil {
		return err
	}
	root = rootLogger
	return nil
}

// GetLogger gets a logger by name
func GetLogger(names ...string) Logger {
	if len(names) == 1 {
		names = strings.Split(names[0], nameSep)
	}

	logger := root
	for _, name := range names {
		child, err := logger.getChild(name)
		if err != nil {
			panic(err)
		}
		logger = child
	}
	return logger
}

// Logger represents an abstract logging interface.
type Logger interface {
	Output

	// Name returns the logger name
	Name() string

	// GetLevel returns the logger's level
	GetLevel() Level

	// SetLevel sets the logger's level
	SetLevel(level Level)
}

func newZapLogger(config Config, loggerConfig LoggerConfig) (*zapLogger, error) {
	var outputs []Output
	outputConfigs := loggerConfig.GetOutputs()
	if len(outputConfigs) > 0 {
		outputs = make([]Output, len(outputConfigs))
		for i, outputConfig := range outputConfigs {
			var sinkConfig SinkConfig
			if outputConfig.Sink == "" {
				sinkConfig = config.GetDefaultSink()
			} else {
				sink, ok := config.GetSink(outputConfig.Sink)
				if !ok {
					panic(fmt.Sprintf("unknown sink %s", outputConfig.Sink))
				}
				sinkConfig = sink
			}
			output, err := newZapOutput(loggerConfig, outputConfig, sinkConfig)
			if err != nil {
				return nil, err
			}
			outputs[i] = output
		}
	} else {
		output, err := newZapOutput(loggerConfig, OutputConfig{Sink: config.GetDefaultSink().Name}, config.GetDefaultSink())
		if err != nil {
			return nil, err
		}
		outputs = []Output{output}
	}

	var level *Level
	if loggerConfig.Level != nil {
		loggerLevel := loggerConfig.GetLevel()
		level = &loggerLevel
	}

	return &zapLogger{
		config:       config,
		loggerConfig: loggerConfig,
		children:     make(map[string]*zapLogger),
		outputs:      outputs,
		level:        level,
	}, nil
}

// zapLogger is the default Logger implementation
type zapLogger struct {
	config       Config
	loggerConfig LoggerConfig
	children     map[string]*zapLogger
	outputs      []Output
	defaultLevel Level
	level        *Level
}

func (l *zapLogger) Name() string {
	return l.loggerConfig.Name
}

func (l *zapLogger) getChild(name string) (*zapLogger, error) {
	child, ok := l.children[name]
	if !ok {
		qualifiedName := strings.Trim(fmt.Sprintf("%s%s%s", l.loggerConfig.Name, nameSep, name), nameSep)
		loggerConfig, ok := l.config.GetLogger(qualifiedName)
		if !ok {
			loggerConfig = l.loggerConfig
			loggerConfig.Name = qualifiedName
			loggerConfig.Level = nil
		}

		logger, err := newZapLogger(l.config, loggerConfig)
		if err != nil {
			return nil, err
		}
		logger.setDefaultLevel(l.GetLevel())
		l.children[name] = logger
		child = logger
	}
	return child, nil
}

func (l *zapLogger) GetLevel() Level {
	level := l.level
	if level != nil {
		return *level
	}
	return l.defaultLevel
}

func (l *zapLogger) SetLevel(level Level) {
	l.level = &level
	l.setDefaultLevel(level)
}

func (l *zapLogger) setDefaultLevel(level Level) {
	l.defaultLevel = level
	for _, child := range l.children {
		child.setDefaultLevel(level)
	}
}

func (l *zapLogger) Debug(args ...interface{}) {
	if l.GetLevel() <= DebugLevel {
		for _, output := range l.outputs {
			output.Debug(args...)
		}
	}
}

func (l *zapLogger) Debugf(template string, args ...interface{}) {
	if l.GetLevel() <= DebugLevel {
		for _, output := range l.outputs {
			output.Debugf(template, args...)
		}
	}
}

func (l *zapLogger) Debugw(msg string, keysAndValues ...interface{}) {
	if l.GetLevel() <= DebugLevel {
		for _, output := range l.outputs {
			output.Debugw(msg, keysAndValues...)
		}
	}
}

func (l *zapLogger) Info(args ...interface{}) {
	if l.GetLevel() <= InfoLevel {
		for _, output := range l.outputs {
			output.Info(args...)
		}
	}
}

func (l *zapLogger) Infof(template string, args ...interface{}) {
	if l.GetLevel() <= InfoLevel {
		for _, output := range l.outputs {
			output.Infof(template, args...)
		}
	}
}

func (l *zapLogger) Infow(msg string, keysAndValues ...interface{}) {
	if l.GetLevel() <= InfoLevel {
		for _, output := range l.outputs {
			output.Infow(msg, keysAndValues...)
		}
	}
}

func (l *zapLogger) Error(args ...interface{}) {
	if l.GetLevel() <= ErrorLevel {
		for _, output := range l.outputs {
			output.Error(args...)
		}
	}
}

func (l *zapLogger) Errorf(template string, args ...interface{}) {
	if l.GetLevel() <= ErrorLevel {
		for _, output := range l.outputs {
			output.Errorf(template, args...)
		}
	}
}

func (l *zapLogger) Errorw(msg string, keysAndValues ...interface{}) {
	if l.GetLevel() <= ErrorLevel {
		for _, output := range l.outputs {
			output.Errorw(msg, keysAndValues...)
		}
	}
}

func (l *zapLogger) Fatal(args ...interface{}) {
	if l.GetLevel() <= FatalLevel {
		for _, output := range l.outputs {
			output.Fatal(args...)
		}
	}
}

func (l *zapLogger) Fatalf(template string, args ...interface{}) {
	if l.GetLevel() <= FatalLevel {
		for _, output := range l.outputs {
			output.Fatalf(template, args...)
		}
	}
}

func (l *zapLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	if l.GetLevel() <= FatalLevel {
		for _, output := range l.outputs {
			output.Fatalw(msg, keysAndValues...)
		}
	}
}

func (l *zapLogger) Panic(args ...interface{}) {
	if l.GetLevel() <= PanicLevel {
		for _, output := range l.outputs {
			output.Panic(args...)
		}
	}
}

func (l *zapLogger) Panicf(template string, args ...interface{}) {
	if l.GetLevel() <= PanicLevel {
		for _, output := range l.outputs {
			output.Panicf(template, args...)
		}
	}
}

func (l *zapLogger) Panicw(msg string, keysAndValues ...interface{}) {
	if l.GetLevel() <= PanicLevel {
		for _, output := range l.outputs {
			output.Panicw(msg, keysAndValues...)
		}
	}
}

func (l *zapLogger) DPanic(args ...interface{}) {
	if l.GetLevel() <= DPanicLevel {
		for _, output := range l.outputs {
			output.DPanic(args...)
		}
	}
}

func (l *zapLogger) DPanicf(template string, args ...interface{}) {
	if l.GetLevel() <= DPanicLevel {
		for _, output := range l.outputs {
			output.DPanicf(template, args...)
		}
	}
}

func (l *zapLogger) DPanicw(msg string, keysAndValues ...interface{}) {
	if l.GetLevel() <= DPanicLevel {
		for _, output := range l.outputs {
			output.DPanicw(msg, keysAndValues...)
		}
	}
}

func (l *zapLogger) Warn(args ...interface{}) {
	if l.GetLevel() <= WarnLevel {
		for _, output := range l.outputs {
			output.Warn(args...)
		}
	}
}

func (l *zapLogger) Warnf(template string, args ...interface{}) {
	if l.GetLevel() <= WarnLevel {
		for _, output := range l.outputs {
			output.Warnf(template, args...)
		}
	}
}

func (l *zapLogger) Warnw(msg string, keysAndValues ...interface{}) {
	if l.GetLevel() <= WarnLevel {
		for _, output := range l.outputs {
			output.Warnw(msg, keysAndValues...)
		}
	}
}

var _ Logger = &zapLogger{}
