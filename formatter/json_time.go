package formatter

import (
	"fmt"
	"time"

	"github.com/mylxsw/asteria/event"
)

// JSONWithTimeFormatter json输格式化
type JSONWithTimeFormatter struct{}

// NewJSONWithTimeFormatter create a new json LogFormatter
func NewJSONWithTimeFormatter() *JSONWithTimeFormatter {
	return &JSONWithTimeFormatter{}
}

// Format 日志格式化
func (formatter JSONWithTimeFormatter) Format(f event.Event) string {
	datetime := f.Time.Format(time.RFC3339)

	res, _ := json.Marshal(jsonOutput{
		DateTime:   datetime,
		Message:    fmt.Sprint(f.Messages...),
		Level:      f.Level,
		ModuleName: f.Module,
		LevelName:  f.Level.GetLevelName(),
		Context:    f.Fields.ToMap(),
	})

	return fmt.Sprintf("[%s] %s", datetime, string(res))
}
