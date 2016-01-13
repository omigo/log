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

	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lprojectfile)

	for i := 0; i < int(log.FatalLevel); i++ {
		log.SetLevel(log.Level(i))
		log.Errorf("level = %s", log.Level(i))
		log.Tracef("%d %s", log.AllLevel, log.AllLevel)
		log.Tracef("%d %s", log.TraceLevel, log.TraceLevel)
		log.Debugf("%d %s", log.DebugLevel, log.DebugLevel)
		log.Infof("%d %s", log.InfoLevel, log.InfoLevel)
		log.Warnf("%d %s", log.WarnLevel, log.WarnLevel)
		log.Errorf("%d %s", log.ErrorLevel, log.ErrorLevel)
		// Fatalf("%d %s", log.FatalLevel, log.FatalLevel)
		log.Error("----------------")
	}
	log.Fatalf("%d %s", log.FatalLevel, log.FatalLevel)
}
