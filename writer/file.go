package writer

import (
	"os"
	"sync"

	"github.com/mylxsw/asteria/level"
)

// FileWriter is a LogWriter which write logs to file
type FileWriter struct {
	filename string
	flag     int
	perm     os.FileMode
	file     *os.File

	lock sync.RWMutex
}

// NewFileWriter create a FileWriter
func NewFileWriter(filename string, flag int, perm os.FileMode) *FileWriter {
	return &FileWriter{filename: filename, flag: flag, perm: perm}
}

// NewDefaultFileWriter create new single file writer with default file mode and flags
func NewDefaultFileWriter(filename string) *FileWriter {
	return NewFileWriter(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
}

// Write the message to file
func (writer *FileWriter) Write(le level.Level, message string) error {
	f, err := writer.open()
	if err != nil {
		return err
	}

	_, err = f.WriteString(message + "\n")
	return err
}

// ReOpen reopen a log file
func (writer *FileWriter) ReOpen() error {
	if err := writer.Close(); err != nil {
		return err
	}

	_, err := writer.open()
	return err
}

// Close a log file
func (writer *FileWriter) Close() error {
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

func (writer *FileWriter) open() (*os.File, error) {
	writer.lock.Lock()
	defer writer.lock.Unlock()

	if writer.file == nil {
		f, err := os.OpenFile(writer.filename, writer.flag, writer.perm)
		if err != nil {
			return nil, err
		}

		writer.file = f
	}

	return writer.file, nil
}

func (writer *FileWriter) GetFilename() string {
	return writer.filename
}

type RotatingFileFn func(le level.Level) string

type RotatingFileWriter struct {
	fn          RotatingFileFn
	flag        int
	perm        os.FileMode
	openedFiles map[level.Level]*FileWriter

	lock sync.Mutex
}

func NewDefaultRotatingFileWriter(fn RotatingFileFn) *RotatingFileWriter {
	return NewRotatingFileWriter(os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666, fn)
}

func NewRotatingFileWriter(flag int, perm os.FileMode, fn RotatingFileFn) *RotatingFileWriter {
	return &RotatingFileWriter{fn: fn, flag: flag, perm: perm, openedFiles: make(map[level.Level]*FileWriter)}
}

func (writer *RotatingFileWriter) Write(le level.Level, message string) error {
	return writer.getWriter(le).Write(le, message)
}

func (writer *RotatingFileWriter) getWriter(le level.Level) *FileWriter {
	destFile := writer.fn(le)

	writer.lock.Lock()
	defer writer.lock.Unlock()

	w, ok := writer.openedFiles[le]
	if ok && w.filename != destFile {
		_ = w.Close()
		ok = false
	}

	if !ok {
		w = NewFileWriter(destFile, writer.flag, writer.perm)
		writer.openedFiles[le] = w
	}

	return w
}

func (writer *RotatingFileWriter) ReOpen() error {
	for _, w := range writer.openedFiles {
		if err := w.ReOpen(); err != nil {
			return err
		}
	}

	return nil
}

func (writer *RotatingFileWriter) Close() error {
	for _, w := range writer.openedFiles {
		if err := w.Close(); err != nil {
			return err
		}
	}

	return nil
}
