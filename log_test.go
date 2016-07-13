package log

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

const uuid = "6ba7b814-9dad-11d1-80b4-00c04fd430c8"

func TestLogLevel(t *testing.T) {
	SetLevel(LevelInfo)
	if IsDebugEnabled() || !IsInfoEnabled() || !IsWarnEnabled() {
		t.FailNow()
	}
	SetLevel(LevelDebug) // 恢复现场，避免影响其他单元测试
}

func TestSetWriter(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 4096))
	SetWriter(buf)

	rand := time.Now().String()
	Info(rand)
	if !bytes.Contains(buf.Bytes(), ([]byte)(rand)) {
		t.FailNow()
	}
}

func TestSetFormat(t *testing.T) {
	format := fmt.Sprintf(`<log><date>%s</date><time>%s</time><level>%s</level><file>%s</file><line>%d</line><msg>%s</msg><log>`,
		"2006-01-02", "15:04:05.000", LevelToken, ProjectToken, LineToken, MessageToken)
	SetFormat(format)

	buf := bytes.NewBuffer(make([]byte, 4096))
	SetWriter(buf)

	rand := time.Now().String()
	Debug(rand)
	if bytes.HasPrefix(buf.Bytes(), ([]byte)("<log><date>")) &&
		!bytes.HasSuffix(buf.Bytes(), ([]byte)("</msg><log>")) {
		t.FailNow()
	}
}

func TestPanicLog(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Fail()
		}
	}()
	Panic("test panic")
}

func TestNormalLog(t *testing.T) {
	SetLevel(LevelAll)

	Trace(LevelAll)
	Trace(LevelTrace)
	Debug(LevelDebug)
	Info(LevelInfo)
	Warn(LevelWarn)
	Error(LevelError)
	func() {
		defer func() {
			if err := recover(); err == nil {
				t.Fail()
			}
		}()
		Panic(LevelPanic)
	}()
	// Fatal( LevelFatal)
	Print(LevelPrint)
	Stack(LevelStack)
}

func TestFormatLog(t *testing.T) {
	SetLevel(LevelAll)

	Tracef("%d %s", LevelAll, LevelAll)
	Tracef("%d %s", LevelTrace, LevelTrace)
	Debugf("%d %s", LevelDebug, LevelDebug)
	Infof("%d %s", LevelInfo, LevelInfo)
	Warnf("%d %s", LevelWarn, LevelWarn)
	Errorf("%d %s", LevelError, LevelError)
	func() {
		defer func() {
			if err := recover(); err == nil {
				t.Fail()
			}
		}()
		Panicf("%d %s", LevelPanic, LevelPanic)
	}()
	// Fatalf("%d %s", LevelFatal, LevelFatal)
	Printf("%d %s", LevelPrint, LevelPrint)
	Stackf("%d %s", LevelStack, LevelStack)
}

func TestNormalLogWithTag(t *testing.T) {
	format := "2006-01-02 15:04:05 tag info examples/main.go:88 message"
	SetFormat(format)
	SetLevel(LevelAll)

	Ttrace(uuid, LevelAll)
	Ttrace(uuid, LevelTrace)
	Tdebug(uuid, LevelDebug)
	Tinfo(uuid, LevelInfo)
	Twarn(uuid, LevelWarn)
	Terror(uuid, LevelError)
	func() {
		defer func() {
			if err := recover(); err == nil {
				t.Fail()
			}
		}()
		Tpanic(uuid, LevelPanic)
	}()
	// Tfatal(uuid, LevelFatal)
	Tprint(uuid, LevelPrint)
	Tstack(uuid, LevelStack)
}

func TestFormatLogWithTag(t *testing.T) {
	format := "2006-01-02 15:04:05 tag info examples/main.go:88 message"
	SetFormat(format)
	SetLevel(LevelAll)

	Ttracef(uuid, "%d %s", LevelAll, LevelAll)
	Ttracef(uuid, "%d %s", LevelTrace, LevelTrace)
	Tdebugf(uuid, "%d %s", LevelDebug, LevelDebug)
	Tinfof(uuid, "%d %s", LevelInfo, LevelInfo)
	Twarnf(uuid, "%d %s", LevelWarn, LevelWarn)
	Terrorf(uuid, "%d %s", LevelError, LevelError)
	func() {
		defer func() {
			if err := recover(); err == nil {
				t.Fail()
			}
		}()
		Tpanicf(uuid, "%d %s", LevelPanic, LevelPanic)
	}()
	// Tfatalf(uuid,"%d %s", LevelFatal, LevelFatal)
	Tprintf(uuid, "%d %s", LevelPrint, LevelPrint)
	Tstackf(uuid, "%d %s", LevelStack, LevelStack)
}
