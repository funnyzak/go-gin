package mygin

import (
	"github.com/gin-gonic/gin"
	"github.com/twinj/uuid"
)

func CORSHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		}
		c.Next()
	}
}

func NoCacheHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
		c.Writer.Header().Set("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
		c.Writer.Header().Set("Last-Modified", "Thu, 01 Jan 1970 00:00:00 GMT")
		c.Writer.Header().Set("Pragma", "no-cache")
		c.Next()
	}
}

func SecureJSONHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}

func GenerateContextIdHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		contextId := uuid.NewV4()
		c.Writer.Header().Set("X-Context-Id", contextId.String())
		c.Next()
	}
}
