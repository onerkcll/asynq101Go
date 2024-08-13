package logHelpers

import (
	"github.com/sirupsen/logrus"
	AppConf "github.com/spf13/viper"
	"io"
	"os"
)

var asynqLogger = logrus.New()

func InitializeLogger() {
	asynqLogger.SetFormatter(&logrus.TextFormatter{})
	asynqLogger.SetOutput(os.Stdout)
	debugMode := AppConf.GetBool("Debug")
	if debugMode {
		asynqLogger.SetLevel(logrus.DebugLevel)
	} else {
		asynqLogger.SetLevel(logrus.InfoLevel)
	}
}

func DiscardOutput() {
	asynqLogger.SetOutput(io.Discard)
}

func GetLogger() *logrus.Logger {
	return asynqLogger
}
