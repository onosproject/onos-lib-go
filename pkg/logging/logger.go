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
	"sync"
	"sync/atomic"
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
	rootLogger, err := newZapLogger(config, config.GetRootLogger(), EmptyLevel)
	if err != nil {
		return err
	}
	root = rootLogger
	return nil
}

// GetLogger gets a logger by name
func GetLogger(names ...string) Logger {
	return root.GetLogger(names...)
}

// Logger represents an abstract logging interface.
type Logger interface {
	// Name returns the logger name
	Name() string

	// GetLogger gets a descendant of this Logger
	GetLogger(names ...string) Logger

	// Level returns the logger's level
	Level() Level

	// SetLevel sets the logger's level
	SetLevel(level Level)

	// WithFields adds fields to the logger
	WithFields(fields ...Field) Logger

	// Sync flushes the logger
	Sync() error

	Debug(...interface{})
	Debugf(template string, args ...interface{})
	Debugw(msg string, fields ...Field)

	Info(...interface{})
	Infof(template string, args ...interface{})
	Infow(msg string, fields ...Field)

	Error(...interface{})
	Errorf(template string, args ...interface{})
	Errorw(msg string, fields ...Field)

	Fatal(...interface{})
	Fatalf(template string, args ...interface{})
	Fatalw(msg string, fields ...Field)

	Panic(...interface{})
	Panicf(template string, args ...interface{})
	Panicw(msg string, fields ...Field)

	DPanic(...interface{})
	DPanicf(template string, args ...interface{})
	DPanicw(msg string, fields ...Field)

	Warn(...interface{})
	Warnf(template string, args ...interface{})
	Warnw(msg string, fields ...Field)
}

func newZapLogger(config Config, loggerConfig LoggerConfig, defaultLevel Level) (*zapLogger, error) {
	var outputs []*zapOutput
	outputConfigs := loggerConfig.GetOutputs()
	outputs = make([]*zapOutput, len(outputConfigs))
	for i, outputConfig := range outputConfigs {
		var sinkConfig SinkConfig
		if outputConfig.Sink == nil {
			return nil, fmt.Errorf("output sink not configured for output %s", outputConfig.Name)
		}
		sink, ok := config.GetSink(*outputConfig.Sink)
		if !ok {
			panic(fmt.Sprintf("unknown sink %s", *outputConfig.Sink))
		}
		sinkConfig = sink
		output, err := newZapOutput(loggerConfig, outputConfig, sinkConfig)
		if err != nil {
			return nil, err
		}
		outputs[i] = output
	}

	logger := &zapLogger{
		config:       config,
		loggerConfig: loggerConfig,
		children:     make(map[string]*zapLogger),
		outputs:      outputs,
		mu:           &sync.RWMutex{},
		defaultLevel: defaultLevel,
	}
	if loggerConfig.Level != nil {
		logger.SetLevel(loggerConfig.GetLevel())
	}
	return logger, nil
}

// zapLogger is the default Logger implementation
type zapLogger struct {
	config       Config
	loggerConfig LoggerConfig
	children     map[string]*zapLogger
	outputs      []*zapOutput
	mu           *sync.RWMutex
	atomicLevel  atomic.Value
	level        *Level
	defaultLevel Level
}

func (l *zapLogger) Name() string {
	return l.loggerConfig.Name
}

func (l *zapLogger) GetLogger(names ...string) Logger {
	if len(names) == 1 {
		names = strings.Split(names[0], nameSep)
	}

	logger := l
	for _, name := range names {
		child, err := logger.getChild(name)
		if err != nil {
			panic(err)
		}
		logger = child
	}
	return logger
}

func (l *zapLogger) getChild(name string) (*zapLogger, error) {
	l.mu.RLock()
	child, ok := l.children[name]
	l.mu.RUnlock()
	if ok {
		return child, nil
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	child, ok = l.children[name]
	if ok {
		return child, nil
	}

	// Compute the name of the child logger
	qualifiedName := strings.Trim(fmt.Sprintf("%s%s%s", l.loggerConfig.Name, nameSep, name), nameSep)

	// Initialize the child logger's configuration if one is not set.
	loggerConfig, ok := l.config.GetLogger(qualifiedName)
	if !ok {
		loggerConfig = l.loggerConfig
		loggerConfig.Name = qualifiedName
		loggerConfig.Level = nil
	}

	// Populate the child logger configuration with outputs inherited from this logger.
	for _, output := range l.outputs {
		outputConfig, ok := loggerConfig.GetOutput(output.config.Name)
		if !ok {
			loggerConfig.Output[output.config.Name] = output.config
		} else {
			if outputConfig.Sink == nil {
				outputConfig.Sink = output.config.Sink
			}
			if outputConfig.Level == nil {
				outputConfig.Level = output.config.Level
			}
			loggerConfig.Output[outputConfig.Name] = outputConfig
		}
	}

	// Create the child logger.
	logger, err := newZapLogger(l.config, loggerConfig, l.Level())
	if err != nil {
		return nil, err
	}

	// Set the default log level on the child.
	l.children[name] = logger
	return logger, nil
}

func (l *zapLogger) Level() Level {
	level := l.atomicLevel.Load()
	if level != nil {
		return *level.(*Level)
	}
	return l.defaultLevel
}

func (l *zapLogger) SetLevel(level Level) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = &level
	l.atomicLevel.Store(&level)
	for _, child := range l.children {
		child.setDefaultLevel(level)
	}
}

func (l *zapLogger) setDefaultLevel(level Level) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.defaultLevel = level
	if l.level == nil {
		l.atomicLevel.Store(&level)
		for _, child := range l.children {
			child.setDefaultLevel(level)
		}
	}
}

func (l *zapLogger) WithFields(fields ...Field) Logger {
	outputs := make([]*zapOutput, len(l.outputs))
	for i, output := range l.outputs {
		outputs[i] = output.WithFields(fields...).(*zapOutput)
	}
	return &zapLogger{
		config:       l.config,
		loggerConfig: l.loggerConfig,
		children:     l.children,
		outputs:      outputs,
		mu:           l.mu,
		atomicLevel:  l.atomicLevel,
		level:        l.level,
		defaultLevel: l.defaultLevel,
	}
}

func (l *zapLogger) Sync() error {
	var err error
	for _, output := range l.outputs {
		err = output.Sync()
	}
	return err
}

func (l *zapLogger) log(level Level, template string, args []interface{}, fields []Field, logger func(output Output, msg string, fields []Field)) {
	if l.Level() > level {
		return
	}

	msg := template
	if msg == "" && len(args) > 0 {
		msg = fmt.Sprint(args...)
	} else if msg != "" && len(args) > 0 {
		msg = fmt.Sprintf(template, args...)
	}

	for _, output := range l.outputs {
		logger(output, msg, fields)
	}
}

func (l *zapLogger) Debug(args ...interface{}) {
	l.log(DebugLevel, "", args, nil, func(output Output, msg string, fields []Field) {
		output.Debug(msg, fields...)
	})
}

func (l *zapLogger) Debugf(template string, args ...interface{}) {
	l.log(DebugLevel, template, args, nil, func(output Output, msg string, fields []Field) {
		output.Debug(msg, fields...)
	})
}

func (l *zapLogger) Debugw(msg string, fields ...Field) {
	l.log(DebugLevel, "", nil, fields, func(output Output, _ string, fields []Field) {
		output.Debug(msg, fields...)
	})
}

func (l *zapLogger) Info(args ...interface{}) {
	l.log(InfoLevel, "", args, nil, func(output Output, msg string, fields []Field) {
		output.Info(msg, fields...)
	})
}

func (l *zapLogger) Infof(template string, args ...interface{}) {
	l.log(InfoLevel, template, args, nil, func(output Output, msg string, fields []Field) {
		output.Info(msg, fields...)
	})
}

func (l *zapLogger) Infow(msg string, fields ...Field) {
	l.log(InfoLevel, "", nil, fields, func(output Output, _ string, fields []Field) {
		output.Info(msg, fields...)
	})
}

func (l *zapLogger) Warn(args ...interface{}) {
	l.log(WarnLevel, "", args, nil, func(output Output, msg string, fields []Field) {
		output.Warn(msg, fields...)
	})
}

func (l *zapLogger) Warnf(template string, args ...interface{}) {
	l.log(WarnLevel, template, args, nil, func(output Output, msg string, fields []Field) {
		output.Warn(msg, fields...)
	})
}

func (l *zapLogger) Warnw(msg string, fields ...Field) {
	l.log(WarnLevel, "", nil, fields, func(output Output, _ string, fields []Field) {
		output.Warn(msg, fields...)
	})
}

func (l *zapLogger) Error(args ...interface{}) {
	l.log(ErrorLevel, "", args, nil, func(output Output, msg string, fields []Field) {
		output.Error(msg, fields...)
	})
}

func (l *zapLogger) Errorf(template string, args ...interface{}) {
	l.log(ErrorLevel, template, args, nil, func(output Output, msg string, fields []Field) {
		output.Error(msg, fields...)
	})
}

func (l *zapLogger) Errorw(msg string, fields ...Field) {
	l.log(ErrorLevel, "", nil, fields, func(output Output, _ string, fields []Field) {
		output.Error(msg, fields...)
	})
}

func (l *zapLogger) Fatal(args ...interface{}) {
	l.log(FatalLevel, "", args, nil, func(output Output, msg string, fields []Field) {
		output.Fatal(msg, fields...)
	})
}

func (l *zapLogger) Fatalf(template string, args ...interface{}) {
	l.log(FatalLevel, template, args, nil, func(output Output, msg string, fields []Field) {
		output.Fatal(msg, fields...)
	})
}

func (l *zapLogger) Fatalw(msg string, fields ...Field) {
	l.log(FatalLevel, "", nil, fields, func(output Output, _ string, fields []Field) {
		output.Fatal(msg, fields...)
	})
}

func (l *zapLogger) Panic(args ...interface{}) {
	l.log(PanicLevel, "", args, nil, func(output Output, msg string, fields []Field) {
		output.Panic(msg, fields...)
	})
}

func (l *zapLogger) Panicf(template string, args ...interface{}) {
	l.log(PanicLevel, template, args, nil, func(output Output, msg string, fields []Field) {
		output.Panic(msg, fields...)
	})
}

func (l *zapLogger) Panicw(msg string, fields ...Field) {
	l.log(PanicLevel, "", nil, fields, func(output Output, _ string, fields []Field) {
		output.Panic(msg, fields...)
	})
}

func (l *zapLogger) DPanic(args ...interface{}) {
	l.log(DPanicLevel, "", args, nil, func(output Output, msg string, fields []Field) {
		output.DPanic(msg, fields...)
	})
}

func (l *zapLogger) DPanicf(template string, args ...interface{}) {
	l.log(DPanicLevel, template, args, nil, func(output Output, msg string, fields []Field) {
		output.DPanic(msg, fields...)
	})
}

func (l *zapLogger) DPanicw(msg string, fields ...Field) {
	l.log(DPanicLevel, "", nil, fields, func(output Output, _ string, fields []Field) {
		output.DPanic(msg, fields...)
	})
}

var _ Logger = &zapLogger{}
