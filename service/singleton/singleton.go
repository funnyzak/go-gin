package singleton

import (
	"fmt"
	"go-gin/internal/config"
	"go-gin/model"

	config_utils "go-gin/pkg/config"
	logger "go-gin/pkg/logger"

	"github.com/rs/zerolog"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Version = "0.0.1"

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

// Initialize the logger
func InitLog(config *config.Config) {
	logPath := config.Log.Path
	if logPath == "" {
		logPath = model.DefaultLogPath
	}
	Log = logger.NewLogger(config.Log.Level, logPath)
}

// InitDBFromPath 从给出的文件路径中加载数据库
func InitDBFromPath(path string) {
	var err error
	DB, err = gorm.Open(sqlite.Open(path), &gorm.Config{
		CreateBatchSize: 200,
	})
	if err != nil {
		panic(err)
	}
	if Config.Debug {
		DB = DB.Debug()
	}
	// err = DB.AutoMigrate(&model.User{}, &model.Role{}, &model.Permission{}, &model.UserRole{}, &model.RolePermission{})
	if err != nil {
		panic(err)
	}
}
