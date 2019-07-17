package color_test

import (
	"fmt"
	"testing"

	"github.com/mylxsw/asteria/color"
)

func TestColorText(t *testing.T) {
	fmt.Println(color.ColorTextWrap(color.TextLightBlue, "Hello, world"))
	fmt.Println(color.ColorBackgroundWrap(color.TextLightCyan, color.TextLightBlue, "中文"))
}
