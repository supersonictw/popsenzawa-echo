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
	CORSSupport           bool
	FrontendHostname      string
	FrontendSSL           bool
	PublishAddress        string
	RefreshInterval       int64
	RefreshDelay          int64
	RedisAddress          string
	RedisPassword         string
	RedisDatabase         int
	CacheNamespacePop     string
	CacheNamespaceGeo     string
	CacheNamespaceRate    string
	MysqlDSN              string
	PopReCaptchaStatus    bool
	PopJWTSecret          []byte
	PopJWTExpired         time.Duration
	PopLimitHttpDuration  time.Duration
	PopLimitHttpRequests  int
	PopLimitRedisDuration time.Duration
	PopLimitRedisPopCount int
	ForceFixRate          bool
)

func init() {
	loadGeneral()
	loadRedis()
	loadMySQL()
	loadJWT()
	loadRecaptcha()
	loadLimit()
	loadFixRate()
}

func loadGeneral() {
	// CORSSupport
	CORSSupport = get(EnvCORSSupport) == "yes"
	// FrontendHostname
	FrontendHostname = get(EnvFrontendHostname)
	// FrontendSSL
	FrontendSSL = get(EnvFrontendSSL) == "yes"
	// PublishAddress
	PublishAddress = get(EnvPublishAddress)
	// RefreshInterval
	refreshIntervalInt, err := strconv.Atoi(get(EnvRefreshInterval))
	if err != nil {
		log.Panicln(err)
	}
	RefreshInterval = int64(refreshIntervalInt)
	// RefreshDelay
	refreshDelayInt, err := strconv.Atoi(get(EnvRefreshDelay))
	if err != nil {
		log.Panicln(err)
	}
	RefreshDelay = int64(refreshDelayInt)
}

func loadRedis() {
	var err error
	// RedisAddress
	RedisAddress = get(EnvRedisAddress)
	// RedisPassword
	RedisPassword = get(EnvRedisPassword)
	// RedisDatabase
	RedisDatabase, err = strconv.Atoi(get(EnvRedisDatabase))
	if err != nil {
		log.Fatal(err)
	}
	// CacheNamespacePop
	CacheNamespacePop = get(EnvRedisNamespacePop)
	// CacheNamespaceGeo
	CacheNamespaceGeo = get(EnvRedisNamespaceGeo)
	// CacheNamespaceRate
	CacheNamespaceRate = get(EnvRedisNamespaceRate)
}

func loadMySQL() {
	// MysqlDSN
	MysqlDSN = get(EnvMysqlDSN)
}

func loadJWT() {
	var err error
	// PopJWTSecret
	if secret := get(EnvPopJWTSecret); secret != "" {
		PopJWTSecret = []byte(secret)
	} else {
		blk := make([]byte, 32)
		if _, err = rand.Read(blk); err != nil {
			log.Panicln(err)
		}
		PopJWTSecret = blk
	}
	// PopJWTExpired
	jwtExpired, err := strconv.Atoi(get(EnvPopJWTExpired))
	if err != nil {
		log.Panicln(err)
	}
	PopJWTExpired = time.Duration(jwtExpired)
}

func loadRecaptcha() {
	// PopReCaptchaStatus
	if secret := get(EnvPopReCaptchaSecret); secret != "" {
		recaptcha.Init(secret)
		PopReCaptchaStatus = true
	} else {
		PopReCaptchaStatus = false
	}
}

func loadLimit() {
	var err error
	// PopLimitHttpDuration
	popLimitHttpDuration, err := strconv.Atoi(get(EnvPopLimitHttpDuration))
	if err != nil {
		log.Panicln(err)
	}
	PopLimitHttpDuration = time.Duration(popLimitHttpDuration)
	// PopLimitHttpRequests
	PopLimitHttpRequests, err = strconv.Atoi(get(EnvPopLimitHttpRequests))
	if err != nil {
		log.Panicln(err)
	}
	// PopLimitRedisDuration
	popLimitRedisDuration, err := strconv.Atoi(get(EnvPopLimitRedisDuration))
	if err != nil {
		log.Panicln(err)
	}
	PopLimitRedisDuration = time.Duration(popLimitRedisDuration)
	// PopLimitRedisPopCount
	PopLimitRedisPopCount, err = strconv.Atoi(get(EnvPopLimitRedisPopCount))
	if err != nil {
		log.Panicln(err)
	}
}

func loadFixRate() {
	// ForceFixRate
	if get(EnvForceFixRate) == "yes" {
		ForceFixRate = true
	} else {
		ForceFixRate = false
	}
}
