package gogin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"go-gin/pkg/mygin"
	"go-gin/service/singleton"
)

func CommonEnvironment(c *gin.Context, data map[string]interface{}) gin.H {
	data["MatchedPath"] = c.MustGet("MatchedPath")
	data["Version"] = singleton.Version
	data["Conf"] = singleton.Conf
	if t, has := data["Title"]; !has {
		data["Title"] = singleton.Conf.Site.Brand
	} else {
		data["Title"] = fmt.Sprintf("%s - %s", t, singleton.Conf.Site.Brand)
	}
	return data
}

func ShowErrorPage(c *gin.Context, i mygin.ErrInfo, isPage bool) {
	if isPage {
		c.HTML(i.Code, "error", CommonEnvironment(c, gin.H{
			"Code":  i.Code,
			"Title": i.Title,
			"Msg":   i.Msg,
			"Link":  i.Link,
			"Btn":   i.Btn,
		}))
	} else {
		c.JSON(http.StatusOK, mygin.Response{
			Code:    i.Code,
			Message: i.Msg,
		})
	}
	c.Abort()
}
