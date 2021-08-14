package internal

const (
	ConfigPublishAddress    = "PUBLISH_ADDRESS"
	ConfigRefreshInterval   = "REFRESH_INTERVAL"
	ConfigRefreshDelay      = "REFRESH_DELAY"
	ConfigRedisAddress      = "REDIS_ADDRESS"
	ConfigRedisPassword     = "REDIS_PASSWORD"
	ConfigRedisDatabase     = "REDIS_DBNAME"
	ConfigRedisNamespacePop = "REDIS_NAMESPACE_POP"
	ConfigRedisNamespaceGeo = "REDIS_NAMESPACE_GEO"
	ConfigMysqlDSN          = "MYSQL_DSN"
	ConfigReCaptchaSecret   = "RECAPTCHA_SECRET"
	ConfigJWTSecret         = "JWT_SECRET"
	ConfigJWTExpired        = "JWT_EXPIRED"
	ErrorInvalidCount       = "invalid count"
	ErrorInvalidCountRange  = "invalid count range"
	ErrorEmptyCaptchaToken  = "empty captcha token"
	ErrorUnsafeCaptchaToken = "unsafe captcha token"
	ErrorUnknownRegionCode  = "unknown region code"
	ErrorUnknownJWTError    = "unknown jwt error"
)
