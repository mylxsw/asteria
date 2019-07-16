package asteria_test

import (
	"fmt"
	"testing"

	"github.com/mylxsw/asteria"
)

func TestColorText(t *testing.T) {
	fmt.Println(asteria.ColorTextWrap(asteria.TextLightBlue, "Hello, world"))
	fmt.Println(asteria.ColorBackgroundWrap(asteria.TextLightCyan, asteria.TextLightBlue, "中文"))
}
