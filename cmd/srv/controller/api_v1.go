package controller

import (
	"fmt"
	"go-gin/internal/gogin"
	"go-gin/mappers"
	"go-gin/model"
	"go-gin/pkg/mygin"
	"go-gin/pkg/utils"
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

	r.POST("/post", v.postPost)         // create post
	r.GET("/post/:id", v.getPost)       // get post
	r.DELETE("/post/:id", v.deletePost) // delete post
	r.GET("/posts", v.getPosts)         // get posts

	user := v.r.Group("user")
	{
		user.GET("/info", v.getUserInfo)
		user.GET("/logout", v.logout)
		user.GET("/refresh", v.refresh)
	}
}

var authModel = model.Auth{}

func (v *apiV1) logout(c *gin.Context) {
	isPage := utils.ParseBool(c.Query("page"), false)
	c.SetCookie(singleton.Conf.Site.CookieName, "", -1, "/", "", false, true)
	if isPage {
		gogin.ShowMessagePage(c, "Logout success", singleton.Conf.Site.BaseURL, "Back to home")
	} else {
		mygin.ResponseJSON(c, 200, gin.H{}, "logout success")
	}
}

func (v *apiV1) refresh(c *gin.Context) {
	var tokenForm mappers.Token
	if err := mygin.BindForm(c, utils.ParseBool("form", false), &tokenForm); err != nil {
		mygin.ResponseJSON(c, 400, gin.H{}, "refresh token is required")
		return
	}
	tk, err := authModel.RefreshToken(tokenForm.RefreshToken, singleton.Conf)
	if err != nil {
		mygin.ResponseJSON(c, 400, gin.H{}, err.Error())
		return
	}
	mygin.ResponseJSON(c, 200, tk, "refresh success")
}

func (v *apiV1) postPost(c *gin.Context) {
	var postForm mappers.PostForm
	isForm := utils.ParseBool(c.Query("form"), false)
	if err := mygin.BindForm(c, isForm, &postForm); err != nil {
		gogin.ShowErrorPage(c, mygin.ErrInfo{
			Code: 400,
			Msg:  "post is required",
			Btn:  "Back",
			Link: "/",
		}, isForm)
		return
	}
	if postForm.CreatedUser == 0 {
		user, _ := gogin.GetCurrentUser(c)
		postForm.CreatedUser = user.ID
	}
	var post model.Post = model.Post{}
	if err := post.Create(postForm, singleton.DB); err != nil {
		gogin.ShowErrorPage(c, mygin.ErrInfo{
			Code: 400,
			Msg:  err.Error(),
			Btn:  "Back",
			Link: "/",
		}, isForm)
		return
	}

	if isForm {
		gogin.ShowMessagePage(c, "Post success", fmt.Sprintf("/post/%d", post.ID), "View post")
	} else {
		mygin.ResponseJSON(c, 200,
			gogin.CommonEnvironment(c, gin.H{
				"post": post,
			}), "post success")
	}
}

func (v *apiV1) getPost(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "post",
	})
}

func (v *apiV1) deletePost(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "post",
	})
}

func (v *apiV1) getPosts(c *gin.Context) {
	posts, _ := postModel.List(singleton.DB, "")
	mygin.ResponseJSON(c, 200, gogin.CommonEnvironment(c, gin.H{
		"Posts": posts,
	}))
}

func (v *apiV1) getUserInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "user info",
	})
}
