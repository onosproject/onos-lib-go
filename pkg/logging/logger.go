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
	"os"
	"strings"

	art "github.com/plar/go-adaptive-radix-tree"

	"go.uber.org/zap"
	zp "go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	zc "go.uber.org/zap/zapcore"
)

// init initialize logger package data structures
func init() {
	loggers = art.New()
	encoderConfig := zap.NewProductionEncoderConfig()
	defaultAtomLevel := zap.NewAtomicLevelAt(zapcore.InfoLevel)
	defaultEncoder := zapcore.NewJSONEncoder(encoderConfig)
	defaultWriter := zapcore.Lock(os.Stdout)
	core := zapcore.NewCore(defaultEncoder, defaultWriter, &defaultAtomLevel)
	zapLogger := zap.New(core).Named("root")
	root = Log{zapLogger, defaultEncoder, defaultWriter}
}

// SetLevel defines a new logger level
func (l *Log) SetLevel(level zc.Level) {
	newLevel := zc.Level(level)
	newLogger := l.stdLogger.WithOptions(
		zp.WrapCore(
			func(zc.Core) zc.Core {
				return zc.NewCore(l.encoder, l.writer, newLevel)
			}))
	l.stdLogger = newLogger
}

// Debug logs a message at Debug level on the sugar logger.
func (l *Log) Debug(args ...interface{}) {
	l.stdLogger.Sugar().Debug(args)
}

// Info logs a message at Info level on the sugar logger
func (l *Log) Info(args ...interface{}) {
	l.stdLogger.Sugar().Info(args)
}

// Error logs a message at Error level on the sugar logger
func (l *Log) Error(args ...interface{}) {
	l.stdLogger.Sugar().Error(args)
}

// Fatal logs a message at Fatal level on the sugar logger
func (l *Log) Fatal(args ...interface{}) {
	l.stdLogger.Sugar().Fatal(args)
}

// Panic logs a message at Panic level on the sugar logger
func (l *Log) Panic(args ...interface{}) {
	l.stdLogger.Sugar().Panic(args)
}

// DPanic logs a message at DPanic level on the sugar logger
func (l *Log) DPanic(args ...interface{}) {
	l.stdLogger.Sugar().DPanic(args)
}

// Warn logs a message at Warn level on the sugar logger
func (l *Log) Warn(args ...interface{}) {
	l.stdLogger.Sugar().Warn(args)
}

func buildTreeName(names ...string) string {
	var treeName string
	var values []string
	values = append(values, names...)
	treeName = strings.Join(values, ".")
	return treeName
}

func findParentsNames(name string) []string {
	names := strings.Split(name, ".")
	if len(names) > 0 {
		names = names[:len(names)-1]
	}
	return names
}

// GetKey returns a key object of radix tree
func GetKey(name string) art.Key {
	return art.Key(name)
}

func assignParentLevelLogger(name string) {
	parentNames := findParentsNames(name)
	for _, parentName := range parentNames {
		parentLogger, found := loggers.Search(art.Key(parentName))
		if found {
			logger := parentLogger.(Log)
			core := logger.stdLogger.Core()
			zapLogger := zap.New(core).Named(name)
			logger.stdLogger = zapLogger
			loggers.Insert(art.Key(name), logger)
			return
		}
	}
	root := GetDefaultLogger()
	core := root.stdLogger.Core()
	zapLogger := zap.New(core).Named(name)
	root.stdLogger = zapLogger
	loggers.Insert(art.Key(name), root)
}

// GetLogger gets a logger based on a give name
func GetLogger(names ...string) (Log, bool) {
	name := buildTreeName(names...)
	value, found := GetLoggers().Search(GetKey(name))
	if found {
		return value.(Log), found
	} else {
		return Log{}, found
	}
}

// AddLogger adds a logger based on a given level and a hierarchy of names
func AddLogger(level string, names ...string) {
	name := buildTreeName(names...)
	if level == "" {
		assignParentLevelLogger(name)
		return
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	atomLevel := zap.AtomicLevel{}
	switch level {
	case zapcore.InfoLevel.String():
		atomLevel = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case zapcore.DebugLevel.String():
		atomLevel = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case zapcore.ErrorLevel.String():
		atomLevel = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	case zapcore.PanicLevel.String():
		atomLevel = zap.NewAtomicLevelAt(zapcore.PanicLevel)
	case zapcore.FatalLevel.String():
		atomLevel = zap.NewAtomicLevelAt(zapcore.FatalLevel)
	case zapcore.WarnLevel.String():
		atomLevel = zap.NewAtomicLevelAt(zapcore.WarnLevel)
	}

	encoder := zapcore.NewJSONEncoder(encoderConfig)
	writer := zapcore.Lock(os.Stdout)
	core := zapcore.NewCore(encoder, writer, &atomLevel)
	zapLogger := zap.New(core).Named(name)
	logger := Log{zapLogger, encoder, writer}
	loggers.Insert(art.Key(name), logger)
}

// GetLoggers get loggers radix tree
func GetLoggers() art.Tree {
	return loggers
}

// GetDefaultLogger gets the default logger
func GetDefaultLogger() Log {
	return root
}
