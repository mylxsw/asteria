package formatter

import (
	"fmt"
	"strings"
	"time"

	"github.com/mylxsw/asteria/color"
	"github.com/mylxsw/asteria/event"
	"github.com/mylxsw/asteria/misc"
)

// DefaultFormatter 默认日志格式化
type DefaultFormatter struct {
	colorful bool
}

// NewDefaultFormatter create a new default LogFormatter
func NewDefaultFormatter(colorful bool) *DefaultFormatter {
	return &DefaultFormatter{colorful: colorful}
}

// Event 日志格式化
func (formatter DefaultFormatter) Format(f event.Event) string {
	var message string
	if formatter.colorful {
		message = fmt.Sprintf(
			"[%s] %s %-20s %s %s",
			f.Time.Format(time.RFC3339),
			misc.ColorfulLevelName(f.Level),
			misc.ModuleNameAbbr(f.Module),
			strings.Trim(fmt.Sprint(f.Messages...), "\n"),
			color.TextWrap(color.LightGrey, f.Fields.String()),
		)
	} else {
		message = fmt.Sprintf(
			"[%s] %s %s %s %s",
			f.Time.Format(time.RFC3339),
			f.Level.GetLevelName(),
			f.Module,
			strings.Trim(fmt.Sprint(f.Messages...), "\n"),
			f.Fields.String(),
		)
	}

	// 将多行内容增加前缀tab，与第一行内容分开
	return strings.Trim(strings.Replace(message, "\n", "\n	", -1), "\n	")
}
