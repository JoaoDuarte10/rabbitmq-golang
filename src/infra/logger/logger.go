package logger

import (
	"log"
	"os"

	logger "github.com/sirupsen/logrus"
)

type Logger interface {
	Info(message string)
	Error(message string)
	Warn(message string)
}

type LoggerAdapter struct {
	ConsoleEnable bool
}

func (l *LoggerAdapter) Info(message string) {
	l.prepareLogger()
	logger.Info(message)
}

func (l *LoggerAdapter) Error(message string) {
	l.prepareLogger()
	logger.Error(message)
}

func (l *LoggerAdapter) Warn(message string) {
	l.prepareLogger()
	logger.Warn(message)
}

func (l *LoggerAdapter) prepareLogger() {
	logger.SetFormatter(&logger.JSONFormatter{
		FieldMap: logger.FieldMap{
			logger.FieldKeyTime: "timestamp",
			logger.FieldKeyMsg:  "message",
		},
	})

	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	if l.ConsoleEnable {
		logger.SetOutput(os.Stdout)
	} else {
		logger.SetOutput(file)
	}
}
