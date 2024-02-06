package middleware

import (
	"net/http"

	"go-gin/internal/config"
	"go-gin/internal/log"
	"go-gin/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	api_utils "go-gin/internal/api"
)

func AuthHanlder() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := api_utils.GetTokenString(c)

		if err != nil {
			log.ZLog.Error().Msgf("Error getting token: %v", err)
			api_utils.ResponseError(c, http.StatusUnauthorized, "Unauthorized")
			return
		}
		claims := &model.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Instance.JWT.Secret), nil
		})
		if err != nil {
			log.ZLog.Error().Msgf("Error parsing token: %v", err)
			api_utils.ResponseError(c, http.StatusUnauthorized, "Unauthorized")
			return
		}
		if !token.Valid {
			log.ZLog.Error().Msgf("Invalid token")
			api_utils.ResponseError(c, http.StatusUnauthorized, "Invalid token")
			return
		}

		c.Next()
	}
}
