package logger

import "github.com/sirupsen/logrus"

type Logger struct {
	*logrus.Logger
}

func NewLogger() *Logger {
	log := logrus.New()
	log.SetReportCaller(true) // include caller info if needed
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:     true,
	})

	return &Logger{log}
}
