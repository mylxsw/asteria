package writer

import (
	"github.com/mylxsw/asteria/level"
)

type stackWriter struct {
	writer Writer
	levels []level.Level
}

func (writer stackWriter) canWrite(le level.Level) bool {
	if len(writer.levels) == 0 {
		return true
	}

	for _, l := range writer.levels {
		if l == le {
			return true
		}
	}

	return false
}

type StackWriter struct {
	writers []stackWriter
}

// NewStackWriter create a new stack writer
func NewStackWriter() *StackWriter {
	return &StackWriter{writers: make([]stackWriter, 0)}
}

// Push add a writer to stacks
func (writer *StackWriter) Push(w Writer, levels ...level.Level) {
	writer.writers = append(writer.writers, stackWriter{
		writer: w,
		levels: levels,
	})
}

func (writer *StackWriter) Write(le level.Level, message string) error {
	for _, w := range writer.writers {
		if !w.canWrite(le) {
			continue
		}

		if err := w.writer.Write(le, message); err != nil {
			return err
		}
	}

	return nil
}

func (writer *StackWriter) ReOpen() error {
	for _, w := range writer.writers {
		if err := w.writer.ReOpen(); err != nil {
			return err
		}
	}

	return nil
}

func (writer *StackWriter) Close() error {
	for _, w := range writer.writers {
		if err := w.writer.Close(); err != nil {
			return err
		}
	}

	return nil
}
