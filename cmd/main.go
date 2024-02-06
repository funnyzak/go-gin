package main

import (
	"fmt"
	"go-gin/cmd/webserver"
	"go-gin/internal/config"
	"go-gin/internal/log"
	"go-gin/model"

	config_utils "go-gin/pkg/config"
	logger "go-gin/pkg/logger"

	flag "github.com/spf13/pflag"
)

type WebServerCliParam struct {
	ConfigName string // 配置文件名称
}

var (
	webServerCliParam WebServerCliParam
)

func main() {
	flag.CommandLine.ParseErrorsWhitelist.UnknownFlags = true
	flag.StringVarP(&webServerCliParam.ConfigName, "config", "c", "config", "config file name")
	flag.Parse()
	flag.Lookup("config").NoOptDefVal = "config"

	initConfig(webServerCliParam.ConfigName)
	initLog(config.Instance)

	webserver.ServerWeb(config.Instance)
}

func initConfig(name string) {
	_config, err := config_utils.ReadViperConfig(name, "yaml", []string{".", "./config", "../"})
	if err != nil {
		panic(fmt.Errorf("unable to read config: %s", err))
	}

	if err := _config.Unmarshal(&config.Instance); err != nil {
		panic(fmt.Errorf("unable to unmarshal config: %s", err))
	}
}

func initLog(config *config.Config) {
	logPath := config.Log.Path
	if logPath == "" {
		logPath = model.DefaultLogPath
	}
	log.ZLog = logger.NewLogger(config.Log.Level, logPath)
}
