package formatter

import (
	"fmt"
	"strings"
	"time"

	"github.com/mylxsw/asteria/color"
	"github.com/mylxsw/asteria/level"
	"github.com/mylxsw/asteria/misc"
)

// DefaultFormatter 默认日志格式化
type DefaultFormatter struct{}

// NewDefaultFormatter create a new default LogFormatter
func NewDefaultFormatter() *DefaultFormatter {
	return &DefaultFormatter{}
}

// Format 日志格式化
func (formatter DefaultFormatter) Format(f Format) string {
	var message string
	if f.Colorful {
		message = fmt.Sprintf(
			"[%s] %s %-20s %s %s",
			f.Time.Format(time.RFC3339),
			misc.ColorfulLevelName(f.Level),
			shortModuleName(f.Module),
			strings.Trim(fmt.Sprint(f.Messages...), "\n"),
			color.TextWrap(color.TextLightGrey, formatContext(createContext(f.Context))),
		)
	} else {
		message = fmt.Sprintf(
			"[%s] %s %s %s %s",
			f.Time.Format(time.RFC3339),
			level.GetLevelName(f.Level),
			f.Module,
			strings.Trim(fmt.Sprint(f.Messages...), "\n"),
			formatContext(createContext(f.Context)),
		)
	}

	// 将多行内容增加前缀tab，与第一行内容分开
	return strings.Trim(strings.Replace(message, "\n", "\n	", -1), "\n	")
}
