package log

import (
	"encoding/json"
	"io"
)

// 默认 debug 级别，方便调试，生产环境可以调用 SetLevel 设置 log 级别
var v Level = DebugLevel

// 默认实现，输出到 os.Std 中，可以重定向到文件中，也可以调用 SetPrinter 其他方式输出
var std Printer

// SetLevel 设置日志级别
func SetLevel(l Level) { v = l }

// Colorized 输出日志是否着色，默认不着色
func Colorized(c bool) { std.Colorized(c) }

// GetLevel 返回设置的日志级别
func GetLevel() (l Level) { return v }

// SetPrinter 切换 Printer 实现
func SetPrinter(p Printer) { std = p }

// SetWriter 改变输出位置，通过这个接口，可以实现日志文件按时间或按大小滚动
func SetWriter(w io.Writer) { std.SetWriter(w) }

// SetFormat 改变日志格式
func SetFormat(format string) { std.SetFormat(format) }

// 判断各种级别的日志是否会被输出
func IsTraceEnabled() bool { return v <= TraceLevel }
func IsDebugEnabled() bool { return v <= DebugLevel }
func IsInfoEnabled() bool  { return v <= InfoLevel }
func IsWarnEnabled() bool  { return v <= WarnLevel }
func IsErrorEnabled() bool { return v <= ErrorLevel }
func IsPanicEnabled() bool { return v <= PanicLevel }
func IsFatalEnabled() bool { return v <= FatalLevel }
func IsPrintEnabled() bool { return v <= PrintLevel }
func IsStackEnabled() bool { return v <= StackLevel }

// 打印日志
func Trace(m ...interface{}) { std.Tprintf(TraceLevel, "", "", m...) }
func Debug(m ...interface{}) { std.Tprintf(DebugLevel, "", "", m...) }
func Info(m ...interface{})  { std.Tprintf(InfoLevel, "", "", m...) }
func Warn(m ...interface{})  { std.Tprintf(WarnLevel, "", "", m...) }
func Error(m ...interface{}) { std.Tprintf(ErrorLevel, "", "", m...) }
func Panic(m ...interface{}) { std.Tprintf(PanicLevel, "", "", m...) }
func Fatal(m ...interface{}) { std.Tprintf(FatalLevel, "", "", m...) }
func Print(m ...interface{}) { std.Tprintf(PrintLevel, "", "", m...) }
func Stack(m ...interface{}) { std.Tprintf(StackLevel, "", "", m...) }

// 按一定格式打印日志
func Tracef(format string, m ...interface{}) { std.Tprintf(TraceLevel, "", format, m...) }
func Debugf(format string, m ...interface{}) { std.Tprintf(DebugLevel, "", format, m...) }
func Infof(format string, m ...interface{})  { std.Tprintf(InfoLevel, "", format, m...) }
func Warnf(format string, m ...interface{})  { std.Tprintf(WarnLevel, "", format, m...) }
func Errorf(format string, m ...interface{}) { std.Tprintf(ErrorLevel, "", format, m...) }
func Panicf(format string, m ...interface{}) { std.Tprintf(PanicLevel, "", format, m...) }
func Fatalf(format string, m ...interface{}) { std.Tprintf(FatalLevel, "", format, m...) }
func Printf(format string, m ...interface{}) { std.Tprintf(PrintLevel, "", format, m...) }
func Stackf(format string, m ...interface{}) { std.Tprintf(StackLevel, "", format, m...) }

// 打印日志时带上 tag
func Ttrace(tag string, m ...interface{}) { std.Tprintf(TraceLevel, tag, "", m...) }
func Tdebug(tag string, m ...interface{}) { std.Tprintf(DebugLevel, tag, "", m...) }
func Tinfo(tag string, m ...interface{})  { std.Tprintf(InfoLevel, tag, "", m...) }
func Twarn(tag string, m ...interface{})  { std.Tprintf(WarnLevel, tag, "", m...) }
func Terror(tag string, m ...interface{}) { std.Tprintf(ErrorLevel, tag, "", m...) }
func Tpanic(tag string, m ...interface{}) { std.Tprintf(PanicLevel, tag, "", m...) }
func Tfatal(tag string, m ...interface{}) { std.Tprintf(FatalLevel, tag, "", m...) }
func Tprint(tag string, m ...interface{}) { std.Tprintf(PrintLevel, tag, "", m...) }
func Tstack(tag string, m ...interface{}) { std.Tprintf(StackLevel, tag, "", m...) }

// 按一定格式打印日志，并在打印日志时带上 tag
func Ttracef(tag string, format string, m ...interface{}) {
	std.Tprintf(TraceLevel, tag, format, m...)
}
func Tdebugf(tag string, format string, m ...interface{}) {
	std.Tprintf(DebugLevel, tag, format, m...)
}
func Tinfof(tag string, format string, m ...interface{}) { std.Tprintf(InfoLevel, tag, format, m...) }
func Twarnf(tag string, format string, m ...interface{}) { std.Tprintf(WarnLevel, tag, format, m...) }
func Terrorf(tag string, format string, m ...interface{}) {
	std.Tprintf(ErrorLevel, tag, format, m...)
}
func Tpanicf(tag string, format string, m ...interface{}) {
	std.Tprintf(PanicLevel, tag, format, m...)
}
func Tfatalf(tag string, format string, m ...interface{}) {
	std.Tprintf(FatalLevel, tag, format, m...)
}
func Tprintf(tag string, format string, m ...interface{}) {
	std.Tprintf(PrintLevel, tag, format, m...)
}
func Tstackf(tag string, format string, m ...interface{}) {
	std.Tprintf(StackLevel, tag, format, m...)
}

// 先转换成 JSON 格式，然后打印
func JSON(m ...interface{}) {
	if v > DebugLevel {
		return
	}
	js, err := json.Marshal(m)
	if err != nil {
		std.Tprintf(DebugLevel, "", "%s", err)
	} else {
		std.Tprintf(DebugLevel, "", "%s", js)
	}
}
func JSONIndent(m ...interface{}) {
	if v > DebugLevel {
		return
	}
	js, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		std.Tprintf(DebugLevel, "", "%s", err)
	} else {
		std.Tprintf(DebugLevel, "", "%s", js)
	}
}
