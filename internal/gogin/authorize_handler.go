package gogin

import (
	"net/http"

	"go-gin/model"
	"go-gin/pkg/mygin"
	"go-gin/service/singleton"

	"github.com/gin-gonic/gin"
)

type AuthorizeOption struct {
	User     bool // if true, only logged user can access
	Guest    bool // if true, only guest can access
	IsPage   bool
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

		token, err := auth.ExtractTokenMetadata(c.Request, singleton.Conf)
		if err != nil {
			singleton.Log.Err(err).Msgf("Error from ExtractTokenMetadata: %v", err)
		}

		isLogin := token != nil

		if isLogin {
			var u model.User

			singleton.ApiLock.RLock()
			err = singleton.DB.Model(&model.User{}).Where("username = ?", token.UserName).First(&u).Error
			singleton.ApiLock.RUnlock()

			if err != nil {
				singleton.Log.Err(err).Msgf("User not found: %v", err)
				isLogin = false
			} else {
				c.Set(model.CtxKeyAuthorizedUser, u)

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
