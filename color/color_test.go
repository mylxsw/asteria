package color_test

import (
	"fmt"
	"testing"

	"github.com/mylxsw/asteria/color"
)

func TestColorText(t *testing.T) {
	fmt.Println(color.TextWrap(color.TextLightBlue, "Hello, world"))
	fmt.Println(color.BackgroundWrap(color.TextLightCyan, color.TextLightBlue, "中文"))
}
