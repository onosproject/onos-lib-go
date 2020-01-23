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
	"path"
	"runtime"
	"strings"

	zp "go.uber.org/zap"
	zc "go.uber.org/zap/zapcore"
)

func getDefaultConfig(outputType string, level Level, defaultFields Fields) zp.Config {
	return zp.Config{
		Level:            intToAtomicLevel(level),
		Encoding:         outputType,
		Development:      true,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		InitialFields:    defaultFields,
		EncoderConfig: zc.EncoderConfig{
			LevelKey:       "level",
			MessageKey:     "msg",
			TimeKey:        "ts",
			StacktraceKey:  "stacktrace",
			LineEnding:     zc.DefaultLineEnding,
			EncodeLevel:    zc.LowercaseLevelEncoder,
			EncodeTime:     zc.ISO8601TimeEncoder,
			EncodeDuration: zc.SecondsDurationEncoder,
			EncodeCaller:   zc.ShortCallerEncoder,
		},
	}
}

// SetLogger needs to be invoked before the logger API can be invoked.  This function
// initialize the default logger (zap's sugaredlogger)
func SetDefaultLogger(outputType string, level Level, defaultFields Fields) (Logger, error) {
	// Build a custom config using zap
	cfg = getDefaultConfig(outputType, level, defaultFields)

	l, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	defaultLogger = &logger{
		log:    l.Sugar(),
		parent: l,
	}

	return defaultLogger, nil
}

// AddPackage registers a package to the log map.  Each package gets its own logger which allows
// its config (loglevel) to be changed dynamically without interacting with the other packages.
// outputType is JSON, level is the lowest level log to output with this logger and defaultFields is a map of
// key-value pairs to always add to the output.
// Note: AddPackage also returns a reference to the actual logger.  If a calling package uses this reference directly
//instead of using the publicly available functions in this log package then a number of functionalities will not
// be available to it, notably log tracing with filename.functionname.linenumber annotation.
//
// pkgNames parameter should be used for testing only as this function detects the caller's package.
func AddPackage(outputType string, level Level, defaultFields Fields, pkgNames ...string) (Logger, error) {
	if cfgs == nil {
		cfgs = make(map[string]zp.Config)
	}
	if loggers == nil {
		loggers = make(map[string]*logger)
	}

	var pkgName string
	for _, name := range pkgNames {
		pkgName = name
		break
	}
	if pkgName == "" {
		pkgName, _, _, _ = getCallerInfo()
	}

	if _, exist := loggers[pkgName]; exist {
		return loggers[pkgName], nil
	}

	cfgs[pkgName] = getDefaultConfig(outputType, level, defaultFields)

	l, err := cfgs[pkgName].Build()
	if err != nil {
		return nil, err
	}

	loggers[pkgName] = &logger{
		log:         l.Sugar(),
		parent:      l,
		packageName: pkgName,
	}
	return loggers[pkgName], nil
}

//UpdateAllLoggers create new loggers for all registered pacakges with the defaultFields.
func UpdateAllLoggers(defaultFields Fields) error {
	for pkgName, cfg := range cfgs {
		for k, v := range defaultFields {
			if cfg.InitialFields == nil {
				cfg.InitialFields = make(map[string]interface{})
			}
			cfg.InitialFields[k] = v
		}
		l, err := cfg.Build()
		if err != nil {
			return err
		}

		loggers[pkgName] = &logger{
			log:         l.Sugar(),
			parent:      l,
			packageName: pkgName,
		}
	}
	return nil
}

// GetPackageNames Return a list of all packages that have individually-configured loggers
func GetPackageNames() []string {
	i := 0
	keys := make([]string, len(loggers))
	for k := range loggers {
		keys[i] = k
		i++
	}
	return keys
}

// UpdateLogger deletes the logger associated with a caller's package and creates a new logger with the
// defaultFields.  If a calling package is holding on to a Logger reference obtained from AddPackage invocation, then
// that package needs to invoke UpdateLogger if it needs to make changes to the default fields and obtain a new logger
// reference
func UpdateLogger(defaultFields Fields) (Logger, error) {
	pkgName, _, _, _ := getCallerInfo()
	if _, exist := loggers[pkgName]; !exist {
		return nil, fmt.Errorf("package-%s-not-registered", pkgName)
	}

	// Build a new logger
	if _, exist := cfgs[pkgName]; !exist {
		return nil, fmt.Errorf("config-%s-not-registered", pkgName)
	}

	cfg := cfgs[pkgName]
	for k, v := range defaultFields {
		if cfg.InitialFields == nil {
			cfg.InitialFields = make(map[string]interface{})
		}
		cfg.InitialFields[k] = v
	}
	l, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	// Set the logger
	loggers[pkgName] = &logger{
		log:         l.Sugar(),
		parent:      l,
		packageName: pkgName,
	}
	return loggers[pkgName], nil
}

func setLevel(cfg zp.Config, level Level) {
	switch level {
	case DebugLevel:
		cfg.Level.SetLevel(zc.DebugLevel)
	case InfoLevel:
		cfg.Level.SetLevel(zc.InfoLevel)
	case WarnLevel:
		cfg.Level.SetLevel(zc.WarnLevel)
	case ErrorLevel:
		cfg.Level.SetLevel(zc.ErrorLevel)
	case FatalLevel:
		cfg.Level.SetLevel(zc.FatalLevel)
	default:
		cfg.Level.SetLevel(zc.ErrorLevel)
	}
}

//SetPackageLogLevel dynamically sets the log level of a given package to level.  This is typically invoked at an
// application level during debugging
func SetPackageLogLevel(packageName string, level Level) {
	// Get proper config
	if cfg, ok := cfgs[packageName]; ok {
		setLevel(cfg, level)
	}
}

//SetAllLogLevel sets the log level of all registered packages to level
func SetAllLogLevel(level Level) {
	// Get proper config
	for _, cfg := range cfgs {
		setLevel(cfg, level)
	}
}

//GetPackageLogLevel returns the current log level of a package.
func GetPackageLogLevel(packageName ...string) (Level, error) {
	var name string
	if len(packageName) == 1 {
		name = packageName[0]
	} else {
		name, _, _, _ = getCallerInfo()
	}
	if cfg, ok := cfgs[name]; ok {
		return levelToInt(cfg.Level.Level()), nil
	}
	return 0, fmt.Errorf("unknown-package-%s", name)
}

//GetDefaultLogLevel gets the log level used for packages that don't have specific loggers
func GetDefaultLogLevel() Level {
	return levelToInt(cfg.Level.Level())
}

//SetLogLevel sets the log level for the logger corresponding to the caller's package
func SetLogLevel(level Level) error {
	pkgName, _, _, _ := getCallerInfo()
	if _, exist := cfgs[pkgName]; !exist {
		return fmt.Errorf("unregistered-package-%s", pkgName)
	}
	cfg := cfgs[pkgName]
	setLevel(cfg, level)
	return nil
}

//SetDefaultLogLevel sets the log level used for packages that don't have specific loggers
func SetDefaultLogLevel(level Level) {
	setLevel(cfg, level)
}

// CleanUp flushed any buffered log entries. Applications should take care to call
// CleanUp before exiting.
func CleanUp() error {
	for _, logger := range loggers {
		if logger != nil {
			if logger.parent != nil {
				if err := logger.parent.Sync(); err != nil {
					return err
				}
			}
		}
	}
	if defaultLogger != nil {
		if defaultLogger.parent != nil {
			if err := defaultLogger.parent.Sync(); err != nil {
				return err
			}
		}
	}
	return nil
}

func getCallerInfo() (string, string, string, int) {
	// Since the caller of a log function is one stack frame before (in terms of stack higher level) the log.go
	// filename, then first look for the last log.go filename and then grab the caller info one level higher.
	maxLevel := 3
	skiplevel := 3 // Level with the most empirical success to see the last log.go stack frame.
	pc := make([]uintptr, maxLevel)
	n := runtime.Callers(skiplevel, pc)
	packageName := ""
	funcName := ""
	fileName := ""
	var line int
	if n == 0 {
		return packageName, fileName, funcName, line
	}
	frames := runtime.CallersFrames(pc[:n])
	var frame runtime.Frame
	var foundFrame runtime.Frame
	more := true
	for more {
		frame, more = frames.Next()
		_, fileName = path.Split(frame.File)
		if fileName != "log.go" {
			foundFrame = frame // First frame after log.go in the frame stack
			break
		}
	}
	parts := strings.Split(foundFrame.Function, ".")
	pl := len(parts)
	if pl >= 2 {
		funcName = parts[pl-1]
		if parts[pl-2][0] == '(' {
			packageName = strings.Join(parts[0:pl-2], ".")
		} else {
			packageName = strings.Join(parts[0:pl-1], ".")
		}
	}

	if strings.HasSuffix(packageName, ".init") {
		packageName = strings.TrimSuffix(packageName, ".init")
	}

	if strings.HasSuffix(fileName, ".go") {
		fileName = strings.TrimSuffix(fileName, ".go")
	}

	return packageName, fileName, funcName, foundFrame.Line
}

func getPackageLevelSugaredLogger() *zp.SugaredLogger {
	pkgName, fileName, funcName, line := getCallerInfo()
	if _, exist := loggers[pkgName]; exist {
		return loggers[pkgName].log.With("caller", fmt.Sprintf("%s.%s:%d", fileName, funcName, line))
	}
	return defaultLogger.log.With("caller", fmt.Sprintf("%s.%s:%d", fileName, funcName, line))
}

func getPackageLevelLogger() Logger {
	pkgName, _, _, _ := getCallerInfo()
	if _, exist := loggers[pkgName]; exist {
		return loggers[pkgName]
	}
	return defaultLogger
}

func serializeMap(fields Fields) []interface{} {
	data := make([]interface{}, len(fields)*2)
	i := 0
	for k, v := range fields {
		data[i] = k
		data[i+1] = v
		i = i + 2
	}
	return data
}

// With returns a logger initialized with the key-value pairs
func (l logger) With(keysAndValues Fields) Logger {
	return logger{log: l.log.With(serializeMap(keysAndValues)...), parent: l.parent}
}

// Debug logs a message at level Debug on the standard logger.
func (l logger) Debug(args ...interface{}) {
	l.log.Debug(args...)
}

// Debugln logs a message at level Debug on the standard logger with a line feed. Default in any case.
func (l logger) Debugln(args ...interface{}) {
	l.log.Debug(args...)
}

// Debugw logs a message at level Debug on the standard logger.
func (l logger) Debugf(format string, args ...interface{}) {
	l.log.Debugf(format, args...)
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (l logger) Debugw(msg string, keysAndValues Fields) {
	l.log.Debugw(msg, serializeMap(keysAndValues)...)
}

// Info logs a message at level Info on the standard logger.
func (l logger) Info(args ...interface{}) {
	l.log.Info(args...)
}

// Infoln logs a message at level Info on the standard logger with a line feed. Default in any case.
func (l logger) Infoln(args ...interface{}) {
	l.log.Info(args...)
	//msg := fmt.Sprintln(args...)
	//l.sourced().Info(msg[:len(msg)-1])
}

// Infof logs a message at level Info on the standard logger.
func (l logger) Infof(format string, args ...interface{}) {
	l.log.Infof(format, args...)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (l logger) Infow(msg string, keysAndValues Fields) {
	l.log.Infow(msg, serializeMap(keysAndValues)...)
}

// Warn logs a message at level Warn on the standard logger.
func (l logger) Warn(args ...interface{}) {
	l.log.Warn(args...)
}

// Warnln logs a message at level Warn on the standard logger with a line feed. Default in any case.
func (l logger) Warnln(args ...interface{}) {
	l.log.Warn(args...)
}

// Warnf logs a message at level Warn on the standard logger.
func (l logger) Warnf(format string, args ...interface{}) {
	l.log.Warnf(format, args...)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (l logger) Warnw(msg string, keysAndValues Fields) {
	l.log.Warnw(msg, serializeMap(keysAndValues)...)
}

// Error logs a message at level Error on the standard logger.
func (l logger) Error(args ...interface{}) {
	l.log.Error(args...)
}

// Errorln logs a message at level Error on the standard logger with a line feed. Default in any case.
func (l logger) Errorln(args ...interface{}) {
	l.log.Error(args...)
}

// Errorf logs a message at level Error on the standard logger.
func (l logger) Errorf(format string, args ...interface{}) {
	l.log.Errorf(format, args...)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (l logger) Errorw(msg string, keysAndValues Fields) {
	l.log.Errorw(msg, serializeMap(keysAndValues)...)
}

// Fatal logs a message at level Fatal on the standard logger.
func (l logger) Fatal(args ...interface{}) {
	l.log.Fatal(args...)
}

// Fatalln logs a message at level Fatal on the standard logger with a line feed. Default in any case.
func (l logger) Fatalln(args ...interface{}) {
	l.log.Fatal(args...)
}

// Fatalf logs a message at level Fatal on the standard logger.
func (l logger) Fatalf(format string, args ...interface{}) {
	l.log.Fatalf(format, args...)
}

// Fatalw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (l logger) Fatalw(msg string, keysAndValues Fields) {
	l.log.Fatalw(msg, serializeMap(keysAndValues)...)
}

// Panic logs a message at level Panic on the standard logger.
func (l logger) Panic(args ...interface{}) {
	l.log.Panic(args...)
}

// Panicln logs a message at level Panic on the standard logger with a line feed. Default in any case.
func (l logger) Panicln(args ...interface{}) {
	l.log.Panic(args...)
}

// Panicf logs a message at level Panic on the standard logger.
func (l logger) Panicf(format string, args ...interface{}) {
	l.log.Panicf(format, args...)
}

// Panicw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (l logger) Panicw(msg string, keysAndValues Fields) {
	l.log.Panicw(msg, serializeMap(keysAndValues)...)
}

// DPanic logs a message at level DPanic on the standard logger.
func (l logger) DPanic(args ...interface{}) {
	l.log.DPanic(args...)
}

// DPanicln logs a message at level DPanic on the standard logger with a line feed. Default in any case.
func (l logger) DPanicln(args ...interface{}) {
	l.log.DPanic(args...)
}

// DPanicf logs a message at level DPanic on the standard logger.
func (l logger) DPanicf(format string, args ...interface{}) {
	l.log.DPanicf(format, args...)
}

// DPanicw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (l logger) DPanicw(msg string, keysAndValues Fields) {
	l.log.DPanicw(msg, serializeMap(keysAndValues)...)
}

// Warning logs a message at level Warn on the standard logger.
func (l logger) Warning(args ...interface{}) {
	l.log.Warn(args...)
}

// Warningln logs a message at level Warn on the standard logger with a line feed. Default in any case.
func (l logger) Warningln(args ...interface{}) {
	l.log.Warn(args...)
}

// Warningf logs a message at level Warn on the standard logger.
func (l logger) Warningf(format string, args ...interface{}) {
	l.log.Warnf(format, args...)
}

// V reports whether verbosity level l is at least the requested verbose level.
func (l logger) V(level Level) bool {
	return l.parent.Core().Enabled(intToLevel(level))
}

// GetLogLevel returns the current level of the logger
func (l logger) GetLogLevel() Level {
	return levelToInt(cfgs[l.packageName].Level.Level())
}

// With returns a logger initialized with the key-value pairs
func With(keysAndValues Fields) Logger {
	return logger{log: getPackageLevelSugaredLogger().With(serializeMap(keysAndValues)...), parent: defaultLogger.parent}
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	getPackageLevelSugaredLogger().Debug(args...)
}

// Debugln logs a message at level Debug on the standard logger.
func Debugln(args ...interface{}) {
	getPackageLevelSugaredLogger().Debug(args...)
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {
	getPackageLevelSugaredLogger().Debugf(format, args...)
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Debugw(msg string, keysAndValues Fields) {
	getPackageLevelSugaredLogger().Debugw(msg, serializeMap(keysAndValues)...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	getPackageLevelSugaredLogger().Info(args...)
}

// Infoln logs a message at level Info on the standard logger.
func Infoln(args ...interface{}) {
	getPackageLevelSugaredLogger().Info(args...)
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
	getPackageLevelSugaredLogger().Infof(format, args...)
}

//Infow logs a message with some additional context. The variadic key-value
//pairs are treated as they are in With.
func Infow(msg string, keysAndValues Fields) {
	getPackageLevelSugaredLogger().Infow(msg, serializeMap(keysAndValues)...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	getPackageLevelSugaredLogger().Warn(args...)
}

// Warnln logs a message at level Warn on the standard logger.
func Warnln(args ...interface{}) {
	getPackageLevelSugaredLogger().Warn(args...)
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {
	getPackageLevelSugaredLogger().Warnf(format, args...)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Warnw(msg string, keysAndValues Fields) {
	getPackageLevelSugaredLogger().Warnw(msg, serializeMap(keysAndValues)...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	getPackageLevelSugaredLogger().Error(args...)
}

// Errorln logs a message at level Error on the standard logger.
func Errorln(args ...interface{}) {
	getPackageLevelSugaredLogger().Error(args...)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
	getPackageLevelSugaredLogger().Errorf(format, args...)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Errorw(msg string, keysAndValues Fields) {
	getPackageLevelSugaredLogger().Errorw(msg, serializeMap(keysAndValues)...)
}

// Fatal logs a message at level Fatal on the standard logger.
func Fatal(args ...interface{}) {
	getPackageLevelSugaredLogger().Fatal(args...)
}

// Fatalln logs a message at level Fatal on the standard logger.
func Fatalln(args ...interface{}) {
	getPackageLevelSugaredLogger().Fatal(args...)
}

// Fatalf logs a message at level Fatal on the standard logger.
func Fatalf(format string, args ...interface{}) {
	getPackageLevelSugaredLogger().Fatalf(format, args...)
}

// Fatalw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Fatalw(msg string, keysAndValues Fields) {
	getPackageLevelSugaredLogger().Fatalw(msg, serializeMap(keysAndValues)...)
}

// Warning logs a message at level Warn on the standard logger.
func Warning(args ...interface{}) {
	getPackageLevelSugaredLogger().Warn(args...)
}

// Warningln logs a message at level Warn on the standard logger.
func Warningln(args ...interface{}) {
	getPackageLevelSugaredLogger().Warn(args...)
}

// Warningf logs a message at level Warn on the standard logger.
func Warningf(format string, args ...interface{}) {
	getPackageLevelSugaredLogger().Warnf(format, args...)
}

// V reports whether verbosity level l is at least the requested verbose level.
func V(level Level) bool {
	return getPackageLevelLogger().V(level)
}

//GetLogLevel returns the log level of the invoking package
func GetLogLevel() Level {
	return getPackageLevelLogger().GetLogLevel()
}
