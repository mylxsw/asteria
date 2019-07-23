package misc_test

import (
	"testing"

	"github.com/mylxsw/asteria/color"
	"github.com/mylxsw/asteria/level"
	"github.com/mylxsw/asteria/misc"
	"github.com/stretchr/testify/assert"
)

func TestColorfulLevelName(t *testing.T) {
	assert.Equal(t, color.BackgroundWrap(color.LightWhite, color.Blue, "[DEBG]"), misc.ColorfulLevelName(level.Debug))

	misc.SetLevelWithColor(level.Debug, color.Green, color.Magenta)
	assert.Equal(t, color.BackgroundWrap(color.Green, color.Magenta, "[DEBG]"), misc.ColorfulLevelName(level.Debug))
}
