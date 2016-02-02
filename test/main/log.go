package main

import g "log"
import "github.com/gotips/log"

func main() {
	gotipslog() //  大约 16w 行每秒

	// golog() // 大约 36.5w 行每秒
}

func gotipslog() {
	for i := 0; i < 200e4; i++ {
		log.Print("can't load package: package lib: cannot find package `xxx` in any of")
	}
}

func golog() {
	g.SetFlags(g.Ldate | g.Ltime | g.Lshortfile)
	for i := 0; i < 200e4; i++ {
		g.Print("can't load package: package lib: cannot find package `xxx` in any of")
	}
}
