package log

import (
	"github.com/sirupsen/logrus"
)

type Fields map[string]interface{}

type Level uint32

type Formatter struct {
	logrus.TextFormatter
}

const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

func UseTimestamp() {
	logrus.SetFormatter(&Formatter{logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}})
}

func SetLevel(level Level) {
	logrus.SetLevel(logrus.Level(level))
}

func Info(m string) {
	logrus.Info(m)
}

func InfoWithFields(m string, f Fields) {
	logrus.WithFields(logrus.Fields(f)).Info(m)
}

func Debug(m string) {
	logrus.Debug(m)
}

func DebugWithFields(m string, f Fields) {
	logrus.WithFields(logrus.Fields(f)).Debug(m)
}

func Warning(m string) {
	logrus.Warning(m)
}

func WarningWithFields(m string, f Fields) {
	logrus.WithFields(logrus.Fields(f)).Warning(m)
}

func Error(m string) {
	logrus.Error(m)
}

func ErrorWithFields(m string, f Fields) {
	logrus.WithFields(logrus.Fields(f)).Error(m)
}
