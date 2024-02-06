package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	api_utils "go-gin/internal/api"
	"go-gin/service/singleton"
)

func GetCreation(c *gin.Context) {
	share_num := c.Param("share_num")
	creation_file := fmt.Sprintf("%s/creation/%s.png", singleton.Config.Upload.Dir, share_num)
	// Check if the file exists
	if _, err := os.Stat(creation_file); os.IsNotExist(err) {
		api_utils.ResponseError(c, http.StatusNotFound, "Creation not found")
		return
	}

	c.HTML(
		http.StatusOK,
		"creation/share",
		gin.H{
			"share_num": share_num,
			"config":    singleton.Config,
		},
	)
}
