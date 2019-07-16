package asteria_test

import (
	"math/rand"
	"os"
	"testing"
	"time"

	"asteria"
)

func TestConcurrentWrite(t *testing.T) {
	var logger = asteria.Module("test.concurrent")

	var logfile = "./test.log"

	logger.Writer(asteria.NewSingleFileWriter(logfile))

	rand.Seed(time.Now().Unix())
	for i := 0; i < 1000; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			logger.Debugf("inner - %d（%d）", i, rand.Intn(10))
		}(i)
	}

	for i := 0; i < 1000; i++ {
		logger.Debugf("outer - %d", i)
	}

	os.Remove(logfile)
}
