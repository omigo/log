package log

import "testing"

func TestLogLevel(t *testing.T) {

	Tracef("%d %s", AllLevel, AllLevel)
	Tracef("%d %s", TraceLevel, TraceLevel)
	Debugf("%d %s", DebugLevel, DebugLevel)
	Infof("%d %s", InfoLevel, InfoLevel)
	Warnf("%d %s", WarnLevel, WarnLevel)
	Errorf("%d %s", ErrorLevel, ErrorLevel)
	// Fatalf("%d %s", FatalLevel, FatalLevel)

	SetFlags(Ldate | Ltime | Lmicroseconds | Lprojectfile)
	for i := 0; i < int(FatalLevel); i++ {
		SetLevel(Level(i))
		Errorf("level = %s", Level(i))
		Tracef("%d %s", AllLevel, AllLevel)
		Tracef("%d %s", TraceLevel, TraceLevel)
		Debugf("%d %s", DebugLevel, DebugLevel)
		Infof("%d %s", InfoLevel, InfoLevel)
		Warnf("%d %s", WarnLevel, WarnLevel)
		Errorf("%d %s", ErrorLevel, ErrorLevel)
		// Fatalf("%d %s", FatalLevel, FatalLevel)
		Error("----------------")
	}
}
