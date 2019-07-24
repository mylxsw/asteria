package writer_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/mylxsw/asteria/level"
	"github.com/mylxsw/asteria/writer"
	"github.com/stretchr/testify/assert"
)

func TestFileWriter_Write(t *testing.T) {
	fw := writer.NewDefaultFileWriter("test/test.log")
	err := fw.Write(level.Debug, "Hello, world")

	assert.Error(t, err)

	fw = writer.NewDefaultFileWriter("test.log")
	defer func() {
		_ = os.Remove("test.log")
	}()

	assert.NoError(t, fw.Write(level.Debug, "Hello, world"))
	assert.NoError(t, fw.Write(level.Debug, "Hello, world"))
	assert.NoError(t, fw.Write(level.Debug, "Hello, world"))
	assert.NoError(t, fw.Write(level.Debug, "Hello, world"))

	rs, err := ioutil.ReadFile("test.log")
	assert.NoError(t, err)
	assert.Equal(t, len(strings.Split(string(rs), "\n")), 4+1)
}

func TestRotatingFileWriter_Write(t *testing.T) {
	fileNo1 := fmt.Sprintf("test-%s.log", time.Now().Format("20060102150405"))
	fileNo2 := fmt.Sprintf("test-%s.log", time.Now().Format("20060102150405")+"-2")

	fileNo := fileNo1

	fw := writer.NewDefaultRotatingFileWriter(func(le level.Level) string {
		return fileNo
	})

	defer func() {
		_ = os.Remove(fileNo1)
		_ = os.Remove(fileNo2)
	}()

	_ = fw.Write(level.Debug, "Hello, world")
	_ = fw.Write(level.Debug, "Hello, world")
	_ = fw.Write(level.Debug, "Hello, world")
	_ = fw.Write(level.Error, "Hello, world")
	_ = fw.Write(level.Error, "Hello, world")
	_ = fw.Write(level.Error, "Hello, world")

	fileNo = fileNo2

	_ = fw.Write(level.Debug, "Hello, world")
	_ = fw.Write(level.Error, "Hello, world")
	_ = fw.Write(level.Error, "Hello, world")
	_ = fw.Write(level.Error, "Hello, world")

	_ = fw.Close()

	rs1, err := ioutil.ReadFile(fileNo1)
	assert.NoError(t, err)
	assert.Equal(t, len(strings.Split(string(rs1), "\n")), 6+1)

	rs2, err := ioutil.ReadFile(fileNo2)
	assert.NoError(t, err)
	assert.Equal(t, len(strings.Split(string(rs2), "\n")), 4+1)
}

func TestRotatingFileWriter_ReOpen(t *testing.T) {
	filename := "test_rotating_file.log"
	fw := writer.NewDefaultRotatingFileWriter(func(le level.Level) string {
		return filename
	})

	defer func() {
		_ = os.Remove(filename)
	}()

	assert.NoError(t, fw.Write(level.Debug, "Hello, world"))
	assert.NoError(t, fw.Write(level.Error, "Hello, world"))

	assert.NoError(t, fw.ReOpen())
	assert.NoError(t, fw.Close())

	assert.NoError(t, fw.Write(level.Error, "Hello, world"))
}
