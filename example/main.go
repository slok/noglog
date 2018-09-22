package main

import (
	"fmt"

	"github.com/google/glog"
)

// GlogReplacement is the replacement for glog using our own
// logging system.
type GlogReplacement struct {
	Debug bool
}

// DebugEnabled satisfies glog.Logger interface.
func (g *GlogReplacement) DebugEnabled() bool {
	return g.Debug
}

// Debugf satisfies glog.Logger interface.
func (g *GlogReplacement) Debugf(format string, args ...interface{}) {
	fmt.Printf("[DEBUG] "+format, args...)
}

// Infof satisfies glog.Logger interface.
func (g *GlogReplacement) Infof(format string, args ...interface{}) {
	fmt.Printf("[INFO] "+format, args...)
}

// Warnf satisfies glog.Logger interface.
func (g *GlogReplacement) Warnf(format string, args ...interface{}) {
	fmt.Printf("[WARN] "+format, args...)
}

// Errorf satisfies glog.Logger interface.
func (g *GlogReplacement) Errorf(format string, args ...interface{}) {
	fmt.Printf("[ERROR] "+format, args...)
}

func main() {
	logger := &GlogReplacement{}
	glog.SetLogger(logger)

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

	// Without debug.
	glog.V(2).Info("I'm batman!")
	glog.V(2).Infoln("I'm batman!")
	glog.V(2).Infof("%s - %s", "I'm", "batman!")

	// With debug.
	logger.Debug = true
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
