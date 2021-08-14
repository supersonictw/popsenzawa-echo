package config

import (
	"crypto/rand"
	"github.com/dpapathanasiou/go-recaptcha"
	_ "github.com/joho/godotenv/autoload"
	"github.com/supersonictw/popcat-echo/internal"
	"strconv"
	"time"
)

var (
	PublishAddress    string
	RefreshInterval   int64
	RefreshDelay      int64
	CacheNamespacePop string
	CacheNamespaceGeo string
	ReCaptchaStatus   bool
	JWTCaptchaSecret  []byte
	JWTExpired        time.Duration
	PopLimit          int
)

func init() {
	PublishAddress = Get(internal.ConfigPublishAddress)

	refreshIntervalInt, err := strconv.Atoi(Get(internal.ConfigRefreshInterval))
	if err != nil {
		panic(err)
	}
	RefreshInterval = int64(refreshIntervalInt)
	refreshDelayInt, err := strconv.Atoi(Get(internal.ConfigRefreshDelay))
	if err != nil {
		panic(err)
	}
	RefreshDelay = int64(refreshDelayInt)
	CacheNamespacePop = Get(internal.ConfigRedisNamespacePop)
	CacheNamespaceGeo = Get(internal.ConfigRedisNamespaceGeo)

	if secret := Get(internal.ConfigReCaptchaSecret); secret != "" {
		recaptcha.Init(secret)
		ReCaptchaStatus = true
	} else {
		ReCaptchaStatus = false
	}

	if secret := Get(internal.ConfigJWTSecret); secret != "" {
		JWTCaptchaSecret = []byte(secret)
	} else {
		blk := make([]byte, 32)
		_, err = rand.Read(blk)
		JWTCaptchaSecret = blk
	}

	jwtExpired, err := strconv.Atoi(Get(internal.ConfigJWTExpired))
	if err != nil {
		panic(err)
	}
	JWTExpired = time.Duration(jwtExpired)

	PopLimit, err = strconv.Atoi(Get(internal.ConfigRefreshDelay))
	if err != nil {
		panic(err)
	}
}
