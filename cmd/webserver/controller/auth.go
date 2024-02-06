package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/funnyzak/gogin/internal/config"
	"github.com/funnyzak/gogin/internal/log"
	"github.com/funnyzak/gogin/model"

	APIUtils "github.com/funnyzak/gogin/internal/api"
)

// Login authenticates the user
func Login(c *gin.Context) {
	var creds model.Credentials
	// get the body of the POST request
	err := c.BindJSON(&creds)
	if err != nil {
		log.ZLog.Log.Error().Msgf("Error binding JSON: %v", err)

		APIUtils.ResponseError(c, http.StatusBadRequest, "Invalid JSON")
		return
	}

	users := config.Instance.Users
	expectedPassword, ok := users[creds.Username]

	if !ok || expectedPassword != creds.Password {
		log.ZLog.Log.Error().Msgf("Invalid credentials: %v", creds)
		APIUtils.ResponseError(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	expirationTime := time.Now().Add(time.Duration(config.Instance.JWT.Expiration) * time.Minute)
	claims := &model.Claims{
		Username: creds.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Instance.JWT.Secret))
	if err != nil {
		log.ZLog.Log.Error().Msgf("Error signing token: %v", err)
		APIUtils.ResponseError(c, http.StatusInternalServerError, "Error signing token")
		return
	}
	c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "", false, true)
	APIUtils.Response(c, gin.H{"token": tokenString})
}

// Refresh refreshes the token
func Refresh(c *gin.Context) {
	tokenString, err := APIUtils.GetTokenString(c)

	if err != nil {
		log.ZLog.Log.Error().Msgf("Error getting token: %v", err)
		c.JSON(http.StatusUnauthorized, &model.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		})
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
		c.JSON(http.StatusUnauthorized, &model.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Invalid token",
		})
		return
	}

	expirationTime := time.Now().Add(time.Duration(config.Instance.JWT.Expiration) * time.Minute)
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(config.Instance.JWT.Secret))
	if err != nil {
		log.ZLog.Log.Error().Msgf("Error signing token: %v", err)
		APIUtils.ResponseError(c, http.StatusInternalServerError, "Error signing token")
		return
	}
	c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "", false, true)
	APIUtils.Response(c, gin.H{"token": tokenString})
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	APIUtils.Response(c, "", "Logged out")
}
