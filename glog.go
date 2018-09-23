// Package glog implements almost [glog](https://github.com/golang/glog) package API.
// The aim of this package is to replace all the glog calls for the ones to a custom
// logger that satisfies this package Logger interface. This way any code using glog
// global logger will use the desired logger.
package glog

import (
	"fmt"
	"os"
	"sync"
)

// Logger is the interface that a glog logger wrapper
// needs to implement with the aim of replacing the logger that
// comes with glog.
type Logger interface {
	// Debugf logs with debug.
	Debugf(format string, args ...interface{})
	// Infof logs with info.
	Infof(format string, args ...interface{})
	// Warnf logs with warning.
	Warnf(format string, args ...interface{})
	// Errorf logs with error.
	Errorf(format string, args ...interface{})
}

// logger is the global logger that will intercept the glog calls
// instead of calling glog it will use this logger.
// By default it will be disabled for security reasons.
var logger Logger = Dummy

var mu = sync.Mutex{}

// SetLogger sets the glog replacement logger.
func SetLogger(l Logger) {
	mu.Lock()
	defer mu.Unlock()
	logger = l
}

// Level replacement for glog.
type Level int32

// Verbose replacement for glog.
type Verbose bool

// Flush replacement for glog.
func Flush() {}

// V replacement for glog.
func V(_ Level) Verbose {
	return Verbose(false) // Doesn't mind.
}

// Info replacement for glog.
func (Verbose) Info(args ...interface{}) {
	s := fmt.Sprint(args...)
	logger.Debugf(s)
}

// Infoln replacement for glog.
func (v Verbose) Infoln(args ...interface{}) {
	args = append(args, "\n")
	v.Info(args...)
}

// Infof replacement for glog.
func (Verbose) Infof(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

// Info replacement for glog.
func Info(args ...interface{}) {
	s := fmt.Sprint(args...)
	logger.Infof(s)
}

// InfoDepth replacement for glog.
func InfoDepth(_ int, args ...interface{}) {
	Info(args...)
}

// Infoln replacement for glog.
func Infoln(args ...interface{}) {
	args = append(args, "\n")
	Info(args...)
}

// Infof replacement for glog.
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Warning replacement for glog.
func Warning(args ...interface{}) {
	s := fmt.Sprint(args...)
	logger.Warnf(s)
}

// WarningDepth replacement for glog.
func WarningDepth(_ int, args ...interface{}) {
	Warning(args...)
}

// Warningln replacement for glog.
func Warningln(args ...interface{}) {
	args = append(args, "\n")
	Warning(args...)
}

// Warningf replacement for glog.
func Warningf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

// Error replacement for glog.
func Error(args ...interface{}) {
	s := fmt.Sprint(args...)
	logger.Errorf(s)
}

// ErrorDepth replacement for glog.
func ErrorDepth(_ int, args ...interface{}) {
	Error(args...)
}

// Errorln replacement for glog.
func Errorln(args ...interface{}) {
	args = append(args, "\n")
	Error(args...)
}

// Errorf replacement for glog.
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

// Fatal replacement for glog.
func Fatal(args ...interface{}) {
	Error(args...)
	os.Exit(255)
}

// FatalDepth replacement for glog.
func FatalDepth(depth int, args ...interface{}) {
	ErrorDepth(depth, args...)
	os.Exit(255)
}

// Fatalln replacement for glog.
func Fatalln(args ...interface{}) {
	Errorln(args...)
	os.Exit(255)
}

// Fatalf replacement for glog.
func Fatalf(format string, args ...interface{}) {
	Errorf(format, args...)
	os.Exit(255)
}

// Exit replacement for glog.
func Exit(args ...interface{}) {
	Error(args...)
	os.Exit(1)
}

// ExitDepth replacement for glog.
func ExitDepth(depth int, args ...interface{}) {
	ErrorDepth(depth, args...)
	os.Exit(1)
}

// Exitln replacement for glog.
func Exitln(args ...interface{}) {
	Errorln(args...)
	os.Exit(1)
}

// Exitf replacement for glog.
func Exitf(format string, args ...interface{}) {
	Errorf(format, args...)
	os.Exit(1)
}

// LoggerFunc is a logger that satisfies Logger interface and
// gets the Logger interface required methods as func fields.
// Basically is a helper to not create a custom type.
type LoggerFunc struct {
	// DebugfFunc is a function for DebugfFunc `Logger.Debugf(format string, args ...interface{})`.
	DebugfFunc func(format string, args ...interface{})
	// InfofFunc is a function for InfofFunc `Logger.Infof(format string, args ...interface{})`.
	InfofFunc func(format string, args ...interface{})
	// WarnfFunc is a function for WarnfFunc `Logger.Warnf(format string, args ...interface{})`.
	WarnfFunc func(format string, args ...interface{})
	// ErrorfFunc is a function for ErrorfFunc `Logger.Errorf(format string, args ...interface{})`.
	ErrorfFunc func(format string, args ...interface{})
}

// Debugf satisfies Logger interface.
func (l *LoggerFunc) Debugf(format string, args ...interface{}) { l.DebugfFunc(format, args...) }

// Infof satisfies Logger interface.
func (l *LoggerFunc) Infof(format string, args ...interface{}) { l.InfofFunc(format, args...) }

// Warnf satisfies Logger interface.
func (l *LoggerFunc) Warnf(format string, args ...interface{}) { l.WarnfFunc(format, args...) }

// Errorf satisfies Logger interface.
func (l *LoggerFunc) Errorf(format string, args ...interface{}) { l.ErrorfFunc(format, args...) }

// Dummy is a dummy logger useful for tests and disabling logging.
var Dummy = &LoggerFunc{
	DebugfFunc: func(format string, args ...interface{}) {},
	InfofFunc:  func(format string, args ...interface{}) {},
	WarnfFunc:  func(format string, args ...interface{}) {},
	ErrorfFunc: func(format string, args ...interface{}) {},
}
