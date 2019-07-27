package misc_test

import (
	"testing"

	"github.com/mylxsw/asteria/misc"
	"github.com/stretchr/testify/assert"
)

func TestModuleNameAbbr(t *testing.T) {
	var testCases = map[string]string{
		"abc.def.ghi": "a.d.ghi",
		"abcd":        "abcd",
		"a.bc":        "a.bc",
		"ab.bc":       "a.bc",
		"中文.布尔":       "中.布尔",
		"xxx..abc":    "x.abc",
	}
	for tc, expected := range testCases {
		assert.Equal(t, expected, misc.ModuleNameAbbr(tc))
	}
}
