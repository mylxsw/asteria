package event

import (
	"encoding/json"
	"time"

	"github.com/mylxsw/asteria/level"
)

type Fields struct {
	CustomFields map[string]interface{}
	GlobalFields map[string]interface{}
}

type Event struct {
	Time     time.Time
	Module   string
	Level    level.Level
	Fields   Fields
	Messages []interface{}
}

func (f Fields) String() string {
	encoded, _ := json.Marshal(f.ToMap())
	return string(encoded)
}

func (f Fields) ToMap() map[string]interface{} {
	cc := f.CustomFields
	if cc == nil {
		cc = make(map[string]interface{}, 0)
	}

	for k, v := range f.GlobalFields {
		cc["#"+k] = v
	}

	return cc
}
