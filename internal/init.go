package internal

import (
	"database/sql"
	"github.com/dpapathanasiou/go-recaptcha"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"
	"github.com/supersonictw/popcat-echo/internal/config"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	DB                *sql.DB
	RDB               *redis.Client
	RefreshInterval   int64
	CacheNamespacePop string
	CacheNamespaceGeo string
	ReCaptchaStatus   bool
	JWTCaptchaSecret  string
)

func init() {
	DB, err := sql.Open("mysql", os.Getenv(config.MysqlDSN))
	if err != nil {
		panic(err)
	}
	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	redisDatabase, err := strconv.Atoi(os.Getenv(config.RedisDatabase))
	if err != nil {
		log.Fatal(err)
	}
	RDB = redis.NewClient(&redis.Options{
		Addr:     os.Getenv(config.RedisHostname),
		Password: os.Getenv(config.RedisPassword),
		DB:       redisDatabase,
	})

	refreshIntervalInt, err := strconv.Atoi(os.Getenv(config.RefreshInterval))
	if err != nil {
		panic(err)
	}
	RefreshInterval = int64(refreshIntervalInt)
	CacheNamespacePop = os.Getenv(config.RedisNamespacePop)
	CacheNamespaceGeo = os.Getenv(config.RedisNamespaceGeo)

	if secret := os.Getenv(config.ReCaptchaSecret); secret != "" {
		recaptcha.Init(secret)
		ReCaptchaStatus = true
	} else {
		ReCaptchaStatus = false
	}

	if secret := os.Getenv(config.JWTSecret); secret != "" {
		JWTCaptchaSecret = secret
	} else {
		JWTCaptchaSecret = uuid.NewString()
	}
}
