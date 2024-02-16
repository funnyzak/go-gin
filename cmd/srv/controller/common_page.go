package controller

import (
	"go-gin/internal/gogin"
	"net/http"

	"github.com/gin-gonic/gin"
)

type commonPage struct {
	r *gin.Engine
}

func (cp *commonPage) serve() {
	cr := cp.r.Group("")
	cr.GET("/", cp.home)
	cr.GET("/ping", cp.ping)
}

func (p *commonPage) home(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gogin.CommonEnvironment(c, gin.H{}))
}

func (p *commonPage) ping(c *gin.Context) {
	c.Writer.WriteString("pong")
}
