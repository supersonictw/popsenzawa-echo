package internal

import (
	"database/sql"
	"github.com/dpapathanasiou/go-recaptcha"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"
	"github.com/supersonictw/popcat-echo/internal/config"
	"log"
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
	DB, err := sql.Open("mysql", config.Get(config.MysqlDSN))
	if err != nil {
		panic(err)
	}
	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	redisDatabase, err := strconv.Atoi(config.Get(config.RedisDatabase))
	if err != nil {
		log.Fatal(err)
	}
	RDB = redis.NewClient(&redis.Options{
		Addr:     config.Get(config.RedisHostname),
		Password: config.Get(config.RedisPassword),
		DB:       redisDatabase,
	})

	refreshIntervalInt, err := strconv.Atoi(config.Get(config.RefreshInterval))
	if err != nil {
		panic(err)
	}
	RefreshInterval = int64(refreshIntervalInt)
	CacheNamespacePop = config.Get(config.RedisNamespacePop)
	CacheNamespaceGeo = config.Get(config.RedisNamespaceGeo)

	if secret := config.Get(config.ReCaptchaSecret); secret != "" {
		recaptcha.Init(secret)
		ReCaptchaStatus = true
	} else {
		ReCaptchaStatus = false
	}

	if secret := config.Get(config.JWTSecret); secret != "" {
		JWTCaptchaSecret = secret
	} else {
		JWTCaptchaSecret = uuid.NewString()
	}
}
