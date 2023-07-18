// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package pop

import (
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
)

var (
	MiddlewareCheckRateLimit gin.HandlerFunc
)

func init() {
	rateLimitFilter := getRateLimitFilter()
	MiddlewareCheckRateLimit = tollbooth_gin.LimitHandler(rateLimitFilter)
}
