package writer

import (
	"context"
	"os"
	"sync"
	"time"

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
func (writer *FileWriter) Write(le level.Level, module string, message string) error {
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
		_ = writer.file.Sync()
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

func (writer *FileWriter) GetFileStat() (os.FileInfo, error) {
	writer.lock.Lock()
	defer writer.lock.Unlock()

	return writer.file.Stat()
}

type RotatingFileFn func(le level.Level, module string) string

type RotatingFileWriter struct {
	fn          RotatingFileFn
	flag        int
	perm        os.FileMode
	openedFiles map[string]*FileWriter

	lock sync.Mutex
}

func NewDefaultRotatingFileWriter(fn RotatingFileFn) *RotatingFileWriter {
	return NewRotatingFileWriter(os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666, fn)
}

func NewRotatingFileWriter(flag int, perm os.FileMode, fn RotatingFileFn) *RotatingFileWriter {
	return &RotatingFileWriter{fn: fn, flag: flag, perm: perm, openedFiles: make(map[string]*FileWriter),}
}

func (writer *RotatingFileWriter) Write(le level.Level, module string, message string) error {
	return writer.getWriter(writer.fn(le, module)).Write(le, module, message)
}

func (writer *RotatingFileWriter) AutoGC(ctx context.Context, interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				writer.GC(interval)
			}
		}
	}()
}

func (writer *RotatingFileWriter) GC(inactiveDuration time.Duration) {
	writer.lock.Lock()
	defer writer.lock.Unlock()

	deleteFiles := make([]string, 0)
	for filename, f := range writer.openedFiles {
		stat, err := f.GetFileStat()
		if err != nil {
			continue
		}

		if time.Now().After(stat.ModTime().Add(inactiveDuration)) {
			deleteFiles = append(deleteFiles, filename)
		}
	}

	for _, filename := range deleteFiles {
		if f, ok := writer.openedFiles[filename]; ok {
			_ = f.Close()
			delete(writer.openedFiles, filename)
		}
	}
}

func (writer *RotatingFileWriter) GetOpenedFiles() []string {
	writer.lock.Lock()
	defer writer.lock.Unlock()

	files := make([]string, 0)
	for filename, _ := range writer.openedFiles {
		files = append(files, filename)
	}

	return files
}

func (writer *RotatingFileWriter) getWriter(destFile string) *FileWriter {
	writer.lock.Lock()
	defer writer.lock.Unlock()

	w, ok := writer.openedFiles[destFile]
	if !ok {
		w = NewFileWriter(destFile, writer.flag, writer.perm)
		writer.openedFiles[destFile] = w
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
