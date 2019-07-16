package asteria

import (
	"fmt"
	"log/syslog"
	"os"
	"sync"
)

// Writer 日志输出接口
type Writer interface {
	Write(level Level, message string) error
	ReOpen() error
	Close() error
}

// SyslogWriter is a log writer for syslog
type SyslogWriter struct {
	syslogWriter *syslog.Writer

	network  string
	raddr    string
	priority syslog.Priority
	tag      string

	lock sync.Mutex
}

// NewSyslogWriter create a new SyslogWriter
func NewSyslogWriter(network, raddr string, priority syslog.Priority, tag string) *SyslogWriter {
	return &SyslogWriter{
		network:  network,
		raddr:    raddr,
		priority: priority,
		tag:      tag,
	}
}

func (w *SyslogWriter) Write(level Level, message string) error {
	writer, err := w.writer()
	if err != nil {
		return err
	}

	switch level {
	case LevelEmergency:
		return writer.Emerg(message)
	case LevelAlert:
		return writer.Alert(message)
	case LevelCritical:
		return writer.Crit(message)
	case LevelError:
		return writer.Err(message)
	case LevelWarning:
		return writer.Warning(message)
	case LevelNotice:
		return writer.Notice(message)
	case LevelInfo:
		return writer.Info(message)
	case LevelDebug:
		return writer.Debug(message)
	}

	return nil
}

func (w *SyslogWriter) writer() (*syslog.Writer, error) {
	w.lock.Lock()
	defer w.lock.Unlock()

	if w.syslogWriter == nil {
		var err error
		w.syslogWriter, err = syslog.Dial(w.network, w.raddr, w.priority, w.tag)
		if err != nil {
			return nil, err
		}
	}

	return w.syslogWriter, nil
}

func (w *SyslogWriter) ReOpen() error {
	return w.Close()
}

func (w *SyslogWriter) Close() error {
	w.lock.Lock()
	defer w.lock.Unlock()

	err := w.syslogWriter.Close()
	w.syslogWriter = nil
	return err
}

// StdoutWriter 默认日志输出器
type StdoutWriter struct{}

// NewDefaultWriter create a new default LogWriter
func NewDefaultWriter() *StdoutWriter {
	return &StdoutWriter{}
}

// Write 日志输出
func (writer StdoutWriter) Write(level Level, message string) error {
	fmt.Println(message)
	return nil
}

// ReOpen reopen a log file
func (writer StdoutWriter) ReOpen() error {
	// do nothing
	return nil
}

// Close close a log file
func (writer StdoutWriter) Close() error {
	// do nothing
	return nil
}

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
func (writer *SingleFileWriter) Write(level Level, message string) error {
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
