package leaderboard

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/r3labs/sse/v2"
	"github.com/supersonictw/popcat-echo/internal/config"
	"strconv"
	"time"
)

func Response(c *gin.Context) {
	server := sse.New()
	ctx, cancel := context.WithCancel(context.Background())
	server.CreateStream("messages")
	go listen(ctx, server)
	server.HTTPHandler(c.Writer, c.Request)
	cancel()
	server.Close()
}

func listen(ctx context.Context, server *sse.Server) {
	for {
		queryCtx := context.Background()
		jsonBytes, err := json.Marshal(fetchRegionPopsFromRedis(queryCtx))
		if err != nil {
			panic(err)
		}
		select {
		case <-ctx.Done():
			return
		default:
			server.Publish("messages", &sse.Event{
				Data: jsonBytes,
			})
		}
		time.Sleep(time.Second)
	}
}

func PrepareCache() {
	ctx := context.Background()
	data := fetchRegionPopsFromMySQL()
	keyGlobal := fmt.Sprintf("%s:%s", config.CacheNamespacePop, "global")
	err := redisClient.Set(ctx, keyGlobal, data["global"], 0).Err()
	if err != nil {
		panic(err)
	}
	keyRegions := fmt.Sprintf("%s:%s", config.CacheNamespacePop, "regions")
	for i, region := range data["regions"].(map[string]int) {
		err := redisClient.HSet(ctx, keyRegions, i, region).Err()
		if err != nil {
			panic(err)
		}
	}
}

func fetchRegionPopsFromRedis(ctx context.Context) map[string]interface{} {
	keyGlobal := fmt.Sprintf("%s:%s", config.CacheNamespacePop, "global")
	resultGlobal := redisClient.Get(ctx, keyGlobal).Val()
	keyRegions := fmt.Sprintf("%s:%s", config.CacheNamespacePop, "regions")
	resultRegions := redisClient.HGetAll(ctx, keyRegions).Val()
	sumGlobal, err := strconv.Atoi(resultGlobal)
	if err != nil {
		panic(err)
	}
	sumRegions := make(map[string]int, len(resultRegions))
	for i, region := range resultRegions {
		sumRegions[i], err = strconv.Atoi(region)
		if err != nil {
			panic(err)
		}
	}
	return map[string]interface{}{
		"global":  sumGlobal,
		"regions": resultRegions,
	}
}

func fetchRegionPopsFromMySQL() map[string]interface{} {
	rows, err := mysqlClient.Query(
		"SELECT `code`, `count` FROM `region`",
	)
	if err != nil {
		panic(err)
	}
	sum := 0
	regions := make(map[string]int)
	for rows.Next() {
		var code string
		var count int
		err = rows.Scan(&code, &count)
		if err != nil {
			panic(err)
		}
		sum += count
		regions[code] = count
	}
	return map[string]interface{}{
		"global":  sum,
		"regions": regions,
	}
}
