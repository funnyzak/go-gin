package singleton

import (
	"fmt"
	"os"
	"time"

	_ "github.com/ncruces/go-sqlite3/embed"
	"github.com/ncruces/go-sqlite3/gormlite"
	"github.com/patrickmn/go-cache"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"go-gin/internal/gconfig"
	"go-gin/model"
	"go-gin/pkg/logger"
	"go-gin/pkg/utils/conf"
	"go-gin/pkg/utils/file"
)

var Version = "0.1.7"

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
	LoadNotifications()
	InitUpload()
}

func InitUpload() {
	if err := file.MkdirAllIfNotExists(Conf.Upload.Dir, os.ModePerm); err != nil {
		panic(err)
	}
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
	ViperConf, err := conf.ReadViperConfig(name, "yaml", []string{".", "./config", "../"})
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
	if err = file.MkdirAllIfNotExists(path, os.ModePerm); err != nil {
		panic(err)
	}
	DB, err = gorm.Open(gormlite.Open(path), &gorm.Config{
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
