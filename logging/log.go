package logging

import (
	"context"
	"io"
)

type Logger interface {
	SimpleLog
	FormatedLog
	LineLog
	CustomLog
}

type SimpleLog interface {
	Log(level string, args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
	Error(args ...interface{})
	Warn(args ...interface{})
}

type FormatedLog interface {
	Logf(level string, format string, args ...interface{})
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
}

type LineLog interface {
	Logln(level string, msg ...interface{})
	Infoln(msg ...interface{})
	Debugln(msg ...interface{})
	Errorln(msg ...interface{})
	Warnln(msg ...interface{})
}

type CustomLog interface {
	WithFields(fs map[string]interface{}) Logger
	WithError(err error) Logger
	WithContext(ctx context.Context) Logger
	WriterLevel(level string) io.Writer
}
