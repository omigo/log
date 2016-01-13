package log

import (
	"io"
	"os"
)

// Printer defined print adaptor
type Printer interface {

	// SetWriter sets the output destination for the logger.
	SetWriter(w io.Writer)

	// SetFlags sets the output flags for the logger.
	SetFlags(flag int)

	// Output writes the output for a logging event.
	Print(l Level, m string) error
}

func init() {
	SetPrinter(NewStandard(os.Stdout, LStdFlags))
}
