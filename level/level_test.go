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
		"alert":     level.Alert,
		"Critical":  level.Critical,
		"errOR":     level.Error,
		"warning":   level.Warning,
		"abc":       level.Level(0),
	}

	for key, val := range testData {
		assert.Equal(t, val, level.GetLevelByName(key))
	}
}

func TestGetLevelName(t *testing.T) {
	var testData = map[level.Level]string{
		level.Alert:      "ALERT",
		level.Emergency:  "EMERGENCY",
		level.Debug:      "DEBUG",
		level.Warning:    "WARNING",
		level.Error:      "ERROR",
		level.Critical:   "CRITICAL",
		level.Notice:     "NOTICE",
		level.Info:       "INFO",
		level.Level(100): "UNKNOWN",
	}

	for key, val := range testData {
		assert.Equal(t, val, key.GetLevelName())
	}
}

func TestGetLevelNameAbbreviation(t *testing.T) {
	var testData = map[level.Level]string{
		level.Emergency:  "EMCY",
		level.Alert:      "ALER",
		level.Critical:   "CRIT",
		level.Error:      "EROR",
		level.Warning:    "WARN",
		level.Notice:     "NOTI",
		level.Info:       "INFO",
		level.Debug:      "DEBG",
		level.Level(100): "UNON",
	}

	for l, exp := range testData {
		assert.Equal(t, exp, l.GetLevelNameAbbreviation())
	}
}
