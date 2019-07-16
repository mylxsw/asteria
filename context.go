package asteria

import "fmt"

type LogContext struct {
	UserContext C
	SysContext C
}

type C map[string]interface{}

// ContextLogger 带有上下文信息的日志输出
type ContextLogger struct {
	logger  *Logger
	context C
}

// Emergency 记录emergency日志
func (logger *ContextLogger) Emergency(v ...interface{}) string {
	return logger.logger.Output(2, LevelEmergency, logger.context, v...)
}

// Alert 记录Alert日志
func (logger *ContextLogger) Alert(v ...interface{}) string {
	return logger.logger.Output(2, LevelAlert, logger.context, v...)
}

// Critical 记录Critical日志
func (logger *ContextLogger) Critical(v ...interface{}) string {
	return logger.logger.Output(2, LevelCritical, logger.context, v...)
}

// Error 记录Error日志
func (logger *ContextLogger) Error(v ...interface{}) string {
	return logger.logger.Output(2, LevelError, logger.context, v...)
}

// Warning 记录Warning日志
func (logger *ContextLogger) Warning(v ...interface{}) string {
	return logger.logger.Output(2, LevelWarning, logger.context, v...)
}

// Notice 记录Notice日志
func (logger *ContextLogger) Notice(v ...interface{}) string {
	return logger.logger.Output(2, LevelNotice, logger.context, v...)
}

// Info 记录Info日志
func (logger *ContextLogger) Info(v ...interface{}) string {
	return logger.logger.Output(2, LevelInfo, logger.context, v...)
}

// Debug 记录Debug日志
func (logger *ContextLogger) Debug(v ...interface{}) string {
	return logger.logger.Output(2, LevelDebug, logger.context, v...)
}

// Emergencyf 记录emergency日志
func (logger *ContextLogger) Emergencyf(format string, v ...interface{}) string {
	return logger.logger.Output(2, LevelEmergency, nil, fmt.Sprintf(format, v...))
}

// Alertf 记录Alert日志
func (logger *ContextLogger) Alertf(format string, v ...interface{}) string {
	return logger.logger.Output(2, LevelAlert, nil, fmt.Sprintf(format, v...))
}

// Criticalf 记录critical日志
func (logger *ContextLogger) Criticalf(format string, v ...interface{}) string {
	return logger.logger.Output(2, LevelCritical, nil, fmt.Sprintf(format, v...))
}

// Errorf 记录error日志
func (logger *ContextLogger) Errorf(format string, v ...interface{}) string {
	return logger.logger.Output(2, LevelError, nil, fmt.Sprintf(format, v...))
}

// Warningf 记录warning日志
func (logger *ContextLogger) Warningf(format string, v ...interface{}) string {
	return logger.logger.Output(2, LevelWarning, nil, fmt.Sprintf(format, v...))
}

// Noticef 记录notice日志
func (logger *ContextLogger) Noticef(format string, v ...interface{}) string {
	return logger.logger.Output(2, LevelNotice, nil, fmt.Sprintf(format, v...))
}

// Infof 记录info日志
func (logger *ContextLogger) Infof(format string, v ...interface{}) string {
	return logger.logger.Output(2, LevelInfo, nil, fmt.Sprintf(format, v...))
}

// Debugf 记录debug日志
func (logger *ContextLogger) Debugf(format string, v ...interface{}) string {
	return logger.logger.Output(2, LevelDebug, nil, fmt.Sprintf(format, v...))
}
