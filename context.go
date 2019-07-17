package asteria

import (
	"fmt"

	"github.com/mylxsw/asteria/level"
)

// ContextLogger 带有上下文信息的日志输出
type ContextLogger struct {
	logger  *Logger
	context C
}

// Emergency 记录emergency日志
func (logger *ContextLogger) Emergency(v ...interface{}) string {
	return logger.logger.Output(3, level.Emergency, logger.context, v...)
}

// Alert 记录Alert日志
func (logger *ContextLogger) Alert(v ...interface{}) string {
	return logger.logger.Output(3, level.Alert, logger.context, v...)
}

// Critical 记录Critical日志
func (logger *ContextLogger) Critical(v ...interface{}) string {
	return logger.logger.Output(3, level.Critical, logger.context, v...)
}

// Error 记录Error日志
func (logger *ContextLogger) Error(v ...interface{}) string {
	return logger.logger.Output(3, level.Error, logger.context, v...)
}

// Warning 记录Warning日志
func (logger *ContextLogger) Warning(v ...interface{}) string {
	return logger.logger.Output(3, level.Warning, logger.context, v...)
}

// Notice 记录Notice日志
func (logger *ContextLogger) Notice(v ...interface{}) string {
	return logger.logger.Output(3, level.Notice, logger.context, v...)
}

// Info 记录Info日志
func (logger *ContextLogger) Info(v ...interface{}) string {
	return logger.logger.Output(3, level.Info, logger.context, v...)
}

// Debug 记录Debug日志
func (logger *ContextLogger) Debug(v ...interface{}) string {
	return logger.logger.Output(3, level.Debug, logger.context, v...)
}

// Emergencyf 记录emergency日志
func (logger *ContextLogger) Emergencyf(format string, v ...interface{}) string {
	return logger.logger.Output(3, level.Emergency, nil, fmt.Sprintf(format, v...))
}

// Alertf 记录Alert日志
func (logger *ContextLogger) Alertf(format string, v ...interface{}) string {
	return logger.logger.Output(3, level.Alert, nil, fmt.Sprintf(format, v...))
}

// Criticalf 记录critical日志
func (logger *ContextLogger) Criticalf(format string, v ...interface{}) string {
	return logger.logger.Output(3, level.Critical, nil, fmt.Sprintf(format, v...))
}

// Errorf 记录error日志
func (logger *ContextLogger) Errorf(format string, v ...interface{}) string {
	return logger.logger.Output(3, level.Error, nil, fmt.Sprintf(format, v...))
}

// Warningf 记录warning日志
func (logger *ContextLogger) Warningf(format string, v ...interface{}) string {
	return logger.logger.Output(3, level.Warning, nil, fmt.Sprintf(format, v...))
}

// Noticef 记录notice日志
func (logger *ContextLogger) Noticef(format string, v ...interface{}) string {
	return logger.logger.Output(3, level.Notice, nil, fmt.Sprintf(format, v...))
}

// Infof 记录info日志
func (logger *ContextLogger) Infof(format string, v ...interface{}) string {
	return logger.logger.Output(3, level.Info, nil, fmt.Sprintf(format, v...))
}

// Debugf 记录debug日志
func (logger *ContextLogger) Debugf(format string, v ...interface{}) string {
	return logger.logger.Output(3, level.Debug, nil, fmt.Sprintf(format, v...))
}
