// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package pop

import (
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/spf13/viper"
)

var (
	configMaxRequestPerSecond  = viper.GetFloat64("rate_limit.max_request_per_second")
	configDefaultExpirationTTL = viper.GetFloat64("rate_limit.default_expiration_ttl")
)

func getRateLimitFilter() *limiter.Limiter {
	filter := tollbooth.NewLimiter(configMaxRequestPerSecond, &limiter.ExpirableOptions{
		DefaultExpirationTTL: time.Duration(configDefaultExpirationTTL) * time.Second,
	})

	filter.SetIPLookups([]string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP"})
	filter.SetMessageContentType("application/json; charset=utf-8")
	filter.SetMessage("{\"message\":\"address rate limited\"}")

	return filter
}
