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
	User     bool // if true, only logged user can access
	Guest    bool // if true, only guest can access
	IsPage   bool
	AllowAPI bool // if true, allow API token
	Msg      string
	Redirect string
	Btn      string
}

var auth = model.Auth{}

func Authorize(opt AuthorizeOption) gin.HandlerFunc {
	return func(c *gin.Context) {
		var code = http.StatusForbidden

		rltErr := mygin.ErrInfo{
			Title: "Unauthorized",
			Code:  code,
			Msg:   opt.Msg,
			Link:  opt.Redirect,
			Btn:   opt.Btn,
		}
		var isLogin bool = false

		token, _ := c.Cookie(singleton.Conf.Site.CookieName)
		token = strings.TrimSpace(token)
		if token != "" {
			var u model.User = model.User{}
			tokenDetail, err := auth.ExtractTokenMetadata(token, singleton.Conf)
			if err != nil {
				singleton.Log.Err(err).Msgf("ExtractTokenMetadata: %v", err)
			} else {
				err = u.GetByUsername(tokenDetail.UserName, singleton.DB)
				if err != nil {
					singleton.Log.Err(err).Msgf("GetByUsername: %v", err)
				} else {
					c.Set(model.CtxKeyAuthorizedUser, u)
					isLogin = true
				}
			}
		}

		if opt.AllowAPI {
			c.Set("isAPI", true)

			if !isLogin {
				apiToken := mygin.RetrieveTokenFromAuthorization(c.Request)
				if apiToken != "" {
					var u model.User = model.User{}
					tokenDetail, err := auth.ExtractTokenMetadata(apiToken, singleton.Conf)
					if err != nil {
						singleton.Log.Err(err).Msgf("ExtractTokenMetadata: %v", err)
					} else {
						err = u.GetByUsername(tokenDetail.UserName, singleton.DB)
						if err != nil {
							singleton.Log.Err(err).Msgf("GetByUsername: %v", err)
						} else {
							c.Set(model.CtxKeyAuthorizedUser, u)
							isLogin = true
						}
					}
				}
			}
		}

		if opt.Guest && isLogin {
			ShowErrorPage(c, rltErr, opt.IsPage)
			return
		}

		if !isLogin && opt.User {
			ShowErrorPage(c, rltErr, opt.IsPage)
			return
		}

		c.Next()
	}
}
