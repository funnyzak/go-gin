package webserver

import (
	"fmt"

	"github.com/funnyzak/go-gin/internal/config"
	"github.com/funnyzak/go-gin/model"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func ServerWeb(config *config.Config) {
	gin.SetMode(gin.ReleaseMode)
	r := NewRoute(config)
	if config.Debug {
		gin.SetMode(gin.DebugMode)
		pprof.Register(r, model.DefaultPprofRoutePath)
	}
	r.Run(fmt.Sprintf(":%v", config.Server.Port))
}
