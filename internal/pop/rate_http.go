// PopCat Echo
// (c) 2021 SuperSonic (https://github.com/supersonictw).

package pop

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
	"github.com/supersonictw/popcat-echo/internal/config"
	"time"
)

func GetHttpLimiter() gin.HandlerFunc {
	filter := tollbooth.NewLimiter(
		float64(config.PopLimitHttpRequests),
		&limiter.ExpirableOptions{DefaultExpirationTTL: config.PopLimitHttpDuration * time.Second},
	)
	filter.SetIPLookups([]string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP"})
	filter.SetMessageContentType("application/json; charset=utf-8")
	filter.SetMessage("{\"message\":\"address rate limited\"}")
	return tollbooth_gin.LimitHandler(filter)
}
