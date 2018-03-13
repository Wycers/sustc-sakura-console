package log

import (
	"fmt"
	"io"
	"os"
	"strings"
	stdlog "log"
)

const(
	Off = iota
	Trace
	Debug
	Info
	Warn
	Error
	Fatal
)

type Logger struct {
	level	int
	logger  *stdlog.Logger
}
func (l *Logger) setLevel(level string) {
	l.level = getLevel(level)
}
func (l *Logger) IsTraceEnabled() bool {
	return l.level <= Trace
}
func (l *Logger) IsDebugEnabled() bool {
	return l.level <= Debug
}
func (l *Logger) IsWarnEnabled() bool {
	return l.level <= Warn
}

func (l *Logger) Trace(v ...interface{}) {
	if Trace < l.level {
		return
	}

	l.logger.SetPrefix("T ")
	l.logger.Output(2, fmt.Sprint(v...))
}

// Tracef prints trace level message with format.
func (l *Logger) Tracef(format string, v ...interface{}) {
	if Trace < l.level {
		return
	}

	l.logger.SetPrefix("T ")
	l.logger.Output(2, fmt.Sprintf(format, v...))
}

// Debug prints debug level message.
func (l *Logger) Debug(v ...interface{}) {
	if Debug < l.level {
		return
	}

	l.logger.SetPrefix("D ")
	l.logger.Output(2, fmt.Sprint(v...))
}

// Debugf prints debug level message with format.
func (l *Logger) Debugf(format string, v ...interface{}) {
	if Debug < l.level {
		return
	}

	l.logger.SetPrefix("D ")
	l.logger.Output(2, fmt.Sprintf(format, v...))
}

// Info prints info level message.
func (l *Logger) Info(v ...interface{}) {
	if Info < l.level {
		return
	}

	l.logger.SetPrefix("I ")
	l.logger.Output(2, fmt.Sprint(v...))
}

// Infof prints info level message with format.
func (l *Logger) Infof(format string, v ...interface{}) {
	if Info < l.level {
		return
	}

	l.logger.SetPrefix("I ")
	l.logger.Output(2, fmt.Sprintf(format, v...))
}

// Warn prints warning level message.
func (l *Logger) Warn(v ...interface{}) {
	if Warn < l.level {
		return
	}

	l.logger.SetPrefix("W ")
	l.logger.Output(2, fmt.Sprint(v...))
}

// Warn prints warning level message with format.
func (l *Logger) Warnf(format string, v ...interface{}) {
	if Warn < l.level {
		return
	}

	l.logger.SetPrefix("W ")
	l.logger.Output(2, fmt.Sprintf(format, v...))
}

// Error prints error level message.
func (l *Logger) Error(v ...interface{}) {
	if Error < l.level {
		return
	}

	l.logger.SetPrefix("E ")
	l.logger.Output(2, fmt.Sprint(v...))
}

// Errorf prints error level message with format.
func (l *Logger) Errorf(format string, v ...interface{}) {
	if Error < l.level {
		return
	}

	l.logger.SetPrefix("E ")
	l.logger.Output(2, fmt.Sprintf(format, v...))
}

// Fatal prints fatal level message.
func (l *Logger) Fatal(v ...interface{}) {
	if Fatal < l.level {
		return
	}

	l.logger.SetPrefix("F ")
	l.logger.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

// Fatalf prints fatal level message with format.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	if Fatal < l.level {
		return
	}

	l.logger.SetPrefix("F ")
	l.logger.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}




var loggers []*Logger
var logLevel = Debug



func NewLogger(out io.Writer) *Logger {
	res := &Logger{
		logLevel,
		stdlog.New(out, "", stdlog.Ldate|stdlog.Ltime|stdlog.Lshortfile),
		}

	loggers = append(loggers, res)

	return res
}
func getLevel(level string) int {
	level = strings.ToLower(level)

	switch level {
	case "off":
		return Off
	case "trace":
		return Trace
	case "debug":
		return Debug
	case "info":
		return Info
	case "warn":
		return Warn
	case "error":
		return Error
	case "fatal":
		return Fatal
	default:
		return Info
	}
}
func SetLevel(level string) {
	logLevel = getLevel(level)

	for _, l := range loggers{
		l.setLevel(level)
	}
}
