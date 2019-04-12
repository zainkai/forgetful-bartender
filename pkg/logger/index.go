package logger

import (
	"os"
	"log"
)

var Logger *log.Logger

func initLogger() *log.Logger {
	Logger = log.New(os.Stdout, "", log.LstdFlags)
	return Logger
}

func Log(packageName string, msg string, data interface{}) {
	if Logger == nil {
		initLogger()
	}

	Logger.SetPrefix(packageName)
	Logger.Printf(msg, data)
}

func Err(packageName string, msg string, data interface{}) {
	if Logger == nil {
		initLogger()
	}

	Logger.SetPrefix(packageName)
	Logger.Fatalf(msg, data)
}