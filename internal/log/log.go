package log

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var _exit = func() { os.Exit(1) }

// Printer is a formatting printer.
type Printer interface {
	Debugf(string, ...interface{})
	Printf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
	Fatalf(string, ...interface{})
}

// New returns a new Logger backed by the logrus library's log package.
func DefaultLogger() Printer {
	return &Logger{
		Printer: logrus.New(),
	}
}

// A Logger writes output to standard error.
type Logger struct {
	Printer
}

// Printf logs a formatted Dancong line.
func (l *Logger) Printf(format string, v ...interface{}) {
	l.Printer.Printf(prepend(format), v...)
}

// Fatalf logs an Dancong line then fatals.
func (l *Logger) Infof(format string, v ...interface{}) {
	l.Printer.Infof(prepend(format), v...)
}

// Fatalf logs an Dancong line then fatals.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Printer.Printf(prepend(format), v...)
}

func prepend(str string) string {
	return fmt.Sprintf("[Dancong] %s", str)
}
