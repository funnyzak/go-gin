package singleton

import (
	"fmt"
	"go-gin/pkg/mygin"

	"github.com/gin-gonic/gin"
)

func CommonEnvironment(c *gin.Context, data map[string]interface{}) gin.H {
	data["MatchedPath"] = c.MustGet("MatchedPath")
	data["Version"] = Version
	data["Conf"] = Conf
	if t, has := data["Title"]; !has {
		data["Title"] = Conf.Site.Brand
	} else {
		data["Title"] = fmt.Sprintf("%s - %s", t, Conf.Site.Brand)
	}
	token, err := mygin.RetrieveToken(c, Conf.JWT.TokenName)
	if err == nil {
		data["Token"] = token
	}
	data["IsLogin"] = token != ""
	return data
}
