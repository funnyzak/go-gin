package middleware

import (
	"time"

	"go-gin/internal/log"

	"github.com/gin-gonic/gin"
)

func LoggingHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		log.ZLog.Log.Info().Msgf("Request Info:\nMethod: %s\nHost: %s\nURL: %s",
			c.Request.Method, c.Request.Host, c.Request.URL)
		log.ZLog.Log.Debug().Msgf("Request Header:\n%v", c.Request.Header)

		c.Next()

		latency := time.Since(t)
		log.ZLog.Log.Info().Msgf("Response Time: %s\nStatus: %d",
			latency.String(), c.Writer.Status())
		log.ZLog.Log.Debug().Msgf("Response Header:\n%v", c.Writer.Header())
	}
}
