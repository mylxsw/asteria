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
func (formatter JSONWithTimeFormatter) Format(colorful bool, currentTime time.Time, moduleName string, le level.Level, logContext LogContext, v ...interface{}) string {
	datetime := currentTime.Format(time.RFC3339)

	res, _ := json.Marshal(jsonOutput{
		DateTime:   datetime,
		Message:    fmt.Sprint(v...),
		Level:      le,
		ModuleName: moduleName,
		LevelName:  level.GetLevelName(le),
		Context:    createContext(logContext),
	})

	message := string(res)
	// if Colorful {
	// 	datetime = ColorTextWrap(TextLightWhite, datetime)
	// 	message = ColorTextWrap(TextLightGrey, message)
	// }

	return fmt.Sprintf("[%s] %s", datetime, message)
}
