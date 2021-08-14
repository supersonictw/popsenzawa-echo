package pop

import (
	"context"
	"fmt"
	"github.com/supersonictw/popcat-echo/internal/config"
	"strconv"
	"time"
)

func GetAddressCountInRefreshInterval(ctx context.Context, address string) int {
	stepTimestamp := getCurrentStepTimestamp()
	key := fmt.Sprintf("%s:%d", config.CacheNamespaceRate, stepTimestamp)
	sumString := redisClient.HGet(ctx, key, address).Val()
	if sumString == "" {
		return 0
	}
	sum, err := strconv.Atoi(sumString)
	if err != nil {
		panic(err)
	}
	return sum
}

func AppendAddressCountInRefreshInterval(ctx context.Context, address string, count int) {
	stepTimestamp := getCurrentStepTimestamp()
	key := fmt.Sprintf("%s:%d", config.CacheNamespaceRate, stepTimestamp)
	previous := GetAddressCountInRefreshInterval(ctx, address)
	count += previous
	err := redisClient.HSet(ctx, key, address, count).Err()
	if err != nil {
		panic(err)
	}
	if previous == 0 {
		redisClient.Expire(ctx, key, config.JWTExpired*time.Second)
	}
}
