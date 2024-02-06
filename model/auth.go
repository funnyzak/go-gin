package model

import "github.com/golang-jwt/jwt/v5"

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (c *Credentials) IsValid() bool {
	return c.Username != "" && c.Password != ""
}
