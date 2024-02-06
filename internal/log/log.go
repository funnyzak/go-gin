package log

import (
	"github.com/funnyzak/go-gin/internal/config"
	"github.com/funnyzak/go-gin/model"
	"github.com/funnyzak/go-gin/pkg/logger"
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
