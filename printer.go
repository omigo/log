package log

import (
	"io"
	"os"
)

// Printer defined print adaptor
type Printer interface {

	// ChangeWriter sets the output destination for the printer.
	ChangeWriter(w io.Writer)

	// SetFormat set output format for the printer
	SetFormat(format string)

	// Output writes the output for a logging event.
	Print(l Level, m string) error
}

func init() {
	SetPrinter(NewStandard(os.Stdout, DefaultFormat))
}
