package formatter

import (
	"fmt"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/mylxsw/asteria/event"
	"github.com/mylxsw/asteria/level"
)

var json = jsoniter.ConfigFastest

type jsonOutput struct {
	ModuleName string                 `json:"module"`
	LevelName  string                 `json:"level_name"`
	Level      level.Level            `json:"level"`
	Context    map[string]interface{} `json:"context"`
	Message    string                 `json:"message"`
	DateTime   string                 `json:"datetime"`
}

// JSONFormatter json输格式化
type JSONFormatter struct{}

// NewJSONFormatter create a new json LogFormatter
func NewJSONFormatter() *JSONFormatter {
	return &JSONFormatter{}
}

// Event 日志格式化
func (formatter JSONFormatter) Format(f event.Event) string {
	datetime := f.Time.Format(time.RFC3339)

	res, _ := json.Marshal(jsonOutput{
		DateTime:   datetime,
		Message:    fmt.Sprint(f.Messages...),
		Level:      f.Level,
		ModuleName: f.Module,
		LevelName:  f.Level.GetLevelName(),
		Context:    f.Fields.ToMap(),
	})

	return string(res)
}
