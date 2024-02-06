package middleware

import (
	"net/http"

	"go-gin/internal/config"
	"go-gin/internal/log"
	"go-gin/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	APIUtils "go-gin/internal/api"
)

func AuthHanlder() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := APIUtils.GetTokenString(c)

		if err != nil {
			log.ZLog.Log.Error().Msgf("Error getting token: %v", err)
			APIUtils.ResponseError(c, http.StatusUnauthorized, "Unauthorized")
			return
		}
		claims := &model.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Instance.JWT.Secret), nil
		})
		if err != nil {
			log.ZLog.Log.Error().Msgf("Error parsing token: %v", err)
			APIUtils.ResponseError(c, http.StatusUnauthorized, "Unauthorized")
			return
		}
		if !token.Valid {
			log.ZLog.Log.Error().Msgf("Invalid token")
			APIUtils.ResponseError(c, http.StatusUnauthorized, "Invalid token")
			return
		}

		c.Next()
	}
}
