package color

import (
	"fmt"
)

// Color is a console color
type Color int

const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

const (
	LightGrey Color = iota + 90
	LightRed
	LightGreen
	LightYellow
	LightBlue
	LightMagenta
	LightCyan
	LightWhite
)

// TextWrap 文字颜色
func TextWrap(color Color, text string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, text)
}

// BackgroundWrap 背景颜色
func BackgroundWrap(color Color, backgroundColor Color, text string) string {
	return fmt.Sprintf("\x1b[%d;%dm%s\x1b[0m", color, backgroundColor+10, text)
}

func BackgroundFunc(color Color, backgroundColor Color) func(text string) string {
	return func(text string) string {
		return BackgroundWrap(color, backgroundColor, text)
	}
}
