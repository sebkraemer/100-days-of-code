package logging

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var AppLog *LoggerClass

type LoggerClass struct {
	Log *logrus.Logger
}

func SetupLogging() {
	AppLog = &LoggerClass{}
	AppLog.Log = logrus.New()
	file, err := os.OpenFile("sudokuserver.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		AppLog.Log.Fatal((err))
	}
	mw := io.MultiWriter(os.Stdout, file)
	AppLog.Log.SetOutput(mw)
	AppLog.Log.SetFormatter(&logrus.JSONFormatter{})
}

func (logger *LoggerClass) WriteLogError(msg string, fields map[string]interface{}) {
	logger.Log.WithFields(fields).Error(msg)
}

func (logger *LoggerClass) WriteLogInfo(msg string, fields map[string]interface{}) {
	logger.Log.WithFields(fields).Info(msg)
}
