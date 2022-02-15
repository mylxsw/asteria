package formatter

import (
	"fmt"
	"strings"
	"time"

	"github.com/mylxsw/asteria/color"
	"github.com/mylxsw/asteria/event"
	"github.com/mylxsw/asteria/misc"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// DefaultFormatter 默认日志格式化
type DefaultFormatter struct {
	colorful                   bool
	stripCtrlAndExtFromUnicode bool
}

// NewDefaultFormatter create a new default LogFormatter
func NewDefaultFormatter(colorful bool) *DefaultFormatter {
	return &DefaultFormatter{colorful: colorful}
}

// NewDefaultCleanFormatter create a new default LogFormatter
func NewDefaultCleanFormatter(colorful bool) *DefaultFormatter {
	return &DefaultFormatter{colorful: colorful, stripCtrlAndExtFromUnicode: true}
}

// Format 日志格式化
func (formatter DefaultFormatter) Format(f event.Event) string {
	var message, messageBody string
	if formatter.stripCtrlAndExtFromUnicode {
		messageBody = StripCtlAndExtFromUnicode(strings.Trim(fmt.Sprint(f.Messages...), "\n"))
	} else {
		messageBody = strings.Trim(fmt.Sprint(f.Messages...), "\n")
	}

	if formatter.colorful {
		message = fmt.Sprintf(
			"[%s] %s %-20s %s %s",
			f.Time.Format(time.RFC3339),
			misc.ColorfulLevelName(f.Level),
			misc.ModuleNameAbbr(f.Module),
			messageBody,
			color.TextWrap(color.LightGrey, f.Fields.String()),
		)
	} else {
		message = fmt.Sprintf(
			"[%s] %s %s %s %s",
			f.Time.Format(time.RFC3339),
			f.Level.GetLevelName(),
			f.Module,
			messageBody,
			f.Fields.String(),
		)
	}

	// 将多行内容增加前缀tab，与第一行内容分开
	return strings.Trim(strings.Replace(message, "\n", "\n	", -1), "\n	")
}

// StripCtlAndExtFromUnicode Advanced Unicode normalization and filtering,
// see http://blog.golang.org/normalization and
// http://godoc.org/golang.org/x/text/unicode/norm for more
// details.
func StripCtlAndExtFromUnicode(str string) string {
	isOk := func(r rune) bool {
		return r < 32 || r >= 127
	}
	// The isOk filter is such that there is no need to chain to norm.NFC
	t := transform.Chain(norm.NFKD, runes.Remove(runes.Predicate(isOk)))
	// This Transformer could also trivially be applied as an io.Reader
	// or io.Writer filter to automatically do such filtering when reading
	// or writing data anywhere.
	str, _, _ = transform.String(t, str)
	return str
}
