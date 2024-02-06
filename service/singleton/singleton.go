package singleton

import (
	"fmt"
	"go-gin/internal/config"
	"go-gin/model"

	config_utils "go-gin/pkg/config"
	logger "go-gin/pkg/logger"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	Log    *zerolog.Logger
	DB     *gorm.DB
)

func InitSingleton() {
	// TOO: init db
}

func InitConfig(name string) {
	_config, err := config_utils.ReadViperConfig(name, "yaml", []string{".", "./config", "../"})
	if err != nil {
		panic(fmt.Errorf("unable to read config: %s", err))
	}

	if err := _config.Unmarshal(&Config); err != nil {
		panic(fmt.Errorf("unable to unmarshal config: %s", err))
	}
}

func InitLog(config *config.Config) {
	logPath := config.Log.Path
	if logPath == "" {
		logPath = model.DefaultLogPath
	}
	Log = logger.NewLogger(config.Log.Level, logPath)
}
