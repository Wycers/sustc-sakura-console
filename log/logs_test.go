package log

import (
	"os"
	"testing"
)

var logger = NewLogger(os.Stdout)

func TestsetLevel(t *testing.T) {
	SetLevel("trace")
}

func TestTrace(t *testing.T) {
	logger.setLevel("trace")
	logger.Trace("trace")
	logger.setLevel("off")
	logger.Trace("trace")
}

func TestTracef(t *testing.T) {
	logger.setLevel("trace")
	logger.Tracef("tracef")
	logger.setLevel("off")
	logger.Tracef("tracef")
}

func TestDebug(t *testing.T) {
	logger.setLevel("debug")
	logger.Debug("debug")
	logger.setLevel("off")
	logger.Debug("debug")
}

func TestDebugf(t *testing.T) {
	logger.setLevel("debug")
	logger.Debugf("debugf")
	logger.setLevel("off")
	logger.Debug("debug")
}

func TestInfo(t *testing.T) {
	logger.setLevel("info")
	logger.Info("info")
	logger.setLevel("off")
	logger.Info("info")
}

func TestInfof(t *testing.T) {
	logger.setLevel("info")
	logger.Infof("infof")
	logger.setLevel("off")
	logger.Infof("infof")
}

func TestWarn(t *testing.T) {
	logger.setLevel("warn")
	logger.Warn("warn")
	logger.setLevel("off")
	logger.Warn("warn")
}

func TestWarnf(t *testing.T) {
	logger.setLevel("warn")
	logger.Warnf("warnf")
	logger.setLevel("off")
	logger.Warnf("warnf")
}

func TestError(t *testing.T) {
	logger.setLevel("error")
	logger.Error("error")
	logger.setLevel("off")
	logger.Error("error")
}

func TestErrorf(t *testing.T) {
	logger.setLevel("error")
	logger.Errorf("errorf")
	logger.setLevel("off")
	logger.Errorf("errorf")
}

func TestGetLevel(t *testing.T) {
	if getLevel("trace") != Trace {
		t.FailNow()

		return
	}

	if getLevel("debug") != Debug {
		t.FailNow()

		return
	}

	if getLevel("info") != Info {
		t.FailNow()

		return
	}

	if getLevel("warn") != Warn {
		t.FailNow()

		return
	}

	if getLevel("error") != Error {
		t.FailNow()

		return
	}

	if getLevel("fatal") != Fatal {
		t.FailNow()

		return
	}
}

func TestLoggersetLevel(t *testing.T) {
	logger.setLevel("trace")

	if logger.level != Trace {
		t.FailNow()

		return
	}
}

func TestisTraceEnabled(t *testing.T) {
	logger.setLevel("trace")

	if !logger.IsTraceEnabled() {
		t.FailNow()

		return
	}
}

func TestisDebugEnabled(t *testing.T) {
	logger.setLevel("debug")

	if !logger.IsDebugEnabled() {
		t.FailNow()

		return
	}
}

func TestisWarnEnabled(t *testing.T) {
	logger.setLevel("warn")

	if !logger.IsWarnEnabled() {
		t.FailNow()

		return
	}
}
