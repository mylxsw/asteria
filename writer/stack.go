package writer

import (
	"github.com/mylxsw/asteria/level"
)

type StackMatchFn func(le level.Level, module string, message string) bool

type stackWriter struct {
	writer Writer
	fn     StackMatchFn
}

func (writer stackWriter) canWrite(le level.Level, module string, message string) bool {
	return writer.fn(le, module, message)
}

type StackWriter struct {
	writers []stackWriter
}

// NewStackWriter create a new stack writer
func NewStackWriter() *StackWriter {
	return &StackWriter{writers: make([]stackWriter, 0)}
}

// Push add a writer to stacks
func (writer *StackWriter) Push(w Writer, fn StackMatchFn) {
	writer.writers = append(writer.writers, stackWriter{
		writer: w,
		fn:     fn,
	})
}

// PushWithLevels add a writer with only specified levels enabled
// if no levels specified, we will use all
func (writer *StackWriter) PushWithLevels(w Writer, levels ...level.Level) {
	writer.Push(w, func(le level.Level, module string, message string) bool {
		if len(levels) == 0 {
			return true
		}

		for _, l := range levels {
			if l == le {
				return true
			}
		}

		return false
	})
}

func (writer *StackWriter) Write(le level.Level, module string, message string) error {
	for _, w := range writer.writers {
		if !w.canWrite(le, module, message) {
			continue
		}

		if err := w.writer.Write(le, module, message); err != nil {
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
