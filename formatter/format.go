package formatter

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/mylxsw/asteria/level"
)

type LogContext struct {
	UserContext map[string]interface{}
	SysContext  map[string]interface{}
}

type Format struct {
	Colorful bool
	Time     time.Time
	Module   string
	Level    level.Level
	Context  LogContext
	Messages []interface{}
}

// Formatter 日志格式化接口
type Formatter interface {
	// Format 日志格式化
	Format(f Format) string
}

func formatContext(context map[string]interface{}) string {
	if context == nil {
		context = make(map[string]interface{})
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

func createContext(logContext LogContext) map[string]interface{} {
	cc := logContext.UserContext
	if cc == nil {
		cc = make(map[string]interface{}, 0)
	}

	for k, v := range logContext.SysContext {
		cc["#"+k] = v
	}
	return cc
}
