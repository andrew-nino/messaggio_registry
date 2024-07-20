package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

func SetLogrus(level string) *logrus.Logger {

	logrusLevel, err := logrus.ParseLevel(level)

	log := logrus.New()

	if err != nil {
		log.SetLevel(logrus.DebugLevel)
	} else {
		log.SetLevel(logrusLevel)
	}

	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	log.SetOutput(os.Stdout)

	return log
}
