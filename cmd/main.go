package main

import (
	"github.com/funnyzak/gogin/cmd/webserver"
	"github.com/funnyzak/gogin/internal/config"
	"github.com/funnyzak/gogin/internal/log"
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

	config.Init(webServerCliParam.ConfigName)
	log.InitLog(config.Instance)

	webserver.ServerWeb(config.Instance)
}
