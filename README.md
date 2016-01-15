[![Build Status](https://travis-ci.org/omigo/log.svg?branch=develop)](https://travis-ci.org/omigo/log)

log
===

Golang 标准库中提供了基本的 log 模块 http://golang.org/pkg/log ，能 print/fatal/panic
日志，但唯独不能像 Java log4j 一样设置日志输出级别， debug 日志开发时输出，生产上关闭。这不能
不说是个巨大的遗憾， gopher 们只能抱怨 Golang 标准库的 log 就是个然并卵，实际项目大多不会使用
它。尽管 print 可以当成 error 使用（标准库确实把日志打印到错误输出流 os.Stdout），但是开发时
 debug/info 就没办法。

或许对标准库的设计 Golang 开发团队有自己的考虑，但是对应用开发者来说，log4j 已经成为事实上的标
准。为了向这个标准库靠近，出现了众多第三方 log 库，有在标准库基础上扩展的（也许 Golang 设计者
们也是想让开发者自己扩展标准库的 log 呢），也就另辟蹊径，也有玩各种花样的。

虽然有那么多的 log 库，但都是大同小异，我们需要的也只是个标准的可以自定义级别的 log 库而已，就
像 slf4j(Simple Logging Facade for Java) 一样，所以这个 log 库的需要完成得任务就是提供一
个标准统一的接口，同时也提供了一个基本的实现，可以自己定义模板格式，输出各种类型的日志，如
csv/json/xml。

使用这个 log 库打印日志，可以随时切换日志级别，可以更换不同的 logger 实现，以打印不同格式的日
志，也可以改变日志输出位置，输出到数据库、消息队列等，者所有的改变都无需修改已经写好的项目源码。


Usage
-----

安装：`go get -v -u github.com/gotips/log`

使用：
``` go
package main

import "github.com/gotips/log"

func main() {
    format = fmt.Sprintf("%s %s %s:%d %s", "2006-01-02 15:04:05.000000",
		log.LevelToken, log.ProjectToken, log.LineToken, log.MessageToken)
	log.SetFormat(format)
	log.Infof("this is a test message, %d", 1111)

	format = fmt.Sprintf(`{"date": "%s", "time": "%s", "level": "%s", "file": "%s", "line": %d, "log": "%s"}`,
		"2006-01-02", "15:04:05.999", log.LevelToken, log.ProjectToken, log.LineToken, log.MessageToken)
	log.SetFormat(format)
	log.Infof("this is a test message, %d", 1111)

	format = fmt.Sprintf(`<log><date>%s</date><time>%s</time><level>%s</level><file>%s</file><line>%d</line><msg>%s</msg><log>`,
		"2006-01-02", "15:04:05.000", log.LevelToken, log.ProjectToken, log.LineToken, log.MessageToken)
	log.SetFormat(format)
	log.Infof("this is a test message, %d", 1111)
}
```
日志输出：
```
2016-01-15 11:40:03.54123 info github.com/gotips/log/examples/main.go:24 this is a test message, 1111
{"date": "2016-01-15", "time": "11:40:03.541", "level": "info", "file": "examples/main.go", "line": 36, "log": "this is a test message, 1111"}
<log><date>2016-01-15</date><time>11:40:03.541</time><level>info</level><file>examples/main.go</file><line>42</line><msg>this is a test message, 1111</msg><log>
```

更多用法 [examples](examples/main.go)


log/Printer/Standard
--------------------

Golang 不同于 Java，非面向对象语言（没有继承，只有组合，不能把组合实例赋给被组合的实例，即 Java
说的 子对象 赋给 父对象），为了方便使用，很多函数都是包封装的，无需创建 struct ，就可以直接调用。

（一般把裸漏的方法称为函数，结构体和其他类型的方法才称为某某的方法）

log 包也一样，使用时，无需 new ，直接用。log 包有所有级别的函数可以调用，所有函数最终都调用了
print 函数。print 函数又调用了包内部变量的 std 的 Print 方法。这个 std 是一个 Printer 接
口类型，定义了打印接口。用不同的实现改变 std 就可以打印出不同格式的日志，也可以输出到不同位置。

Printer 有个基本的实现 Standard，如果不改变，默认使用这个实现打印日志。

Standard 实现了的 Printer 接口，以简洁的格式把日志打印到 Stdout。


TraceID 问题
------------

对有些项目，需要在各个系统之间跟踪请求链路，往往会产生一个请求的唯一标识 traceID，以下简称 tid，
需要在日志中打印出 tid。

Golang 1.4 之前可以取到 goroutine ID(goid)，但之后就取不到了，不然 log 库就可以通过上下文
直接取到 tid。目前来说，只能有两个折中的办法：

* tid 透传，所有方法都带上 tid 参数，或者使用 google context 库，传递更多参数，推荐后者，log
日志时，也带上这个 tid 参数；

* 如果系统按功能模块划分，而不是按层次划分，可以定义一个 struct ，比如 Foo ，把 TID 作为它的
一个属性，每个请求进入时，new 一个 struct ，传入 tid ，`Foo{tid}`，方法体内可以使用

``` go
func (f *Foo)Bar(){
    log.Infof("%s %s", f.TID, "something")
}
```

如果也不想在打印日志时传递 tid，可以定义一个 log 实现 Logger2 ，组合到 Foo 中，Logger 就可
以取到 tid 了，new Foo 时，也要 new Logger2 ，调用 `f.Infof("%s", "something")`。但这
样需要实现所有的 Trace/Debug/Info/Warn/Error/Fatal[f] 方法，而且无法切换 log 实现了，除
非改 Logger2 源码，这叫封装，就不是扩展了。不推荐这种极端做法！

TODO
----

* Benchmark Test
* 目前还不支持各种格式的日期
* 处理秒和毫秒，如1:1:02.9
* 实现日志文件按一定规则自动滚动
* 错误日志着色，开发阶段非常有用


Others
------

最近更新请移步 https://github.com/omigo/log
