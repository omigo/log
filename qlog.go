package log

import (
	"io"

	"git.qutoutiao.net/gopher/qms/pkg/qlog"
	//_ "unsafe"
	dlog "git.qutoutiao.net/pedestal/discovery/logger"
)

// 替换 Logrus，打印日志更简单，看起来清晰一些

////go:linkname m git.qutoutiao.net/gopher/qms/pkg/qlog.logger
//var logger qlog.Logger = log.NewMyLogger()
//
////go:linkname n git.qutoutiao.net/pedestal/discovery/logger.globalLogger
//var logger2 dlog.Interface = log.NewMyLogger()

func init() {
	qlog.SetLogger(NewMyLogger())
	dlog.SetLogger(NewMyLogger())
}

type MyLogger struct{}

func NewMyLogger() *MyLogger {
	return &MyLogger{}
}

func (m *MyLogger) SetOutput(output io.Writer)                                { panic("unimplemented") } // 设置输出
func (m *MyLogger) GetOutput() io.Writer                                      { panic("unimplemented") } // 获取输出
func (m *MyLogger) SetLevel(level qlog.Level)                                 {}                         // 设置log等级
func (m *MyLogger) GetLevel() qlog.Level                                      { panic("unimplemented") } // 获取log等级
func (m *MyLogger) Log(level qlog.Level, args ...interface{})                 { panic("unimplemented") } // 记录对应级别的日志
func (m *MyLogger) Logf(level qlog.Level, format string, args ...interface{}) { panic("unimplemented") } // 记录对应级别的日志
func (m *MyLogger) WithField(key string, value interface{}) qlog.Logger       { return m }               // 为日志添加一个上下文数据
func (m *MyLogger) WithFields(fields qlog.Fields) qlog.Logger                 { return m }               // 为日志添加多个上下文数据
func (m *MyLogger) WithError(err error) qlog.Logger                           { return m }               // 为日志添加标准错误上下文数据

// 打印日志
func (*MyLogger) Trace(m ...interface{}) { std.Tprintf(Ltrace, "", "", m...) }
func (*MyLogger) Debug(m ...interface{}) { std.Tprintf(Ldebug, "", "", m...) }
func (*MyLogger) Info(m ...interface{})  { std.Tprintf(Linfo, "", "", m...) }
func (*MyLogger) Warn(m ...interface{})  { std.Tprintf(Lwarn, "", "", m...) }
func (*MyLogger) Error(m ...interface{}) { std.Tprintf(Lerror, "", "", m...) }
func (*MyLogger) Panic(m ...interface{}) { std.Tprintf(Lpanic, "", "", m...) }
func (*MyLogger) Fatal(m ...interface{}) { std.Tprintf(Lfatal, "", "", m...) }
func (*MyLogger) Print(m ...interface{}) { std.Tprintf(Lprint, "", "", m...) }
func (*MyLogger) Stack(m ...interface{}) { std.Tprintf(Lstack, "", "", m...) }

// 按一定格式打印日志
func (*MyLogger) Tracef(format string, m ...interface{}) { std.Tprintf(Ltrace, "", format, m...) }
func (*MyLogger) Debugf(format string, m ...interface{}) { std.Tprintf(Ldebug, "", format, m...) }
func (*MyLogger) Infof(format string, m ...interface{})  { std.Tprintf(Linfo, "", format, m...) }
func (*MyLogger) Warnf(format string, m ...interface{})  { std.Tprintf(Lwarn, "", format, m...) }
func (*MyLogger) Errorf(format string, m ...interface{}) { std.Tprintf(Lerror, "", format, m...) }
func (*MyLogger) Panicf(format string, m ...interface{}) { std.Tprintf(Lpanic, "", format, m...) }
func (*MyLogger) Fatalf(format string, m ...interface{}) { std.Tprintf(Lfatal, "", format, m...) }
func (*MyLogger) Printf(format string, m ...interface{}) { std.Tprintf(Lprint, "", format, m...) }
func (*MyLogger) Stackf(format string, m ...interface{}) { std.Tprintf(Lstack, "", format, m...) }
