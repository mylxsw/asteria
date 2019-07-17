package formatter

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/mylxsw/asteria/level"
)

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

// Format 日志格式化
func (formatter JSONFormatter) Format(colorful bool, currentTime time.Time, moduleName string, le level.Level, logContext LogContext, v ...interface{}) string {
	datetime := currentTime.Format(time.RFC3339)

	res, _ := json.Marshal(jsonOutput{
		DateTime:   datetime,
		Message:    fmt.Sprint(v...),
		Level:      le,
		ModuleName: moduleName,
		LevelName:  level.GetLevelName(le),
		Context:    createContext(logContext),
	})

	return string(res)
}
