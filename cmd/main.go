package main

import (
	"go-gin/cmd/srv"
	"go-gin/service/singleton"

	flag "github.com/spf13/pflag"
)

type CliParam struct {
	ConfigName string // 配置文件名称
}

var (
	svrCliParam CliParam
)

func main() {
	flag.CommandLine.ParseErrorsWhitelist.UnknownFlags = true
	flag.StringVarP(&svrCliParam.ConfigName, "config", "c", "config", "config file name")
	flag.Parse()
	flag.Lookup("config").NoOptDefVal = "config"

	singleton.InitConfig(svrCliParam.ConfigName)
	singleton.InitLog(singleton.Config)
	// singleton.InitDBFromPath(singleton.Config.DB_Path)
	initService()

	srv.ServerWeb(singleton.Config)
}

func initService() {
	singleton.InitSingleton()
}
