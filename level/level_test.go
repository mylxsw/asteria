package level_test

import (
	"testing"

	"github.com/mylxsw/asteria/level"
	"github.com/stretchr/testify/assert"
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
		assert.Equal(t, val, level.GetLevelByName(key))
	}
}

func TestGetLevelName(t *testing.T) {
	var testData = map[level.Level]string{
		level.Alert:     "ALERT",
		level.Emergency: "EMERGENCY",
		level.Debug:     "DEBUG",
	}

	for key, val := range testData {
		assert.Equal(t, val, key.GetLevelName())
	}
}

func TestGetLevelNameAbbreviation(t *testing.T) {
	var testData = map[level.Level]string{
		level.Emergency: "EMCY",
		level.Alert:     "ALER",
		level.Critical:  "CRIT",
		level.Error:     "EROR",
		level.Warning:   "WARN",
		level.Notice:    "NOTI",
		level.Info:      "INFO",
		level.Debug:     "DEBG",
	}

	for l, exp := range testData {
		assert.Equal(t, exp, l.GetLevelNameAbbreviation())
	}
}
