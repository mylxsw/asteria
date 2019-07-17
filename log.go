package asteria

import (
	"fmt"

	"github.com/mylxsw/asteria/formatter"
	"github.com/mylxsw/asteria/level"
	"github.com/mylxsw/asteria/writer"
)

type C map[string]interface{}

// ReOpenAll reopen all logger
func ReOpenAll() map[string]error {
	errors := make(map[string]error, len(loggers))
	for name, l := range loggers {
		errors[name] = l.ReOpen()
	}

	return errors
}

// ReOpen reopen default log file
func ReOpen() error {
	return Default().ReOpen()
}

// CloseAll close all loggers
func CloseAll() map[string]error {
	errors := make(map[string]error, len(loggers))
	for name, l := range loggers {
		errors[name] = l.Close()
	}

	return errors
}

// Close default log file
func Close() error {
	return Default().Close()
}

// LogLevel 设置日志输出级别
func SetLevel(le level.Level) *Logger {
	return Default().LogLevel(le)
}

// Formatter 设置日志格式化器
func SetFormatter(f formatter.Formatter) *Logger {
	return Default().Formatter(f)
}

// Writer 设置日志输出器
func SetWriter(w writer.Writer) *Logger {
	return Default().Writer(w)
}

func WithContext(context C) *ContextLogger {
	return Default().WithContext(context)
}

func Emergency(v ...interface{}) {
	Default().Output(3, level.Emergency, nil, v...)
}

func Alert(v ...interface{}) {
	Default().Output(3, level.Alert, nil, v...)
}

func Critical(v ...interface{}) {
	Default().Output(3, level.Critical, nil, v...)
}

func Error(v ...interface{}) {
	Default().Output(3, level.Error, nil, v...)
}

func Warning(v ...interface{}) {
	Default().Output(3, level.Warning, nil, v...)
}

func Notice(v ...interface{}) {
	Default().Output(3, level.Notice, nil, v...)
}

func Info(v ...interface{}) {
	Default().Output(3, level.Info, nil, v...)
}

func Debug(v ...interface{}) {
	Default().Output(3, level.Debug, nil, v...)
}

func Emergencyf(format string, v ...interface{}) {
	Default().Output(3, level.Emergency, nil, fmt.Sprintf(format, v...))
}

func Alertf(format string, v ...interface{}) {
	Default().Output(3, level.Alert, nil, fmt.Sprintf(format, v...))
}

func Criticalf(format string, v ...interface{}) {
	Default().Output(3, level.Critical, nil, fmt.Sprintf(format, v...))
}

func Errorf(format string, v ...interface{}) {
	Default().Output(3, level.Error, nil, fmt.Sprintf(format, v...))
}

func Warningf(format string, v ...interface{}) {
	Default().Output(3, level.Warning, nil, fmt.Sprintf(format, v...))
}

func Noticef(format string, v ...interface{}) {
	Default().Output(3, level.Notice, nil, fmt.Sprintf(format, v...))
}

func Infof(format string, v ...interface{}) {
	Default().Output(3, level.Info, nil, fmt.Sprintf(format, v...))
}

func Debugf(format string, v ...interface{}) {
	Default().Output(3, level.Debug, nil, fmt.Sprintf(format, v...))
}
