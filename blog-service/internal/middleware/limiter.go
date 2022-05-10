package middleware

import (
	"GoprogrammingTourBook/blog-service/pkg/app"
	"GoprogrammingTourBook/blog-service/pkg/errcode"
	"GoprogrammingTourBook/blog-service/pkg/limiter"

	"github.com/gin-gonic/gin"
)

func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooMantRequests)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}