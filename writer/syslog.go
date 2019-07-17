package writer

import (
	"log/syslog"
	"sync"

	"github.com/mylxsw/asteria/level"
)

// SyslogWriter is a log writer for syslog
type SyslogWriter struct {
	syslogWriter *syslog.Writer

	network  string
	raddr    string
	priority syslog.Priority
	tag      string

	lock sync.Mutex
}

// NewSyslogWriter create a new SyslogWriter
func NewSyslogWriter(network, raddr string, priority syslog.Priority, tag string) *SyslogWriter {
	return &SyslogWriter{
		network:  network,
		raddr:    raddr,
		priority: priority,
		tag:      tag,
	}
}

func (w *SyslogWriter) Write(le level.Level, message string) error {
	writer, err := w.writer()
	if err != nil {
		return err
	}

	switch le {
	case level.Emergency:
		return writer.Emerg(message)
	case level.Alert:
		return writer.Alert(message)
	case level.Critical:
		return writer.Crit(message)
	case level.Error:
		return writer.Err(message)
	case level.Warning:
		return writer.Warning(message)
	case level.Notice:
		return writer.Notice(message)
	case level.Info:
		return writer.Info(message)
	case level.Debug:
		return writer.Debug(message)
	}

	return nil
}

func (w *SyslogWriter) writer() (*syslog.Writer, error) {
	w.lock.Lock()
	defer w.lock.Unlock()

	if w.syslogWriter == nil {
		var err error
		w.syslogWriter, err = syslog.Dial(w.network, w.raddr, w.priority, w.tag)
		if err != nil {
			return nil, err
		}
	}

	return w.syslogWriter, nil
}

func (w *SyslogWriter) ReOpen() error {
	return w.Close()
}

func (w *SyslogWriter) Close() error {
	w.lock.Lock()
	defer w.lock.Unlock()

	err := w.syslogWriter.Close()
	w.syslogWriter = nil
	return err
}
