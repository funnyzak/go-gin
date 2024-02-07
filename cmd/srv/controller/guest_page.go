package controller

import (
	"go-gin/internal/gogin"
	"go-gin/service/singleton"
	"net/http"

	"github.com/gin-gonic/gin"
)

type guestPage struct {
	r *gin.Engine
}

func (gp *guestPage) serve() {
	gr := gp.r.Group("")
	gr.Use(gogin.Authorize(gogin.AuthorizeOption{
		Guest:    true,
		IsPage:   true,
		Msg:      "You are already logged in",
		Btn:      "Return to home",
		Redirect: singleton.Conf.Site.BaseURL,
	}))
	gr.GET("/register", gp.register)
	gr.GET("/login", gp.login)
}

func (gp *guestPage) register(c *gin.Context) {
	c.HTML(http.StatusOK, "register", gin.H{})
}

func (gp *guestPage) login(c *gin.Context) {
	c.HTML(http.StatusOK, "login", gin.H{})
}
