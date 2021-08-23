package pop

import (
	"context"
	"fmt"
	"github.com/supersonictw/popcat-echo/internal/config"
	"log"
	"strconv"
)

func GetGlobalCount(ctx context.Context) int {
	key := fmt.Sprintf("%s:%s", config.CacheNamespacePop, "global")
	sumString := redisClient.Get(ctx, key).Val()
	if sumString == "" {
		return 0
	}
	sum, err := strconv.Atoi(sumString)
	if err != nil {
		log.Panicln(err)
	}
	return sum
}

func GetRegionCount(ctx context.Context, region string) int {
	key := fmt.Sprintf("%s:%s", config.CacheNamespacePop, "regions")
	sumString := redisClient.HGet(ctx, key, region).Val()
	if sumString == "" {
		return 0
	}
	sum, err := strconv.Atoi(sumString)
	if err != nil {
		log.Panicln(err)
	}
	return sum
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
