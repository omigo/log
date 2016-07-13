package log

import (
	"encoding/json"
	"io"
)

// 默认 debug 级别，方便调试，生产环境可以调用 LevelSet 设置 log 级别
var v Level = LevelDebug

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
func IsTraceEnabled() bool { return v <= LevelTrace }
func IsDebugEnabled() bool { return v <= LevelDebug }
func IsInfoEnabled() bool  { return v <= LevelInfo }
func IsWarnEnabled() bool  { return v <= LevelWarn }
func IsErrorEnabled() bool { return v <= LevelError }
func IsPanicEnabled() bool { return v <= LevelPanic }
func IsFatalEnabled() bool { return v <= LevelFatal }
func IsPrintEnabled() bool { return v <= LevelPrint }
func IsStackEnabled() bool { return v <= LevelStack }

// 打印日志
func Trace(m ...interface{}) { std.Tprintf(LevelTrace, "", "", m...) }
func Debug(m ...interface{}) { std.Tprintf(LevelDebug, "", "", m...) }
func Info(m ...interface{})  { std.Tprintf(LevelInfo, "", "", m...) }
func Warn(m ...interface{})  { std.Tprintf(LevelWarn, "", "", m...) }
func Error(m ...interface{}) { std.Tprintf(LevelError, "", "", m...) }
func Panic(m ...interface{}) { std.Tprintf(LevelPanic, "", "", m...) }
func Fatal(m ...interface{}) { std.Tprintf(LevelFatal, "", "", m...) }
func Print(m ...interface{}) { std.Tprintf(LevelPrint, "", "", m...) }
func Stack(m ...interface{}) { std.Tprintf(LevelStack, "", "", m...) }

// 按一定格式打印日志
func Tracef(format string, m ...interface{}) { std.Tprintf(LevelTrace, "", format, m...) }
func Debugf(format string, m ...interface{}) { std.Tprintf(LevelDebug, "", format, m...) }
func Infof(format string, m ...interface{})  { std.Tprintf(LevelInfo, "", format, m...) }
func Warnf(format string, m ...interface{})  { std.Tprintf(LevelWarn, "", format, m...) }
func Errorf(format string, m ...interface{}) { std.Tprintf(LevelError, "", format, m...) }
func Panicf(format string, m ...interface{}) { std.Tprintf(LevelPanic, "", format, m...) }
func Fatalf(format string, m ...interface{}) { std.Tprintf(LevelFatal, "", format, m...) }
func Printf(format string, m ...interface{}) { std.Tprintf(LevelPrint, "", format, m...) }
func Stackf(format string, m ...interface{}) { std.Tprintf(LevelStack, "", format, m...) }

// 打印日志时带上 tag
func Ttrace(tag string, m ...interface{}) { std.Tprintf(LevelTrace, tag, "", m...) }
func Tdebug(tag string, m ...interface{}) { std.Tprintf(LevelDebug, tag, "", m...) }
func Tinfo(tag string, m ...interface{})  { std.Tprintf(LevelInfo, tag, "", m...) }
func Twarn(tag string, m ...interface{})  { std.Tprintf(LevelWarn, tag, "", m...) }
func Terror(tag string, m ...interface{}) { std.Tprintf(LevelError, tag, "", m...) }
func Tpanic(tag string, m ...interface{}) { std.Tprintf(LevelPanic, tag, "", m...) }
func Tfatal(tag string, m ...interface{}) { std.Tprintf(LevelFatal, tag, "", m...) }
func Tprint(tag string, m ...interface{}) { std.Tprintf(LevelPrint, tag, "", m...) }
func Tstack(tag string, m ...interface{}) { std.Tprintf(LevelStack, tag, "", m...) }

// 按一定格式打印日志，并在打印日志时带上 tag
func Ttracef(tag string, format string, m ...interface{}) {
	std.Tprintf(LevelTrace, tag, format, m...)
}
func Tdebugf(tag string, format string, m ...interface{}) {
	std.Tprintf(LevelDebug, tag, format, m...)
}
func Tinfof(tag string, format string, m ...interface{}) { std.Tprintf(LevelInfo, tag, format, m...) }
func Twarnf(tag string, format string, m ...interface{}) { std.Tprintf(LevelWarn, tag, format, m...) }
func Terrorf(tag string, format string, m ...interface{}) {
	std.Tprintf(LevelError, tag, format, m...)
}
func Tpanicf(tag string, format string, m ...interface{}) {
	std.Tprintf(LevelPanic, tag, format, m...)
}
func Tfatalf(tag string, format string, m ...interface{}) {
	std.Tprintf(LevelFatal, tag, format, m...)
}
func Tprintf(tag string, format string, m ...interface{}) {
	std.Tprintf(LevelPrint, tag, format, m...)
}
func Tstackf(tag string, format string, m ...interface{}) {
	std.Tprintf(LevelStack, tag, format, m...)
}

// 先转换成 JSON 格式，然后打印
func JSON(m ...interface{}) {
	if v > LevelDebug {
		return
	}
	js, err := json.Marshal(m)
	if err != nil {
		std.Tprintf(LevelDebug, "", "%s", err)
	} else {
		std.Tprintf(LevelDebug, "", "%s", js)
	}
}
func JSONIndent(m ...interface{}) {
	if v > LevelDebug {
		return
	}
	js, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		std.Tprintf(LevelDebug, "", "%s", err)
	} else {
		std.Tprintf(LevelDebug, "", "%s", js)
	}
}
