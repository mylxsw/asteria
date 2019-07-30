package writer_test

import (
	"testing"

	"github.com/mylxsw/asteria/level"
	"github.com/mylxsw/asteria/writer"
	"github.com/stretchr/testify/assert"
)

type MockWriter struct {
	WriteCount  int
	CloseCount  int
	ReOpenCount int
}

func (m *MockWriter) Write(le level.Level, module string, message string) error {
	m.WriteCount++
	return nil
}

func (m *MockWriter) ReOpen() error {
	m.ReOpenCount++
	return nil
}

func (m *MockWriter) Close() error {
	m.CloseCount++
	return nil
}

func TestStackWriter_Write(t *testing.T) {
	m1 := &MockWriter{}
	m2 := &MockWriter{}

	stack := writer.NewStackWriter()

	stack.PushWithLevels(m1)
	stack.PushWithLevels(m2)

	_ = stack.Write(level.Debug, "", "Hello, world")
	assert.Equal(t, m1.WriteCount, m2.WriteCount)
	assert.Equal(t, m1.WriteCount, 1)

	_ = stack.Write(level.Error, "", "Hello, error")
	assert.Equal(t, m1.WriteCount, m2.WriteCount)
	assert.Equal(t, m1.WriteCount, 2)

	_ = stack.ReOpen()
	assert.Equal(t, m1.ReOpenCount, m2.ReOpenCount)
	assert.Equal(t, m1.ReOpenCount, 1)

	_ = stack.ReOpen()
	assert.Equal(t, m1.ReOpenCount, m2.ReOpenCount)
	assert.Equal(t, m1.ReOpenCount, 2)

	_ = stack.Close()
	assert.Equal(t, m1.CloseCount, m2.CloseCount)
	assert.Equal(t, m1.CloseCount, 1)

	_ = stack.Close()
	assert.Equal(t, m1.CloseCount, m2.CloseCount)
	assert.Equal(t, m1.CloseCount, 2)
}

func TestStackWriter_Distribute(t *testing.T) {
	m1 := &MockWriter{}
	m2 := &MockWriter{}
	m3 := &MockWriter{}
	m4 := &MockWriter{}

	stack := writer.NewStackWriter()
	stack.PushWithLevels(m1, level.Debug)
	stack.PushWithLevels(m2, level.Error)
	stack.PushWithLevels(m3)

	_ = stack.Write(level.Debug, "", "hello")
	_ = stack.Write(level.Debug, "", "hello")
	_ = stack.Write(level.Debug, "", "hello")
	_ = stack.Write(level.Error, "", "hello")
	_ = stack.Write(level.Warning, "", "hello")
	_ = stack.Write(level.Error, "", "hello")

	assert.Equal(t, 3, m1.WriteCount)
	assert.Equal(t, 2, m2.WriteCount)
	assert.Equal(t, 6, m3.WriteCount)

	stack.Push(m4, func(le level.Level, module string, message string) bool {
		return module == "write_log"
	})

	_ = stack.Write(level.Debug, "", "hello")
	_ = stack.Write(level.Warning, "write_log", "hello")

	assert.Equal(t, 1, m4.WriteCount)
	assert.Equal(t, 4, m1.WriteCount)
	assert.Equal(t, 2, m2.WriteCount)
	assert.Equal(t, 8, m3.WriteCount)
}
