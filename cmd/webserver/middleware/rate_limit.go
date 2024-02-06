package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	APIUtils "github.com/funnyzak/go-gin/internal/api"
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
			APIUtils.ResponseError(c, http.StatusTooManyRequests, "too many requests")
		}
		c.Next()
	}
}
