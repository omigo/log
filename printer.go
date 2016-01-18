package log

import (
	"io"
	"os"
)

func init() {
	SetPrinter(NewStandard(os.Stdout, DefaultFormat))
}

// Printer defined print interface
type Printer interface {
	// SetFormat set output format for the printer
	SetFormat(format string)

	Tprintf(v, l Level, tid string, format string, m ...interface{})

	// ChangeWriter sets the output destination for the printer.
	ChangeWriter(w io.Writer)
}
