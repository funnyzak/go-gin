package gogin

import (
	"go-gin/pkg/mygin"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimiterHandler(reqsPerMin int) gin.HandlerFunc {
	var limiter *rate.Limiter
	if reqsPerMin > 0 {
		limiter = rate.NewLimiter(rate.Every(time.Minute), reqsPerMin)
	} else {
		limiter = rate.NewLimiter(rate.Inf, 0)
	}

	return func(c *gin.Context) {
		if !limiter.Allow() {
			ShowErrorPage(c, mygin.ErrInfo{
				Code: http.StatusTooManyRequests,
				Msg:  "too many requests",
			}, false)
		}
		c.Next()
	}
}
