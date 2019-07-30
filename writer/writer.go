package writer

import (
	"github.com/mylxsw/asteria/level"
)

// Writer 日志输出接口
type Writer interface {
	Write(le level.Level, module string, message string) error
	ReOpen() error
	Close() error
}
