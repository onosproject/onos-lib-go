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

	"fmt"

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
		SetECEncoderCaller("caller", zc.ShortCallerEncoder).
		SetECTimeKey("ts").
		SetECTimeEncoder(zc.ISO8601TimeEncoder).
		SetECEncodeLevel(zc.CapitalLevelEncoder).
		SetName(name).
		Build()
	return cfg
}

// init initialize logger package data structures
func init() {
	defaultLoggerName := "root"
	cfg := getDefaultConfig(defaultLoggerName, levelToInt(zc.InfoLevel))
	rootLogger, _ := cfg.GetZapConfig().Build(zap.AddCallerSkip(1))
	defaultAtomLevel := zap.NewAtomicLevelAt(zc.InfoLevel)
	defaultEncoder := zc.NewJSONEncoder(cfg.GetZapConfig().EncoderConfig)
	defaultWriter := zc.Lock(os.Stdout)
	newLogger := rootLogger.WithOptions(
		zap.WrapCore(
			func(zc.Core) zc.Core {
				return zc.NewCore(defaultEncoder, defaultWriter, &defaultAtomLevel)
			}))

	rootLogger = newLogger.Named(defaultLoggerName)
	loggers = art.New()
	err := zap.RegisterSink("kafka", GetSink)
	if err != nil {
		fmt.Println("Kafka Sink cannot be registered", err)
	}
	root = Log{rootLogger, defaultEncoder, defaultWriter, defaultLoggerName}
}

// SetLevel defines a new logger level and propagate the change its children
func (l *Log) SetLevel(level Level) {
	logger := Log{}
	loggers.ForEachPrefix(art.Key(l.name), func(node art.Node) bool {
		if node.Key() != nil {
			loggerNode := node.Value().(Log)
			newLevel := intToAtomicLevel(level)
			newLogger := loggerNode.stdLogger.WithOptions(
				zp.WrapCore(
					func(zc.Core) zc.Core {
						return zc.NewCore(loggerNode.encoder, loggerNode.writer, &newLevel)
					}))
			newLogger.Named(string(node.Key()))
			loggerNode.stdLogger = newLogger
			l.stdLogger = newLogger
			logger = Log{newLogger, loggerNode.encoder, loggerNode.writer, string(node.Key())}
			loggers.Insert(node.Key(), logger)
		}
		return true
	})
}

func (l *Log) RemoveSink() {
	key := art.Key(l.name)
	value, found := loggers.Search(key)
	if found {
		loggerNode := value.(Log)
		newLevel := intToAtomicLevel(InfoLevel)
		defaultWriter := zc.Lock(os.Stdout)
		newLogger := loggerNode.stdLogger.WithOptions(
			zp.WrapCore(
				func(zc.Core) zc.Core {
					return zc.NewCore(loggerNode.encoder, defaultWriter, &newLevel)
				}))
		newLogger.Named(string(key))
		loggerNode.stdLogger = newLogger
		l.stdLogger = newLogger
		logger := Log{newLogger, loggerNode.encoder, defaultWriter, string(key)}
		loggers.Insert(key, logger)
	}

}

func (l *Log) AddSink(sink SinkURL) {
	key := art.Key(l.name)
	value, found := loggers.Search(key)
	if found {
		loggerNode := value.(Log)
		ws, _, err := zap.Open(sink.String())
		if err != nil {
			fmt.Println("Cannot open sink", err)
		}
		newLevel := intToAtomicLevel(InfoLevel)
		newLogger := loggerNode.stdLogger.WithOptions(
			zp.WrapCore(
				func(zc.Core) zc.Core {
					return zc.NewCore(loggerNode.encoder, ws, &newLevel)
				}))
		newLogger.Named(string(key))
		loggerNode.stdLogger = newLogger
		l.stdLogger = newLogger
		logger := Log{newLogger, loggerNode.encoder, ws, string(key)}
		loggers.Insert(key, logger)
	}

}

// Debug logs a message at Debug level on the sugar logger.
func (l *Log) Debug(args ...interface{}) {
	l.stdLogger.Sugar().Debug(args...)
}

// Debugf logs a message at Debugf level on the sugar logger.
func (l *Log) Debugf(template string, args ...interface{}) {
	l.stdLogger.Sugar().Debugf(template, args...)
}

// Debugw logs a message at Debugw level on the sugar logger.
func (l *Log) Debugw(msg string, keysAndValues ...interface{}) {
	l.stdLogger.Sugar().Debugw(msg, keysAndValues...)
}

// Info logs a message at Info level on the sugar logger
func (l *Log) Info(args ...interface{}) {
	l.stdLogger.Sugar().Info(args...)
}

// Infof logs a message at Infof level on the sugar logger.
func (l *Log) Infof(template string, args ...interface{}) {
	l.stdLogger.Sugar().Infof(template, args...)
}

// Infow logs a message at Infow level on the sugar logger.
func (l *Log) Infow(msg string, keysAndValues ...interface{}) {
	l.stdLogger.Sugar().Infow(msg, keysAndValues...)
}

// Error logs a message at Error level on the sugar logger
func (l *Log) Error(args ...interface{}) {
	l.stdLogger.Sugar().Error(args...)
}

// Errorf logs a message at Errorf level on the sugar logger.
func (l *Log) Errorf(template string, args ...interface{}) {
	l.stdLogger.Sugar().Errorf(template, args...)
}

// Errorw logs a message at Errorw level on the sugar logger.
func (l *Log) Errorw(msg string, keysAndValues ...interface{}) {
	l.stdLogger.Sugar().Errorw(msg, keysAndValues...)
}

// Fatal logs a message at Fatal level on the sugar logger
func (l *Log) Fatal(args ...interface{}) {
	l.stdLogger.Sugar().Fatal(args...)
}

// Fatalf logs a message at Fatalf level on the sugar logger.
func (l *Log) Fatalf(template string, args ...interface{}) {
	l.stdLogger.Sugar().Fatalf(template, args)
}

// Fatalw logs a message at Fatalw level on the sugar logger.
func (l *Log) Fatalw(msg string, keysAndValues ...interface{}) {
	l.stdLogger.Sugar().Fatalw(msg, keysAndValues...)
}

// Panic logs a message at Panic level on the sugar logger
func (l *Log) Panic(args ...interface{}) {
	l.stdLogger.Sugar().Panic(args...)
}

// Panicf logs a message at Panicf level on the sugar logger.
func (l *Log) Panicf(template string, args ...interface{}) {
	l.stdLogger.Sugar().Panicf(template, args...)
}

// Panicw logs a message at Panicw level on the sugar logger.
func (l *Log) Panicw(msg string, keysAndValues ...interface{}) {
	l.stdLogger.Sugar().Panicw(msg, keysAndValues...)
}

// DPanic logs a message at DPanic level on the sugar logger
func (l *Log) DPanic(args ...interface{}) {
	l.stdLogger.Sugar().DPanic(args...)
}

// DPanicf logs a message at DPanicf level on the sugar logger.
func (l *Log) DPanicf(template string, args ...interface{}) {
	l.stdLogger.Sugar().DPanicf(template, args...)
}

// Panicw logs a message at DPanicw level on the sugar logger.
func (l *Log) DPanicw(msg string, keysAndValues ...interface{}) {
	l.stdLogger.Sugar().DPanicw(msg, keysAndValues...)
}

// Warn logs a message at Warn level on the sugar logger
func (l *Log) Warn(args ...interface{}) {
	l.stdLogger.Sugar().Warn(args...)
}

// Warnf logs a message at Warnf level on the sugar logger.
func (l *Log) Warnf(template string, args ...interface{}) {
	l.stdLogger.Sugar().Warnf(template, args...)
}

// Warnw logs a message at Warnw level on the sugar logger.
func (l *Log) Warnw(msg string, keysAndValues ...interface{}) {
	l.stdLogger.Sugar().Warnw(msg, keysAndValues...)
}

// GetKey returns a key object of radix tree
func GetKey(name string) art.Key {
	return art.Key(name)
}

func assignParentLevelLogger(name string) Log {
	parentNames := findParentsNames(name)
	for _, parentName := range parentNames {
		parentLogger, found := loggers.Search(art.Key(parentName))
		if found {
			logger := parentLogger.(Log)
			core := logger.stdLogger.Core()
			zapLogger := zap.New(core).Named(name)
			logger.stdLogger = zapLogger
			loggers.Insert(art.Key(name), logger)
			return logger
		}
	}
	root := GetDefaultLogger()
	core := root.stdLogger.Core()
	zapLogger := zap.New(core).Named(name)
	root.stdLogger = zapLogger
	loggers.Insert(art.Key(name), root)
	return root
}

// GetLogger gets a logger based on a give name
func GetLogger(names ...string) Log {
	name := buildTreeName(names...)
	value, found := GetLoggers().Search(GetKey(name))
	if found {
		return value.(Log)
	} else {
		return AddLogger(InfoLevel, names...)
	}
}

func (c *Configuration) GetLogger() Log {
	level := c.zapConfig.Level.Level().String()
	name := c.zapConfig.EncoderConfig.NameKey
	if level == "" {
		return assignParentLevelLogger(name)
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

	sinkURLs := c.GetSinkURLs()
	var urls []string
	if len(sinkURLs) > 0 {
		for _, url := range sinkURLs {
			urls = append(urls, url.String())

		}

		ws, _, err := zap.Open(urls...)
		if err != nil {
			return Log{}
		}

		cfg := c.zapConfig
		configLogger, _ := cfg.Build(zap.AddCallerSkip(1))
		encoder := zc.NewJSONEncoder(cfg.EncoderConfig)
		newLogger := configLogger.WithOptions(
			zap.WrapCore(
				func(zc.Core) zc.Core {
					return zc.NewCore(encoder, ws, &atomLevel)
				}))

		configLogger = newLogger.Named(name)
		logger := Log{configLogger, encoder, ws, name}
		loggers.Insert(art.Key(name), logger)
		return logger

	} else {

		cfg := c.zapConfig
		configLogger, _ := cfg.Build(zap.AddCallerSkip(1))
		encoder := zc.NewJSONEncoder(cfg.EncoderConfig)
		writer := zc.Lock(os.Stdout)
		newLogger := configLogger.WithOptions(
			zap.WrapCore(
				func(zc.Core) zc.Core {
					return zc.NewCore(encoder, writer, &atomLevel)
				}))

		configLogger = newLogger.Named(name)
		logger := Log{configLogger, encoder, writer, name}
		loggers.Insert(art.Key(name), logger)
		return logger
	}

}

// SetLevel change level of a logger and propagates the change to its children
func SetLevel(level Level, names ...string) Log {
	name := buildTreeName(names...)
	logger := Log{}
	loggers.ForEachPrefix(art.Key(name), func(node art.Node) bool {
		if node.Key() != nil {
			loggerNode := node.Value().(Log)
			newLevel := intToAtomicLevel(level)
			newLogger := loggerNode.stdLogger.WithOptions(
				zp.WrapCore(
					func(zc.Core) zc.Core {
						return zc.NewCore(loggerNode.encoder, loggerNode.writer, &newLevel)
					}))
			newLogger.Named(string(node.Key()))
			loggerNode.stdLogger = newLogger
			logger = Log{newLogger, loggerNode.encoder, loggerNode.writer, string(node.Key())}
			loggers.Insert(node.Key(), logger)
		}
		return true
	})
	return logger
}

// AddLogger adds a logger based on a given level and a hierarchy of names
func AddLogger(level Level, names ...string) Log {
	name := buildTreeName(names...)
	if level.String() == "" {
		return assignParentLevelLogger(name)
	}

	atomLevel := zap.AtomicLevel{}
	var internalLevel Level
	switch level {
	case InfoLevel:
		atomLevel = zap.NewAtomicLevelAt(zc.InfoLevel)
		internalLevel = InfoLevel
	case DebugLevel:
		atomLevel = zap.NewAtomicLevelAt(zc.DebugLevel)
		internalLevel = DebugLevel
	case ErrorLevel:
		atomLevel = zap.NewAtomicLevelAt(zc.ErrorLevel)
		internalLevel = ErrorLevel
	case PanicLevel:
		atomLevel = zap.NewAtomicLevelAt(zc.PanicLevel)
		internalLevel = PanicLevel
	case FatalLevel:
		atomLevel = zap.NewAtomicLevelAt(zc.FatalLevel)
		internalLevel = FatalLevel
	case WarnLevel:
		atomLevel = zap.NewAtomicLevelAt(zc.WarnLevel)
		internalLevel = WarnLevel
	}

	cfg := getDefaultConfig(name, internalLevel)
	configLogger, _ := cfg.GetZapConfig().Build(zap.AddCallerSkip(1))
	defaultEncoder := zc.NewJSONEncoder(cfg.GetZapConfig().EncoderConfig)
	defaultWriter := zc.Lock(os.Stdout)
	newLogger := configLogger.WithOptions(
		zap.WrapCore(
			func(zc.Core) zc.Core {
				return zc.NewCore(defaultEncoder, defaultWriter, &atomLevel)
			}))

	configLogger = newLogger.Named(name)
	logger := Log{configLogger, defaultEncoder, defaultWriter, name}
	loggers.Insert(art.Key(name), logger)
	return logger
}

// GetLoggers get loggers radix tree
func GetLoggers() art.Tree {
	return loggers
}

// GetDefaultLogger gets the default logger
func GetDefaultLogger() Log {
	return root
}
