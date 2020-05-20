package log

type Logger interface {
	WithFields(c Fields) Logger
	Emergency(v ...interface{})
	Alert(v ...interface{})
	Critical(v ...interface{})
	Error(v ...interface{})
	Warning(v ...interface{})
	Notice(v ...interface{})
	Info(v ...interface{})
	Debug(v ...interface{})
	Emergencyf(format string, v ...interface{})
	Alertf(format string, v ...interface{})
	Criticalf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Warningf(format string, v ...interface{})
	Noticef(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Debugf(format string, v ...interface{})
}
