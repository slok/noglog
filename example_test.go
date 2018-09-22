package glog_test

import (
	"fmt"

	glog "github.com/slok/noglog"
)

// LoggerFunc shows how you would set a custom logger. In this case the LoggerFunc helper has been used,
// so we don't need to create a new type for the example.  But you could create a new type that satisfies
// noglog.Logger interface.
func Example_loggerFunc() {
	debug := false

	// Create our logger using the helper funcs so we don't need to create a new type.
	mlf := &glog.LoggerFunc{
		DebugEnabledFunc: func() bool { return debug },
		DebugfFunc:       func(format string, args ...interface{}) { fmt.Printf("[DEBUG] "+format, args...) },
		InfofFunc:        func(format string, args ...interface{}) { fmt.Printf("[INFO] "+format, args...) },
		WarnfFunc:        func(format string, args ...interface{}) { fmt.Printf("[WARN] "+format, args...) },
		ErrorfFunc:       func(format string, args ...interface{}) { fmt.Printf("[ERROR] "+format, args...) },
	}

	// Set the custom logger
	glog.SetLogger(mlf)

	// Test it.
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
	debug = true
	glog.V(2).Info("I'm batman!")
	glog.V(2).Infoln("I'm batman!")
	glog.V(2).Infof("%s - %s", "I'm", "batman!")

}
