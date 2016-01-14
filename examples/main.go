package main

import "github.com/gotips/log"

func main() {
	log.Error("level = debug")
	log.Tracef("%d %s", log.AllLevel, log.AllLevel)
	log.Tracef("%d %s", log.TraceLevel, log.TraceLevel)
	log.Debugf("%d %s", log.DebugLevel, log.DebugLevel)
	log.Infof("%d %s", log.InfoLevel, log.InfoLevel)
	log.Warnf("%d %s", log.WarnLevel, log.WarnLevel)
	log.Errorf("%d %s", log.ErrorLevel, log.ErrorLevel)
	// Fatalf("%d %s", log.FatalLevel, log.FatalLevel)
	log.Error("----------------")

	log.Errorf("this is a test message, %d", 1111)
	log.Errorf("this is another test message, %d", 22222)

	log.Error("----------------")
	format := "2006-01-02 15:04:05.999 info /go/src/github.com/gotips/log/examples/main.go:88 message"
	log.SetFormat(format)
	log.Errorf("this is a test message, %d", 1111)
	log.Errorf("this is another test message, %d", 22222)
	log.Error("----------------")

	log.Error("----------------")
	format = "2006-01-02 15:01:02.999999 info github.com/gotips/log/examples/main.go:88 message"
	log.SetFormat(format)
	log.Errorf("this is a test message, %d", 1111)
	log.Errorf("this is another test message, %d", 22222)
	log.Error("----------------")

	format = "2006-01-02 15:01:02.000000 info examples/main.go:88 message"
	log.SetFormat(format)
	log.Errorf("this is a test message, %d", 1111)
	log.Errorf("this is another test message, %d", 22222)
	log.Error("----------------")

	format = `{"date": "2006-01-02", "time": "15:04:05.999", "level": "info", "file": "log/main.go", "line":88, "log": "message"}`
	log.SetFormat(format)
	log.Errorf("this is a test message, %d", 1111)
	log.Errorf("this is another test message, %d", 22222)
	log.Error("----------------")

	format = `<log><date>2006-01-02</date><time>15:04:05.999</time><level>info</level><file>log/main.go</file><line>88</line><msg>message</msg><log>`
	log.SetFormat(format)
	log.Errorf("this is a test message, %d", 1111)
	log.Errorf("this is another test message, %d", 22222)
	log.Error("----------------")
}
