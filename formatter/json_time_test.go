package formatter_test

import (
	"regexp"
	"testing"
	"time"

	"github.com/mylxsw/asteria/event"
	"github.com/mylxsw/asteria/formatter"
	"github.com/mylxsw/asteria/level"
	"github.com/stretchr/testify/assert"
)

func TestJSONWithTimeFormatter_Format(t *testing.T) {
	now := time.Now()
	f := formatter.NewJSONWithTimeFormatter()

	fm := event.Event{
		Time:   now,
		Module: "test",
		Level:  level.Alert,
		Fields: event.Fields{
			GlobalFields: map[string]interface{}{"abc": "def",},
			CustomFields: map[string]interface{}{"uid": 134,},
		},
		Messages: []interface{}{"Hello, world"},
	}
	res := f.Format(fm)

	assert.NotEmpty(t, res)
	assert.Regexp(t, regexp.MustCompile("^\\[.*?\\] {.*}$"), res)
}
