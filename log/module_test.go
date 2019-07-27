package log_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/mylxsw/asteria/formatter"
	"github.com/mylxsw/asteria/level"
	"github.com/mylxsw/asteria/log"
	"github.com/stretchr/testify/assert"
)

type ErrorWriter struct{}

func (e ErrorWriter) Write(le level.Level, message string) error {
	return fmt.Errorf("has some error")
}

func (e ErrorWriter) ReOpen() error {
	panic("implement me")
}

func (e ErrorWriter) Close() error {
	panic("implement me")
}

func TestWriteFailed(t *testing.T) {
	log.Reset()

	mockWriter := &ErrorWriter{}
	log.DefaultLogWriter(mockWriter)

	panicTriggered := false

	func() {
		defer func() {
			if err := recover(); err != nil {
				panicTriggered = true
			}
		}()

		log.Debug("hello")
	}()

	assert.True(t, panicTriggered)
}

func TestDefaultWithFileLine(t *testing.T) {
	log.Reset()

	mockWriter := &MockWriter{}
	log.DefaultLogWriter(mockWriter)
	log.DefaultLogFormatter(formatter.NewDefaultFormatter(false))

	log.DefaultWithFileLine(true)
	log.Debug("hello")

	assert.Regexp(t, regexp.MustCompile(`"#file"`), mockWriter.LastMessage)
	assert.Regexp(t, regexp.MustCompile(`"#line"`), mockWriter.LastMessage)
	assert.Regexp(t, regexp.MustCompile(`"#package"`), mockWriter.LastMessage)

	log.DefaultWithFileLine(false)

	log.Debug("hello")
	assert.NotRegexp(t, regexp.MustCompile(`"#file"`), mockWriter.LastMessage)
	assert.NotRegexp(t, regexp.MustCompile(`"#line"`), mockWriter.LastMessage)
	assert.NotRegexp(t, regexp.MustCompile(`"#package"`), mockWriter.LastMessage)
}

func TestDefaultLocation(t *testing.T) {
	log.Reset()

	mockWriter := &MockWriter{}
	log.DefaultLogWriter(mockWriter)
	log.DefaultLogFormatter(formatter.NewDefaultFormatter(false))

	// testing UTC
	loc, _ := time.LoadLocation("UTC")
	log.DefaultLocation(loc)

	log.Warning("hello")

	logTimeStr := strings.Trim(strings.Split(mockWriter.LastMessage, " ")[0], "[]")
	realTimeStr := time.Now().In(loc).Format(time.RFC3339)

	assert.Equal(t, logTimeStr[:13], realTimeStr[:13])

	// testing Local
	loc, _ = time.LoadLocation("Local")
	log.DefaultLocation(loc)

	log.Warning("hello")

	logTimeStr = strings.Trim(strings.Split(mockWriter.LastMessage, " ")[0], "[]")
	realTimeStr = time.Now().In(loc).Format(time.RFC3339)

	assert.Equal(t, logTimeStr[:13], realTimeStr[:13])
}

func TestDefaultLogLevel(t *testing.T) {
	log.Reset()

	mockWriter := &MockWriter{}
	log.DefaultLogWriter(mockWriter)
	log.DefaultLogFormatter(formatter.NewDefaultFormatter(false))

	log.DefaultLogLevel(level.Warning)

	log.Debug("hello")
	assert.Equal(t, level.Level(0), mockWriter.LastLevel)

	log.Info("hello")
	assert.Equal(t, level.Level(0), mockWriter.LastLevel)

	log.Warning("hello")
	assert.Equal(t, level.Warning, mockWriter.LastLevel)

	log.Error("hello")
	assert.Equal(t, level.Error, mockWriter.LastLevel)

}

func TestLogger_Location(t *testing.T) {
	log.Reset()

	mockWriter := &MockWriter{}
	logger := log.Module("test").
		Writer(mockWriter).
		Formatter(formatter.NewDefaultFormatter(false))

	// testing UTC
	loc, _ := time.LoadLocation("UTC")
	logger.Location(loc)

	logger.Warning("hello")

	logTimeStr := strings.Trim(strings.Split(mockWriter.LastMessage, " ")[0], "[]")
	realTimeStr := time.Now().In(loc).Format(time.RFC3339)

	assert.Equal(t, logTimeStr[:13], realTimeStr[:13])

	// testing Local
	loc, _ = time.LoadLocation("Local")
	logger.Location(loc)

	logger.Warning("hello")

	logTimeStr = strings.Trim(strings.Split(mockWriter.LastMessage, " ")[0], "[]")
	realTimeStr = time.Now().In(loc).Format(time.RFC3339)

	assert.Equal(t, logTimeStr[:13], realTimeStr[:13])
}

func TestLogger_WithFileLine(t *testing.T) {
	log.Reset()

	mockWriter := &MockWriter{}
	logger := log.Module("test").
		Writer(mockWriter).
		Formatter(formatter.NewDefaultFormatter(false))

	logger.WithFileLine(true)
	logger.Debug("hello")

	assert.Regexp(t, regexp.MustCompile(`"#file"`), mockWriter.LastMessage)
	assert.Regexp(t, regexp.MustCompile(`"#line"`), mockWriter.LastMessage)
	assert.Regexp(t, regexp.MustCompile(`"#package"`), mockWriter.LastMessage)

	logger.WithFileLine(false)

	logger.Debug("hello")
	assert.NotRegexp(t, regexp.MustCompile(`"#file"`), mockWriter.LastMessage)
	assert.NotRegexp(t, regexp.MustCompile(`"#line"`), mockWriter.LastMessage)
	assert.NotRegexp(t, regexp.MustCompile(`"#package"`), mockWriter.LastMessage)
}

func TestModuleLogger(t *testing.T) {
	log.Reset()

	mockWriter := &MockWriter{}
	log.DefaultLogWriter(mockWriter)
	log.DefaultLogFormatter(formatter.NewDefaultFormatter(false))

	var logger = log.Module("test")

	logger.Emergency("hello")
	assert.Equal(t, level.Emergency, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello {}", level.Emergency.GetLevelName())), mockWriter.LastMessage)

	logger.Alert("hello")
	assert.Equal(t, level.Alert, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello {}", level.Alert.GetLevelName())), mockWriter.LastMessage)

	logger.Critical("hello")
	assert.Equal(t, level.Critical, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello {}", level.Critical.GetLevelName())), mockWriter.LastMessage)

	logger.Error("hello")
	assert.Equal(t, level.Error, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello {}", level.Error.GetLevelName())), mockWriter.LastMessage)

	logger.Warning("hello")
	assert.Equal(t, level.Warning, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello {}", level.Warning.GetLevelName())), mockWriter.LastMessage)

	logger.Notice("hello")
	assert.Equal(t, level.Notice, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello {}", level.Notice.GetLevelName())), mockWriter.LastMessage)

	logger.Info("hello")
	assert.Equal(t, level.Info, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello {}", level.Info.GetLevelName())), mockWriter.LastMessage)

	logger.Debug("hello")
	assert.Equal(t, level.Debug, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello {}", level.Debug.GetLevelName())), mockWriter.LastMessage)

	// logf

	logger.Emergencyf("hello %s", "world")
	assert.Equal(t, level.Emergency, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello world {}", level.Emergency.GetLevelName())), mockWriter.LastMessage)

	logger.Alertf("hello %s", "world")
	assert.Equal(t, level.Alert, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello world {}", level.Alert.GetLevelName())), mockWriter.LastMessage)

	logger.Criticalf("hello %s", "world")
	assert.Equal(t, level.Critical, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello world {}", level.Critical.GetLevelName())), mockWriter.LastMessage)

	logger.Errorf("hello %s", "world")
	assert.Equal(t, level.Error, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello world {}", level.Error.GetLevelName())), mockWriter.LastMessage)

	logger.Warningf("hello %s", "world")
	assert.Equal(t, level.Warning, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello world {}", level.Warning.GetLevelName())), mockWriter.LastMessage)

	logger.Noticef("hello %s", "world")
	assert.Equal(t, level.Notice, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello world {}", level.Notice.GetLevelName())), mockWriter.LastMessage)

	logger.Infof("hello %s", "world")
	assert.Equal(t, level.Info, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello world {}", level.Info.GetLevelName())), mockWriter.LastMessage)

	logger.Debugf("hello %s", "world")
	assert.Equal(t, level.Debug, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello world {}", level.Debug.GetLevelName())), mockWriter.LastMessage)

	logger.Print("hello")
	assert.Equal(t, level.Debug, mockWriter.LastLevel)
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^\\[.*?\\] %s .*? hello {}", level.Debug.GetLevelName())), mockWriter.LastMessage)

}
