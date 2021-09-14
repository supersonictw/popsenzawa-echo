// PopCat Echo
// (c) 2021 SuperSonic (https://github.com/supersonictw).

package config

import (
	"crypto/rand"
	"github.com/dpapathanasiou/go-recaptcha"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"strconv"
	"time"
)

var (
	PublishAddress     string
	RefreshInterval    int64
	RefreshDelay       int64
	CacheNamespacePop  string
	CacheNamespaceGeo  string
	CacheNamespaceRate string
	ReCaptchaStatus    bool
	JWTCaptchaSecret   []byte
	JWTExpired         time.Duration
	PopLimit           int
	RateLimit          int
	ForceFixRate       bool
)

func init() {
	PublishAddress = Get(EnvPublishAddress)

	refreshIntervalInt, err := strconv.Atoi(Get(EnvRefreshInterval))
	if err != nil {
		log.Panicln(err)
	}
	RefreshInterval = int64(refreshIntervalInt)
	refreshDelayInt, err := strconv.Atoi(Get(EnvRefreshDelay))
	if err != nil {
		log.Panicln(err)
	}
	RefreshDelay = int64(refreshDelayInt)

	CacheNamespacePop = Get(EnvRedisNamespacePop)
	CacheNamespaceGeo = Get(EnvRedisNamespaceGeo)
	CacheNamespaceRate = Get(EnvRedisNamespaceRate)

	if secret := Get(EnvReCaptchaSecret); secret != "" {
		recaptcha.Init(secret)
		ReCaptchaStatus = true
	} else {
		ReCaptchaStatus = false
	}

	if secret := Get(EnvJWTSecret); secret != "" {
		JWTCaptchaSecret = []byte(secret)
	} else {
		blk := make([]byte, 32)
		if _, err = rand.Read(blk); err != nil {
			log.Panicln(err)
		}
		JWTCaptchaSecret = blk
	}

	jwtExpired, err := strconv.Atoi(Get(EnvJWTExpired))
	if err != nil {
		log.Panicln(err)
	}
	JWTExpired = time.Duration(jwtExpired)

	PopLimit, err = strconv.Atoi(Get(EnvPopLimit))
	if err != nil {
		log.Panicln(err)
	}
	RateLimit, err = strconv.Atoi(Get(EnvRateLimit))
	if err != nil {
		log.Panicln(err)
	}

	if Get(EnvForceFixRate) == "yes" {
		ForceFixRate = true
	} else {
		ForceFixRate = false
	}
}
