package asteria

import (
	"fmt"
	"strings"
	"sync"
)

type Level int

// 日志输出级别
const (
	LevelEmergency Level = 600
	LevelAlert           = 550
	LevelCritical        = 500
	LevelError           = 400
	LevelWarning         = 300
	LevelNotice          = 250
	LevelInfo            = 200
	LevelDebug           = 100
)

// GetLevelNameAbbreviation 获取日志级别缩写
func GetLevelNameAbbreviation(level Level) string {
	switch level {
	case LevelEmergency:
		return "EMCY"
	case LevelAlert:
		return "ALER"
	case LevelCritical:
		return "CRIT"
	case LevelError:
		return "EROR"
	case LevelWarning:
		return "WARN"
	case LevelNotice:
		return "NOTI"
	case LevelInfo:
		return "INFO"
	case LevelDebug:
		return "DEBG"
	}

	return "UNON"
}

// GetLevelName 获取日志级别名称
func GetLevelName(level Level) string {
	switch level {
	case LevelEmergency:
		return "EMERGENCY"
	case LevelAlert:
		return "ALERT"
	case LevelCritical:
		return "CRITICAL"
	case LevelError:
		return "ERROR"
	case LevelWarning:
		return "WARNING"
	case LevelNotice:
		return "NOTICE"
	case LevelInfo:
		return "INFO"
	case LevelDebug:
		return "DEBUG"
	}

	return "UNKNOWN"
}

// GetLevelByName 使用名称获取Level真实的数值
func GetLevelByName(levelName string) Level {
	switch strings.ToUpper(levelName) {
	case "EMERGENCY":
		return LevelEmergency
	case "ALERT":
		return LevelAlert
	case "CRITICAL":
		return LevelCritical
	case "ERROR":
		return LevelError
	case "WARNING":
		return LevelWarning
	case "NOTICE":
		return LevelNotice
	case "INFO":
		return LevelInfo
	case "DEBUG":
		return LevelDebug
	}

	return 0
}

var levelColorRef = map[Level][]Color{
	LevelDebug:     {TextLightWhite, TextBlue},
	LevelInfo:      {TextLightWhite, TextCyan},
	LevelNotice:    {TextLightWhite, TextYellow},
	LevelWarning:   {TextRed, TextYellow},
	LevelError:     {TextLightWhite, TextRed},
	LevelCritical:  {TextLightWhite, TextLightRed},
	LevelAlert:     {TextLightWhite, TextLightRed},
	LevelEmergency: {TextLightWhite, TextLightRed},
}

var levelColorRefLock sync.RWMutex

func colorfulLevelName(level Level) string {
	levelColorRefLock.RLock()
	defer levelColorRefLock.RUnlock()

	levelName := fmt.Sprintf("[%s]", GetLevelNameAbbreviation(level))

	if lc, ok := levelColorRef[level]; ok {
		return ColorBackgroundFunc(lc[0], lc[1])(levelName)
	}

	return levelName
}

// SetLevelWithColor specify the color for level
func SetLevelWithColor(level Level, textColor Color, backgroundColor Color) {
	levelColorRefLock.Lock()
	defer levelColorRefLock.Unlock()

	levelColorRef[level] = []Color{textColor, backgroundColor}
}
