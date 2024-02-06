package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/funnyzak/go-gin/internal/config"
	"github.com/funnyzak/go-gin/pkg/utils"
	"github.com/gin-gonic/gin"

	APIUtils "github.com/funnyzak/go-gin/internal/api"
)

func UploadCreation(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		APIUtils.ResponseError(c, http.StatusBadRequest, "File not found")
		return
	}
	if !strings.Contains(file.Header.Get("Content-Type"), "image") {
		APIUtils.ResponseError(c, http.StatusBadRequest, "File is not an image")
		return
	}
	new_id := utils.GenHexStr(32)
	c.SaveUploadedFile(file, fmt.Sprintf("%s/creation/%s.png", config.Instance.Upload.Dir, new_id))

	APIUtils.Response(c, gin.H{
		"creation_id": new_id,
		"share_url":   fmt.Sprintf("%s/share/creation/%s", config.Instance.Server.BaseUrl, new_id),
	})
}
