package middleware

import (
	"net/http"

	api_utils "go-gin/internal/api"
	"go-gin/model"
	"go-gin/pkg/mygin"
	"go-gin/service/singleton"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthorizeOption struct {
	Guest    bool // If true, allow guest
	User     bool
	IsPage   bool
	AllowAPI bool
	Msg      string
	Redirect string
	Btn      string
}

func Authorize(opt AuthorizeOption) gin.HandlerFunc {
	return func(c *gin.Context) {
		var code = http.StatusForbidden
		if opt.Guest {
			code = http.StatusBadRequest
		}

		rltErr := mygin.ErrInfo{
			Title: "Unauthorized",
			Code:  code,
			Msg:   opt.Msg,
			Link:  opt.Redirect,
			Btn:   opt.Btn,
		}

		var isLogin = false
		token, _ := c.Cookie( singleton.Conf.Site.CookieName)
		if token != "" {
			claims := &model.Claims{}
			_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(singleton.Conf.JWT.Secret), nil
			})
			if err == nil {
				isLogin = true
				c.Set(model.CtxKeyAuthorizedUser, claims.Username)
			}
		}

		if opt.AllowAPI {

		}

		if err != nil {
			singleton.Log.Error().Msgf("Error getting token: %v", err)
			api_utils.ResponseError(c, http.StatusUnauthorized, "Unauthorized")
			return
		}
		claims := &model.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(singleton.Conf.JWT.Secret), nil
		})
		if err != nil {
			singleton.Log.Error().Msgf("Error parsing token: %v", err)
			api_utils.ResponseError(c, http.StatusUnauthorized, "Unauthorized")
			return
		}
		if !token.Valid {
			singleton.Log.Error().Msgf("Invalid token")
			api_utils.ResponseError(c, http.StatusUnauthorized, "Invalid token")
			return
		}

		c.Next()
	}
}
