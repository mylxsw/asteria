package log

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/mylxsw/asteria/color"
	"github.com/mylxsw/asteria/formatter"
	"github.com/mylxsw/asteria/level"
	"github.com/mylxsw/asteria/misc"
	"github.com/mylxsw/asteria/writer"
)

type Filter func(f formatter.Format)
type FilterChain func(filter Filter) Filter

// Logger 日志对象
type Logger struct {
	moduleName    string
	level         func() level.Level
	formatter     formatter.Formatter
	writer        writer.Writer
	timeLocation  func() *time.Location
	colorful      func() bool
	fileLine      func() bool
	globalContext func() func(c formatter.LogContext)
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
	Colorful      bool
	WithFileLine  bool
	GlobalContext func(c formatter.LogContext)
	GlobalFilters []FilterChain
}

// 默认配置信息
var defaultLogConfig = DefaultConfig{
	LogLevel:      level.Debug,
	LogFormatter:  formatter.NewDefaultFormatter(),
	LogWriter:     writer.NewStdoutWriter(),
	TimeLocation:  time.Local,
	Colorful:      true,
	WithFileLine:  false,
	GlobalFilters: make([]FilterChain, 0),
}

// GetDefaultConfig return default log config
func GetDefaultConfig() DefaultConfig {
	return defaultLogConfig
}

// SetLevelWithColor specify the color for level
func SetLevelWithColor(le level.Level, textColor color.Color, backgroundColor color.Color) {
	misc.SetLevelWithColor(le, textColor, backgroundColor)
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

// DefaultWithColor set default Colorful property
func DefaultWithColor(colorful bool) {
	moduleLock.Lock()
	defer moduleLock.Unlock()

	defaultLogConfig.Colorful = colorful
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

// GlobalContext set a global context
func GlobalContext(f func(c formatter.LogContext)) {
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
		globalContext: func() func(c formatter.LogContext) {
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

// WithFileLine set whether output file & Line
func (module *Logger) WithFileLine(enable bool) *Logger {
	module.lock.Lock()
	defer module.lock.Unlock()

	module.fileLine = func() bool {
		return enable
	}

	return module
}

// GlobalContext set a global context
func (module *Logger) GlobalContext(f func(c formatter.LogContext)) *Logger {
	module.lock.Lock()
	defer module.lock.Unlock()

	module.globalContext = func() func(c formatter.LogContext) {
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

func (module *Logger) Output(callDepth int, le level.Level, userContext C, v ...interface{}) {
	if le < module.level() {
		return
	}

	if userContext == nil {
		userContext = C{}
	}

	logCtx := formatter.LogContext{
		UserContext: userContext,
		SysContext:  C{},
	}

	moduleName := module.moduleName

	if moduleName == "" || module.fileLine() {
		cg := misc.CallGraph(callDepth)
		if module.fileLine() {
			logCtx.SysContext["file"] = cg.FileName
			logCtx.SysContext["line"] = cg.Line
			logCtx.SysContext["package"] = cg.PackageName
			logCtx.SysContext["func"] = cg.FuncName
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

	f := formatter.Format{
		Colorful: module.colorful(),
		Time:     time.Now().In(module.timeLocation()),
		Module:   moduleName,
		Level:    le,
		Context:  logCtx,
		Messages: v,
	}

	var chain Filter = func(f formatter.Format) {
		message := module.getFormatter().Format(f)
		if err := module.getWriter().Write(le, message); err != nil {
			fmt.Printf("can not write to output: %s", err)
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

// WithContext 带有上下文信息的日志输出
func (module *Logger) WithContext(c C) *ContextLogger {
	return &ContextLogger{
		logger:  module,
		context: c,
	}
}

// Emergency 记录emergency日志
func (module *Logger) Emergency(v ...interface{}) {
	module.Output(3, level.Emergency, nil, v...)
}

// Alert 记录Alert日志
func (module *Logger) Alert(v ...interface{}) {
	module.Output(3, level.Alert, nil, v...)
}

// Critical 记录Critical日志
func (module *Logger) Critical(v ...interface{}) {
	module.Output(3, level.Critical, nil, v...)
}

// Error 记录Error日志
func (module *Logger) Error(v ...interface{}) {
	module.Output(3, level.Error, nil, v...)
}

// Warning 记录Warning日志
func (module *Logger) Warning(v ...interface{}) {
	module.Output(3, level.Warning, nil, v...)
}

// Notice 记录Notice日志
func (module *Logger) Notice(v ...interface{}) {
	module.Output(3, level.Notice, nil, v...)
}

// Info 记录Info日志
func (module *Logger) Info(v ...interface{}) {
	module.Output(3, level.Info, nil, v...)
}

// Debug 记录Debug日志
func (module *Logger) Debug(v ...interface{}) {
	module.Output(3, level.Debug, nil, v...)
}

// Emergencyf 记录emergency日志
func (module *Logger) Emergencyf(format string, v ...interface{}) {
	module.Output(3, level.Emergency, nil, fmt.Sprintf(format, v...))
}

// Alertf 记录Alert日志
func (module *Logger) Alertf(format string, v ...interface{}) {
	module.Output(3, level.Alert, nil, fmt.Sprintf(format, v...))
}

// Criticalf 记录critical日志
func (module *Logger) Criticalf(format string, v ...interface{}) {
	module.Output(3, level.Critical, nil, fmt.Sprintf(format, v...))
}

// Errorf 记录error日志
func (module *Logger) Errorf(format string, v ...interface{}) {
	module.Output(3, level.Error, nil, fmt.Sprintf(format, v...))
}

// Warningf 记录warning日志
func (module *Logger) Warningf(format string, v ...interface{}) {
	module.Output(3, level.Warning, nil, fmt.Sprintf(format, v...))
}

// Noticef 记录notice日志
func (module *Logger) Noticef(format string, v ...interface{}) {
	module.Output(3, level.Notice, nil, fmt.Sprintf(format, v...))
}

// Infof 记录info日志
func (module *Logger) Infof(format string, v ...interface{}) {
	module.Output(3, level.Info, nil, fmt.Sprintf(format, v...))
}

// Debugf 记录debug日志
func (module *Logger) Debugf(format string, v ...interface{}) {
	module.Output(3, level.Debug, nil, fmt.Sprintf(format, v...))
}

// Print 使用debug模式输出日志，为了兼容其它项目框架等
func (module *Logger) Print(v ...interface{}) {
	module.Output(3, level.Debug, nil, v...)
}
