package controller

import (
	"go-gin/internal/gogin"
	"go-gin/pkg/mygin"
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
		Redirect: singleton.Conf.Site.BaseURL + "/user/profile",
	}))
	gr.GET("/register", gp.register)
	gr.GET("/login", gp.login)
}

func (gp *guestPage) register(c *gin.Context) {
	if !singleton.Conf.EnableUserRegistration {
		gogin.ShowErrorPage(c, mygin.ErrInfo{
			Title: "Registration is disabled",
			Code:  http.StatusNotAcceptable,
			Msg:   "Please contact the administrator",
			Link:  singleton.Conf.Site.BaseURL,
			Btn:   "Return to home",
		}, true)
		return
	}
	c.HTML(http.StatusOK, "register", gogin.CommonEnvironment(
		c, gin.H{
			"Title": "Register",
		},
	))
}

func (gp *guestPage) login(c *gin.Context) {
	c.HTML(http.StatusOK, "login", gogin.CommonEnvironment(
		c, gin.H{
			"Title": "Login",
		},
	))
}
