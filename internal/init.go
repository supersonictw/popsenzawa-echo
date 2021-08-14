package internal

import (
	"database/sql"
	"fmt"
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
	PublishAddress    string
	DB                *sql.DB
	RDB               *redis.Client
	RefreshInterval   int64
	RefreshDelay      int64
	CacheNamespacePop string
	CacheNamespaceGeo string
	ReCaptchaStatus   bool
	JWTCaptchaSecret  string
	JWTExpired        time.Duration
)

func init() {
	fmt.Println("PopCat Echo")
	fmt.Println("===")
	fmt.Println("The server reproduce of https://popcat.click with improvement.")
	fmt.Println("License: MIT LICENSE")
	fmt.Println("Repository: https://github.com/supersonictw/popcat-echo")
	fmt.Println("(c) 2021 SuperSonic. https://github.com/supersonictw")
	fmt.Println()

	PublishAddress = config.Get(config.PublishAddress)

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
	refreshDelayInt, err := strconv.Atoi(config.Get(config.RefreshDelay))
	if err != nil {
		panic(err)
	}
	RefreshDelay = int64(refreshDelayInt)
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

	jwtExpired, err := strconv.Atoi(config.Get(config.JWTExpired))
	if err != nil {
		panic(err)
	}
	JWTExpired = time.Duration(jwtExpired)
}
