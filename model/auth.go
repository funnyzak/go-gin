package model

import (
	"fmt"
	"go-gin/internal/gconfig"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/twinj/uuid"
)

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}

type AccessDetails struct {
	AccessUUID string
	UserName   string
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Auth struct{}

// CreateToken by username
func (m Auth) CreateToken(username string, conf *gconfig.Config) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * time.Duration(conf.JWT.TokenExpiration)).Unix()
	td.AccessUUID = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Minute * time.Duration(conf.JWT.RefreshTokenExpiration)).Unix()
	td.RefreshUUID = uuid.NewV4().String()

	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUUID
	atClaims["user_id"] = username
	atClaims["exp"] = td.AtExpires

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(conf.JWT.AccessSecret))
	if err != nil {
		return nil, err
	}
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUUID
	rtClaims["user_id"] = username
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(conf.JWT.RefreshSecret))
	if err != nil {
		return nil, err
	}
	return td, nil
}

// Extract the token from the request
func (m Auth) ExtractTokenFromAuthorization(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// Verify the token from the request
func (m Auth) VerifyToken(r *http.Request, conf *gconfig.Config) (*jwt.Token, error) {
	tokenString := m.ExtractTokenFromAuthorization(r)
	if tokenString == "" {
		return nil, fmt.Errorf("Can't find token from request")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(conf.JWT.AccessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// TokenValid from the request
func (m Auth) TokenValid(r *http.Request, conf *gconfig.Config) error {
	token, err := m.VerifyToken(r, conf)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

// Get the token metadata from the request
func (m Auth) ExtractTokenMetadata(r *http.Request, conf *gconfig.Config) (*AccessDetails, error) {
	token, err := m.VerifyToken(r, conf)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userName := claims["user_id"].(string)
		return &AccessDetails{
			AccessUUID: accessUUID,
			UserName:   userName,
		}, nil
	}
	return nil, err
}
