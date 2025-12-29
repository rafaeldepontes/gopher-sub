package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var logger *log.Logger

func GetLogger() *log.Logger {
	if logger != nil {
		return logger
	}
	initLogger()
	return logger
}

func initLogger() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	logger = &log.Logger{
		Out:   os.Stderr,
		Level: log.DebugLevel,
		Formatter: &prefixed.TextFormatter{
			DisableColors:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			ForceFormatting: true,
		},
	}
}
