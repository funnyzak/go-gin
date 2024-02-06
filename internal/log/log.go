package log

import (
	"github.com/funnyzak/gogin/internal/config"
	"github.com/funnyzak/gogin/model"
	"github.com/funnyzak/gogin/pkg/logger"
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
