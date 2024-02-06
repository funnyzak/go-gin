package logger

import (
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"

	"os"
)

func NewLogger(logLevel string, logPath string) *zerolog.Logger {
	if logPath == "" {
		logPath = "logs/log.log"
	}
	logger := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    50,
		MaxBackups: 15,
		MaxAge:     15,
		LocalTime:  true,
		Compress:   true,
	}
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05"}
	multi := zerolog.MultiLevelWriter(logger, consoleWriter)
	zlog := zerolog.New(multi).With().Timestamp().Logger()

	if logLevel == "" || logLevel == "none" || logLevel == "null" || logLevel == "nil" || logLevel == "off" {
		zlog.Level(zerolog.NoLevel)
		return &zlog
	}

	switch logLevel {
	case "debug":
		zlog = zlog.Level(zerolog.DebugLevel)
	case "warn":
		zlog = zlog.Level(zerolog.WarnLevel)
	case "error":
		zlog = zlog.Level(zerolog.ErrorLevel)
	case "fatal":
		zlog = zlog.Level(zerolog.FatalLevel)
	case "panic":
		zlog = zlog.Level(zerolog.PanicLevel)
	case "trace":
		zlog = zlog.Level(zerolog.TraceLevel)
	default:
		zlog = zlog.Level(zerolog.InfoLevel)
	}

	return &zlog
}
