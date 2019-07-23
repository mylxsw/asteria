package formatter

import (
	"github.com/mylxsw/asteria/event"
)

// Formatter 日志格式化接口
type Formatter interface {
	// Event 日志格式化
	Format(f event.Event) string
}
