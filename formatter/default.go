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
func (formatter DefaultFormatter) Format(colorful bool, currentTime time.Time, moduleName string, le level.Level, logContext LogContext, v ...interface{}) string {
	var message string
	if colorful {
		message = fmt.Sprintf(
			"[%s] %s %-20s %s %s",
			currentTime.Format(time.RFC3339),
			misc.ColorfulLevelName(le),
			shortModuleName(moduleName),
			strings.Trim(fmt.Sprint(v...), "\n"),
			color.ColorTextWrap(color.TextLightGrey, formatContext(createContext(logContext))),
		)
	} else {
		message = fmt.Sprintf(
			"[%s] %s %s %s %s",
			currentTime.Format(time.RFC3339),
			level.GetLevelName(le),
			moduleName,
			strings.Trim(fmt.Sprint(v...), "\n"),
			formatContext(createContext(logContext)),
		)
	}

	// 将多行内容增加前缀tab，与第一行内容分开
	return strings.Trim(strings.Replace(message, "\n", "\n	", -1), "\n	")
}
