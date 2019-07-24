package writer_test

import (
	"testing"

	"github.com/mylxsw/asteria/level"
	"github.com/mylxsw/asteria/writer"
	"github.com/stretchr/testify/assert"
)

type MockIOWriter struct {
	count int
}

func (w *MockIOWriter) Write(p []byte) (n int, err error) {
	w.count++

	return 0, nil
}

func TestStreamWriter_Write(t *testing.T) {
	mockWriter := &MockIOWriter{}
	sw := writer.NewStreamWriter(mockWriter)

	_ = sw.Write(level.Debug, "Hello, world")
	_ = sw.Write(level.Warning, "Yes, you are")

	assert.Equal(t, 2, mockWriter.count)
	assert.NoError(t, sw.Close())
	assert.NoError(t, sw.ReOpen())
}
