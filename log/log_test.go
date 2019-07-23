package log_test

import (
	"regexp"
	"testing"

	"github.com/mylxsw/asteria/event"
	"github.com/mylxsw/asteria/formatter"
	"github.com/mylxsw/asteria/level"
	"github.com/mylxsw/asteria/log"
	"github.com/stretchr/testify/assert"
)

type MockWriter struct {
	LastLevel   level.Level
	LastMessage string
}

func (w *MockWriter) Write(le level.Level, message string) error {
	w.LastLevel = le
	w.LastMessage = message

	return nil
}

func (w *MockWriter) ReOpen() error {
	panic("implement me")
}

func (w *MockWriter) Close() error {
	panic("implement me")
}

func TestGlobalFilters(t *testing.T) {
	mockWriter := &MockWriter{}

	log.DefaultLogWriter(mockWriter)
	log.AddGlobalFilter(func(filter log.Filter) log.Filter {
		return func(f event.Event) {
			f.Fields.CustomFields["user_id"] = 123
			filter(f)
		}
	})

	log.Debug("Hello")
	assert.Equal(t, level.Debug, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile("^\\[.*?\\] .*?[DEBUG].*?Hello.*?\"user_id\":123"), mockWriter.LastMessage)
}

func TestModule(t *testing.T) {
	log.GlobalFields(func(c event.Fields) {
		c.GlobalFields["ref"] = "190101931"
	})
	log.Default().AddFilter(func(filter log.Filter) log.Filter {
		return func(f event.Event) {
			f.Fields.GlobalFields["ref"] = "6f96cfdfe"
			filter(f)
		}
	})

	log.Default().LogLevel(level.Debug)
	log.Debug("细雨微风岸，危樯独夜舟")
	log.Error("月上柳梢头，人约黄昏后")
	log.WithFields(log.Fields{
		"user_id":  123,
		"username": "Tom",
	}).Warningf("君子坦荡荡，小人常戚戚")

	log.Module("order.test.scheduler").Noticef("order %s created", "1234592")
	log.Module("order.scheduler.module1.module2").Infof("order %s created", "1234592")
	log.Module("order.apple").Debugf("order %s created", "1234592")
	log.Module("order").Errorf("order %s created", "1234592")
	log.Module("order").Emergencyf("order %s created", "1234592")
	log.Module("order").Warningf("order %s created", "1234592")
	log.Module("order").Alertf("order %s created", "1234592")
	log.Module("order").Criticalf("order %s created\n", "1234592")

	log.Module("user").Formatter(formatter.NewJSONWithTimeFormatter()).Error("user create failed")

	log.WithFields(nil).Debug("error occur")
	log.Module("purchase").Formatter(formatter.NewJSONWithTimeFormatter()).WithFields(nil).Infof("用户 %s 已创建", "mylxsw")

	userLog := log.Module("user")
	userLog.WithFields(log.Fields{
		"id":   123,
		"name": "lixiaoyao",
	}).Debugf("Hello, %s", "world")

	taskLogger := log.Module("log.user.tasks").WithFileLine(true).GlobalFields(func(c event.Fields) {
		log.GetDefaultConfig().GlobalFields(c)
		c.GlobalFields["enterprise_id"] = 15
	})
	taskLogger.Debug("Hello, world")

	enterpriseJobLogger := log.Module("log.user.enterprise.jobs").WithFileLine(true).Formatter(formatter.NewJSONFormatter())

	enterpriseJobLogger.AddFilter(func(filter log.Filter) log.Filter {
		return func(f event.Event) {
			// filter(f)
			f.Level = level.Emergency
			filter(f)
		}
	})

	enterpriseJobLogger.AddFilter(func(filter log.Filter) log.Filter {
		return func(f event.Event) {
			filter(f)
		}
	})

	enterpriseJobLogger.Debug("He remembered the count of Monte cristo")
	enterpriseJobLogger.Info("You are mistaken💯 I am not the Count of Monte Cristo")
	enterpriseJobLogger.Error("The noiseless door again turned on its hinges, and the Count of Monte Cristo reappeared")
	enterpriseJobLogger.WithFields(log.Fields{
		"user_id":  123,
		"username": "Tom",
	}).Warningf("君子坦荡荡，小人常戚戚")
}
