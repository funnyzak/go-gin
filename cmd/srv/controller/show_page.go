package controller

import (
	"go-gin/internal/gogin"
	"go-gin/model"
	"go-gin/pkg/mygin"
	"go-gin/pkg/utils"
	"go-gin/service/singleton"
	"net/http"

	"github.com/gin-gonic/gin"
)

type showPage struct {
	r *gin.Engine
}

var postModel = model.Post{}

func (sp *showPage) serve() {
	gr := sp.r.Group("")
	gr.GET("/post/:id", sp.postDetail)
	gr.GET("/post/list", sp.postList)
}

func (sp *showPage) postDetail(c *gin.Context) {
	postId := utils.ParseInt(c.Param("id"), 0)
	if postId <= 0 {
		gogin.ShowErrorPage(c, mygin.ErrInfo{
			Msg:  "Post not found",
			Code: http.StatusNotFound}, true)
		return
	}
	post, err := postModel.Get(postId, singleton.DB)
	if err != nil {
		gogin.ShowErrorPage(c, mygin.ErrInfo{
			Msg:  "Post not found",
			Code: http.StatusNotFound}, true)
		return
	}
	c.HTML(http.StatusOK, "post/detail", gogin.CommonEnvironment(c, gin.H{
		"Title": post.Title,
		"Post":  post,
	}))
}

func (sp *showPage) postList(c *gin.Context) {
	posts, _ := postModel.List(singleton.DB, "id > ?", 0)
	c.HTML(http.StatusOK, "post/list", gogin.CommonEnvironment(c, gin.H{
		"Posts": posts,
	}))
}
