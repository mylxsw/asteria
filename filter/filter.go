package filter

import (
	"github.com/mylxsw/asteria/event"
	"github.com/mylxsw/asteria/level"
	"os"
	"runtime/debug"
)

type Filter func(evt event.Event)
type Chain func(filter Filter) Filter

// EmergencyExit 设置该 Filter 为 GlobalFilter 后，当发生 Emergency 级别的日志后，自动退出当前系统
func EmergencyExit(exitCode int) func(Filter) Filter {
	return func(filter Filter) Filter {
		return func(evt event.Event) {
			filter(evt)
			if evt.Level == level.Emergency {
				os.Exit(exitCode)
			}
		}
	}
}

// WithStacktrace 将日志级别为匹配的级别时，输出 stacktrace 信息到日志中
func WithStacktrace(levels ...level.Level) func(Filter) Filter {
	return func(filter Filter) Filter {
		return func(evt event.Event) {
			if matchLevels(evt.Level, levels) {
				evt.Fields.GlobalFields["stacktrace"] = string(debug.Stack())
			}
			filter(evt)
		}
	}
}

func matchLevels(le level.Level, levels []level.Level) bool {
	for _, l := range levels {
		if le == l {
			return true
		}
	}

	return false
}
