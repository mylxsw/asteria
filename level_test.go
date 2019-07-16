package asteria

import "testing"

func TestGetLevelByName(t *testing.T) {
	var testData = map[string]Level{
		"debug":     LevelDebug,
		"info":      LevelInfo,
		"emergency": LevelEmergency,
		"DEBUG":     LevelDebug,
		"NOTice":    LevelNotice,
	}

	for key, val := range testData {
		if GetLevelByName(key) != val {
			t.Errorf("测试结果出错: GetLevelByName(%s) != %d", key, val)
		}
	}

}

func TestGetLevelName(t *testing.T) {
	var testData = map[Level]string{
		LevelAlert:     "ALERT",
		LevelEmergency: "EMERGENCY",
		LevelDebug:     "DEBUG",
	}

	for key, val := range testData {
		if GetLevelName(key) != val {
			t.Errorf("测试结果出错: GetLevelName(%d) != %s", key, val)
		}
	}
}
