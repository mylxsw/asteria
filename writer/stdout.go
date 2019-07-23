package writer

import (
	"fmt"

	"github.com/mylxsw/asteria/level"
)

// StdoutWriter 默认日志输出器
type StdoutWriter struct{}

// NewStdoutWriter create a new default LogWriter
func NewStdoutWriter() *StdoutWriter {
	return &StdoutWriter{}
}

// Write 日志输出
func (writer StdoutWriter) Write(le level.Level, message string) error {
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
