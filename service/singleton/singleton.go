package singleton

import (
	"fmt"
	"os"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"go-gin/internal/gconfig"
	"go-gin/model"
	"go-gin/pkg/logger"
	"go-gin/pkg/utils"
)

var Version = "0.0.1"

var (
	ViperConf *viper.Viper    // Viper config for the application
	Conf      *gconfig.Config // Global config for the application
	Log       *zerolog.Logger // Global logger for the application
	DB        *gorm.DB        // Global db for the application
	Cache     *cache.Cache    // Global cache for the application
	Loc       *time.Location  // Global location for the application
)

func LoadSingleton() {
	LoadCronTasks()
	// TODO: Add your initialization code here, eg Service, Task, etc.
}

func InitTimezoneAndCache() {
	var err error
	Loc, err = time.LoadLocation(Conf.Location)
	if err != nil {
		panic(err)
	}

	Cache = cache.New(5*time.Minute, 10*time.Minute)
}

func InitConfig(name string) {
	ViperConf, err := utils.ReadViperConfig(name, "yaml", []string{".", "./config", "../"})
	if err != nil {
		panic(fmt.Errorf("unable to read config: %s", err))
	}

	if err := ViperConf.Unmarshal(&Conf); err != nil {
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
