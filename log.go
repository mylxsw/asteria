package asteria

import (
	"fmt"
)

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
	return GetDefaultModule().ReOpen()
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
	return GetDefaultModule().Close()
}

// LogLevel 设置日志输出级别
func SetLevel(level Level) *Logger {
	return GetDefaultModule().LogLevel(level)
}

// Formatter 设置日志格式化器
func SetFormatter(formatter Formatter) *Logger {
	return GetDefaultModule().Formatter(formatter)
}

// Writer 设置日志输出器
func SetWriter(writer Writer) *Logger {
	return GetDefaultModule().Writer(writer)
}

func WithContext(context C) *ContextLogger {
	return GetDefaultModule().WithContext(context)
}

func Emergency(v ...interface{}) string {
	return GetDefaultModule().Output(2, LevelEmergency, nil, v...)
}

func Alert(v ...interface{}) string {
	return GetDefaultModule().Output(2, LevelAlert, nil, v...)
}

func Critical(v ...interface{}) string {
	return GetDefaultModule().Output(2, LevelCritical, nil, v...)
}

func Error(v ...interface{}) string {
	return GetDefaultModule().Output(2, LevelError, nil, v...)
}

func Warning(v ...interface{}) string {
	return GetDefaultModule().Output(2, LevelWarning, nil, v...)
}

func Notice(v ...interface{}) string {
	return GetDefaultModule().Output(2, LevelNotice, nil, v...)
}

func Info(v ...interface{}) string {
	return GetDefaultModule().Output(2, LevelInfo, nil, v...)
}

func Debug(v ...interface{}) string {
	return GetDefaultModule().Output(2, LevelDebug, nil, v...)
}

func Emergencyf(format string, v ...interface{}) string {
	return GetDefaultModule().Output(2, LevelEmergency, nil, fmt.Sprintf(format, v...))
}

func Alertf(format string, v ...interface{}) string {
	return GetDefaultModule().Output(2, LevelAlert, nil, fmt.Sprintf(format, v...))
}

func Criticalf(format string, v ...interface{}) string {
	return GetDefaultModule().Output(2, LevelCritical, nil, fmt.Sprintf(format, v...))
}

func Errorf(format string, v ...interface{}) string {
	return GetDefaultModule().Output(2, LevelError, nil, fmt.Sprintf(format, v...))
}

func Warningf(format string, v ...interface{}) string {
	return GetDefaultModule().Output(2, LevelWarning, nil, fmt.Sprintf(format, v...))
}

func Noticef(format string, v ...interface{}) string {
	return GetDefaultModule().Output(2, LevelNotice, nil, fmt.Sprintf(format, v...))
}

func Infof(format string, v ...interface{}) string {
	return GetDefaultModule().Output(2, LevelInfo, nil, fmt.Sprintf(format, v...))
}

func Debugf(format string, v ...interface{}) string {
	return GetDefaultModule().Output(2, LevelDebug, nil, fmt.Sprintf(format, v...))
}
