package writer_test

import (
	"log/syslog"
	"testing"

	"github.com/mylxsw/asteria/level"
	"github.com/mylxsw/asteria/writer"
	"github.com/stretchr/testify/assert"
)

func TestSyslogWriter_Write(t *testing.T) {
	fw := writer.NewSyslogWriter("", "", syslog.LOG_ERR, "asteria")

	assert.NoError(t, fw.Write(level.Error, "", "hello, world"))
	assert.NoError(t, fw.Write(level.Debug, "", "hello, world"))
	assert.NoError(t, fw.Write(level.Warning, "", "hello, world"))
	assert.NoError(t, fw.Write(level.Info, "", "hello, world"))
	assert.NoError(t, fw.Write(level.Alert, "", "hello, world"))
	assert.NoError(t, fw.Write(level.Emergency, "", "hello, world"))
	assert.NoError(t, fw.Write(level.Notice, "", "hello, world"))
	assert.NoError(t, fw.Write(level.Critical, "", "hello, world"))

	assert.NoError(t, fw.Close())
	assert.NoError(t, fw.Write(level.Emergency, "", "hello, world"))

	assert.NoError(t, fw.ReOpen())
	assert.NoError(t, fw.Write(level.Error, "", "hello, world"))
}
