package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type showPage struct {
	r *gin.Engine
}

func (sp *showPage) serve() {
	gr := sp.r.Group("")
	gr.GET("/post/:num", sp.postDetail)
}

func (sp *showPage) postDetail(c *gin.Context) {
	num := c.Param("num")
	c.HTML(http.StatusOK, "post/detail", gin.H{
		"Num": num,
	})
}
