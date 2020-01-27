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
	art "github.com/plar/go-adaptive-radix-tree"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

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
	Info(...interface{})
	Error(...interface{})
	Fatal(...interface{})
	Panic(...interface{})
	DPanic(...interface{})
	Warn(...interface{})
}

type Log struct {
	stdLogger *zap.Logger
	encoder   zapcore.Encoder
	writer    zapcore.WriteSyncer
}

var loggers art.Tree
var root Log
