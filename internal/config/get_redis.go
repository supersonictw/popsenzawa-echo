package config

import (
	"github.com/go-redis/redis/v8"
	"github.com/supersonictw/popcat-echo/internal"
	"log"
	"strconv"
)

func GetRedis() *redis.Client {
	redisDatabase, err := strconv.Atoi(Get(internal.ConfigRedisDatabase))
	if err != nil {
		log.Fatal(err)
	}
	return redis.NewClient(&redis.Options{
		Addr:     Get(internal.ConfigRedisAddress),
		Password: Get(internal.ConfigRedisPassword),
		DB:       redisDatabase,
	})
}
