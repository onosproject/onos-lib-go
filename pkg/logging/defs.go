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
	"net/url"

	art "github.com/plar/go-adaptive-radix-tree"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type SinkType int

const (
	// Kafka kafka sink type
	Kafka SinkType = iota

	// Stdout sink type
	Stdout
)

func (s SinkType) String() string {
	return [...]string{"kafka", "stdout"}[s]
}

type Level int

const (
	// DebugLevel logs a message at debug level
	DebugLevel Level = iota
	// InfoLevel logs a message at info level
	InfoLevel
	// WarnLevel logs a message at warning level
	WarnLevel
	// ErrorLevel logs a message at error level
	ErrorLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel
	// PanicLevel logs a message, then panics.
	PanicLevel
	// DPanicLevel logs at PanicLevel; otherwise, it logs at ErrorLevel
	DPanicLevel

	EmptyLevel
)

func (l Level) String() string {
	return [...]string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL", "PANIC", "DPANIC", ""}[l]
}

// Logger represents an abstract logging interface.
type Logger interface {
	Debug(...interface{})
	Debugf(template string, args ...interface{})
	Debugw(msg string, keysAndValues ...interface{})

	Info(...interface{})
	Infof(template string, args ...interface{})
	Infow(msg string, keysAndValues ...interface{})

	Error(...interface{})
	Errorf(template string, args ...interface{})
	Errorw(msg string, keysAndValues ...interface{})

	Fatal(...interface{})
	Fatalf(template string, args ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})

	Panic(...interface{})
	Panicf(template string, args ...interface{})
	Panicw(msg string, keysAndValues ...interface{})

	DPanic(...interface{})
	DPanicf(template string, args ...interface{})
	DPanicw(msg string, keysAndValues ...interface{})

	Warn(...interface{})
	Warnf(template string, args ...interface{})
	Warnw(msg string, keysAndValues ...interface{})

	SetLevel(level Level)
}

type Log struct {
	stdLogger *zap.Logger
	encoder   *zapcore.Encoder
	writer    *zapcore.WriteSyncer
	name      string
	level     zap.AtomicLevel
}

var loggers art.Tree
var root Log

// SinkURL sink url
type SinkURL struct {
	url.URL
}

var dbg Debug
