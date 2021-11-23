// PopCat Echo
// (c) 2021 SuperSonic (https://github.com/supersonictw).

package pop

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/supersonictw/popcat-echo/internal/config"
	"log"
	"time"
)

func GetAddressCountInRefreshInterval(ctx context.Context, address string) int {
	stepTimestamp := getCurrentStepTimestamp()
	key := fmt.Sprintf("%s:%d", config.CacheNamespaceRate, stepTimestamp)
	value, err := redisClient.HGet(ctx, key, address).Int()
	if err != nil && !errors.Is(err, redis.Nil) {
		log.Panicln(err)
	}
	return value
}

func AppendAddressCountInRefreshInterval(ctx context.Context, address string, count int) {
	stepTimestamp := getCurrentStepTimestamp()
	key := fmt.Sprintf("%s:%d", config.CacheNamespaceRate, stepTimestamp)
	previous := GetAddressCountInRefreshInterval(ctx, address)
	count += previous
	err := redisClient.HSet(ctx, key, address, count).Err()
	if err != nil {
		log.Panicln(err)
	}
	if previous == 0 {
		redisClient.Expire(ctx, key, config.PopLimitRedisDuration*time.Second)
	}
}
