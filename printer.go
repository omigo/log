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

	// ChangeFormat 改变日志格式
	ChangeFormat(format string)

	// ChangeWriter 改变输出流
	ChangeWriter(w io.Writer)
}
