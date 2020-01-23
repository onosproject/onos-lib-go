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
	// DPanicLevel
	DPanicLevel
)

func (l Level) String() string {
	return [...]string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL", "PANIC", "DPANIC"}[l]
}

// CONSOLE formats the log for the console, mostly used during development
const CONSOLE = "console"

// JSON formats the log using json format, mostly used by an automated logging system consumption
const JSON = "json"

// Logger represents an abstract logging interface.  Any logging implementation used
// will need to abide by this interface
type Logger interface {
	Debug(...interface{})
	Debugln(...interface{})
	Debugf(string, ...interface{})
	Debugw(string, Fields)

	Info(...interface{})
	Infoln(...interface{})
	Infof(string, ...interface{})
	Infow(string, Fields)

	Warn(...interface{})
	Warnln(...interface{})
	Warnf(string, ...interface{})
	Warnw(string, Fields)

	Error(...interface{})
	Errorln(...interface{})
	Errorf(string, ...interface{})
	Errorw(string, Fields)

	Fatal(...interface{})
	Fatalln(...interface{})
	Fatalf(string, ...interface{})
	Fatalw(string, Fields)

	Panic(...interface{})
	Panicln(...interface{})
	Panicf(string, ...interface{})
	Panicw(string, Fields)

	DPanic(...interface{})
	DPanicln(...interface{})
	DPanicf(string, ...interface{})
	DPanicw(string, Fields)

	With(Fields) Logger

	// The following are added to be able to use this logger as a gRPC LoggerV2 if needed
	//
	Warning(...interface{})
	Warningln(...interface{})
	Warningf(string, ...interface{})

	// V reports whether verbosity level l is at least the requested verbose level.
	V(l Level) bool

	//Returns the log level of this specific logger
	GetLogLevel() Level
}

// Fields is used as key-value pairs for structured logging
type Fields map[string]interface{}

var defaultLogger *logger
var cfg zp.Config

var loggers map[string]*logger
var cfgs map[string]zp.Config

type logger struct {
	log         *zp.SugaredLogger
	parent      *zp.Logger
	packageName string
}
