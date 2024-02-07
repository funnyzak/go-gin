package controller

import (
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
	c.HTML(http.StatusOK, "index", gin.H{})
}

func (p *commonPage) ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
