package writer

import (
	"fmt"
	"io"
	"os"

	"github.com/mylxsw/asteria/level"
)

type StreamWriter struct {
	w io.Writer
}

func NewStreamWriter(w io.Writer) *StreamWriter {
	return &StreamWriter{w: w}
}

func NewStdoutWriter() *StreamWriter {
	return NewStreamWriter(os.Stdout)
}

func (writer *StreamWriter) Write(le level.Level, message string) error {
	_, err := fmt.Fprintln(writer.w, message)
	return err
}

func (writer *StreamWriter) ReOpen() error {
	return nil
}

func (writer *StreamWriter) Close() error {
	return nil
}
