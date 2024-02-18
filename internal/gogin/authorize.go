package gogin

import (
	"net/http"
	"strings"

	"go-gin/model"
	"go-gin/pkg/mygin"
	"go-gin/service/singleton"

	"github.com/gin-gonic/gin"
)

type AuthorizeOption struct {
	User     bool // if true, only logined user can access
	Guest    bool // if true, only guest can access
	IsPage   bool
	AllowAPI bool // if true, allow API token
	Msg      string
	Redirect string
	Btn      string
}

var auth = model.Auth{}

func getUserFromToken(token string, c *gin.Context) (model.User, bool) {
	var u model.User
	tokenDetail, err := auth.ExtractTokenMetadata(token, singleton.Conf)
	if err != nil {
		singleton.Log.Err(err).Msgf("ExtractTokenMetadata: %v", err)
		return u, false
	}

	err = u.GetByUsername(tokenDetail.UserName, singleton.DB)
	if err != nil {
		singleton.Log.Err(err).Msgf("GetByUsername: %v", err)
		return u, false
	}

	c.Set(model.CtxKeyAuthorizedUser, u)
	return u, true
}

func Authorize(opt AuthorizeOption) gin.HandlerFunc {
	return func(c *gin.Context) {
		var code = http.StatusForbidden

		unauthorizedErr := mygin.ErrInfo{
			Title: "Unauthorized",
			Code:  code,
			Msg:   opt.Msg,
			Link:  opt.Redirect,
			Btn:   opt.Btn,
		}
		var isLogin bool = false

		token, _ := c.Cookie(singleton.Conf.JWT.AccessTokenCookieName)
		token = strings.TrimSpace(token)
		if token != "" {
			_, isLogin = getUserFromToken(token, c)
		}

		if !isLogin {
			refreshToken, err := c.Cookie(singleton.Conf.JWT.RefreshTokenCookieName)
			if refreshToken != "" && err == nil {
				newToken, err := auth.RefreshToken(refreshToken, singleton.Conf)
				if err != nil {
					singleton.Log.Err(err).Msgf("RefreshToken: %v", err)
					ShowErrorPage(c, unauthorizedErr, opt.IsPage)
					return
				}

				UserLoginSuccess(c, newToken)
				token = newToken.AccessToken
				_, isLogin = getUserFromToken(token, c)
			}
		}

		if opt.AllowAPI && !isLogin {
			c.Set("isAPI", true)
			apiToken := mygin.RetrieveTokenFromAuthorization(c.Request)
			if apiToken != "" {
				_, isLogin = getUserFromToken(apiToken, c)
			}
		}

		if opt.Guest && isLogin {
			ShowErrorPage(c, unauthorizedErr, opt.IsPage)
			return
		}

		if !isLogin && opt.User {
			ShowErrorPage(c, unauthorizedErr, opt.IsPage)
			return
		}

		c.Next()
	}
}
