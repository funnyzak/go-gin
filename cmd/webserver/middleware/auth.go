package middleware

import (
	"net/http"

	"github.com/funnyzak/gogin/internal/config"
	"github.com/funnyzak/gogin/internal/log"
	"github.com/funnyzak/gogin/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	APIUtils "github.com/funnyzak/gogin/internal/api"
)

func AuthHanlder() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := APIUtils.GetTokenString(c)

		if err != nil {
			log.ZLog.Log.Error().Msgf("Error getting token: %v", err)
			c.JSON(http.StatusUnauthorized, &model.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized",
			})
			c.Abort()
			return
		}

		claims := &model.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Instance.JWT.Secret), nil
		})
		if err != nil {
			log.ZLog.Log.Error().Msgf("Error parsing token: %v", err)
			c.JSON(http.StatusUnauthorized, &model.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized",
			})
			c.Abort()
			return
		}
		if !token.Valid {
			log.ZLog.Log.Error().Msgf("Invalid token")
			c.JSON(http.StatusUnauthorized, &model.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "Invalid token",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
