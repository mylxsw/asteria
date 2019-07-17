package color

import (
	"fmt"
)

// Color is a console color
type Color int

const (
	// TextBlack 黑色
	TextBlack Color = iota + 30
	// TextRed 红色
	TextRed
	// TextGreen 绿色
	TextGreen
	// TextYellow 黄色
	TextYellow
	// TextBlue 蓝色
	TextBlue
	// TextMagenta 洋红
	TextMagenta
	// TextCyan 青色
	TextCyan
	// TextWhite 白色
	TextWhite
)

const (
	// TextLightGrey 亮灰色
	TextLightGrey Color = iota + 90
	// TextLightRed 亮红色
	TextLightRed
	// TextLightGreen 亮绿色
	TextLightGreen
	// TextLightYellow 亮黄色
	TextLightYellow
	// TextLightBlue 亮蓝色
	TextLightBlue
	// TextLightMagenta 亮洋红
	TextLightMagenta
	// TextLightCyan 亮青色
	TextLightCyan
	// TextLightWhite 亮白色
	TextLightWhite
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