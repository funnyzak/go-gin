package log

import (
	"go-gin/internal/config"
	"go-gin/model"
	"go-gin/pkg/logger"
)

var ZLog *logger.Logger

func InitLog(config *config.Config) error {
	logPath := config.Log.Path
	if logPath == "" {
		logPath = model.DefaultLogPath
	}
	ZLog = logger.NewLogger(config.Log.Level, logPath)
	return nil
}
