package health

import (
	gLog "github.com/InVisionApp/go-logger"
	"github.com/go-seidon/provider/logging"
)

type healthLog struct {
	client logging.Logger
}

func (l *healthLog) Info(args ...interface{}) {
	l.client.Info(args...)
}

func (l *healthLog) Debug(args ...interface{}) {
	l.client.Debug(args...)
}

func (l *healthLog) Error(args ...interface{}) {
	l.client.Error(args...)
}

func (l *healthLog) Warn(args ...interface{}) {
	l.client.Warn(args...)
}

func (l *healthLog) Infof(format string, args ...interface{}) {
	l.client.Infof(format, args...)
}

func (l *healthLog) Debugf(format string, args ...interface{}) {
	l.client.Debugf(format, args...)
}

func (l *healthLog) Errorf(format string, args ...interface{}) {
	l.client.Errorf(format, args...)
}

func (l *healthLog) Warnf(format string, args ...interface{}) {
	l.client.Warnf(format, args...)
}

func (l *healthLog) Infoln(args ...interface{}) {
	l.client.Infoln(args...)
}

func (l *healthLog) Debugln(args ...interface{}) {
	l.client.Debugln(args...)
}

func (l *healthLog) Errorln(args ...interface{}) {
	l.client.Errorln(args...)
}

func (l *healthLog) Warnln(args ...interface{}) {
	l.client.Warnln(args...)
}

func (l *healthLog) WithFields(fs gLog.Fields) gLog.Logger {
	client := l.client.WithFields(fs)
	return &healthLog{
		client: client,
	}
}
