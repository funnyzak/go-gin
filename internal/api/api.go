package api

import (
	"fmt"

	"github.com/funnyzak/gogin/model"
	"github.com/gin-gonic/gin"
)

func GetTokenString(c *gin.Context) (string, error) {
	tokenString, err := c.Cookie("token")
	if err == nil {
		return tokenString, nil
	}
	// Get the token from the header X-JWT-Token
	tokenString = c.GetHeader("X-JWT-Token")
	if tokenString != "" {
		return tokenString, nil
	}
	// Get the token from the query
	tokenString = c.Query("token")
	if tokenString != "" {
		return tokenString, nil
	}
	return "", fmt.Errorf("no token found")
}

func ResponseError(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, &model.ErrorResponse{
		Code:    code,
		Message: message,
	})
}

func Response(c *gin.Context, data interface{}, messages ...string) {
	var message string

	if len(messages) > 0 {
		message = messages[0]
	}
	c.AbortWithStatusJSON(200, &model.Response{
		Message: message,
		Data:    data,
	})
}
