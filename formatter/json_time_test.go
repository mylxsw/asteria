package formatter_test

import (
	"regexp"
	"testing"
	"time"

	"github.com/mylxsw/asteria/formatter"
	"github.com/mylxsw/asteria/level"
	"github.com/stretchr/testify/assert"
)

func TestJSONWithTimeFormatter_Format(t *testing.T) {
	now := time.Now()
	f := formatter.NewJSONWithTimeFormatter()

	fm := formatter.Format{
		Colorful: false,
		Time:     now,
		Module:   "test",
		Level:    level.Alert,
		Context: formatter.LogContext{
			SysContext:  map[string]interface{}{"abc": "def",},
			UserContext: map[string]interface{}{"uid": 134,},
		},
		Messages: []interface{}{"Hello, world"},
	}
	res := f.Format(fm)

	assert.NotEmpty(t, res)
	assert.Regexp(t, regexp.MustCompile("^\\[.*?\\] {.*}$"), res)
}
