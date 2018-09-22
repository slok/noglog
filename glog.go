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
	// DebugEnabled returns if the Debug level of the logger is enabled.
	DebugEnabled() bool
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
	return Verbose(logger.DebugEnabled())
}

// Info replacement for glog.
func (v Verbose) Info(args ...interface{}) {
	if v {
		s := fmt.Sprint(args...)
		logger.Debugf(s)
	}
}

// Infoln replacement for glog.
func (v Verbose) Infoln(args ...interface{}) {
	args = append(args, "\n")
	v.Info(args...)
}

// Infof replacement for glog.
func (v Verbose) Infof(format string, args ...interface{}) {
	if v {
		logger.Debugf(format, args...)
	}
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

// Dummy is a dummy logger useful for tests and disabling logging.
var Dummy = &dummy{}

type dummy struct{}

func (*dummy) DebugEnabled() bool                        { return true }
func (*dummy) Debugf(format string, args ...interface{}) {}
func (*dummy) Infof(format string, args ...interface{})  {}
func (*dummy) Warnf(format string, args ...interface{})  {}
func (*dummy) Errorf(format string, args ...interface{}) {}
