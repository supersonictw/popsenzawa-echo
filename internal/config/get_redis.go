// PopCat Echo
// (c) 2021 SuperSonic (https://github.com/supersonictw).

package config

import (
	"github.com/go-redis/redis/v8"
)

func GetRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     RedisAddress,
		Password: RedisPassword,
		DB:       RedisDatabase,
	})
}
