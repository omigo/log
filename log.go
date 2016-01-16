package log

import "io"

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
	PanicLevel
	FatalLevel
	PrintLevel
	StackLevel
)

// Labels level label
var Labels = [...]string{"all", "trace", "debug", "info", "warn", "error", "panic", "fatal", "print", "stack"}

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

// log
func Trace(m ...interface{}) { std.Tprintf(v, TraceLevel, "", "", m...) }
func Debug(m ...interface{}) { std.Tprintf(v, DebugLevel, "", "", m...) }
func Info(m ...interface{})  { std.Tprintf(v, InfoLevel, "", "", m...) }
func Warn(m ...interface{})  { std.Tprintf(v, WarnLevel, "", "", m...) }
func Error(m ...interface{}) { std.Tprintf(v, ErrorLevel, "", "", m...) }
func Panic(m ...interface{}) { std.Tprintf(v, PanicLevel, "", "", m...) }
func Fatal(m ...interface{}) { std.Tprintf(v, FatalLevel, "", "", m...) }
func Print(m ...interface{}) { std.Tprintf(v, PrintLevel, "", "", m...) }
func Stack(m ...interface{}) { std.Tprintf(v, StackLevel, "", "", m...) }

// log with format
func Tracef(format string, m ...interface{}) { std.Tprintf(v, TraceLevel, "", format, m...) }
func Debugf(format string, m ...interface{}) { std.Tprintf(v, DebugLevel, "", format, m...) }
func Infof(format string, m ...interface{})  { std.Tprintf(v, InfoLevel, "", format, m...) }
func Warnf(format string, m ...interface{})  { std.Tprintf(v, WarnLevel, "", format, m...) }
func Errorf(format string, m ...interface{}) { std.Tprintf(v, ErrorLevel, "", format, m...) }
func Panicf(format string, m ...interface{}) { std.Tprintf(v, PanicLevel, "", format, m...) }
func Fatalf(format string, m ...interface{}) { std.Tprintf(v, FatalLevel, "", format, m...) }
func Printf(format string, m ...interface{}) { std.Tprintf(v, PrintLevel, "", format, m...) }
func Stackf(format string, m ...interface{}) { std.Tprintf(v, StackLevel, "", format, m...) }

// log with traceID
func Ttrace(tid string, m ...interface{}) { std.Tprintf(v, TraceLevel, tid, "", m...) }
func Tdebug(tid string, m ...interface{}) { std.Tprintf(v, DebugLevel, tid, "", m...) }
func Tinfo(tid string, m ...interface{})  { std.Tprintf(v, InfoLevel, tid, "", m...) }
func Twarn(tid string, m ...interface{})  { std.Tprintf(v, WarnLevel, tid, "", m...) }
func Terror(tid string, m ...interface{}) { std.Tprintf(v, ErrorLevel, tid, "", m...) }
func Tpanic(tid string, m ...interface{}) { std.Tprintf(v, PanicLevel, tid, "", m...) }
func Tfatal(tid string, m ...interface{}) { std.Tprintf(v, FatalLevel, tid, "", m...) }
func Tprint(tid string, m ...interface{}) { std.Tprintf(v, PrintLevel, tid, "", m...) }
func Tstack(tid string, m ...interface{}) { std.Tprintf(v, StackLevel, tid, "", m...) }

// log with traceID and format
func Ttracef(tid string, format string, m ...interface{}) {
	std.Tprintf(v, TraceLevel, tid, format, m...)
}
func Tdebugf(tid string, format string, m ...interface{}) {
	std.Tprintf(v, DebugLevel, tid, format, m...)
}
func Tinfof(tid string, format string, m ...interface{}) { std.Tprintf(v, InfoLevel, tid, format, m...) }
func Twarnf(tid string, format string, m ...interface{}) { std.Tprintf(v, WarnLevel, tid, format, m...) }
func Terrorf(tid string, format string, m ...interface{}) {
	std.Tprintf(v, ErrorLevel, tid, format, m...)
}
func Tpanicf(tid string, format string, m ...interface{}) {
	std.Tprintf(v, PanicLevel, tid, format, m...)
}
func Tfatalf(tid string, format string, m ...interface{}) {
	std.Tprintf(v, FatalLevel, tid, format, m...)
}
func Tprintf(tid string, format string, m ...interface{}) {
	std.Tprintf(v, PrintLevel, tid, format, m...)
}
func Tstackf(tid string, format string, m ...interface{}) {
	std.Tprintf(v, StackLevel, tid, format, m...)
}

// ======== 兼容 wothing/log ===============

// log with traceID
func TraceT(tid string, m ...interface{}) { Ttrace(tid, m...) }
func DebugT(tid string, m ...interface{}) { Tdebug(tid, m...) }
func InfoT(tid string, m ...interface{})  { Tinfo(tid, m...) }
func WarnT(tid string, m ...interface{})  { Twarn(tid, m...) }
func ErrorT(tid string, m ...interface{}) { Terror(tid, m...) }
func PanicT(tid string, m ...interface{}) { Tpanic(tid, m...) }
func FatalT(tid string, m ...interface{}) { Tfatal(tid, m...) }
func PrintT(tid string, m ...interface{}) { Tprint(tid, m...) }
func StackT(tid string, m ...interface{}) { Tstack(tid, m...) }

// log with traceID and format
func TracefT(tid string, format string, m ...interface{}) { Ttracef(tid, format, m...) }
func DebugfT(tid string, format string, m ...interface{}) { Tdebugf(tid, format, m...) }
func InfofT(tid string, format string, m ...interface{})  { Tinfof(tid, format, m...) }
func WarnfT(tid string, format string, m ...interface{})  { Twarnf(tid, format, m...) }
func ErrorfT(tid string, format string, m ...interface{}) { Terrorf(tid, format, m...) }
func PanicfT(tid string, format string, m ...interface{}) { Tpanicf(tid, format, m...) }
func FatalfT(tid string, format string, m ...interface{}) { Tfatalf(tid, format, m...) }
func PrintfT(tid string, format string, m ...interface{}) { Tprintf(tid, format, m...) }
func StackfT(tid string, format string, m ...interface{}) { Tstackf(tid, format, m...) }
