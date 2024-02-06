package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/funnyzak/go-gin/internal/config"
	"github.com/gin-gonic/gin"

	APIUtils "github.com/funnyzak/go-gin/internal/api"
)

func GetCreation(c *gin.Context) {
	share_num := c.Param("share_num")
	creation_file := fmt.Sprintf("%s/creation/%s.png", config.Instance.Upload.Dir, share_num)
	// Check if the file exists
	if _, err := os.Stat(creation_file); os.IsNotExist(err) {
		APIUtils.ResponseError(c, http.StatusNotFound, "Creation not found")
		return
	}

	c.HTML(
		http.StatusOK,
		"creation/share",
		gin.H{
			"share_num": share_num,
			"config":    config.Instance,
		},
	)
}
