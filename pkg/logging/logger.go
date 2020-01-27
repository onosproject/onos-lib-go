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
	zc "go.uber.org/zap/zapcore"
)

func getDefaultConfig(name string, level Level) Configuration {
	cfg := Configuration{}
	cfg.SetEncoding("json").
		SetLevel(level).
		SetOutputPaths([]string{"stdout"}).
		SetErrorOutputPaths([]string{"stderr"}).
		SetECMsgKey("msg").
		SetECLevelKey("level").
		SetECTimeKey("ts").
		SetECTimeEncoder(zc.ISO8601TimeEncoder).
		SetECEncodeLevel(zc.CapitalLevelEncoder).
		SetName(name).
		Build()
	return cfg
}

// init initialize logger package data structures
func init() {

	cfg := getDefaultConfig("root", levelToInt(zc.InfoLevel))
	rootLogger, _ := cfg.GetZapConfig().Build()
	defaultAtomLevel := zap.NewAtomicLevelAt(zc.InfoLevel)
	defaultEncoder := zc.NewJSONEncoder(cfg.GetZapConfig().EncoderConfig)
	defaultWriter := zc.Lock(os.Stdout)
	newLogger := rootLogger.WithOptions(
		zap.WrapCore(
			func(zc.Core) zc.Core {
				return zc.NewCore(defaultEncoder, defaultWriter, &defaultAtomLevel)
			}))

	rootLogger = newLogger.Named("root")
	loggers = art.New()
	root = Log{rootLogger, defaultEncoder, defaultWriter}
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

func (c *Configuration) AddLogger() {
	level := c.zapConfig.Level.Level().String()
	name := c.zapConfig.EncoderConfig.NameKey
	if level == "" {
		assignParentLevelLogger(name)
		return
	}

	atomLevel := zap.AtomicLevel{}
	switch level {
	case zc.InfoLevel.String():
		atomLevel = zap.NewAtomicLevelAt(zc.InfoLevel)
	case zc.DebugLevel.String():
		atomLevel = zap.NewAtomicLevelAt(zc.DebugLevel)
	case zc.ErrorLevel.String():
		atomLevel = zap.NewAtomicLevelAt(zc.ErrorLevel)
	case zc.PanicLevel.String():
		atomLevel = zap.NewAtomicLevelAt(zc.PanicLevel)
	case zc.FatalLevel.String():
		atomLevel = zap.NewAtomicLevelAt(zc.FatalLevel)
	case zc.WarnLevel.String():
		atomLevel = zap.NewAtomicLevelAt(zc.WarnLevel)
	}

	cfg := c.zapConfig
	configLogger, _ := cfg.Build()
	encoder := zc.NewJSONEncoder(cfg.EncoderConfig)
	writer := zc.Lock(os.Stdout)
	newLogger := configLogger.WithOptions(
		zap.WrapCore(
			func(zc.Core) zc.Core {
				return zc.NewCore(encoder, writer, &atomLevel)
			}))

	configLogger = newLogger.Named(name)
	logger := Log{configLogger, encoder, writer}
	loggers.Insert(art.Key(name), logger)

}

// AddLogger adds a logger based on a given level and a hierarchy of names
func AddLogger(level string, names ...string) {
	name := buildTreeName(names...)
	if level == "" {
		assignParentLevelLogger(name)
		return
	}

	atomLevel := zap.AtomicLevel{}
	var internalLevel Level
	switch level {
	case zc.InfoLevel.String():
		atomLevel = zap.NewAtomicLevelAt(zc.InfoLevel)
		internalLevel = InfoLevel
	case zc.DebugLevel.String():
		atomLevel = zap.NewAtomicLevelAt(zc.DebugLevel)
		internalLevel = DebugLevel
	case zc.ErrorLevel.String():
		atomLevel = zap.NewAtomicLevelAt(zc.ErrorLevel)
		internalLevel = ErrorLevel
	case zc.PanicLevel.String():
		atomLevel = zap.NewAtomicLevelAt(zc.PanicLevel)
		internalLevel = PanicLevel
	case zc.FatalLevel.String():
		atomLevel = zap.NewAtomicLevelAt(zc.FatalLevel)
		internalLevel = FatalLevel
	case zc.WarnLevel.String():
		atomLevel = zap.NewAtomicLevelAt(zc.WarnLevel)
		internalLevel = WarnLevel
	}

	cfg := getDefaultConfig(name, internalLevel)
	configLogger, _ := cfg.GetZapConfig().Build()
	defaultEncoder := zc.NewJSONEncoder(cfg.GetZapConfig().EncoderConfig)
	defaultWriter := zc.Lock(os.Stdout)
	newLogger := configLogger.WithOptions(
		zap.WrapCore(
			func(zc.Core) zc.Core {
				return zc.NewCore(defaultEncoder, defaultWriter, &atomLevel)
			}))

	configLogger = newLogger.Named(name)
	logger := Log{configLogger, defaultEncoder, defaultWriter}
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
