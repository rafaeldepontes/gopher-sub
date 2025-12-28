package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func InitLogger() *log.Logger {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	return &log.Logger{
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
