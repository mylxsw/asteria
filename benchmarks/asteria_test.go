package benchmarks

import (
	"io/ioutil"
	"testing"

	"github.com/mylxsw/asteria/formatter"
	"github.com/mylxsw/asteria/log"
	"github.com/mylxsw/asteria/writer"
)

func init() {
	log.DefaultLogWriter(writer.NewStreamWriter(ioutil.Discard))
	log.DefaultLogFormatter(formatter.NewJSONFormatter())
}

func BenchmarkAsteriaSimple(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Debug("Hello, world")
	}
}

func BenchmarkAsteriaWithFields(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.WithFields(log.Fields{
			"int":    12,
			"string": "Hello",
			"float":  33.4,
			"nested": log.Fields{
				"int":    12,
				"string": "Hello",
				"float":  33.4,
			},
		}).Debug("Hello, world")
	}
}

func BenchmarkAsteriaWithFieldsAndStack(b *testing.B) {
	log.DefaultWithFileLine(true)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.WithFields(log.Fields{
			"int":    12,
			"string": "Hello",
			"float":  33.4,
			"nested": log.Fields{
				"int":    12,
				"string": "Hello",
				"float":  33.4,
			},
		}).Debug("Hello, world")
	}
}

func BenchmarkAsteriaWithModule(b *testing.B) {
	var logger = log.Module("test")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.WithFields(log.Fields{
			"int":    12,
			"string": "Hello",
			"float":  33.4,
			"nested": log.Fields{
				"int":    12,
				"string": "Hello",
				"float":  33.4,
			},
		}).Debug("Hello, world")
	}
}

func BenchmarkAsteriaWithJsonFormatter(b *testing.B) {
	log.SetFormatter(formatter.NewJSONWithTimeFormatter())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.WithFields(log.Fields{
			"int":    12,
			"string": "Hello",
			"float":  33.4,
			"nested": log.Fields{
				"int":    12,
				"string": "Hello",
				"float":  33.4,
			},
		}).Debug("Hello, world")
	}
}
