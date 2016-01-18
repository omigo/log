package log

import "testing"

func TestSetFormat(t *testing.T) {
	format := "2006-01-02 15:04:05.999999999 info log/main.go:88 message"
	SetFormat(format)

	Errorf("%s %s", "6ba7b814-9dad-11d1-80b4-00c04fd430c8", "this is a test long long long long message")
}

func TestPanicLog(t *testing.T) {
	panic("test")
}

func TestNormalLog(t *testing.T) {
	SetLevel(AllLevel)

	Trace(AllLevel)
	Trace(TraceLevel)
	Debug(DebugLevel)
	Info(InfoLevel)
	Warn(WarnLevel)
	Error(ErrorLevel)
	func() {
		defer func() {
			if err := recover(); err != nil {
				// fmt.Println("panic recover")
			}
		}()
		Panic(PanicLevel)
	}()
	// Fatal( FatalLevel)
	Print(PrintLevel)
	Stack(StackLevel)
}

func TestFormatLog(t *testing.T) {
	SetLevel(AllLevel)

	Tracef("%d %s", AllLevel, AllLevel)
	Tracef("%d %s", TraceLevel, TraceLevel)
	Debugf("%d %s", DebugLevel, DebugLevel)
	Infof("%d %s", InfoLevel, InfoLevel)
	Warnf("%d %s", WarnLevel, WarnLevel)
	Errorf("%d %s", ErrorLevel, ErrorLevel)
	func() {
		defer func() {
			if err := recover(); err != nil {
				// do nothing
			}
		}()
		Panicf("%d %s", PanicLevel, PanicLevel)
	}()
	// Fatalf("%d %s", FatalLevel, FatalLevel)
	Printf("%d %s", PrintLevel, PrintLevel)
	Stackf("%d %s", StackLevel, StackLevel)
}

func TestNormalLogWithTraceID(t *testing.T) {
	format := "2006-01-02 15:04:05 tid info examples/main.go:88 message"
	SetFormat(format)
	SetLevel(AllLevel)

	Ttrace("6ba7b814-9dad-11d1-80b4-00c04fd430c8", AllLevel)
	Ttrace("6ba7b814-9dad-11d1-80b4-00c04fd430c8", TraceLevel)
	Tdebug("6ba7b814-9dad-11d1-80b4-00c04fd430c8", DebugLevel)
	Tinfo("6ba7b814-9dad-11d1-80b4-00c04fd430c8", InfoLevel)
	Twarn("6ba7b814-9dad-11d1-80b4-00c04fd430c8", WarnLevel)
	Terror("6ba7b814-9dad-11d1-80b4-00c04fd430c8", ErrorLevel)
	func() {
		defer func() {
			if err := recover(); err != nil {
				// do nothing
			}
		}()
		Tpanic("6ba7b814-9dad-11d1-80b4-00c04fd430c8", PanicLevel)
	}()
	// Tfatal("6ba7b814-9dad-11d1-80b4-00c04fd430c8", FatalLevel)
	Tprint("6ba7b814-9dad-11d1-80b4-00c04fd430c8", PrintLevel)
	Tstack("6ba7b814-9dad-11d1-80b4-00c04fd430c8", StackLevel)
}

func TestFormatLogWithTraceID(t *testing.T) {
	format := "2006-01-02 15:04:05 tid info examples/main.go:88 message"
	SetFormat(format)
	SetLevel(AllLevel)

	Ttracef("6ba7b814-9dad-11d1-80b4-00c04fd430c8", "%d %s", AllLevel, AllLevel)
	Ttracef("6ba7b814-9dad-11d1-80b4-00c04fd430c8", "%d %s", TraceLevel, TraceLevel)
	Tdebugf("6ba7b814-9dad-11d1-80b4-00c04fd430c8", "%d %s", DebugLevel, DebugLevel)
	Tinfof("6ba7b814-9dad-11d1-80b4-00c04fd430c8", "%d %s", InfoLevel, InfoLevel)
	Twarnf("6ba7b814-9dad-11d1-80b4-00c04fd430c8", "%d %s", WarnLevel, WarnLevel)
	Terrorf("6ba7b814-9dad-11d1-80b4-00c04fd430c8", "%d %s", ErrorLevel, ErrorLevel)
	func() {
		defer func() {
			if err := recover(); err != nil {
				// do nothing
			}
		}()
		Tpanicf("6ba7b814-9dad-11d1-80b4-00c04fd430c8", "%d %s", PanicLevel, PanicLevel)
	}()
	// Tfatalf("6ba7b814-9dad-11d1-80b4-00c04fd430c8","%d %s", FatalLevel, FatalLevel)
	Tprintf("6ba7b814-9dad-11d1-80b4-00c04fd430c8", "%d %s", PrintLevel, PrintLevel)
	Tstackf("6ba7b814-9dad-11d1-80b4-00c04fd430c8", "%d %s", StackLevel, StackLevel)
}

func TestWothingNormalLogWithTraceID(t *testing.T) {
	format := "2006-01-02 15:04:05 tid info examples/main.go:88 message"
	SetFormat(format)
	SetLevel(AllLevel)

	TraceT("6ba7b814-9dad-11d1-80b4-00c04fd430c8", AllLevel)
	TraceT("6ba7b814-9dad-11d1-80b4-00c04fd430c8", TraceLevel)
	DebugT("6ba7b814-9dad-11d1-80b4-00c04fd430c8", DebugLevel)
	InfoT("6ba7b814-9dad-11d1-80b4-00c04fd430c8", InfoLevel)
	WarnT("6ba7b814-9dad-11d1-80b4-00c04fd430c8", WarnLevel)
	ErrorT("6ba7b814-9dad-11d1-80b4-00c04fd430c8", ErrorLevel)
	func() {
		defer func() {
			if err := recover(); err != nil {
				// do nothing
			}
		}()
		PanicT("6ba7b814-9dad-11d1-80b4-00c04fd430c8", PanicLevel)
	}()
	// FatalT("6ba7b814-9dad-11d1-80b4-00c04fd430c8", FatalLevel)
	PrintT("6ba7b814-9dad-11d1-80b4-00c04fd430c8", PrintLevel)
	StackT("6ba7b814-9dad-11d1-80b4-00c04fd430c8", StackLevel)
}

func TestWothingFormatLogWithTraceID(t *testing.T) {
	format := "2006-01-02 15:04:05 tid info examples/main.go:88 message"
	SetFormat(format)
	SetLevel(AllLevel)

	TracefT("6ba7b814-9dad-11d1-80b4-00c04fd430c8", "%d %s", AllLevel, AllLevel)
	TracefT("6ba7b814-9dad-11d1-80b4-00c04fd430c8", "%d %s", TraceLevel, TraceLevel)
	DebugfT("6ba7b814-9dad-11d1-80b4-00c04fd430c8", "%d %s", DebugLevel, DebugLevel)
	InfofT("6ba7b814-9dad-11d1-80b4-00c04fd430c8", "%d %s", InfoLevel, InfoLevel)
	WarnfT("6ba7b814-9dad-11d1-80b4-00c04fd430c8", "%d %s", WarnLevel, WarnLevel)
	ErrorfT("6ba7b814-9dad-11d1-80b4-00c04fd430c8", "%d %s", ErrorLevel, ErrorLevel)
	func() {
		defer func() {
			if err := recover(); err != nil {
				// do nothing
			}
		}()
		PanicfT("6ba7b814-9dad-11d1-80b4-00c04fd430c8", "%d %s", PanicLevel, PanicLevel)
	}()
	// FatalfT("6ba7b814-9dad-11d1-80b4-00c04fd430c8","%d %s", FatalLevel, FatalLevel)
	PrintfT("6ba7b814-9dad-11d1-80b4-00c04fd430c8", "%d %s", PrintLevel, PrintLevel)
	StackfT("6ba7b814-9dad-11d1-80b4-00c04fd430c8", "%d %s", StackLevel, StackLevel)
}
