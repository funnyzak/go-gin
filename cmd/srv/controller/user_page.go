package controller

import (
	"go-gin/internal/gogin"
	"go-gin/model"
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
	gr.GET("/user/mine", up.mine)
	gr.GET("/user/post", up.userPost)
}

func (sp *userPage) mine(c *gin.Context) {
	user, _ := gogin.GetCurrentUser(c)
	posts, _ := model.NewPost().List(singleton.DB, "created_user = ?", user.ID)
	c.HTML(http.StatusOK, "user/mine", gogin.CommonEnvironment(c, gin.H{
		"Title": "Mine",
		"Posts": posts,
	}))
}

func (sp *userPage) userPost(c *gin.Context) {
	c.HTML(http.StatusOK, "user/post", gogin.CommonEnvironment(c, gin.H{
		"Title": "Post a new article",
	}))
}
