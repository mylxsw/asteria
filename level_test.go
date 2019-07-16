package asteria_test

import (
	"testing"

	"github.com/mylxsw/asteria"
)

func TestGetLevelByName(t *testing.T) {
	var testData = map[string]asteria.Level{
		"debug":     asteria.LevelDebug,
		"info":      asteria.LevelInfo,
		"emergency": asteria.LevelEmergency,
		"DEBUG":     asteria.LevelDebug,
		"NOTice":    asteria.LevelNotice,
	}

	for key, val := range testData {
		if asteria.GetLevelByName(key) != val {
			t.Errorf("测试结果出错: GetLevelByName(%s) != %d", key, val)
		}
	}

}

func TestGetLevelName(t *testing.T) {
	var testData = map[asteria.Level]string{
		asteria.LevelAlert:     "ALERT",
		asteria.LevelEmergency: "EMERGENCY",
		asteria.LevelDebug:     "DEBUG",
	}

	for key, val := range testData {
		if asteria.GetLevelName(key) != val {
			t.Errorf("测试结果出错: GetLevelName(%d) != %s", key, val)
		}
	}
}
