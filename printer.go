package log

import (
	"io"
	"os"
)

func init() {
	// 默认实现标准格式标准输出
	SetPrinter(NewStandard(os.Stdout, DefaultFormat))
}

// Printer 定义了打印接口
type Printer interface {

	// 所有方法最终归为这个方法，真正打印日志
	Tprintf(v, l Level, tag string, format string, m ...interface{})

	// 貌似有问题，下面两个应该是默认实现的方法，而不是通用接口的方法
	// 但其他实现也可能有这两个方法，所以。。。纠结

	// SetFormat 设置日志格式
	SetFormat(format string)

	// ChangeWriter 改变输出流
	ChangeWriter(w io.Writer)
}
