package gogin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"go-gin/model"
	"go-gin/pkg/mygin"
	"go-gin/service/singleton"
)

func CommonEnvironment(c *gin.Context, data map[string]interface{}) gin.H {
	data["MatchedPath"] = c.MustGet("MatchedPath")
	data["Version"] = singleton.Version
	data["Conf"] = singleton.Conf
	if val, ok := c.Get(model.CtxKeyAuthorizedUser); ok {
		data["User"] = val
	}
	if t, has := data["Title"]; !has {
		data["Title"] = singleton.Conf.Site.Brand
	} else {
		data["Title"] = fmt.Sprintf("%s - %s", t, singleton.Conf.Site.Brand)
	}
	return data
}

func GetCurrentUser(c *gin.Context) (user model.User, err error) {
	val, ok := c.Get(model.CtxKeyAuthorizedUser)
	if !ok {
		return user, fmt.Errorf("user not found")
	}
	user = val.(model.User)
	return user, nil
}

func GetCurrentUserId(c *gin.Context) uint64 {
	user, err := GetCurrentUser(c)
	if err != nil {
		return 0
	}
	return user.ID
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

func ShowMessagePage(c *gin.Context, msg, link, btn string) {
	c.HTML(http.StatusOK, "message", CommonEnvironment(c, gin.H{
		"Msg":  msg,
		"Link": link,
		"Btn":  btn,
	}))
	c.Abort()
}

func UserLoginSuccess(c *gin.Context, token *model.Token) {
	c.SetCookie(singleton.Conf.JWT.AccessTokenCookieName, token.AccessToken, singleton.Conf.JWT.AccessTokenExpiration*60, "/", "", false, true)
	c.SetCookie(singleton.Conf.JWT.RefreshTokenCookieName, token.RefreshToken, singleton.Conf.JWT.RefreshTokenExpiration*60, "/", "", false, true)
}

func UserLogout(c *gin.Context) {
	c.SetCookie(singleton.Conf.JWT.AccessTokenCookieName, "", -1, "/", "", false, true)
	c.SetCookie(singleton.Conf.JWT.RefreshTokenCookieName, "", -1, "/", "", false, true)
}
