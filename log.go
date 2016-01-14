package log

import (
	"fmt"
	"io"
	"os"
)

// Level log level type
type Level uint8

// levels enum
const (
	AllLevel Level = iota
	TraceLevel
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

// Labels level label
var Labels = [...]string{
	"all", "trace", "debug", "info", "warn", "error", "fatal"}

// String return level label
func (v Level) String() string {
	return Labels[v]
}

// convenient for development
var v Level = DebugLevel

// SetLevel change log level
func SetLevel(l Level) {
	v = l
}

// std default output
var std Printer

// SetPrinter change default logger
func SetPrinter(out Printer) {
	std = out
}

// ChangeWriter change logger output file
func ChangeWriter(w io.Writer) {
	std.ChangeWriter(w)
}

// SetFormat sets the output format
func SetFormat(format string) {
	std.SetFormat(format)
}

// Trace log a message at the Trace level
func Trace(m ...interface{}) {
	output(TraceLevel, m...)
}

// Tracef log a message at the Trace level according to the specified format and arguments
func Tracef(format string, m ...interface{}) {
	outputf(TraceLevel, format, m...)
}

// Debug log a message at the Debug level
func Debug(m ...interface{}) {
	output(DebugLevel, m...)
}

// Debugf log a message at the Debug level according to the specified format and arguments
func Debugf(format string, m ...interface{}) {
	outputf(DebugLevel, format, m...)
}

// Info log a message at the Info Info
func Info(m ...interface{}) {
	output(InfoLevel, m...)
}

// Infof log a message at the Info Info according to the specified format and arguments
func Infof(format string, m ...interface{}) {
	outputf(InfoLevel, format, m...)
}

// Warn log a message at the Warn level
func Warn(m ...interface{}) {
	output(WarnLevel, m...)
}

// Warnf log a message at the Warn level according to the specified format and arguments
func Warnf(format string, m ...interface{}) {
	outputf(WarnLevel, format, m...)
}

// Error log a message at the Error level
func Error(m ...interface{}) {
	output(ErrorLevel, m...)
}

// Errorf log a message at the Error level according to the specified format and arguments
func Errorf(format string, m ...interface{}) {
	outputf(ErrorLevel, format, m...)
}

// Fatal is equivalent to Error() followed by a call to os.Exit(1).
func Fatal(m ...interface{}) {
	output(FatalLevel, m...)
	os.Exit(1)
}

// Fatalf is equivalent to Errorf() followed by a call to os.Exit(1).
func Fatalf(format string, m ...interface{}) {
	outputf(FatalLevel, format, m...)
	os.Exit(1)
}

func output(l Level, m ...interface{}) {
	// if log level heigh than method level, no output
	if v > l {
		return
	}

	std.Print(l, fmt.Sprint(m...))
}

func outputf(l Level, format string, m ...interface{}) {
	// if log level heigh than method level, no output
	if v > l {
		return
	}

	std.Print(l, fmt.Sprintf(format, m...))
}
