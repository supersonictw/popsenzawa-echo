package config

type EnvKey string

const (
	EnvPublishAddress     EnvKey = "PUBLISH_ADDRESS"
	EnvRefreshInterval    EnvKey = "REFRESH_INTERVAL"
	EnvRefreshDelay       EnvKey = "REFRESH_DELAY"
	EnvRedisAddress       EnvKey = "REDIS_ADDRESS"
	EnvRedisPassword      EnvKey = "REDIS_PASSWORD"
	EnvRedisDatabase      EnvKey = "REDIS_DBNAME"
	EnvRedisNamespacePop  EnvKey = "REDIS_NAMESPACE_POP"
	EnvRedisNamespaceGeo  EnvKey = "REDIS_NAMESPACE_GEO"
	EnvRedisNamespaceRate EnvKey = "REDIS_NAMESPACE_RATE"
	EnvMysqlDSN           EnvKey = "MYSQL_DSN"
	EnvReCaptchaSecret    EnvKey = "RECAPTCHA_SECRET"
	EnvJWTSecret          EnvKey = "JWT_SECRET"
	EnvJWTExpired         EnvKey = "JWT_EXPIRED"
	EnvPopLimit           EnvKey = "POP_LIMIT"
	EnvRateLimit          EnvKey = "RATE_LIMIT"
	EnvForceFixRate       EnvKey = "FORCE_FIX_RATE"
)
