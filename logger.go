package log

import (
	"io"
	"os"
)

// Logger defined adapt logger
type Logger interface {

	// SetOutput sets the output destination for the logger.
	SetOutput(w io.Writer)

	// SetFlags sets the output flags for the logger.
	SetFlags(flag int)

	// Output writes the output for a logging event.
	Output(l Level, m string) error
}

func init() {
	SetLogger(NewStandard(os.Stderr, LStdFlags))
}
