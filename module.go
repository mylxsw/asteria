package asteria

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Logger 日志对象
type Logger struct {
	moduleName    string
	level         func() Level
	formatter     Formatter
	writer        Writer
	timeLocation  func() *time.Location
	colorful      func() bool
	fileLine      func() bool
	globalContext func() func(c LogContext)

	lock sync.RWMutex
}

var loggers = make(map[string]*Logger)
var moduleLock sync.RWMutex

// DefaultConfig 默认配置对象
type DefaultConfig struct {
	LogLevel      Level
	LogFormatter  Formatter
	LogWriter     Writer
	TimeLocation  *time.Location
	Colorful      bool
	WithFileLine  bool
	GlobalContext func(c LogContext)
}

// 默认配置信息
var defaultLogConfig = DefaultConfig{
	LogLevel:     LevelDebug,
	LogFormatter: NewDefaultFormatter(),
	LogWriter:    NewDefaultWriter(),
	TimeLocation: time.Local,
	Colorful:     true,
	WithFileLine: false,
}

// Default return default log config
func Default() DefaultConfig {
	return defaultLogConfig
}

// DefaultWithFileLine set whether output file & line
func DefaultWithFileLine(enable bool) {
	moduleLock.Lock()
	defer moduleLock.Unlock()

	defaultLogConfig.WithFileLine = enable
}

// DefaultLocation set default time location
func DefaultLocation(loc *time.Location) {
	moduleLock.Lock()
	defer moduleLock.Unlock()

	defaultLogConfig.TimeLocation = loc
}

// DefaultWithColor set default Colorful property
func DefaultWithColor(colorful bool) {
	moduleLock.Lock()
	defer moduleLock.Unlock()

	defaultLogConfig.Colorful = colorful
}

// DefaultLogLevel 设置全局默认日志输出级别
func DefaultLogLevel(level Level) {
	moduleLock.Lock()
	defer moduleLock.Unlock()

	defaultLogConfig.LogLevel = level
}

// DefaultLogFormatter 设置全局默认的日志输出格式化器
func DefaultLogFormatter(formatter Formatter) {
	moduleLock.Lock()
	defer moduleLock.Unlock()

	defaultLogConfig.LogFormatter = formatter
}

// DefaultLogWriter 设置全局默认的日志输出器
func DefaultLogWriter(writer Writer) {
	moduleLock.Lock()
	defer moduleLock.Unlock()

	defaultLogConfig.LogWriter = writer
}

// GlobalContext set a global context
func GlobalContext(f func(c LogContext)) {
	moduleLock.Lock()
	defer moduleLock.Unlock()

	defaultLogConfig.GlobalContext = f
}

// Module 获取指定模块的日志输出对象
func Module(moduleName string) *Logger {
	moduleLock.Lock()
	defer moduleLock.Unlock()

	if logger, ok := loggers[moduleName]; ok {
		return logger
	}

	logger := &Logger{
		moduleName: moduleName,
		formatter:  defaultLogConfig.LogFormatter,
		writer:     defaultLogConfig.LogWriter,
		level: func() Level {
			moduleLock.RLock()
			defer moduleLock.RUnlock()

			return defaultLogConfig.LogLevel
		},
		timeLocation: func() *time.Location {
			moduleLock.RLock()
			defer moduleLock.RUnlock()

			return defaultLogConfig.TimeLocation
		},
		colorful: func() bool {
			moduleLock.RLock()
			defer moduleLock.RUnlock()

			return defaultLogConfig.Colorful
		},
		fileLine: func() bool {
			moduleLock.RLock()
			defer moduleLock.RUnlock()

			return defaultLogConfig.WithFileLine
		},
		globalContext: func() func(c LogContext) {
			moduleLock.RLock()
			defer moduleLock.RUnlock()

			return defaultLogConfig.GlobalContext
		},
	}

	loggers[moduleName] = logger

	return logger
}

// Location set time location for module
func (module *Logger) Location(loc *time.Location) *Logger {
	module.lock.Lock()
	defer module.lock.Unlock()

	module.timeLocation = func() *time.Location {
		return loc
	}

	return module
}

// WithFileLine set whether output file & line
func (module *Logger) WithFileLine(enable bool) *Logger {
	module.lock.Lock()
	defer module.lock.Unlock()

	module.fileLine = func() bool {
		return enable
	}

	return module
}

// GlobalContext set a global context
func (module *Logger) GlobalContext(f func(c LogContext)) *Logger {
	module.lock.Lock()
	defer module.lock.Unlock()

	module.globalContext = func() func(c LogContext) {
		return f
	}

	return module
}

// WithColor set Colorful property
func (module *Logger) WithColor(colorful bool) *Logger {
	module.lock.Lock()
	defer module.lock.Unlock()

	module.colorful = func() bool {
		return colorful
	}

	return module
}

func (module *Logger) Output(callDepth int, level Level, userContext C, v ...interface{}) string {
	if userContext == nil {
		userContext = C{}
	}

	logCtx := LogContext{
		UserContext: userContext,
		SysContext:  C{},
	}

	if module.fileLine() {
		_, f, line, _ := runtime.Caller(callDepth)
		logCtx.SysContext["file"] = f
		logCtx.SysContext["line"] = line
	}

	if module.globalContext != nil {
		cf := module.globalContext()
		if cf != nil {
			cf(logCtx)
		}
	}

	message := module.getFormatter().Format(module.colorful(), time.Now().In(module.timeLocation()), module.moduleName, level, logCtx, v...)
	// 低于设定日志级别的日志不会输出
	if level >= module.level() {
		if err := module.getWriter().Write(level, message); err != nil {
			fmt.Printf("can not write to Output: %s", err)
		}
	}

	return message
}

// GetDefaultModule 获取默认的模块日志
func GetDefaultModule() *Logger {
	return Module("default")
}

// LogLevel 设置日志输出级别
func (module *Logger) LogLevel(level Level) *Logger {
	module.lock.Lock()
	defer module.lock.Unlock()

	module.level = func() Level {
		return level
	}

	return module
}

// Formatter 设置日志格式化器
func (module *Logger) Formatter(formatter Formatter) *Logger {
	module.lock.Lock()
	defer module.lock.Unlock()

	module.formatter = formatter
	return module
}

func (module *Logger) getFormatter() Formatter {
	module.lock.RLock()
	defer module.lock.RUnlock()

	return module.formatter
}

// Writer 设置日志输出器
func (module *Logger) Writer(writer Writer) *Logger {
	module.lock.Lock()
	defer module.lock.Unlock()

	module.writer = writer
	return module
}

func (module *Logger) getWriter() Writer {
	module.lock.RLock()
	defer module.lock.RUnlock()

	return module.writer
}

// ReOpen reopen a log file
func (module *Logger) ReOpen() error {
	return module.getWriter().ReOpen()
}

// Close close a log LogWriter
func (module *Logger) Close() error {
	return module.getWriter().Close()
}

// WithContext 带有上下文信息的日志输出
func (module *Logger) WithContext(context C) *ContextLogger {
	return &ContextLogger{
		logger:  module,
		context: context,
	}
}

// Emergency 记录emergency日志
func (module *Logger) Emergency(v ...interface{}) string {
	return module.Output(2, LevelEmergency, nil, v...)
}

// Alert 记录Alert日志
func (module *Logger) Alert(v ...interface{}) string {
	return module.Output(2, LevelAlert, nil, v...)
}

// Critical 记录Critical日志
func (module *Logger) Critical(v ...interface{}) string {
	return module.Output(2, LevelCritical, nil, v...)
}

// Error 记录Error日志
func (module *Logger) Error(v ...interface{}) string {
	return module.Output(2, LevelError, nil, v...)
}

// Warning 记录Warning日志
func (module *Logger) Warning(v ...interface{}) string {
	return module.Output(2, LevelWarning, nil, v...)
}

// Notice 记录Notice日志
func (module *Logger) Notice(v ...interface{}) string {
	return module.Output(2, LevelNotice, nil, v...)
}

// Info 记录Info日志
func (module *Logger) Info(v ...interface{}) string {
	return module.Output(2, LevelInfo, nil, v...)
}

// Debug 记录Debug日志
func (module *Logger) Debug(v ...interface{}) string {
	return module.Output(2, LevelDebug, nil, v...)
}

// Emergencyf 记录emergency日志
func (module *Logger) Emergencyf(format string, v ...interface{}) string {
	return module.Output(2, LevelEmergency, nil, fmt.Sprintf(format, v...))
}

// Alertf 记录Alert日志
func (module *Logger) Alertf(format string, v ...interface{}) string {
	return module.Output(2, LevelAlert, nil, fmt.Sprintf(format, v...))
}

// Criticalf 记录critical日志
func (module *Logger) Criticalf(format string, v ...interface{}) string {
	return module.Output(2, LevelCritical, nil, fmt.Sprintf(format, v...))
}

// Errorf 记录error日志
func (module *Logger) Errorf(format string, v ...interface{}) string {
	return module.Output(2, LevelError, nil, fmt.Sprintf(format, v...))
}

// Warningf 记录warning日志
func (module *Logger) Warningf(format string, v ...interface{}) string {
	return module.Output(2, LevelWarning, nil, fmt.Sprintf(format, v...))
}

// Noticef 记录notice日志
func (module *Logger) Noticef(format string, v ...interface{}) string {
	return module.Output(2, LevelNotice, nil, fmt.Sprintf(format, v...))
}

// Infof 记录info日志
func (module *Logger) Infof(format string, v ...interface{}) string {
	return module.Output(2, LevelInfo, nil, fmt.Sprintf(format, v...))
}

// Debugf 记录debug日志
func (module *Logger) Debugf(format string, v ...interface{}) string {
	return module.Output(2, LevelDebug, nil, fmt.Sprintf(format, v...))
}

// Print 使用debug模式输出日志，为了兼容其它项目框架等
func (module *Logger) Print(v ...interface{}) {
	module.Output(2, LevelDebug, nil, v...)
}
