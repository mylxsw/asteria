package log

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/mylxsw/asteria/filter"

	"github.com/mylxsw/asteria/event"
	"github.com/mylxsw/asteria/formatter"
	"github.com/mylxsw/asteria/level"
	"github.com/mylxsw/asteria/misc"
	"github.com/mylxsw/asteria/writer"
)

var loggers = make(Loggers)
var moduleLock sync.RWMutex

// Loggers is a map holds all loggers
type Loggers map[string]*AsteriaLogger

// All return all loggers
func All() Loggers {
	return loggers
}

// DynamicModuleName set whether enable dynamic module name generate
func (loggers Loggers) DynamicModuleName(enable bool) {
	moduleLock.Lock()
	defer moduleLock.Unlock()

	for _, l := range loggers {
		l.dynamicModuleName = enable
	}

	defaultLogConfig.DynamicModuleName = enable
}

// WithFileLine set whether output file & Line
func (loggers Loggers) WithFileLine(enable bool) {
	moduleLock.Lock()
	defer moduleLock.Unlock()

	for _, l := range loggers {
		l.fileLine = enable
	}

	defaultLogConfig.WithFileLine = enable
}

// Location set default time location
func (loggers Loggers) Location(loc *time.Location) {
	moduleLock.Lock()
	defer moduleLock.Unlock()

	for _, l := range loggers {
		l.timeLocation = loc
	}

	defaultLogConfig.TimeLocation = loc
}

// LogLevel 设置全局默认日志输出级别
func (loggers Loggers) LogLevel(le level.Level) {
	moduleLock.Lock()
	defer moduleLock.Unlock()

	for _, l := range loggers {
		l.level = le
	}

	defaultLogConfig.LogLevel = le
}

// LogFormatter 设置全局默认的日志输出格式化器
func (loggers Loggers) LogFormatter(f formatter.Formatter) {
	moduleLock.Lock()
	defer moduleLock.Unlock()

	for _, l := range loggers {
		l.formatter = f
	}

	defaultLogConfig.LogFormatter = f
}

// LogWriter 设置全局默认的日志输出器
func (loggers Loggers) LogWriter(w writer.Writer) {
	moduleLock.Lock()
	defer moduleLock.Unlock()

	for _, l := range loggers {
		l.writer = w
	}

	defaultLogConfig.LogWriter = w
}

// AsteriaLogger 日志对象
type AsteriaLogger struct {
	moduleName        string
	level             level.Level
	formatter         formatter.Formatter
	writer            writer.Writer
	timeLocation      *time.Location
	dynamicModuleName bool
	fileLine          bool
	globalContext     func(c event.Fields)
	filters           []filter.Chain

	lock sync.RWMutex
}

func (module *AsteriaLogger) DebugEnabled() bool {
	return level.Debug <= module.level
}

func (module *AsteriaLogger) InfoEnabled() bool {
	return level.Info <= module.level
}

func (module *AsteriaLogger) NoticeEnabled() bool {
	return level.Notice <= module.level
}

func (module *AsteriaLogger) WarningEnabled() bool {
	return level.Warning <= module.level
}

func (module *AsteriaLogger) ErrorEnabled() bool {
	return level.Error <= module.level
}

func (module *AsteriaLogger) CriticalEnabled() bool {
	return level.Critical <= module.level
}

func (module *AsteriaLogger) AlertEnabled() bool {
	return level.Alert <= module.level
}

func (module *AsteriaLogger) EmergencyEnabled() bool {
	return level.Emergency <= module.level
}

// AddFilter append a filter to logger
func (module *AsteriaLogger) AddFilter(f ...filter.Chain) {
	module.lock.Lock()
	defer module.lock.Unlock()

	module.filters = append(module.filters, f...)
}

// Filters return all filters
func (module *AsteriaLogger) Filters() []filter.Chain {
	module.lock.RLock()
	defer module.lock.RUnlock()

	return module.filters
}

// AddGlobalFilter add a global filter
func AddGlobalFilter(f ...filter.Chain) {
	moduleLock.Lock()
	defer moduleLock.Unlock()

	defaultLogConfig.GlobalFilters = append(defaultLogConfig.GlobalFilters, f...)
}

// GlobalFilters return all global filters
func GlobalFilters() []filter.Chain {
	moduleLock.RLock()
	defer moduleLock.RUnlock()

	return defaultLogConfig.GlobalFilters
}

// DefaultConfig 默认配置对象
type DefaultConfig struct {
	LogLevel          level.Level
	LogFormatter      formatter.Formatter
	LogWriter         writer.Writer
	TimeLocation      *time.Location
	WithFileLine      bool
	DynamicModuleName bool
	GlobalFields      func(c event.Fields)
	GlobalFilters     []filter.Chain
}

// 默认配置信息
var defaultLogConfig DefaultConfig

// Reset all configuration for logger
func Reset() {
	moduleLock.Lock()
	defer moduleLock.Unlock()

	defaultLogConfig = DefaultConfig{
		LogLevel:          level.Debug,
		LogFormatter:      formatter.NewDefaultFormatter(true),
		LogWriter:         writer.NewStdoutWriter(),
		TimeLocation:      time.Local,
		WithFileLine:      false,
		DynamicModuleName: false,
		GlobalFilters:     make([]filter.Chain, 0),
	}

	loggers = make(Loggers)
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

// DefaultDynamicModuleName set if enable dynamic module name generate
func DefaultDynamicModuleName(enable bool) {
	moduleLock.Lock()
	defer moduleLock.Unlock()

	defaultLogConfig.DynamicModuleName = enable
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
func Module(moduleName string) *AsteriaLogger {
	moduleLock.Lock()
	defer moduleLock.Unlock()

	if logger, ok := loggers[moduleName]; ok {
		return logger
	}

	logger := &AsteriaLogger{
		moduleName:        moduleName,
		formatter:         defaultLogConfig.LogFormatter,
		writer:            defaultLogConfig.LogWriter,
		level:             defaultLogConfig.LogLevel,
		timeLocation:      defaultLogConfig.TimeLocation,
		dynamicModuleName: defaultLogConfig.DynamicModuleName,
		fileLine:          defaultLogConfig.WithFileLine,
		globalContext:     defaultLogConfig.GlobalFields,
	}

	loggers[moduleName] = logger

	return logger
}

// Location set time location for module
func (module *AsteriaLogger) Location(loc *time.Location) *AsteriaLogger {
	module.lock.Lock()
	defer module.lock.Unlock()

	module.timeLocation = loc

	return module
}

// WithFileLine set whether output file & Line
func (module *AsteriaLogger) WithFileLine(enable bool) *AsteriaLogger {
	module.lock.Lock()
	defer module.lock.Unlock()

	module.fileLine = enable
	return module
}

// DynamicModuleName set whether enable dynamic module name generate
func (module *AsteriaLogger) DynamicModuleName(enable bool) *AsteriaLogger {
	module.lock.Lock()
	defer module.lock.Unlock()

	module.dynamicModuleName = enable
	return module
}

// GlobalFields set global fields
func (module *AsteriaLogger) GlobalFields(f func(c event.Fields)) *AsteriaLogger {
	module.lock.Lock()
	defer module.lock.Unlock()

	module.globalContext = f

	return module
}

func (module *AsteriaLogger) Output(callDepth int, le level.Level, userContext Fields, v ...interface{}) {
	if le > module.level {
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
	if module.dynamicModuleName || module.fileLine || !level.In(le, []level.Level{level.Debug, level.Info, level.Notice, level.Warning}) {
		cg := misc.CallGraph(callDepth)
		if module.fileLine || !level.In(le, []level.Level{level.Debug, level.Info, level.Notice, level.Warning}) {
			logCtx.GlobalFields["file"] = cg.FileName
			logCtx.GlobalFields["line"] = cg.Line
			logCtx.GlobalFields["package"] = cg.PackageName
		}

		if module.dynamicModuleName {
			moduleName = strings.Replace(cg.PackageName, "/", ".", -1)
		}
	}

	if module.globalContext != nil {
		cf := module.globalContext
		if cf != nil {
			cf(logCtx)
		}
	}

	f := event.Event{
		Time:     time.Now().In(module.timeLocation),
		Module:   moduleName,
		Level:    le,
		Fields:   logCtx,
		Messages: v,
	}

	var chain filter.Filter = func(f event.Event) {
		message := module.getFormatter().Format(f)
		if err := module.getWriter().Write(le, f.Module, message); err != nil {
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
func Default() *AsteriaLogger {
	return Module("main")
}

// LogLevel 设置日志输出级别
func (module *AsteriaLogger) LogLevel(le level.Level) *AsteriaLogger {
	module.lock.Lock()
	defer module.lock.Unlock()

	module.level = le

	return module
}

// Formatter 设置日志格式化器
func (module *AsteriaLogger) Formatter(f formatter.Formatter) *AsteriaLogger {
	module.lock.Lock()
	defer module.lock.Unlock()

	module.formatter = f
	return module
}

func (module *AsteriaLogger) getFormatter() formatter.Formatter {
	module.lock.RLock()
	defer module.lock.RUnlock()

	return module.formatter
}

// Writer 设置日志输出器
func (module *AsteriaLogger) Writer(w writer.Writer) *AsteriaLogger {
	module.lock.Lock()
	defer module.lock.Unlock()

	module.writer = w
	return module
}

func (module *AsteriaLogger) getWriter() writer.Writer {
	module.lock.RLock()
	defer module.lock.RUnlock()

	return module.writer
}

// ReOpen reopen a log file
func (module *AsteriaLogger) ReOpen() error {
	return module.getWriter().ReOpen()
}

// Close close a log LogWriter
func (module *AsteriaLogger) Close() error {
	return module.getWriter().Close()
}

func (module *AsteriaLogger) F(fields M) Logger {
	return module.WithFields(Fields(fields))
}

func (module *AsteriaLogger) KV(kvs ...interface{}) Logger {
	f := Fields{}
	for i := 0; i < len(kvs); i += 2 {
		if i+1 >= len(kvs) {
			f[fmt.Sprintf("%v", kvs[i])] = nil
		} else {
			f[fmt.Sprintf("%v", kvs[i])] = kvs[i+1]
		}
	}

	return module.WithFields(f)
}

// WithFields 带有上下文信息的日志输出
func (module *AsteriaLogger) WithFields(c Fields) Logger {
	return &ContextLogger{
		logger:  module,
		context: c,
	}
}

func (module *AsteriaLogger) With(data interface{}) Logger {
	return module.WithFields(Fields{
		"@": data,
	})
}

func (module *AsteriaLogger) Emergency(v ...interface{}) {
	module.Output(3, level.Emergency, nil, v...)
}

func (module *AsteriaLogger) Alert(v ...interface{}) {
	module.Output(3, level.Alert, nil, v...)
}

func (module *AsteriaLogger) Critical(v ...interface{}) {
	module.Output(3, level.Critical, nil, v...)
}

func (module *AsteriaLogger) Error(v ...interface{}) {
	module.Output(3, level.Error, nil, v...)
}

func (module *AsteriaLogger) Warning(v ...interface{}) {
	module.Output(3, level.Warning, nil, v...)
}

func (module *AsteriaLogger) Notice(v ...interface{}) {
	module.Output(3, level.Notice, nil, v...)
}

func (module *AsteriaLogger) Info(v ...interface{}) {
	module.Output(3, level.Info, nil, v...)
}

func (module *AsteriaLogger) Debug(v ...interface{}) {
	module.Output(3, level.Debug, nil, v...)
}

func (module *AsteriaLogger) Emergencyf(format string, v ...interface{}) {
	module.Output(3, level.Emergency, nil, fmt.Sprintf(format, v...))
}

func (module *AsteriaLogger) Alertf(format string, v ...interface{}) {
	module.Output(3, level.Alert, nil, fmt.Sprintf(format, v...))
}

func (module *AsteriaLogger) Criticalf(format string, v ...interface{}) {
	module.Output(3, level.Critical, nil, fmt.Sprintf(format, v...))
}

func (module *AsteriaLogger) Errorf(format string, v ...interface{}) {
	module.Output(3, level.Error, nil, fmt.Sprintf(format, v...))
}

func (module *AsteriaLogger) Warningf(format string, v ...interface{}) {
	module.Output(3, level.Warning, nil, fmt.Sprintf(format, v...))
}

func (module *AsteriaLogger) Noticef(format string, v ...interface{}) {
	module.Output(3, level.Notice, nil, fmt.Sprintf(format, v...))
}

func (module *AsteriaLogger) Infof(format string, v ...interface{}) {
	module.Output(3, level.Info, nil, fmt.Sprintf(format, v...))
}

func (module *AsteriaLogger) Debugf(format string, v ...interface{}) {
	module.Output(3, level.Debug, nil, fmt.Sprintf(format, v...))
}

func (module *AsteriaLogger) Print(v ...interface{}) {
	module.Output(3, level.Debug, nil, v...)
}
