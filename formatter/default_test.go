package formatter_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/mylxsw/asteria/event"
	"github.com/mylxsw/asteria/formatter"
	"github.com/mylxsw/asteria/level"
	"github.com/stretchr/testify/assert"
)

func TestDefaultFormatter_Format(t *testing.T) {

	now := time.Now()

	var testcases = map[string]event.Event{
		fmt.Sprintf(`[%s] ALERT test Hello, world {"uid":134,"#abc":"def"}`, now.Format(time.RFC3339)): {
			Time:   now,
			Module: "test",
			Level:  level.Alert,
			Fields: event.Fields{
				GlobalFields: map[string]interface{}{"abc": "def",},
				CustomFields: map[string]interface{}{"uid": 134,},
			},
			Messages: []interface{}{"Hello, world"},
		},
		fmt.Sprintf(`[%s] DEBUG test Hello, world {"#abc":"def"}`, now.Format(time.RFC3339)): {
			Time:   now,
			Module: "test",
			Level:  level.Debug,
			Fields: event.Fields{
				GlobalFields: map[string]interface{}{"abc": "def",},
			},
			Messages: []interface{}{"Hello, world"},
		},
		fmt.Sprintf(`[%s] `+"DEBUG test2 Hello, world {\"uid\":134,\"#abc\":\"def\"}", now.Format(time.RFC3339)): {
			Time:   now,
			Module: "test2",
			Level:  level.Debug,
			Fields: event.Fields{
				GlobalFields: map[string]interface{}{"abc": "def",},
				CustomFields: map[string]interface{}{"uid": 134,},
			},
			Messages: []interface{}{"Hello, world"},
		},
	}

	f := formatter.NewDefaultFormatter(false)
	for expected, ts := range testcases {
		assert.Equal(t, f.Format(ts), expected)
	}
}

func TestDefaultFormatter_ColorfulFormat(t *testing.T) {
	now := time.Now()

	var testcases = map[string]event.Event{
		fmt.Sprintf(`[%s] ` + "\x1b[97;101m[ALER]\x1b[0m test                 Hello, world \x1b[90m{\"#abc\":\"def\"}\x1b[0m", now.Format(time.RFC3339)): {
			Time:   now,
			Module: "test",
			Level:  level.Alert,
			Fields: event.Fields{
				GlobalFields: map[string]interface{}{"abc": "def",},
			},
			Messages: []interface{}{"Hello, world"},
		},
	}

	f := formatter.NewDefaultFormatter(true)
	for expected, ts := range testcases {
		assert.Equal(t, expected, f.Format(ts))
	}
}
