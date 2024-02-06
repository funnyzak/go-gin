package logger

import (
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"

	"os"
)

type Logger struct {
	Log zerolog.Logger
}

func NewLogger(logLevel string, logPath string) *Logger {
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
		return &Logger{Log: zerolog.New(zerolog.Nop())}
	}

	switch logLevel {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	return &Logger{Log: zlog}
}
