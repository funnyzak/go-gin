package controller

import (
	"go-gin/internal/gogin"
	"go-gin/service/singleton"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userPage struct {
	r *gin.Engine
}

func (up *userPage) serve() {
	gr := up.r.Group("")
	gr.Use(gogin.Authorize(
		gogin.AuthorizeOption{
			User:     true,
			IsPage:   true,
			Msg:      "Please login to access this page",
			Redirect: "/login",
			Btn:      "Login",
		},
	))
	gr.GET("/user/profile", up.userPage)
}

func (sp *userPage) userPage(c *gin.Context) {
	user, _ := gogin.GetCurrentUser(c)
	posts, _ := postModel.List(singleton.DB, "created_user = ?", user.ID)
	c.HTML(http.StatusOK, "user/profile", gogin.CommonEnvironment(c, gin.H{
		"Posts": posts,
	}))
}
