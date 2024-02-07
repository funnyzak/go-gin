package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type sharePage struct {
	r *gin.Engine
}

func (cp *sharePage) serve() {
	cr := cp.r.Group("share")
	cr.GET("/post", cp.post)
}

func (p *sharePage) post(c *gin.Context) {
	c.HTML(http.StatusOK, "share/post", gin.H{})
}
