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

func TestJSONFormatter_Format(t *testing.T) {
	now := time.Now()
	f := formatter.NewJSONFormatter()

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
	assert.Regexp(t, regexp.MustCompile("^{.*}$"), res)
}

type User struct {
	ID   int
	Name string
	Role []Role
}

type Role struct {
	ID         int
	Name       string
	Privileges []Privilege
}

type Privilege struct {
	ID   int
	Name string
}

func TestJSONFormatter_Complex(t *testing.T) {
	var user = User{
		ID:   123,
		Name: "Lixiaoyao",
		Role: []Role{
			{
				ID:   444,
				Name: "Admin",
				Privileges: []Privilege{
					{
						ID:   555,
						Name: "user_create: ABC\\DLA",
					},
				},
			},
		},
	}

	now := time.Now()
	f := formatter.NewJSONFormatter()

	fm := event.Event{
		Time:   now,
		Module: "test",
		Level:  level.Alert,
		Fields: event.Fields{
			GlobalFields: map[string]interface{}{"abc": "def",},
			CustomFields: map[string]interface{}{"user": user,},
		},
		Messages: []interface{}{"Hello, world"},
	}
	res := f.Format(fm)

	assert.NotEmpty(t, res)
}
