package middleware

import (
	"net/http"

	"go-gin/model"
	"go-gin/service/singleton"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	api_utils "go-gin/internal/api"
)

func AuthHanlder() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := api_utils.GetTokenString(c)

		if err != nil {
			singleton.Log.Error().Msgf("Error getting token: %v", err)
			api_utils.ResponseError(c, http.StatusUnauthorized, "Unauthorized")
			return
		}
		claims := &model.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(singleton.Config.JWT.Secret), nil
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
