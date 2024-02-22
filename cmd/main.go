package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ory/graceful"
	flag "github.com/spf13/pflag"

	"go-gin/cmd/srv/controller"
	"go-gin/pkg/mygin"
	"go-gin/pkg/utils"
	"go-gin/pkg/utils/ip"
	"go-gin/service/singleton"
)

type ClIParam struct {
	Version    bool   // Show version
	ConfigName string // Config file name
	Port       uint   // Server port
	Debug      bool   // Debug mode
}

var cliParam ClIParam

func parseCommandLineParams() (cliParam ClIParam) {
	flag.CommandLine.ParseErrorsWhitelist.UnknownFlags = true
	flag.BoolVarP(&cliParam.Version, "version", "v", false, "show version")
	flag.StringVarP(&cliParam.ConfigName, "config", "c", "config", "config file name")
	flag.UintVarP(&cliParam.Port, "port", "p", 0, "server port")
	flag.BoolVarP(&cliParam.Debug, "debug", "d", false, "debug mode")
	flag.Parse()
	flag.Lookup("config").NoOptDefVal = "config"
	return cliParam
}

func loadConfig() {
	cliParam = parseCommandLineParams()

	singleton.InitConfig(cliParam.ConfigName)
	if cliParam.Port > 0 && cliParam.Port < 65536 {
		singleton.Conf.Server.Port = cliParam.Port
	}
	if cliParam.Debug {
		singleton.Conf.Debug = cliParam.Debug
	}
}

func startupOutput(httpserver *http.Server) {
	if singleton.Conf.Debug {
		fmt.Println()
		fmt.Printf("Service version: %s\n", utils.Colorize(utils.ColorGreen, singleton.Version))

		fmt.Println()
		fmt.Println("Server available routes:")
		mygin.PrintRoute(httpserver.Handler.(*gin.Engine))

		fmt.Println()
		fmt.Println("Server running with config:")
		utils.PrintStructFieldsAndValues(singleton.Conf, "")
	}

	fmt.Println()
	fmt.Println("Server is running at:")
	fmt.Printf(" - %-7s: %s\n", "Local", utils.Colorize(utils.ColorGreen, fmt.Sprintf("http://127.0.0.1:%d", singleton.Conf.Server.Port)))
	ipv4s, err := ip.GetIPv4NetworkIPs()
	if ipv4s != nil && err == nil {
		for _, ip := range ipv4s {
			fmt.Printf(" - %-7s: %s\n", "Network", utils.Colorize(utils.ColorGreen, fmt.Sprintf("http://%s:%d", ip, singleton.Conf.Server.Port)))
		}
	}
	fmt.Println()
}

func init() {
	loadConfig()
	singleton.InitLog(singleton.Conf)
	singleton.InitTimezoneAndCache()
	singleton.InitDBFromPath(singleton.Conf.DBPath)
	initService()
}

func main() {
	if cliParam.Version {
		fmt.Println(singleton.Version)
		return
	}

	srv := controller.ServerWeb(singleton.Conf.Server.Port)
	if err := graceful.Graceful(func() error {
		startupOutput(srv)
		return srv.ListenAndServe()
	}, func(c context.Context) error {
		fmt.Print(utils.Colorize("Server is shutting down", utils.ColorRed))
		srv.Shutdown(c)
		return nil
	}); err != nil {
		fmt.Printf("Server is shutting down with error: %s", utils.Colorize(utils.ColorRed, err.Error()))
	}
}

func initService() {
	singleton.LoadSingleton()

	// if _, err := singleton.Cron.AddFunc("0 * * * * *", sayHello); err != nil {
	// 	panic(err)
	// }
}

// func sayHello() {
// 	singleton.Log.Info().Msg("Hello world, I am a cron task")
// 	// singleton.SendNotificationByType("wecom", "Hello world", "I am a cron task")
// }
