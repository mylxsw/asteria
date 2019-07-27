package log

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/mylxsw/asteria/event"
	"github.com/mylxsw/asteria/formatter"
	"github.com/mylxsw/asteria/level"
	"github.com/mylxsw/asteria/misc"
	"github.com/mylxsw/asteria/writer"
)

type Filter func(f event.Event)
type FilterChain func(filter Filter) Filter

// Logger 日志对象
type Logger struct {
	moduleName    string
	level         func() level.Level
	formatter     formatter.Formatter
	writer        writer.Writer
	timeLocation  func() *time.Location
	fileLine      func() bool
	globalContext func() func(c event.Fields)
	filters       []FilterChain

	lock sync.RWMutex
}

// AddFilter append a filter to logger
func (module *Logger) AddFilter(f ...FilterChain) {
	module.lock.Lock()
	defer module.lock.Unlock()

	module.filters = append(module.filters, f...)
}

// Filters return all filters
func (module *Logger) Filters() []FilterChain {
	module.lock.RLock()
	defer module.lock.RUnlock()

	return module.filters
}

// AddGlobalFilter add a global filter
func AddGlobalFilter(f ...FilterChain) {
	moduleLock.Lock()
	defer moduleLock.Unlock()

	defaultLogConfig.GlobalFilters = append(defaultLogConfig.GlobalFilters, f...)
}

// GlobalFilters return all global filters
func GlobalFilters() []FilterChain {
	moduleLock.RLock()
	defer moduleLock.RUnlock()

	return defaultLogConfig.GlobalFilters
}

var loggers = make(map[string]*Logger)
var moduleLock sync.RWMutex

// DefaultConfig 默认配置对象
type DefaultConfig struct {
	LogLevel      level.Level
	LogFormatter  formatter.Formatter
	LogWriter     writer.Writer
	TimeLocation  *time.Location
	WithFileLine  bool
	GlobalFields  func(c event.Fields)
	GlobalFilters []FilterChain
}

// 默认配置信息
var defaultLogConfig DefaultConfig

// Reset all configuration for logger
func Reset() {
	moduleLock.Lock()
	defer moduleLock.Unlock()

	defaultLogConfig = DefaultConfig{
		LogLevel:      level.Debug,
		LogFormatter:  formatter.NewDefaultFormatter(true),
		LogWriter:     writer.NewStdoutWriter(),
		TimeLocation:  time.Local,
		WithFileLine:  false,
		GlobalFilters: make([]FilterChain, 0),
	}

	loggers = make(map[string]*Logger)
}

// GetDefaultConfig return default log config
func GetDefaultConfig() DefaultConfig {
	return defaultLogConfig
}

// DefaultWithFileLine set whether output file & Line
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

// DefaultLogLevel 设置全局默认日志输出级别
func DefaultLogLevel(l level.Level) {
	moduleLock.Lock()
	defer moduleLock.Unlock()

	defaultLogConfig.LogLevel = l
}

// DefaultLogFormatter 设置全局默认的日志输出格式化器
func DefaultLogFormatter(f formatter.Formatter) {
	moduleLock.Lock()
	defer moduleLock.Unlock()

	defaultLogConfig.LogFormatter = f
}

// DefaultLogWriter 设置全局默认的日志输出器
func DefaultLogWriter(w writer.Writer) {
	moduleLock.Lock()
	defer moduleLock.Unlock()

	defaultLogConfig.LogWriter = w
}

// GlobalFields set global fields
func GlobalFields(f func(c event.Fields)) {
	moduleLock.Lock()
	defer moduleLock.Unlock()

	defaultLogConfig.GlobalFields = f
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
		level: func() level.Level {
			moduleLock.RLock()
			defer moduleLock.RUnlock()

			return defaultLogConfig.LogLevel
		},
		timeLocation: func() *time.Location {
			moduleLock.RLock()
			defer moduleLock.RUnlock()

			return defaultLogConfig.TimeLocation
		},
		fileLine: func() bool {
			moduleLock.RLock()
			defer moduleLock.RUnlock()

			return defaultLogConfig.WithFileLine
		},
		globalContext: func() func(c event.Fields) {
			moduleLock.RLock()
			defer moduleLock.RUnlock()

			return defaultLogConfig.GlobalFields
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

// WithFileLine set whether output file & Line
func (module *Logger) WithFileLine(enable bool) *Logger {
	module.lock.Lock()
	defer module.lock.Unlock()

	module.fileLine = func() bool {
		return enable
	}

	return module
}

// GlobalFields set global fields
func (module *Logger) GlobalFields(f func(c event.Fields)) *Logger {
	module.lock.Lock()
	defer module.lock.Unlock()

	module.globalContext = func() func(c event.Fields) {
		return f
	}

	return module
}

func (module *Logger) Output(callDepth int, le level.Level, userContext Fields, v ...interface{}) {
	if le > module.level() {
		return
	}

	if userContext == nil {
		userContext = Fields{}
	}

	logCtx := event.Fields{
		CustomFields: userContext,
		GlobalFields: Fields{},
	}

	moduleName := module.moduleName

	if moduleName == "" || module.fileLine() {
		cg := misc.CallGraph(callDepth)
		if module.fileLine() {
			logCtx.GlobalFields["file"] = cg.FileName
			logCtx.GlobalFields["line"] = cg.Line
			logCtx.GlobalFields["package"] = cg.PackageName
		}

		if moduleName == "" {
			moduleName = strings.Replace(cg.PackageName, "/", ".", -1)
		}
	}

	if module.globalContext != nil {
		cf := module.globalContext()
		if cf != nil {
			cf(logCtx)
		}
	}

	f := event.Event{
		Time:     time.Now().In(module.timeLocation()),
		Module:   moduleName,
		Level:    le,
		Fields:   logCtx,
		Messages: v,
	}

	var chain Filter = func(f event.Event) {
		message := module.getFormatter().Format(f)
		if err := module.getWriter().Write(le, message); err != nil {
			panic(fmt.Sprintf("can not write to output: %s", err))
		}
	}

	filters := append(GlobalFilters(), module.Filters()...)
	for i := range filters {
		ff := filters[len(filters)-i-1]
		chain = ff(chain)
	}

	chain(f)
}

// Default 获取默认的模块日志
func Default() *Logger {
	return Module("")
}

// LogLevel 设置日志输出级别
func (module *Logger) LogLevel(le level.Level) *Logger {
	module.lock.Lock()
	defer module.lock.Unlock()

	module.level = func() level.Level {
		return le
	}

	return module
}

// Formatter 设置日志格式化器
func (module *Logger) Formatter(f formatter.Formatter) *Logger {
	module.lock.Lock()
	defer module.lock.Unlock()

	module.formatter = f
	return module
}

func (module *Logger) getFormatter() formatter.Formatter {
	module.lock.RLock()
	defer module.lock.RUnlock()

	return module.formatter
}

// Writer 设置日志输出器
func (module *Logger) Writer(w writer.Writer) *Logger {
	module.lock.Lock()
	defer module.lock.Unlock()

	module.writer = w
	return module
}

func (module *Logger) getWriter() writer.Writer {
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

// WithFields 带有上下文信息的日志输出
func (module *Logger) WithFields(c Fields) *ContextLogger {
	return &ContextLogger{
		logger:  module,
		context: c,
	}
}

func (module *Logger) Emergency(v ...interface{}) {
	module.Output(3, level.Emergency, nil, v...)
}

func (module *Logger) Alert(v ...interface{}) {
	module.Output(3, level.Alert, nil, v...)
}

func (module *Logger) Critical(v ...interface{}) {
	module.Output(3, level.Critical, nil, v...)
}

func (module *Logger) Error(v ...interface{}) {
	module.Output(3, level.Error, nil, v...)
}

func (module *Logger) Warning(v ...interface{}) {
	module.Output(3, level.Warning, nil, v...)
}

func (module *Logger) Notice(v ...interface{}) {
	module.Output(3, level.Notice, nil, v...)
}

func (module *Logger) Info(v ...interface{}) {
	module.Output(3, level.Info, nil, v...)
}

func (module *Logger) Debug(v ...interface{}) {
	module.Output(3, level.Debug, nil, v...)
}

func (module *Logger) Emergencyf(format string, v ...interface{}) {
	module.Output(3, level.Emergency, nil, fmt.Sprintf(format, v...))
}

func (module *Logger) Alertf(format string, v ...interface{}) {
	module.Output(3, level.Alert, nil, fmt.Sprintf(format, v...))
}

func (module *Logger) Criticalf(format string, v ...interface{}) {
	module.Output(3, level.Critical, nil, fmt.Sprintf(format, v...))
}

func (module *Logger) Errorf(format string, v ...interface{}) {
	module.Output(3, level.Error, nil, fmt.Sprintf(format, v...))
}

func (module *Logger) Warningf(format string, v ...interface{}) {
	module.Output(3, level.Warning, nil, fmt.Sprintf(format, v...))
}

func (module *Logger) Noticef(format string, v ...interface{}) {
	module.Output(3, level.Notice, nil, fmt.Sprintf(format, v...))
}

func (module *Logger) Infof(format string, v ...interface{}) {
	module.Output(3, level.Info, nil, fmt.Sprintf(format, v...))
}

func (module *Logger) Debugf(format string, v ...interface{}) {
	module.Output(3, level.Debug, nil, fmt.Sprintf(format, v...))
}

func (module *Logger) Print(v ...interface{}) {
	module.Output(3, level.Debug, nil, v...)
}
