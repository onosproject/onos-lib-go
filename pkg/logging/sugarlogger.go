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

// DPanicw logs a message at DPanicw level on the sugar logger.
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
