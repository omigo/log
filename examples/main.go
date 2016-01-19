package main

import (
	"fmt"

	"github.com/gotips/log"
)

func main() {
	log.Debugf("this is a test message, %d", 1111)

	format := fmt.Sprintf("%s %s %s %s:%d %s", "2006-01-02 15:04:05.000000", log.TagToken,
		log.LevelToken, log.ProjectToken, log.LineToken, log.MessageToken)
	log.ChangeFormat(format)
	log.Tinfof("6ba7b814-9dad-11d1-80b4-00c04fd430c8", "this is a test message, %d", 1111)

	format = fmt.Sprintf(`{"date": "%s", "time": "%s", "level": "%s", "file": "%s", "line": %d, "log": "%s"}`,
		"2006-01-02", "15:04:05.999", log.LevelToken, log.ProjectToken, log.LineToken, log.MessageToken)
	log.ChangeFormat(format)
	log.Infof("this is a test message, %d", 1111)

	format = fmt.Sprintf(`<log><date>%s</date><time>%s</time><level>%s</level><file>%s</file><line>%d</line><msg>%s</msg><log>`,
		"2006-01-02", "15:04:05.000", log.LevelToken, log.ProjectToken, log.LineToken, log.MessageToken)
	log.ChangeFormat(format)
	log.Tinfof("6ba7b814-9dad-11d1-80b4-00c04fd430c8", "this is a test message, %d", 1111)

	log.Error("level = debug")
	log.Infof("this is a test message, %d", 1111)
	log.Errorf("this is another test message, %d", 22222)
	// Fatalf("%d %s", log.FatalLevel, log.FatalLevel)

	format = fmt.Sprintf("%s %s %s %s:%d %s", "2006-1-2", "3:4:05.9",
		log.LevelToken, log.PathToken, log.LineToken, log.MessageToken)
	log.ChangeFormat(format)
	log.Infof("this is a test message, %d", 1111)

	format = fmt.Sprintf("%s %s %s %s:%d %s", "2006-01-02", "15:04:05.999999",
		log.LevelToken, log.PackageToken, log.LineToken, log.MessageToken)
	log.ChangeFormat(format)
	log.Infof("this is a test message, %d", 1111)

	format = fmt.Sprintf("%s %s %s:%d %s", "2006-01-02 15:04:05.000000",
		log.LevelToken, log.ProjectToken, log.LineToken, log.MessageToken)
	log.ChangeFormat(format)
	log.Infof("this is a test message, %d", 1111)

	format = fmt.Sprintf(`{"date": "%s", "time": "%s", "level": "%s", "file": "%s", "line": %d, "log": "%s"}`,
		"2006-01-02", "15:04:05.999", log.LevelToken, log.ProjectToken, log.LineToken, log.MessageToken)
	log.ChangeFormat(format)
	log.Infof("this is a test message, %d", 1111)

	format = fmt.Sprintf(`<log><date>%s</date><time>%s</time><level>%s</level><file>%s</file><line>%d</line><msg>%s</msg><log>`,
		"2006-01-02", "15:04:05.000", log.LevelToken, log.ProjectToken, log.LineToken, log.MessageToken)
	log.ChangeFormat(format)
	log.Infof("this is a test message, %d", 1111)
}
