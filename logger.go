package log


import (
	"io"
)
import "git.qutoutiao.net/gopher/qms/pkg/qlog"

type MyLogger struct{}

func NewMyLogger() *MyLogger {
	return &MyLogger{}
}

func (m *MyLogger) SetOutput(output io.Writer)  {}                         // 设置输出
func (m *MyLogger)  GetOutput() io.Writer    {panic("unimplemented")}                             // 获取输出
func (m *MyLogger)  SetLevel(level qlog.Level)     {}                              // 设置log等级
func (m *MyLogger)  GetLevel() qlog.Level       {panic("unimplemented")}                                 // 获取log等级
func (m *MyLogger)  Log(level qlog.Level, args ...interface{})     {panic("unimplemented")}              // 记录对应级别的日志
func (m *MyLogger)  Logf(level qlog.Level, format string, args ...interface{})  {panic("unimplemented")} // 记录对应级别的日志
func (m *MyLogger)  Trace(args ...interface{}) {Trace(args...)}   // 记录 TraceLevel 级别的日志
func (m *MyLogger)  Tracef(format string, args ...interface{}) {Tracef(format, args...)}   // 格式化并记录 TraceLevel 级别的日志
func (m *MyLogger)  Debug(args ...interface{}) {Debug(args...)}   // 记录 DebugLevel 级别的日志
func (m *MyLogger)  Debugf(format string, args ...interface{}) {Debugf(format, args...)}   // 格式化并记录 DebugLevel 级别的日志
func (m *MyLogger)  Info(args ...interface{}) {Info(args...)}   // 记录 InfoLevel 级别的日志
func (m *MyLogger)  Infof(format string, args ...interface{}) {Infof(format, args...)}   // 格式化并记录 InfoLevel 级别的日志
func (m *MyLogger)  Print(args ...interface{}) {Print(args...)}   // 记录 InfoLevel 级别的日志[gorm logger扩展]
func (m *MyLogger)  Printf(format string, args ...interface{}) {Printf(format, args...)}   // 格式化并记录 InfoLevel 级别的日志[gorm logger扩展]
func (m *MyLogger)  Warn(args ...interface{}) {Warn(args...)}   // 记录 WarnLevel 级别的日志
func (m *MyLogger)  Warnf(format string, args ...interface{}) {Warnf(format, args...)}   // 格式化并记录 WarnLevel 级别的日志
func (m *MyLogger)  Error(args ...interface{}) {Error(args...)}   // 记录 ErrorLevel 级别的日志
func (m *MyLogger)  Errorf(format string, args ...interface{}) {Errorf(format, args...)}   // 格式化并记录 ErrorLevel 级别的日志
func (m *MyLogger)  Fatal(args ...interface{}) {Fatal(args...)}   // 记录 FatalLevel 级别的日志
func (m *MyLogger)  Fatalf(format string, args ...interface{}) {Fatalf(format, args...)}   // 格式化并记录 FatalLevel 级别的日志
func (m *MyLogger)  Panic(args ...interface{}) {Panic(args...)}   // 记录 PanicLevel 级别的日志
func (m *MyLogger)  Panicf(format string, args ...interface{}) {Panicf(format, args...)}   // 格式化并记录 PanicLevel 级别的日志
func (m *MyLogger)  WithField(key string, value interface{}) qlog.Logger {panic("unimplemented")}   // 为日志添加一个上下文数据
func (m *MyLogger)  WithFields(fields qlog.Fields) qlog.Logger {panic("unimplemented")}   // 为日志添加多个上下文数据
func (m *MyLogger)  WithError(err error) qlog.Logger {panic("unimplemented")}   // 为日志添加标准错误上下文数据
