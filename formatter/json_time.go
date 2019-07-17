package formatter

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/mylxsw/asteria/level"
)

// JSONWithTimeFormatter json输格式化
type JSONWithTimeFormatter struct{}

// NewJSONFormatter create a new json LogFormatter
func NewJSONWithTimeFormatter() *JSONWithTimeFormatter {
	return &JSONWithTimeFormatter{}
}

// Format 日志格式化
func (formatter JSONWithTimeFormatter) Format(f Format) string {
	datetime := f.Time.Format(time.RFC3339)

	res, _ := json.Marshal(jsonOutput{
		DateTime:   datetime,
		Message:    fmt.Sprint(f.Messages...),
		Level:      f.Level,
		ModuleName: f.Module,
		LevelName:  level.GetLevelName(f.Level),
		Context:    createContext(f.Context),
	})

	return fmt.Sprintf("[%s] %s", datetime, string(res))
}
