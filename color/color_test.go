package color_test

import (
	"testing"

	"github.com/mylxsw/asteria/color"
	"github.com/stretchr/testify/assert"
)

func TestTextWrap(t *testing.T) {
	assert.Equal(t, color.TextWrap(color.Green, "Hello"), "\x1b[32mHello\x1b[0m", "color not match")
	assert.Equal(t, color.TextWrap(color.Red, "Hello"), "\x1b[31mHello\x1b[0m", "color not match")
	assert.Equal(t, color.TextWrap(color.LightBlue, "Hello"), "\x1b[94mHello\x1b[0m", "color not match")
}

func TestBackgroundWrap(t *testing.T) {
	assert.Equal(t, color.BackgroundWrap(color.LightBlue, color.White, "Hello"), "\x1b[94;47mHello\x1b[0m")
}

func TestBackgroundFunc(t *testing.T) {
	assert.Equal(t, color.BackgroundFunc(color.LightBlue, color.White)("Hello"), "\x1b[94;47mHello\x1b[0m")
}
