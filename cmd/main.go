package main

import (
	"go-gin/cmd/webserver"
	"go-gin/service/singleton"

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

	singleton.InitConfig(webServerCliParam.ConfigName)
	singleton.InitLog(singleton.Config)
	initService()

	webserver.ServerWeb(singleton.Config)
}

func initService() {
	singleton.InitSingleton()
}
