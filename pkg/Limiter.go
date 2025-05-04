package pkg

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
)

var limiter = rate.NewLimiter(1, 5)

func RateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !limiter.Allow() {
			LogAction(c, "TOO_MANY_REQUESTS", 429)
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too Many Requests",
			})
			return
		}
		c.Next()
	}
}
