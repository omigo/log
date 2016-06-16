package log

import (
	"time"

	"google.golang.org/grpc/metadata"

	"golang.org/x/net/context"
)

// ======== 兼容 qiniu/log   ===============
const (
	Ldebug int = int(DebugLevel)
	Linfo      = int(InfoLevel)
	Lwarn      = int(WarnLevel)
	Lerror     = int(ErrorLevel)
	Lpanic     = int(PanicLevel)
	Lfatal     = int(FatalLevel)
)

// ======== 兼容 qiniu/log   ===============
func SetOutputLevel(l int) { v = Level(l) }

// ======== 兼容 wothing/log ===============

// TraceIn and TraceOut use in function in and out,reduce code line
// Example:
//	func test() {
//		user := User{Name: "zhangsan", Age: 21, School: "xayddx"}
//		service := "verification.GetVerifiCode"
//		defer log.TraceOut(log.TraceIn("12345", service, "user:%v", user))
//		....
//	}

// TraceIn 方法入口打印日志
func TraceIn(tag string, method string, format string, m ...interface{}) (string, string, time.Time) {
	startTime := time.Now()
	std.Tprintf(v, InfoLevel, tag, "calling "+method+", "+format, m...)
	return tag, method, startTime
}

// TraceCtx 方法入口打印日志
func TraceCtx(ctx context.Context, method string, format string, m ...interface{}) (string, string, time.Time) {
	tag := "-"
	if md, ok := metadata.FromContext(ctx); ok {
		if md["tid"] != nil && len(md["tid"]) > 0 {
			tag = md["tid"][0]
		}
	}
	startTime := time.Now()
	std.Tprintf(v, InfoLevel, tag, "calling "+method+", "+format, m...)
	return tag, method, startTime
}

// TraceOut 方法退出记录下消耗时间
func TraceOut(tag string, method string, startTime time.Time) {
	std.Tprintf(v, InfoLevel, tag, "finished "+method+", took %v", time.Since(startTime))
}
