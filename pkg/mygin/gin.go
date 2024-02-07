package mygin

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

// RetrieveToken retrieves token from cookie, header, query, or post form
func RetrieveToken(c *gin.Context, tokenName string) (string, error) {
	tokenString, err := c.Cookie(tokenName)
	if err == nil {
		return tokenString, nil
	}
	tokenString = c.GetHeader(tokenName)
	if tokenString != "" {
		return tokenString, nil
	}
	tokenString = c.Query(tokenName)
	if tokenString != "" {
		return tokenString, nil
	}
	tokenString = c.PostForm(tokenName)
	if tokenString != "" {
		return tokenString, nil
	}
	// get token from json body
	var jsonBody map[string]interface{}
	if err := c.BindJSON(&jsonBody); err == nil {
		if token, ok := jsonBody[tokenName]; ok {
			if tokenString, ok := token.(string); ok {
				return tokenString, nil
			}
		}
	}
	return "", fmt.Errorf("token not found")
}

// MatchedPath returns the matched path of the request
func RecordPath(c *gin.Context) {
	url := c.Request.URL.String()
	for _, p := range c.Params {
		url = strings.Replace(url, p.Value, ":"+p.Key, 1)
	}
	c.Set("MatchedPath", url)
}
