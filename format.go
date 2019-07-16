package asteria

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Formatter 日志格式化接口
type Formatter interface {
	// Format 日志格式化
	Format(colorful bool, currentTime time.Time, moduleName string, level Level, context LogContext, v ...interface{}) string
}

// DefaultFormatter 默认日志格式化
type DefaultFormatter struct{}

// NewDefaultFormatter create a new default formatter
func NewDefaultFormatter() *DefaultFormatter {
	return &DefaultFormatter{}
}

// Format 日志格式化
func (formatter DefaultFormatter) Format(colorful bool, currentTime time.Time, moduleName string, level Level, logContext LogContext, v ...interface{}) string {
	var message string
	if colorful {
		message = fmt.Sprintf(
			"[%s] %s %-20s %s %s",
			currentTime.Format(time.RFC3339),
			colorfulLevelName(level),
			shortModuleName(moduleName),
			strings.Trim(fmt.Sprint(v...), "\n"),
			ColorTextWrap(TextLightGrey, formatContext(createContext(logContext))),
		)
	} else {
		message = fmt.Sprintf(
			"[%s] %s %s %s %s",
			currentTime.Format(time.RFC3339),
			GetLevelName(level),
			moduleName,
			strings.Trim(fmt.Sprint(v...), "\n"),
			formatContext(createContext(logContext)),
		)
	}

	// 将多行内容增加前缀tab，与第一行内容分开
	return strings.Trim(strings.Replace(message, "\n", "\n	", -1), "\n	")
}

// JSONFormatter json输格式化
type JSONFormatter struct{}

// NewJSONFormatter create a new json formatter
func NewJSONFormatter() *JSONFormatter {
	return &JSONFormatter{}
}

type jsonOutput struct {
	ModuleName string `json:"module"`
	LevelName  string `json:"level_name"`
	Level      Level  `json:"level"`
	Context    C      `json:"context"`
	Message    string `json:"message"`
	DateTime   string `json:"datetime"`
}

// Format 日志格式化
func (formatter JSONFormatter) Format(colorful bool, currentTime time.Time, moduleName string, level Level, logContext LogContext, v ...interface{}) string {
	datetime := currentTime.Format(time.RFC3339)

	res, _ := json.Marshal(jsonOutput{
		DateTime:   datetime,
		Message:    fmt.Sprint(v...),
		Level:      level,
		ModuleName: moduleName,
		LevelName:  GetLevelName(level),
		Context:    createContext(logContext),
	})

	message := string(res)
	// if colorful {
	// 	datetime = ColorTextWrap(TextLightWhite, datetime)
	// 	message = ColorTextWrap(TextLightGrey, message)
	// }

	return fmt.Sprintf("[%s] %s", datetime, message)
}

func formatContext(context C) string {
	if context == nil {
		context = make(C)
	}

	contextJSON, _ := json.Marshal(context)
	return string(contextJSON)
}

func shortModuleName(moduleName string) string {
	segs := strings.Split(moduleName, ".")
	if len(segs) > 1 {
		ss := make([]string, 0)
		for _, s := range segs[:len(segs)-1] {
			if len(s) == 0 {
				continue
			}

			ss = append(ss, s[:1])
		}

		moduleName = strings.Join(append(ss, segs[len(segs)-1]), ".")
	} else {
		moduleName = strings.Join(segs, ".")
	}
	return moduleName
}

func createContext(logContext LogContext) C {
	cc := logContext.UserContext
	for k, v := range logContext.SysContext {
		cc["#" + k] = v
	}
	return cc
}
