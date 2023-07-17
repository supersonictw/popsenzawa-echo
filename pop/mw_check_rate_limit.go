// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package pop

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MiddlewareCheckRateLimit(c *gin.Context) {
	ipAddress := c.ClientIP()

	ctx := context.Background()
	if err := ValidateRateLimitIpAddress(ctx, ipAddress); err != nil {
		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Next()
}
