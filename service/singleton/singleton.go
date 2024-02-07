package singleton

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"go-gin/internal/gconfig"
	"go-gin/model"
	"go-gin/pkg/logger"
	"go-gin/pkg/utils"
)

var Version = "0.0.1"

var (
	Conf *gconfig.Config
	Log  *zerolog.Logger
	DB   *gorm.DB
)

func InitSingleton() {
	// TOO: init db
}

func InitConfig(name string) {
	_config, err := utils.ReadViperConfig(name, "yaml", []string{".", "./config", "../"})
	if err != nil {
		panic(fmt.Errorf("unable to read config: %s", err))
	}

	if err := _config.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("unable to unmarshal config: %s", err))
	}
}

// Initialize the logger
func InitLog(conf *gconfig.Config) {
	logPath := conf.Log.Path
	if logPath == "" {
		logPath = Conf.DBPath
	}
	Log = logger.NewLogger(conf.Log.Level, logPath)
}

// InitDBFromPath initialize the database from the given path
func InitDBFromPath(path string) {
	var err error
	if err = utils.MkdirAllIfNotExists(path, os.ModePerm); err != nil {
		panic(err)
	}
	DB, err = gorm.Open(sqlite.Open(path), &gorm.Config{
		CreateBatchSize: 200,
	})
	if err != nil {
		panic(err)
	}
	if Conf.Debug {
		DB = DB.Debug()
	}
	err = DB.AutoMigrate(&model.User{}, &model.Post{})
	if err != nil {
		panic(err)
	}
}
