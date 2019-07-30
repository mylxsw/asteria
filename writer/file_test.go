package writer_test

import (
	"context"
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
	err := fw.Write(level.Debug, "", "Hello, world")

	assert.Error(t, err)

	fw = writer.NewDefaultFileWriter("test.log")
	defer func() {
		_ = os.Remove("test.log")
	}()

	assert.NoError(t, fw.Write(level.Debug, "", "Hello, world"))
	assert.NoError(t, fw.Write(level.Debug, "", "Hello, world"))
	assert.NoError(t, fw.Write(level.Debug, "", "Hello, world"))
	assert.NoError(t, fw.Write(level.Debug, "", "Hello, world"))

	assert.Equal(t, "test.log", fw.GetFilename())

	rs, err := ioutil.ReadFile("test.log")
	assert.NoError(t, err)
	assert.Equal(t, len(strings.Split(string(rs), "\n")), 4+1)
}

func TestRotatingFileWriter_Write(t *testing.T) {
	fileNo1 := fmt.Sprintf("test-%s.log", time.Now().Format("20060102150405"))
	fileNo2 := fmt.Sprintf("test-%s.log", time.Now().Format("20060102150405")+"-2")

	fileNo := fileNo1

	fw := writer.NewDefaultRotatingFileWriter(func(le level.Level, module string) string {
		return fileNo
	})

	defer func() {
		_ = os.Remove(fileNo1)
		_ = os.Remove(fileNo2)
	}()

	_ = fw.Write(level.Debug, "", "Hello, world")
	_ = fw.Write(level.Debug, "", "Hello, world")
	_ = fw.Write(level.Debug, "", "Hello, world")
	_ = fw.Write(level.Error, "", "Hello, world")
	_ = fw.Write(level.Error, "", "Hello, world")
	_ = fw.Write(level.Error, "", "Hello, world")

	fileNo = fileNo2

	_ = fw.Write(level.Debug, "", "Hello, world")
	_ = fw.Write(level.Error, "", "Hello, world")
	_ = fw.Write(level.Error, "", "Hello, world")
	_ = fw.Write(level.Error, "", "Hello, world")

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
	fw := writer.NewDefaultRotatingFileWriter(func(le level.Level, module string) string {
		return filename
	})

	defer func() {
		_ = os.Remove(filename)
	}()

	assert.NoError(t, fw.Write(level.Debug, "", "Hello, world"))
	assert.NoError(t, fw.Write(level.Error, "", "Hello, world"))

	assert.NoError(t, fw.ReOpen())
	assert.NoError(t, fw.Close())

	assert.NoError(t, fw.Write(level.Error, "", "Hello, world"))
}

func TestRotatingFileWriter_GC(t *testing.T) {
	fw := writer.NewDefaultRotatingFileWriter(func(le level.Level, module string) string {
		return fmt.Sprintf("%s-%s.log", module, le.GetLevelName())
	})

	defer func() {
		_ = os.Remove("test1-DEBUG.log")
		_ = os.Remove("test2-WARNING.log")
		_ = os.Remove("test3-DEBUG.log")
		_ = os.Remove("test4-ERROR.log")
	}()

	_ = fw.Write(level.Debug, "test1", "hello")
	_ = fw.Write(level.Warning, "test2", "hello")

	assert.ElementsMatch(t, []string{"test1-DEBUG.log", "test2-WARNING.log"}, fw.GetOpenedFiles())

	time.Sleep(2 * time.Millisecond)
	fw.GC(time.Millisecond)

	_ = fw.Write(level.Debug, "test3", "hello, world")
	_ = fw.Write(level.Error, "test4", "hello, world")

	assert.ElementsMatch(t, []string{"test3-DEBUG.log", "test4-ERROR.log"}, fw.GetOpenedFiles())
}

func TestRotatingFileWriter_AutoGC(t *testing.T) {
	fw := writer.NewDefaultRotatingFileWriter(func(le level.Level, module string) string {
		return fmt.Sprintf("%s-%s.log", module, le.GetLevelName())
	})

	fw.AutoGC(context.Background(), time.Millisecond)

	defer func() {
		_ = os.Remove("test1-DEBUG.log")
		_ = os.Remove("test2-WARNING.log")
		_ = os.Remove("test3-DEBUG.log")
		_ = os.Remove("test4-ERROR.log")
	}()

	_ = fw.Write(level.Debug, "test1", "hello")
	_ = fw.Write(level.Warning, "test2", "hello")

	assert.ElementsMatch(t, []string{"test1-DEBUG.log", "test2-WARNING.log"}, fw.GetOpenedFiles())

	time.Sleep(2 * time.Millisecond)

	_ = fw.Write(level.Debug, "test3", "hello, world")
	_ = fw.Write(level.Error, "test4", "hello, world")

	assert.ElementsMatch(t, []string{"test3-DEBUG.log", "test4-ERROR.log"}, fw.GetOpenedFiles())
}
