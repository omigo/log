// Package log 实现了一个标准的可以自定义级别的 log 库，就像 slf4j(Simple Logging Facade for Java)
// 一样。这个 log 库的需要完成得任务就是提供一个标准统一的接口，同时也提供了一个基本的实现。
// 使用这个 log 库打印日志，可以随时切换日志级别，可以更换不同的 logger 实现，以打印不同格式的日
// 志，也可以改变日志输出位置，输出到数据库、消息队列等，者所有的改变都无需修改已经写好的项目源码。
//
// 安装：`go get -v -u github.com/gotips/log`
//
// 使用：
// ``` go
// package main
//
// import "github.com/gotips/log"
//
// func main() {
// 	 log.Info("level = %s", log.DebugLevel)
//   log.Error("this is a error message")
// }
// ```
// 日志输出：
// ```
// 2016-01-13 11:39:29.055566 info examples/main.go:6 level = debug
// 2016-01-13 11:39:29.055566 error examples/main.go:7 this is a error message
// ```
package log
