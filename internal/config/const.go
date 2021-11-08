// PopCat Echo
// (c) 2021 SuperSonic (https://github.com/supersonictw).

package config

type EnvKey string

const (
	EnvCORSSupport           EnvKey = "CORS_SUPPORT"
	EnvFrontendHostname      EnvKey = "FRONTEND_HOSTNAME"
	EnvFrontendSSL           EnvKey = "FRONTEND_SSL"
	EnvPublishAddress        EnvKey = "PUBLISH_ADDRESS"
	EnvRefreshInterval       EnvKey = "REFRESH_INTERVAL"
	EnvRefreshDelay          EnvKey = "REFRESH_DELAY"
	EnvRedisAddress          EnvKey = "REDIS_ADDRESS"
	EnvRedisPassword         EnvKey = "REDIS_PASSWORD"
	EnvRedisDatabase         EnvKey = "REDIS_DBNAME"
	EnvRedisNamespacePop     EnvKey = "REDIS_NAMESPACE_POP"
	EnvRedisNamespaceGeo     EnvKey = "REDIS_NAMESPACE_GEO"
	EnvRedisNamespaceRate    EnvKey = "REDIS_NAMESPACE_RATE"
	EnvMysqlDSN              EnvKey = "MYSQL_DSN"
	EnvPopJWTSecret          EnvKey = "POP_JWT_SECRET"
	EnvPopJWTExpired         EnvKey = "POP_JWT_EXPIRED"
	EnvPopReCaptchaSecret    EnvKey = "POP_RECAPTCHA_SECRET"
	EnvPopLimitHttpDuration  EnvKey = "POP_LIMIT_HTTP_DURATION"
	EnvPopLimitHttpRequests  EnvKey = "POP_LIMIT_HTTP_REQUESTS"
	EnvPopLimitRedisDuration EnvKey = "POP_LIMIT_REDIS_DURATION"
	EnvPopLimitRedisPopCount EnvKey = "POP_LIMIT_REDIS_POP_COUNT"
	EnvForceFixRate          EnvKey = "FORCE_FIX_RATE"
)
