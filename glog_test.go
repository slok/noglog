package glog_test

import (
	"bytes"
	"fmt"
	"testing"

	glog "github.com/slok/noglog"
)

type mockLogger struct {
	debug bool
	rw    bytes.Buffer
}

func (m *mockLogger) DebugEnabled() bool {
	return m.debug
}

func (m *mockLogger) Debugf(format string, args ...interface{}) {
	if m.debug {
		s := fmt.Sprintf(format, args...)
		m.rw.WriteString("debugf: " + s)
	}
}
func (m *mockLogger) Infof(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	m.rw.WriteString("infof: " + s)
}
func (m *mockLogger) Warnf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	m.rw.WriteString("warnf: " + s)
}
func (m *mockLogger) Errorf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	m.rw.WriteString("errorf: " + s)
}

// TestGlog will test the glog replaced logger. This tests uses globals don't run in parallel.
func TestGlog(t *testing.T) {
	tests := []struct {
		name        string
		enableDebug bool
		glogActions func()
		expLog      string
	}{
		{
			name: "Info.",
			glogActions: func() {
				glog.Info("I'm Batman!")
			},
			expLog: "infof: I'm Batman!",
		},
		{
			name: "InfoDepth.",
			glogActions: func() {
				glog.InfoDepth(0, "I'm Batman!")
			},
			expLog: "infof: I'm Batman!",
		},
		{
			name: "Infoln.",
			glogActions: func() {
				glog.Infoln("I'm Batman!")
			},
			expLog: "infof: I'm Batman!\n",
		},
		{
			name: "Infof.",
			glogActions: func() {
				glog.Infof("%s %s", "I'm", "Batman!")
			},
			expLog: "infof: I'm Batman!",
		},

		{
			name: "Warning.",
			glogActions: func() {
				glog.Warning("I'm Batman!")
			},
			expLog: "warnf: I'm Batman!",
		},
		{
			name: "WarningDepth.",
			glogActions: func() {
				glog.WarningDepth(0, "I'm Batman!")
			},
			expLog: "warnf: I'm Batman!",
		},
		{
			name: "Warningln.",
			glogActions: func() {
				glog.Warningln("I'm Batman!")
			},
			expLog: "warnf: I'm Batman!\n",
		},
		{
			name: "Warningf.",
			glogActions: func() {
				glog.Warningf("%s %s", "I'm", "Batman!")
			},
			expLog: "warnf: I'm Batman!",
		},

		{
			name: "Error.",
			glogActions: func() {
				glog.Error("I'm Batman!")
			},
			expLog: "errorf: I'm Batman!",
		},
		{
			name: "ErrorDepth.",
			glogActions: func() {
				glog.ErrorDepth(0, "I'm Batman!")
			},
			expLog: "errorf: I'm Batman!",
		},
		{
			name: "Errorln.",
			glogActions: func() {
				glog.Errorln("I'm Batman!")
			},
			expLog: "errorf: I'm Batman!\n",
		},
		{
			name: "Errorf.",
			glogActions: func() {
				glog.Errorf("%s %s", "I'm", "Batman!")
			},
			expLog: "errorf: I'm Batman!",
		},

		{
			name: "Verbose Info (not enabled).",
			glogActions: func() {
				glog.V(2).Info("I'm Batman!")
			},
			expLog: "",
		},
		{
			name: "Verbose Infoln (not enabled).",
			glogActions: func() {
				glog.V(2).Infoln("I'm Batman!")
			},
			expLog: "",
		},
		{
			name: "Verbose Infof (enabled).",
			glogActions: func() {
				glog.V(2).Infof("%s %s", "I'm", "Batman!")
			},
			expLog: "",
		},
		{
			name:        "Verbose Info (enabled).",
			enableDebug: true,
			glogActions: func() {
				glog.V(2).Info("I'm Batman!")
			},
			expLog: "debugf: I'm Batman!",
		},
		{
			name:        "Verbose Infoln (enabled).",
			enableDebug: true,
			glogActions: func() {
				glog.V(2).Infoln("I'm Batman!")
			},
			expLog: "debugf: I'm Batman!\n",
		},
		{
			name:        "Verbose Infof (enabled).",
			enableDebug: true,
			glogActions: func() {
				glog.V(2).Infof("%s %s", "I'm", "Batman!")
			},
			expLog: "debugf: I'm Batman!",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ml := &mockLogger{
				debug: test.enableDebug,
			}
			glog.SetLogger(ml)
			test.glogActions()

			gotLog := ml.rw.String()
			if test.expLog != gotLog {
				t.Errorf("expected log doesn't match: \n exp: %s\n got: %s", test.expLog, gotLog)
			}
		})
	}
}
