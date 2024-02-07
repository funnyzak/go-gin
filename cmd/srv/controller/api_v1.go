package controller

import (
	"fmt"
	"go-gin/internal/gogin"
	"go-gin/service/singleton"

	"github.com/gin-gonic/gin"
)

type apiV1 struct {
	r gin.IRouter
}

func (v *apiV1) serve() {
	r := v.r.Group("")
	// API
	r.Use(gogin.Authorize(gogin.AuthorizeOption{
		User:     true,
		IsPage:   false,
		Msg:      "Please log in first",
		Btn:      "Log in",
		Redirect: fmt.Sprintf("%s/login", singleton.Conf.Site.BaseURL),
	}))
	r.PUT("/post", v.putPost)

	user := v.r.Group("user")
	{
		user.GET("/info", v.getUserInfo)
		user.GET("/logout", v.logout)
		user.GET("/refresh", v.refresh)
	}
}

func (v *apiV1) putPost(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "post",
	})
}

func (v *apiV1) getUserInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "user info",
	})
}

func (v *apiV1) logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "logout",
	})
}

func (v *apiV1) refresh(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "refresh",
	})
}
