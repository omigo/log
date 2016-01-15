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
var Labels = [...]string{"all", "trace", "debug", "info", "warn", "error", "fatal"}

// String return level label
func (v Level) String() string {
	return Labels[v]
}

var (
	// convenient for development
	v Level = DebugLevel
	// std default output
	std Printer
)

// SetLevel change log level
func SetLevel(l Level) { v = l }

// SetPrinter change default logger
func SetPrinter(out Printer) { std = out }

// ChangeWriter change logger output file
func ChangeWriter(w io.Writer) { std.ChangeWriter(w) }

// SetFormat sets the output format
func SetFormat(format string) { std.SetFormat(format) }

func Trace(m ...interface{})                 { output(TraceLevel, m...) }
func Tracef(format string, m ...interface{}) { outputf(TraceLevel, format, m...) }
func Debug(m ...interface{})                 { output(DebugLevel, m...) }
func Debugf(format string, m ...interface{}) { outputf(DebugLevel, format, m...) }
func Info(m ...interface{})                  { output(InfoLevel, m...) }
func Infof(format string, m ...interface{})  { outputf(InfoLevel, format, m...) }
func Warn(m ...interface{})                  { output(WarnLevel, m...) }
func Warnf(format string, m ...interface{})  { outputf(WarnLevel, format, m...) }
func Error(m ...interface{})                 { output(ErrorLevel, m...) }
func Errorf(format string, m ...interface{}) { outputf(ErrorLevel, format, m...) }
func Fatal(m ...interface{})                 { output(FatalLevel, m...); os.Exit(1) }
func Fatalf(format string, m ...interface{}) { outputf(FatalLevel, format, m...); os.Exit(1) }

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
