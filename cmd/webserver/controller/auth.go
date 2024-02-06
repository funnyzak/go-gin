package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"go-gin/model"
	"go-gin/service/singleton"

	api_utils "go-gin/internal/api"
)

// Login authenticates the user
func Login(c *gin.Context) {
	var creds model.Credentials
	// get the body of the POST request
	err := c.BindJSON(&creds)
	if err != nil {
		singleton.Log.Error().Msgf("Error binding JSON: %v", err)
		api_utils.ResponseError(c, http.StatusBadRequest, "Invalid JSON")
		return
	}

	users := singleton.Config.Users
	expectedPassword, ok := users[creds.Username]

	if !ok || expectedPassword != creds.Password {
		singleton.Log.Error().Msgf("Invalid credentials: %v", creds)
		api_utils.ResponseError(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	expirationTime := time.Now().Add(time.Duration(singleton.Config.JWT.Expiration) * time.Minute)
	claims := &model.Claims{
		Username: creds.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(singleton.Config.JWT.Secret))
	if err != nil {
		singleton.Log.Error().Msgf("Error signing token: %v", err)
		api_utils.ResponseError(c, http.StatusInternalServerError, "Error signing token")
		return
	}
	c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "", false, true)
	api_utils.Response(c, gin.H{"token": tokenString})
}

// Refresh refreshes the token
func Refresh(c *gin.Context) {
	tokenString, err := api_utils.GetTokenString(c)

	if err != nil {
		singleton.Log.Error().Msgf("Error getting token: %v", err)
		c.JSON(http.StatusUnauthorized, &model.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		})
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
		c.JSON(http.StatusUnauthorized, &model.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Invalid token",
		})
		return
	}

	expirationTime := time.Now().Add(time.Duration(singleton.Config.JWT.Expiration) * time.Minute)
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(singleton.Config.JWT.Secret))
	if err != nil {
		singleton.Log.Error().Msgf("Error signing token: %v", err)
		api_utils.ResponseError(c, http.StatusInternalServerError, "Error signing token")
		return
	}
	c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "", false, true)
	api_utils.Response(c, gin.H{"token": tokenString})
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	api_utils.Response(c, "", "Logged out")
}
