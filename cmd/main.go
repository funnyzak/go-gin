package main

import (
	"context"
	"fmt"
	"go-gin/cmd/srv/controller"
	"go-gin/pkg/utils"
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

	startOutput := func() {
		fmt.Println()
		fmt.Println("Server is running with config:")
		utils.PrintStructFieldsAndValues(singleton.Conf, "")

		fmt.Println()
		fmt.Println("Server is running at:")
		fmt.Printf(" - %-7s: %s\n", "Local", utils.Colorize(utils.ColorGreen, fmt.Sprintf("http://127.0.0.1:%d", port)))
		ipv4s, err := utils.GetIPv4NetworkIPs()
		if ipv4s != nil && err == nil {
			for _, ip := range ipv4s {
				fmt.Printf(" - %-7s: %s\n", "Network", utils.Colorize(utils.ColorGreen, fmt.Sprintf("http://%s:%d", ip, port)))
			}
		}
		fmt.Println()
	}

	if err := graceful.Graceful(func() error {
		startOutput()
		return srv.ListenAndServe()
	}, func(c context.Context) error {
		fmt.Print(utils.Colorize("Server is shutting down", utils.ColorRed))
		srv.Shutdown(c)
		return nil
	}); err != nil {
		fmt.Println(utils.Colorize("Server is shutting down with error: %s", utils.ColorRed), err)
	}
}

func initService() {
	singleton.InitSingleton()
}
