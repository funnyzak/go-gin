package srv

import (
	"fmt"

	"go-gin/internal/config"
	"go-gin/model"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func Server(config *config.Config) {
	gin.SetMode(gin.ReleaseMode)
	r := NewRoute(config)
	if config.Debug {
		gin.SetMode(gin.DebugMode)
		pprof.Register(r, model.DefaultPprofRoutePath)
	}
	r.Run(fmt.Sprintf(":%v", config.Server.Port))
}
