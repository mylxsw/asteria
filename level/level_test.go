package level_test

import (
	"testing"

	"github.com/mylxsw/asteria/level"
)

func TestGetLevelByName(t *testing.T) {
	var testData = map[string]level.Level{
		"debug":     level.Debug,
		"info":      level.Info,
		"emergency": level.Emergency,
		"DEBUG":     level.Debug,
		"NOTice":    level.Notice,
	}

	for key, val := range testData {
		if level.GetLevelByName(key) != val {
			t.Errorf("Test Failed: GetLevelByName(%s) != %d", key, val)
		}
	}

}

func TestGetLevelName(t *testing.T) {
	var testData = map[level.Level]string{
		level.Alert:     "ALERT",
		level.Emergency: "EMERGENCY",
		level.Debug:     "DEBUG",
	}

	for key, val := range testData {
		if level.GetLevelName(key) != val {
			t.Errorf("Test Failed: GetLevelName(%d) != %s", key, val)
		}
	}
}
