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
)

func GetGlobalCount(ctx context.Context) int {
	key := fmt.Sprintf("%s:%s", config.CacheNamespacePop, "global")
	value, err := redisClient.Get(ctx, key).Int()
	if err != nil && !errors.Is(err, redis.Nil) {
		log.Panicln(err)
	}
	return value
}

func GetRegionCount(ctx context.Context, region string) int {
	key := fmt.Sprintf("%s:%s", config.CacheNamespacePop, "regions")
	value, err := redisClient.HGet(ctx, key, region).Int()
	if err != nil && !errors.Is(err, redis.Nil) {
		log.Panicln(err)
	}
	return value
}

func AppendRegionCount(ctx context.Context, region string, count int) {
	keyGlobal := fmt.Sprintf("%s:%s", config.CacheNamespacePop, "global")
	sumGlobal := GetGlobalCount(ctx) + count
	err := redisClient.Set(ctx, keyGlobal, sumGlobal, 0).Err()
	if err != nil {
		log.Panicln(err)
	}
	keyRegions := fmt.Sprintf("%s:%s", config.CacheNamespacePop, "regions")
	sumRegion := GetRegionCount(ctx, region) + count
	err = redisClient.HSet(ctx, keyRegions, region, sumRegion).Err()
	if err != nil {
		log.Panicln(err)
	}
}
