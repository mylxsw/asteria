package event

import (
	"time"

	"github.com/json-iterator/go"
	"github.com/mylxsw/asteria/level"
)

var json = jsoniter.ConfigFastest

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

func (f Fields) String(excludes ...string) string {
	encoded, _ := json.Marshal(f.ToMap(excludes...))
	return string(encoded)
}

func (f Fields) ToMap(excludes ...string) map[string]interface{} {
	cc := f.CustomFields
	if cc == nil {
		cc = make(map[string]interface{}, 0)
	}

	for k, v := range f.GlobalFields {
		if strIn(k, excludes) {
			continue
		}

		cc["#"+k] = v
	}

	return cc
}

func strIn(str string, strArray []string) bool {
	for _, s := range strArray {
		if s == str {
			return true
		}
	}

	return false
}
