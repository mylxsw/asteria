package filter

import (
	"github.com/mylxsw/asteria/event"
	"github.com/mylxsw/asteria/level"
	"os"
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
