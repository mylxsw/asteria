package log_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/mylxsw/asteria/formatter"
	"github.com/mylxsw/asteria/level"
	"github.com/mylxsw/asteria/log"
	"github.com/stretchr/testify/assert"
)

func TestContextLogger(t *testing.T) {
	log.Reset()

	mockWriter := &MockWriter{}
	log.DefaultLogFormatter(formatter.NewDefaultFormatter(false))
	log.DefaultLogWriter(mockWriter)

	log.WithFields(log.Fields{"user_id": 123}).Emergency("hello")
	assert.Equal(t, level.Emergency, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello {\"user_id\":123}", level.Emergency.GetLevelName())), mockWriter.LastMessage)

	log.WithFields(log.Fields{"user_id": 123}).Alert("hello")
	assert.Equal(t, level.Alert, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello {\"user_id\":123}", level.Alert.GetLevelName())), mockWriter.LastMessage)

	log.WithFields(log.Fields{"user_id": 123}).Critical("hello")
	assert.Equal(t, level.Critical, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello {\"user_id\":123}", level.Critical.GetLevelName())), mockWriter.LastMessage)

	log.WithFields(log.Fields{"user_id": 123}).Error("hello")
	assert.Equal(t, level.Error, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello {\"user_id\":123}", level.Error.GetLevelName())), mockWriter.LastMessage)

	log.WithFields(log.Fields{"user_id": 123}).Warning("hello")
	assert.Equal(t, level.Warning, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello {\"user_id\":123}", level.Warning.GetLevelName())), mockWriter.LastMessage)

	log.WithFields(log.Fields{"user_id": 123}).Notice("hello")
	assert.Equal(t, level.Notice, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello {\"user_id\":123}", level.Notice.GetLevelName())), mockWriter.LastMessage)

	log.WithFields(log.Fields{"user_id": 123}).Info("hello")
	assert.Equal(t, level.Info, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello {\"user_id\":123}", level.Info.GetLevelName())), mockWriter.LastMessage)

	log.WithFields(log.Fields{"user_id": 123}).Debug("hello")
	assert.Equal(t, level.Debug, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello {\"user_id\":123}", level.Debug.GetLevelName())), mockWriter.LastMessage)

	log.WithFields(log.Fields{"user_id": 123}).Emergencyf("hello %s", "world")
	assert.Equal(t, level.Emergency, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello world {\"user_id\":123}", level.Emergency.GetLevelName())), mockWriter.LastMessage)

	log.WithFields(log.Fields{"user_id": 123}).Alertf("hello %s", "world")
	assert.Equal(t, level.Alert, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello world {\"user_id\":123}", level.Alert.GetLevelName())), mockWriter.LastMessage)

	log.WithFields(log.Fields{"user_id": 123}).Criticalf("hello %s", "world")
	assert.Equal(t, level.Critical, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello world {\"user_id\":123}", level.Critical.GetLevelName())), mockWriter.LastMessage)

	log.WithFields(log.Fields{"user_id": 123}).Errorf("hello %s", "world")
	assert.Equal(t, level.Error, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello world {\"user_id\":123}", level.Error.GetLevelName())), mockWriter.LastMessage)

	log.WithFields(log.Fields{"user_id": 123}).Warningf("hello %s", "world")
	assert.Equal(t, level.Warning, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello world {\"user_id\":123}", level.Warning.GetLevelName())), mockWriter.LastMessage)

	log.WithFields(log.Fields{"user_id": 123}).Noticef("hello %s", "world")
	assert.Equal(t, level.Notice, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello world {\"user_id\":123}", level.Notice.GetLevelName())), mockWriter.LastMessage)

	log.WithFields(log.Fields{"user_id": 123}).Infof("hello %s", "world")
	assert.Equal(t, level.Info, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello world {\"user_id\":123}", level.Info.GetLevelName())), mockWriter.LastMessage)

	log.WithFields(log.Fields{"user_id": 123}).Debugf("hello %s", "world")
	assert.Equal(t, level.Debug, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello world {\"user_id\":123}", level.Debug.GetLevelName())), mockWriter.LastMessage)

}
