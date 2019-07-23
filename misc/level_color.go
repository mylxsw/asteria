package misc

import (
	"fmt"
	"sync"

	"github.com/mylxsw/asteria/color"
	"github.com/mylxsw/asteria/level"
)

var levelColorRef = map[level.Level][]color.Color{
	level.Debug:     {color.LightWhite, color.Blue},
	level.Info:      {color.LightWhite, color.Cyan},
	level.Notice:    {color.LightWhite, color.Yellow},
	level.Warning:   {color.Red, color.Yellow},
	level.Error:     {color.LightWhite, color.Red},
	level.Critical:  {color.LightWhite, color.LightRed},
	level.Alert:     {color.LightWhite, color.LightRed},
	level.Emergency: {color.LightWhite, color.LightRed},
}

var levelColorRefLock sync.RWMutex

func ColorfulLevelName(le level.Level) string {
	levelColorRefLock.RLock()
	defer levelColorRefLock.RUnlock()

	levelName := fmt.Sprintf("[%s]", le.GetLevelNameAbbreviation())

	if lc, ok := levelColorRef[le]; ok {
		return color.BackgroundFunc(lc[0], lc[1])(levelName)
	}

	return levelName
}

// SetLevelWithColor specify the color for level
func SetLevelWithColor(le level.Level, textColor color.Color, backgroundColor color.Color) {
	levelColorRefLock.Lock()
	defer levelColorRefLock.Unlock()

	levelColorRef[le] = []color.Color{textColor, backgroundColor}
}
