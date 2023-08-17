package log

import (
	"github.com/sirupsen/logrus"
)

type Fields map[string]interface{}

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
