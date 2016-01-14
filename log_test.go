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

	format := "2006-01-02 15:04:05.999999999 info log/main.go:88 message"
	SetFormat(format)
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

func TestSetFormat(t *testing.T) {
	format := "2006-01-02 15:04:05.999999999 info log/main.go:88 message"
	SetFormat(format)

	Errorf("%s %s", "6ba7b814-9dad-11d1-80b4-00c04fd430c8", "this is a test long long long long message")
}
