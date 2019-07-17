package misc

import (
	"fmt"
	"sync"

	"github.com/mylxsw/asteria/color"
	"github.com/mylxsw/asteria/level"
)

var levelColorRef = map[level.Level][]color.Color{
	level.Debug:     {color.TextLightWhite, color.TextBlue},
	level.Info:      {color.TextLightWhite, color.TextCyan},
	level.Notice:    {color.TextLightWhite, color.TextYellow},
	level.Warning:   {color.TextRed, color.TextYellow},
	level.Error:     {color.TextLightWhite, color.TextRed},
	level.Critical:  {color.TextLightWhite, color.TextLightRed},
	level.Alert:     {color.TextLightWhite, color.TextLightRed},
	level.Emergency: {color.TextLightWhite, color.TextLightRed},
}

var levelColorRefLock sync.RWMutex

func ColorfulLevelName(le level.Level) string {
	levelColorRefLock.RLock()
	defer levelColorRefLock.RUnlock()

	levelName := fmt.Sprintf("[%s]", level.GetLevelNameAbbreviation(le))

	if lc, ok := levelColorRef[le]; ok {
		return color.ColorBackgroundFunc(lc[0], lc[1])(levelName)
	}

	return levelName
}

// SetLevelWithColor specify the color for level
func SetLevelWithColor(le level.Level, textColor color.Color, backgroundColor color.Color) {
	levelColorRefLock.Lock()
	defer levelColorRefLock.Unlock()

	levelColorRef[le] = []color.Color{textColor, backgroundColor}
}
