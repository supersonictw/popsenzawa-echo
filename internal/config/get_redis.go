// PopCat Echo
// (c) 2021 SuperSonic (https://github.com/supersonictw).

package config

import (
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
)

func GetRedis() *redis.Client {
	redisDatabase, err := strconv.Atoi(Get(EnvRedisDatabase))
	if err != nil {
		log.Fatal(err)
	}
	return redis.NewClient(&redis.Options{
		Addr:     Get(EnvRedisAddress),
		Password: Get(EnvRedisPassword),
		DB:       redisDatabase,
	})
}
