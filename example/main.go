package main

import (
	"fmt"
	"os"

	"github.com/google/glog"
	"github.com/rs/zerolog"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

// stdLogger is the replacement for glog using our own
// logging system.
type stdLogger struct {
	debug bool
}

// Debugf satisfies glog.Logger interface.
func (s *stdLogger) Debugf(format string, args ...interface{}) {
	if s.debug {
		fmt.Printf("[DEBUG] "+format, args...)
	}
}

// Infof satisfies glog.Logger interface.
func (*stdLogger) Infof(format string, args ...interface{}) {
	fmt.Printf("[INFO] "+format, args...)
}

// Warnf satisfies glog.Logger interface.
func (*stdLogger) Warnf(format string, args ...interface{}) {
	fmt.Printf("[WARN] "+format, args...)
}

// Errorf satisfies glog.Logger interface.
func (*stdLogger) Errorf(format string, args ...interface{}) {
	fmt.Printf("[ERROR] "+format, args...)
}

func main() {
	loggerSetters := map[string]func() glog.Logger{
		"std": func() glog.Logger {
			// Example using a type.
			return &stdLogger{debug: true}
		},
		"logrus": func() glog.Logger {
			// Example using LoggerFunc.
			logger := logrus.New()
			logger.SetLevel(logrus.DebugLevel)
			return &glog.LoggerFunc{
				DebugfFunc: func(f string, a ...interface{}) { logger.Debugf(f, a...) },
				InfofFunc:  func(f string, a ...interface{}) { logger.Infof(f, a...) },
				WarnfFunc:  func(f string, a ...interface{}) { logger.Warnf(f, a...) },
				ErrorfFunc: func(f string, a ...interface{}) { logger.Errorf(f, a...) },
			}
		},
		"zap": func() glog.Logger {
			// Example using LoggerFunc.
			logger, _ := zap.NewProduction()
			slogger := logger.Sugar()
			return &glog.LoggerFunc{
				DebugfFunc: func(f string, a ...interface{}) { slogger.Debugf(f, a...) },
				InfofFunc:  func(f string, a ...interface{}) { slogger.Infof(f, a...) },
				WarnfFunc:  func(f string, a ...interface{}) { slogger.Warnf(f, a...) },
				ErrorfFunc: func(f string, a ...interface{}) { slogger.Errorf(f, a...) },
			}
		},
		"zerolog": func() glog.Logger {
			// Example using LoggerFunc.
			w := zerolog.ConsoleWriter{Out: os.Stderr}
			logger := zerolog.New(w).With().Timestamp().Logger()
			return &glog.LoggerFunc{
				DebugfFunc: func(f string, a ...interface{}) { logger.Debug().Msgf(f, a...) },
				InfofFunc:  func(f string, a ...interface{}) { logger.Info().Msgf(f, a...) },
				WarnfFunc:  func(f string, a ...interface{}) { logger.Warn().Msgf(f, a...) },
				ErrorfFunc: func(f string, a ...interface{}) { logger.Error().Msgf(f, a...) },
			}
		},
	}

	for name, f := range loggerSetters {
		fmt.Printf("\n---------%s----------\n\n", name)
		glog.SetLogger(f())
		logWithGlog()
		fmt.Println()
	}
}

func logWithGlog() {
	glog.Info("I'm batman!")
	glog.InfoDepth(1, "I'm batman!")
	glog.Infoln("I'm batman!")
	glog.Infof("%s - %s", "I'm", "batman!")

	glog.Warning("I'm batman!")
	glog.WarningDepth(1, "I'm batman!")
	glog.Warningln("I'm batman!")
	glog.Warningf("%s - %s", "I'm", "batman!")

	glog.Error("I'm batman!")
	glog.ErrorDepth(1, "I'm batman!")
	glog.Errorln("I'm batman!")
	glog.Errorf("%s - %s", "I'm", "batman!")

	glog.V(2).Info("I'm batman!")
	glog.V(2).Infoln("I'm batman!")
	glog.V(2).Infof("%s - %s", "I'm", "batman!")

	// glog.Exit("I'm batman!")
	// glog.ExitDepth(1, "I'm batman!")
	// glog.Exitln("I'm batman!")
	// glog.Exitf("%s - %s", "I'm", "batman!")

	// glog.Fatal("I'm batman!")
	// glog.FatalDepth(1, "I'm batman!")
	// glog.Fatalln("I'm batman!")
	// glog.Fatalf("%s - %s", "I'm", "batman!")
}
