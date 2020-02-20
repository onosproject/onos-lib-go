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

	zp "go.uber.org/zap"
	zc "go.uber.org/zap/zapcore"
)

func intToAtomicLevel(l Level) zp.AtomicLevel {
	switch l {
	case DebugLevel:
		return zp.NewAtomicLevelAt(zc.DebugLevel)
	case InfoLevel:
		return zp.NewAtomicLevelAt(zc.InfoLevel)
	case WarnLevel:
		return zp.NewAtomicLevelAt(zc.WarnLevel)
	case ErrorLevel:
		return zp.NewAtomicLevelAt(zc.ErrorLevel)
	case FatalLevel:
		return zp.NewAtomicLevelAt(zc.FatalLevel)
	case PanicLevel:
		return zp.NewAtomicLevelAt(zc.PanicLevel)
	case DPanicLevel:
		return zp.NewAtomicLevelAt(zc.DPanicLevel)
	}
	return zp.NewAtomicLevelAt(zc.ErrorLevel)
}

func levelToInt(l zc.Level) Level {
	switch l {
	case zc.DebugLevel:
		return DebugLevel
	case zc.InfoLevel:
		return InfoLevel
	case zc.WarnLevel:
		return WarnLevel
	case zc.ErrorLevel:
		return ErrorLevel
	case zc.FatalLevel:
		return FatalLevel
	case zc.PanicLevel:
		return PanicLevel
	case zc.DPanicLevel:
		return DPanicLevel

	}
	return ErrorLevel
}

func StringToInt(l string) Level {
	switch l {
	case DebugLevel.String():
		return DebugLevel
	case InfoLevel.String():
		return InfoLevel
	case WarnLevel.String():
		return WarnLevel
	case ErrorLevel.String():
		return ErrorLevel
	case FatalLevel.String():
		return FatalLevel
	case PanicLevel.String():
		return PanicLevel
	case DPanicLevel.String():
		return DPanicLevel
	}
	return ErrorLevel
}

// Errors concatenates multiple error into one error buf
type Errors []error

func (e Errors) Error() string {
	var errBuf bytes.Buffer
	for _, err := range e {
		errBuf.WriteString(err.Error())
		errBuf.WriteByte('\n')
	}
	return errBuf.String()
}
