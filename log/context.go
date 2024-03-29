package log

import (
	"fmt"

	"github.com/mylxsw/asteria/level"
)

type ContextLogger struct {
	logger  *AsteriaLogger
	context Fields
}

func (logger *ContextLogger) DebugEnabled() bool {
	return logger.logger.DebugEnabled()
}

func (logger *ContextLogger) InfoEnabled() bool {
	return logger.logger.InfoEnabled()
}

func (logger *ContextLogger) NoticeEnabled() bool {
	return logger.logger.NoticeEnabled()
}

func (logger *ContextLogger) WarningEnabled() bool {
	return logger.logger.WarningEnabled()
}

func (logger *ContextLogger) ErrorEnabled() bool {
	return logger.logger.ErrorEnabled()
}

func (logger *ContextLogger) CriticalEnabled() bool {
	return logger.logger.CriticalEnabled()
}

func (logger *ContextLogger) AlertEnabled() bool {
	return logger.logger.AlertEnabled()
}

func (logger *ContextLogger) EmergencyEnabled() bool {
	return logger.logger.EmergencyEnabled()
}

func (logger *ContextLogger) With(data interface{}) Logger {
	return logger.WithFields(Fields{
		"@": data,
	})
}

func (logger *ContextLogger) F(fields M) Logger {
	return logger.WithFields(Fields(fields))
}

func (logger *ContextLogger) KV(kvs ...interface{}) Logger {
	return logger.logger.KV(kvs...)
}

// WithFields 带有上下文信息的日志输出
func (logger *ContextLogger) WithFields(c Fields) Logger {
	c2 := make(Fields)
	for k, v := range logger.context {
		c2[k] = v
	}
	for k, v := range c {
		c2[k] = v
	}

	return &ContextLogger{
		logger:  logger.logger,
		context: c2,
	}
}

func (logger *ContextLogger) Emergency(v ...interface{}) {
	logger.logger.Output(3, level.Emergency, logger.context, v...)
}

func (logger *ContextLogger) Alert(v ...interface{}) {
	logger.logger.Output(3, level.Alert, logger.context, v...)
}

func (logger *ContextLogger) Critical(v ...interface{}) {
	logger.logger.Output(3, level.Critical, logger.context, v...)
}

func (logger *ContextLogger) Error(v ...interface{}) {
	logger.logger.Output(3, level.Error, logger.context, v...)
}

func (logger *ContextLogger) Warning(v ...interface{}) {
	logger.logger.Output(3, level.Warning, logger.context, v...)
}

func (logger *ContextLogger) Notice(v ...interface{}) {
	logger.logger.Output(3, level.Notice, logger.context, v...)
}

func (logger *ContextLogger) Info(v ...interface{}) {
	logger.logger.Output(3, level.Info, logger.context, v...)
}

func (logger *ContextLogger) Debug(v ...interface{}) {
	logger.logger.Output(3, level.Debug, logger.context, v...)
}

func (logger *ContextLogger) Emergencyf(format string, v ...interface{}) {
	logger.logger.Output(3, level.Emergency, logger.context, fmt.Sprintf(format, v...))
}

func (logger *ContextLogger) Alertf(format string, v ...interface{}) {
	logger.logger.Output(3, level.Alert, logger.context, fmt.Sprintf(format, v...))
}

func (logger *ContextLogger) Criticalf(format string, v ...interface{}) {
	logger.logger.Output(3, level.Critical, logger.context, fmt.Sprintf(format, v...))
}

func (logger *ContextLogger) Errorf(format string, v ...interface{}) {
	logger.logger.Output(3, level.Error, logger.context, fmt.Sprintf(format, v...))
}

func (logger *ContextLogger) Warningf(format string, v ...interface{}) {
	logger.logger.Output(3, level.Warning, logger.context, fmt.Sprintf(format, v...))
}

func (logger *ContextLogger) Noticef(format string, v ...interface{}) {
	logger.logger.Output(3, level.Notice, logger.context, fmt.Sprintf(format, v...))
}

func (logger *ContextLogger) Infof(format string, v ...interface{}) {
	logger.logger.Output(3, level.Info, logger.context, fmt.Sprintf(format, v...))
}

func (logger *ContextLogger) Debugf(format string, v ...interface{}) {
	logger.logger.Output(3, level.Debug, logger.context, fmt.Sprintf(format, v...))
}
