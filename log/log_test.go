package log_test

import (
	"fmt"
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
	ReOpenCount int
	CloseCount  int
}

func (w *MockWriter) Write(le level.Level, message string) error {
	w.LastLevel = le
	w.LastMessage = message

	return nil
}

func (w *MockWriter) ReOpen() error {
	w.ReOpenCount++
	return nil
}

func (w *MockWriter) Close() error {
	w.CloseCount++
	return nil
}

func TestGlobalFilters(t *testing.T) {
	log.Reset()

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

func TestFilters(t *testing.T) {
	log.Reset()

	mockWriter := &MockWriter{}
	log.DefaultLogWriter(mockWriter)
	log.DefaultLogFormatter(formatter.NewDefaultFormatter(false))

	log.Module("test").AddFilter(func(filter log.Filter) log.Filter {
		return func(f event.Event) {
			f.Fields.CustomFields["user_id"] = 123
			filter(f)
		}
	})

	log.Module("test").Debug("hello")
	assert.Regexp(t, regexp.MustCompile("\"user_id\":123"), mockWriter.LastMessage)

	log.Module("test2").Debug("hello")
	assert.NotRegexp(t, regexp.MustCompile("\"user_id\":123"), mockWriter.LastMessage)
}

func TestGlobalFields(t *testing.T) {
	log.Reset()

	mockWriter := &MockWriter{}
	log.DefaultLogFormatter(formatter.NewDefaultFormatter(false))
	log.DefaultLogWriter(mockWriter)

	log.GlobalFields(func(c event.Fields) {
		c.GlobalFields["ref"] = "abcd"
		c.CustomFields["user_id"] = 123
	})

	log.Module("test").Debug("hello")
	assert.Regexp(t, regexp.MustCompile("#ref"), mockWriter.LastMessage)
	assert.Regexp(t, regexp.MustCompile("\"user_id\""), mockWriter.LastMessage)

	log.Module("test").GlobalFields(func(c event.Fields) {
		c.CustomFields["enterprise_id"] = 123
	})

	log.Module("test").Debug("hello")
	assert.NotRegexp(t, regexp.MustCompile("#ref"), mockWriter.LastMessage)
	assert.NotRegexp(t, regexp.MustCompile("\"user_id\""), mockWriter.LastMessage)
	assert.Regexp(t, regexp.MustCompile("\"enterprise_id\""), mockWriter.LastMessage)

	log.Module("test").GlobalFields(func(c event.Fields) {
		c.CustomFields["enterprise_id"] = 123
		log.GetDefaultConfig().GlobalFields(c)
	})

	log.Module("test").Debug("hello")
	assert.Regexp(t, regexp.MustCompile("#ref"), mockWriter.LastMessage)
	assert.Regexp(t, regexp.MustCompile("\"user_id\""), mockWriter.LastMessage)
	assert.Regexp(t, regexp.MustCompile("\"enterprise_id\""), mockWriter.LastMessage)
}

func TestBasicLog(t *testing.T) {
	log.Reset()

	mockWriter := &MockWriter{}
	log.DefaultLogFormatter(formatter.NewDefaultFormatter(false))
	log.DefaultLogWriter(mockWriter)

	log.Emergency("hello")
	assert.Equal(t, level.Emergency, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello {}", level.Emergency.GetLevelName())), mockWriter.LastMessage)

	log.Alert("hello")
	assert.Equal(t, level.Alert, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello {}", level.Alert.GetLevelName())), mockWriter.LastMessage)

	log.Critical("hello")
	assert.Equal(t, level.Critical, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello {}", level.Critical.GetLevelName())), mockWriter.LastMessage)

	log.Error("hello")
	assert.Equal(t, level.Error, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello {}", level.Error.GetLevelName())), mockWriter.LastMessage)

	log.Warning("hello")
	assert.Equal(t, level.Warning, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello {}", level.Warning.GetLevelName())), mockWriter.LastMessage)

	log.Notice("hello")
	assert.Equal(t, level.Notice, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello {}", level.Notice.GetLevelName())), mockWriter.LastMessage)

	log.Info("hello")
	assert.Equal(t, level.Info, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello {}", level.Info.GetLevelName())), mockWriter.LastMessage)

	log.Debug("hello")
	assert.Equal(t, level.Debug, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello {}", level.Debug.GetLevelName())), mockWriter.LastMessage)
}

func TestBasicLogf(t *testing.T) {
	log.Reset()

	mockWriter := &MockWriter{}
	log.DefaultLogFormatter(formatter.NewDefaultFormatter(false))
	log.DefaultLogWriter(mockWriter)

	log.Emergencyf("hello %s", "world")
	assert.Equal(t, level.Emergency, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello world {}", level.Emergency.GetLevelName())), mockWriter.LastMessage)

	log.Alertf("hello %s", "world")
	assert.Equal(t, level.Alert, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello world {}", level.Alert.GetLevelName())), mockWriter.LastMessage)

	log.Criticalf("hello %s", "world")
	assert.Equal(t, level.Critical, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello world {}", level.Critical.GetLevelName())), mockWriter.LastMessage)

	log.Errorf("hello %s", "world")
	assert.Equal(t, level.Error, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello world {}", level.Error.GetLevelName())), mockWriter.LastMessage)

	log.Warningf("hello %s", "world")
	assert.Equal(t, level.Warning, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello world {}", level.Warning.GetLevelName())), mockWriter.LastMessage)

	log.Noticef("hello %s", "world")
	assert.Equal(t, level.Notice, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello world {}", level.Notice.GetLevelName())), mockWriter.LastMessage)

	log.Infof("hello %s", "world")
	assert.Equal(t, level.Info, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello world {}", level.Info.GetLevelName())), mockWriter.LastMessage)

	log.Debugf("hello %s", "world")
	assert.Equal(t, level.Debug, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello world {}", level.Debug.GetLevelName())), mockWriter.LastMessage)
}

func TestLogger_ReOpenClose(t *testing.T) {
	log.Reset()

	mockWriter := &MockWriter{}
	log.SetWriter(mockWriter)

	assert.NoError(t, log.Close())
	assert.Equal(t, 1, mockWriter.CloseCount)

	assert.NoError(t, log.ReOpen())
	assert.Equal(t, 1, mockWriter.ReOpenCount)
	assert.Equal(t, 1, mockWriter.CloseCount)

	for _, err := range log.CloseAll() {
		assert.NoError(t, err)
	}
	assert.Equal(t, 2, mockWriter.CloseCount)
	assert.Equal(t, 1, mockWriter.ReOpenCount)

	for _, err := range log.ReOpenAll() {
		assert.NoError(t, err)
	}
	assert.Equal(t, 2, mockWriter.CloseCount)
	assert.Equal(t, 2, mockWriter.ReOpenCount)

}
