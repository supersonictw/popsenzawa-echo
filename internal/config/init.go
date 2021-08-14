package config

import (
	"crypto/rand"
	"database/sql"
	"github.com/dpapathanasiou/go-recaptcha"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	"github.com/supersonictw/popcat-echo/internal"
	"log"
	"strconv"
	"time"
)

var (
	PublishAddress    string
	DB                *sql.DB
	RDB               *redis.Client
	RefreshInterval   int64
	RefreshDelay      int64
	CacheNamespacePop string
	CacheNamespaceGeo string
	ReCaptchaStatus   bool
	JWTCaptchaSecret  []byte
	JWTExpired        time.Duration
)

func init() {
	PublishAddress = Get(internal.ConfigPublishAddress)

	DB, err := sql.Open("mysql", Get(internal.ConfigMysqlDSN))
	if err != nil {
		panic(err)
	}
	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	redisDatabase, err := strconv.Atoi(Get(internal.ConfigRedisDatabase))
	if err != nil {
		log.Fatal(err)
	}
	RDB = redis.NewClient(&redis.Options{
		Addr:     Get(internal.ConfigRedisHostname),
		Password: Get(internal.ConfigRedisPassword),
		DB:       redisDatabase,
	})

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
}
