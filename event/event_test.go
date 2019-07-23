package event_test

import (
	"testing"
	"time"

	"github.com/mylxsw/asteria/event"
	"github.com/mylxsw/asteria/level"
	"github.com/stretchr/testify/assert"
)

var ev = event.Event{
	Time:   time.Now(),
	Module: "test",
	Level:  level.Debug,
	Fields: struct {
		CustomFields map[string]interface{}
		GlobalFields map[string]interface{}
	}{
		CustomFields: map[string]interface{}{
			"user_id": 123,
		},
		GlobalFields: map[string]interface{}{
			"ref": "abcdef",
		},
	},
	Messages: []interface{}{
		"Hello, world",
	},
}

func TestFields_ToMap(t *testing.T) {
	res := ev.Fields.ToMap()

	assert.Equal(t, "abcdef", res["#ref"])
	assert.Equal(t, 123, res["user_id"])
}

func TestFields_String(t *testing.T) {
	assert.JSONEq(t, `{"#ref":"abcdef","user_id":123}`, ev.Fields.String())

	ev.Fields.CustomFields = nil
	assert.JSONEq(t, `{"#ref":"abcdef"}`, ev.Fields.String())
}
