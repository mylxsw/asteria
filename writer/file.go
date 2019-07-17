package writer

import (
	"os"
	"sync"

	"github.com/mylxsw/asteria/level"
)

// SingleFileWriter is a LogWriter which write logs to file
type SingleFileWriter struct {
	filename string
	file     *os.File

	lock sync.RWMutex
}

// NewSingleFileWriter create a SingleFileWriter
func NewSingleFileWriter(filename string) *SingleFileWriter {
	return &SingleFileWriter{filename: filename,}
}

// Write the message to file
func (writer *SingleFileWriter) Write(le level.Level, message string) error {
	f, err := writer.open()
	if err != nil {
		return err
	}

	_, err = f.WriteString(message + "\n")
	return err
}

// ReOpen reopen a log file
func (writer *SingleFileWriter) ReOpen() error {
	if err := writer.Close(); err != nil {
		return err
	}

	_, err := writer.open()
	return err
}

// Close a log file
func (writer *SingleFileWriter) Close() error {
	writer.lock.Lock()
	defer writer.lock.Unlock()

	if writer.file != nil {
		err := writer.file.Close()
		if err != nil {
			return err
		}

		writer.file = nil
	}

	return nil
}

func (writer *SingleFileWriter) open() (*os.File, error) {
	writer.lock.Lock()
	defer writer.lock.Unlock()

	if writer.file == nil {
		f, err := os.OpenFile(writer.filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0660)
		if err != nil {
			return nil, err
		}

		writer.file = f
	}

	return writer.file, nil
}
