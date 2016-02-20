package main

import "github.com/gotips/log"

func execColorizedExamples() {
	log.SetLevel(log.AllLevel)
	log.Info("default config")

	log.Colorized(true)
	log.Info("colorized config")

	log.Colorized(false)
	log.Error("close colorized config")
}
