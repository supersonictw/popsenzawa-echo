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
	RedisAddress       string
	RedisPassword      string
	RedisDatabase      int
	CacheNamespacePop  string
	CacheNamespaceGeo  string
	CacheNamespaceRate string
	MysqlDSN           string
	ReCaptchaStatus    bool
	JWTCaptchaSecret   []byte
	JWTExpired         time.Duration
	PopLimit           int
	RateLimit          int
	ForceFixRate       bool
)

func init() {
	loadGeneral()
	loadRedis()
	loadMySQL()
	loadRecaptcha()
	loadJWT()
	loadLimit()
	loadFixRate()
}

func loadGeneral() {
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
}

func loadRedis() {
	var err error

	RedisAddress = Get(EnvRedisAddress)
	RedisPassword = Get(EnvRedisPassword)
	RedisDatabase, err = strconv.Atoi(Get(EnvRedisDatabase))
	if err != nil {
		log.Fatal(err)
	}

	CacheNamespacePop = Get(EnvRedisNamespacePop)
	CacheNamespaceGeo = Get(EnvRedisNamespaceGeo)
	CacheNamespaceRate = Get(EnvRedisNamespaceRate)
}

func loadMySQL() {
	MysqlDSN = Get(EnvMysqlDSN)
}

func loadRecaptcha() {
	if secret := Get(EnvReCaptchaSecret); secret != "" {
		recaptcha.Init(secret)
		ReCaptchaStatus = true
	} else {
		ReCaptchaStatus = false
	}
}

func loadJWT() {
	var err error

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
}

func loadLimit() {
	var err error

	PopLimit, err = strconv.Atoi(Get(EnvPopLimit))
	if err != nil {
		log.Panicln(err)
	}
	RateLimit, err = strconv.Atoi(Get(EnvRateLimit))
	if err != nil {
		log.Panicln(err)
	}
}

func loadFixRate() {
	if Get(EnvForceFixRate) == "yes" {
		ForceFixRate = true
	} else {
		ForceFixRate = false
	}
}
