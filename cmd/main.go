package main

import (
	"context"
	"go-gin/cmd/srv/controller"
	"go-gin/service/singleton"

	"github.com/ory/graceful"
	flag "github.com/spf13/pflag"
)

type CliParam struct {
	ConfigName string // Config file name
	Port       uint   // Server port
}

var (
	cliParam CliParam
)

func main() {
	flag.CommandLine.ParseErrorsWhitelist.UnknownFlags = true
	flag.StringVarP(&cliParam.ConfigName, "config", "c", "config", "config file name")
	flag.UintVarP(&cliParam.Port, "port", "p", 0, "server port")
	flag.Parse()
	flag.Lookup("config").NoOptDefVal = "config"

	singleton.InitConfig(cliParam.ConfigName)
	singleton.InitLog(singleton.Conf)
	singleton.InitDBFromPath(singleton.Conf.DBPath)
	initService()

	port := singleton.Conf.Server.Port
	if cliParam.Port != 0 {
		port = cliParam.Port
	}

	srv := controller.ServerWeb(port)

	if err := graceful.Graceful(func() error {
		return srv.ListenAndServe()
	}, func(c context.Context) error {
		singleton.Log.Info().Msg("Graceful::START")
		srv.Shutdown(c)
		return nil
	}); err != nil {
		singleton.Log.Err(err).Msg("Graceful::Error")
	}

}

func initService() {
	singleton.InitSingleton()
}
