package formatter_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/mylxsw/asteria/formatter"
	"github.com/mylxsw/asteria/level"
	"github.com/stretchr/testify/assert"
)

func TestDefaultFormatter_Format(t *testing.T) {

	now := time.Now()

	var testcases = map[string]formatter.Format{
		fmt.Sprintf(`[%s] ALERT test Hello, world {"#abc":"def","uid":134}`, now.Format(time.RFC3339)): {
			Colorful: false,
			Time:     now,
			Module:   "test",
			Level:    level.Alert,
			Context: formatter.LogContext{
				SysContext:  map[string]interface{}{"abc": "def",},
				UserContext: map[string]interface{}{"uid": 134,},
			},
			Messages: []interface{}{"Hello, world"},
		},
		fmt.Sprintf(`[%s] DEBUG test Hello, world {"#abc":"def"}`, now.Format(time.RFC3339)): {
			Colorful: false,
			Time:     now,
			Module:   "test",
			Level:    level.Debug,
			Context: formatter.LogContext{
				SysContext:  map[string]interface{}{"abc": "def",},
			},
			Messages: []interface{}{"Hello, world"},
		},
		fmt.Sprintf(`[%s] ` + "\x1b[97;44m[DEBG]\x1b[0m test2                Hello, world \x1b[90m{\"#abc\":\"def\",\"uid\":134}\x1b[0m", now.Format(time.RFC3339)): {
			Colorful: true,
			Time:     now,
			Module:   "test2",
			Level:    level.Debug,
			Context: formatter.LogContext{
				SysContext:  map[string]interface{}{"abc": "def",},
				UserContext: map[string]interface{}{"uid": 134,},
			},
			Messages: []interface{}{"Hello, world"},
		},
	}

	f := formatter.NewDefaultFormatter()
	for expected, ts := range testcases {
		assert.Equal(t, f.Format(ts), expected)
	}
}
