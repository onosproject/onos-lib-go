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
	"bytes"
	"net/url"
	"os"

	"github.com/onosproject/onos-lib-go/pkg/logging/config"

	art "github.com/plar/go-adaptive-radix-tree"

	"go.uber.org/zap"
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

// SinkInfo sink information
type SinkInfo struct {
	name  string
	_type string
	uri   string
	topic string
	key   string
}

// GetSinks get sinks info from the configuration
func GetSinks(config *config.Config) []SinkInfo {
	sinksList := config.Logging.Sinks
	sinks := make([]SinkInfo, len(sinksList))
	for _, sink := range sinksList {
		sinkInfo := SinkInfo{
			name:  sink.Name,
			_type: sink.Type,
			topic: sink.Topic,
			key:   sink.Key,
			uri:   sink.URI,
		}
		sinks = append(sinks, sinkInfo)
	}
	return sinks
}

func ContainSink(sinks []SinkInfo, sinkName string) (SinkInfo, bool) {
	for _, sink := range sinks {
		if sink.name == sinkName {
			return sink, true
		}
	}
	return SinkInfo{}, false
}

// AddConfiguredLoggers adds configured loggers
func AddConfiguredLoggers(config *config.Config) {
	loggersList := config.Logging.Loggers
	sinks := GetSinks(config)
	for _, logger := range loggersList {
		loggerSinkInfo, found := ContainSink(sinks, logger.Sink)
		if found {
			switch loggerSinkInfo._type {
			case Kafka.String():
				var urls []SinkURL
				var rawQuery bytes.Buffer
				if loggerSinkInfo.topic != "" {
					rawQuery.WriteString("topic=")
					rawQuery.WriteString(loggerSinkInfo.topic)
				}
				rawQuery.WriteString(",")
				if loggerSinkInfo.key != "" {
					rawQuery.WriteString("key=")
					rawQuery.WriteString(loggerSinkInfo.key)
				}
				urlInfo := SinkURL{
					URL: url.URL{Scheme: Kafka.String(), Host: loggerSinkInfo.uri, RawQuery: rawQuery.String()},
				}
				urls = append(urls, urlInfo)

				cfg := Configuration{}
				cfg.SetEncoding(logger.Encoding).
					SetLevel(StringToInt(logger.Level)).
					SetSinkURLs(urls).
					SetName(logger.Name).
					SetECMsgKey("msg").
					SetECLevelKey("level").
					SetECTimeKey("ts").
					SetECTimeEncoder(zc.ISO8601TimeEncoder).
					SetECEncodeLevel(zc.CapitalLevelEncoder).
					Build()
				cfg.GetLogger()

			case Stdout.String():
				cfg := Configuration{}
				cfg.SetEncoding(logger.Encoding).
					SetLevel(StringToInt(logger.Level)).
					SetOutputPaths([]string{Stdout.String()}).
					SetName(logger.Name).
					SetECMsgKey("msg").
					SetECLevelKey("level").
					SetECTimeKey("ts").
					SetECTimeEncoder(zc.ISO8601TimeEncoder).
					SetECEncodeLevel(zc.CapitalLevelEncoder).
					Build()
				cfg.GetLogger()
			}

		}
	}
}

// init initialize logger package data structures
func init() {
	dbg = true
	loggers = art.New()

	// Adds default logger (i.e. root logger)
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
	/*err := zap.RegisterSink("kafka", InitSink)
	if err != nil {
		fmt.Println("Kafka Sink cannot be registered", err)
	}*/
	root = Log{rootLogger, &defaultEncoder, &defaultWriter, defaultLoggerName, defaultAtomLevel}
	dbg = false

	err := zap.RegisterSink("kafka", InitSink)
	if err != nil {
		dbg.Println("Kafka Sink cannot be registered %s", err)
	}
	loggersConfig := config.GetConfig()
	AddConfiguredLoggers(loggersConfig)

}

// SetLevel defines a new logger level and propagate the change its children
func (l *Log) SetLevel(level Level) {
	newLevel := intToAtomicLevel(level)
	l.level.SetLevel(newLevel.Level())
	dbg.Println("Logger %s writer, encoder, stdlogger memory addresses are 0X%08x, 0X%08x, 0X%08x", l.name, &l.writer, &l.encoder, &l.stdLogger)
	loggers.ForEachPrefix(art.Key(l.name), func(node art.Node) bool {
		if node.Key() != nil {
			loggerNode := node.Value().(*Log)
			loggerNode.level.SetLevel(newLevel.Level())
		}
		return true
	})
}

// GetKey returns a key object of radix tree
func GetKey(name string) art.Key {
	return art.Key(name)
}

func assignParentLevelLogger(name string) *Log {
	parentNames := findParentsNames(name)
	for _, parentName := range parentNames {
		parentLogger, found := loggers.Search(art.Key(parentName))
		if found {
			logger := parentLogger.(Log)
			core := logger.stdLogger.Core()
			zapLogger := zap.New(core).Named(name)
			logger.stdLogger = zapLogger
			loggers.Insert(art.Key(name), &logger)
			return &logger
		}
	}
	root := GetDefaultLogger()
	core := root.stdLogger.Core()
	zapLogger := zap.New(core).Named(name)
	root.stdLogger = zapLogger
	loggers.Insert(art.Key(name), &root)
	return &root
}

func FindLogger(names ...string) (*Log, bool) {
	name := buildTreeName(names...)
	value, found := GetLoggers().Search(GetKey(name))
	if found {
		dbg.Println("found logger: %s", names)
		return value.(*Log), found
	}
	return &Log{}, found
}

// GetLogger gets a logger based on a give name
func GetLogger(names ...string) *Log {
	name := buildTreeName(names...)
	value, found := GetLoggers().Search(GetKey(name))
	if found {
		dbg.Println("found logger: %s", names)
		return value.(*Log)
	} else {
		dbg.Println("not found logger: %s, creating it", name)
		return AddLogger(InfoLevel, names...)
	}
}

func (c *Configuration) GetLogger() *Log {
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
			return &Log{}
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
		logger := Log{configLogger, &encoder, &ws, name, atomLevel}
		loggers.Insert(art.Key(name), &logger)
		return &logger

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
		logger := Log{configLogger, &encoder, &writer, name, atomLevel}
		loggers.Insert(art.Key(name), &logger)
		return &logger
	}

}

// SetLevel change level of a logger and propagates the change to its children
func SetLevel(level Level, names ...string) {
	name := buildTreeName(names...)
	loggers.ForEachPrefix(art.Key(name), func(node art.Node) bool {
		if node.Key() != nil {
			loggerNode := node.Value().(*Log)
			newLevel := intToAtomicLevel(level)
			loggerNode.level.SetLevel(newLevel.Level())
		}
		return true
	})
}

// AddLogger adds a logger based on a given level and a hierarchy of names
func AddLogger(level Level, names ...string) *Log {
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
	logger := Log{configLogger, &defaultEncoder, &defaultWriter, name, atomLevel}
	loggers.Insert(art.Key(name), &logger)
	return &logger
}

// GetLoggers get loggers radix tree
func GetLoggers() art.Tree {
	return loggers
}

// GetDefaultLogger gets the default logger
func GetDefaultLogger() Log {
	return root
}
