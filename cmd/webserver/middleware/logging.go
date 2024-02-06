package middleware

import (
	"go-gin/service/singleton"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		singleton.Log.Info().Msgf("Request Info:\nMethod: %s\nHost: %s\nURL: %s",
			c.Request.Method, c.Request.Host, c.Request.URL)
		singleton.Log.Debug().Msgf("Request Header:\n%v", c.Request.Header)

		c.Next()

		latency := time.Since(t)
		singleton.Log.Info().Msgf("Response Time: %s\nStatus: %d",
			latency.String(), c.Writer.Status())
		singleton.Log.Debug().Msgf("Response Header:\n%v", c.Writer.Header())
	}
}
