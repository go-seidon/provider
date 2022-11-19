package logrus

import (
	"context"
	"io"
	"os"

	"github.com/go-seidon/provider/logging"
	"github.com/sirupsen/logrus"
)

const (
	FIELD_SERVICE = "service"
	FIELD_ERROR   = "error"
)

type LogMessage struct {
	Timestamp      string `json:"timestamp"`
	Message        string `json:"message"`
	Severity       string `json:"severity"`
	ReportLocation struct {
		FilePath     string `json:"filePath,omitempty"`
		LineNumber   int    `json:"lineNumber,omitempty"`
		FunctionName string `json:"functionName,omitempty"`
	} `json:"reportLocation,omitempty"`
	Service  interface{}            `json:"service,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type logrusLog struct {
	client *logrus.Entry
}

func (l *logrusLog) Log(level string, args ...interface{}) {
	lv, err := logrus.ParseLevel(level)
	if err != nil {
		lv = logrus.InfoLevel
	}

	l.client.Log(lv, args...)
}

func (l *logrusLog) Info(args ...interface{}) {
	l.client.Info(args...)
}

func (l *logrusLog) Debug(args ...interface{}) {
	l.client.Debug(args...)
}

func (l *logrusLog) Error(args ...interface{}) {
	l.client.Error(args...)
}

func (l *logrusLog) Warn(args ...interface{}) {
	l.client.Warn(args...)
}

func (l *logrusLog) Logf(level string, format string, args ...interface{}) {
	lv, err := logrus.ParseLevel(level)
	if err != nil {
		lv = logrus.InfoLevel
	}

	l.client.Logf(lv, format, args...)
}

func (l *logrusLog) Infof(format string, args ...interface{}) {
	l.client.Infof(format, args...)
}

func (l *logrusLog) Debugf(format string, args ...interface{}) {
	l.client.Debugf(format, args...)
}

func (l *logrusLog) Errorf(format string, args ...interface{}) {
	l.client.Errorf(format, args...)
}

func (l *logrusLog) Warnf(format string, args ...interface{}) {
	l.client.Warnf(format, args...)
}

func (l *logrusLog) Logln(level string, msg ...interface{}) {
	lv, err := logrus.ParseLevel(level)
	if err != nil {
		lv = logrus.InfoLevel
	}

	l.client.Logln(lv, msg...)
}

func (l *logrusLog) Infoln(msg ...interface{}) {
	l.client.Infoln(msg...)
}

func (l *logrusLog) Debugln(msg ...interface{}) {
	l.client.Debugln(msg...)
}

func (l *logrusLog) Errorln(msg ...interface{}) {
	l.client.Errorln(msg...)
}

func (l *logrusLog) Warnln(msg ...interface{}) {
	l.client.Warnln(msg...)
}

func (l *logrusLog) WithFields(fs map[string]interface{}) logging.Logger {
	entry := l.client.WithFields(fs)
	nl := &logrusLog{
		client: entry,
	}
	return nl
}

func (l *logrusLog) WithError(err error) logging.Logger {
	entry := l.client.WithField(FIELD_ERROR, err)
	nl := &logrusLog{
		client: entry,
	}
	return nl
}

func (l *logrusLog) WithContext(ctx context.Context) logging.Logger {
	entry := l.client.WithContext(ctx)
	nl := &logrusLog{
		client: entry,
	}
	return nl
}

func (l *logrusLog) WriterLevel(level string) io.Writer {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		lvl = logrus.ErrorLevel
	}
	return l.client.WriterLevel(lvl)
}

func NewLogger(opts ...LogOption) *logrusLog {
	p := LogParam{
		StackSkip: []string{
			"github.com/sirupsen/logrus",
		},
	}
	for _, opt := range opts {
		opt(&p)
	}

	client := logrus.New()
	client.SetOutput(os.Stdout)
	client.SetFormatter(&GoFormatter{
		PrettyPrint: p.PrettyPrintEnabled,
		StackSkip:   p.StackSkip,
	})
	if p.DebuggingEnabled {
		client.SetLevel(logrus.DebugLevel)
	}

	appCtx := logrus.Fields{}
	if p.AppCtxEnabled {
		appCtx = logrus.Fields{
			FIELD_SERVICE: map[string]interface{}{
				"name":    p.AppName,
				"version": p.AppVersion,
			},
		}
	}
	entry := client.WithFields(appCtx)

	l := &logrusLog{
		client: entry,
	}
	return l
}
