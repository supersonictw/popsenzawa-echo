// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package data

import (
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

const (
	redisNamespace = "popsenzawa_redis"
)

const (
	redisKeyBroker = "broker"
)

var (
	configRedisNetwork  = viper.GetString("redis.network")
	configRedisAddress  = viper.GetString("redis.address")
	configRedisDatabase = viper.GetInt("redis.database")
)

var (
	redisClient *redis.Client
)

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Network: configRedisNetwork,
		Addr:    configRedisAddress,
		DB:      configRedisDatabase,
	})
}

func redisKey(key string) string {
	return strings.Join([]string{redisNamespace, key}, ":")
}
