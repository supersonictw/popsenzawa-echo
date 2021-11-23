// PopCat Echo
// (c) 2021 SuperSonic (https://github.com/supersonictw).

package leaderboard

import (
	"context"
	"fmt"
	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
	"github.com/supersonictw/popcat-echo/internal/config"
	"log"
	"net"
	"strconv"
	"time"
)

func Response(c *gin.Context) {
	err := sse.Event{
		Event: "handshake",
		Data:  "This is PopCat Echo, ùwú.",
	}.Render(c.Writer)
	if err != nil {
		log.Panicln(err)
	}
	ctx := context.Background()
	for {
		err = sse.Encode(c.Writer, sse.Event{
			Event: "message",
			Data:  fetchRegionPopsFromRedis(ctx),
		})
		if status, ok := err.(*net.OpError); ok && status.Err.Error() == "write: broken pipe" {
			return
		}
		if err != nil {
			log.Panicln(err)
		}
		c.Writer.Flush()
		time.Sleep(time.Second)
	}
}

func PrepareCache() {
	ctx := context.Background()
	data := fetchRegionPopsFromMySQL()
	keyGlobal := fmt.Sprintf("%s:%s", config.CacheNamespacePop, "global")
	err := redisClient.Set(ctx, keyGlobal, data["global"], 0).Err()
	if err != nil {
		log.Panicln(err)
	}
	keyRegions := fmt.Sprintf("%s:%s", config.CacheNamespacePop, "regions")
	for i, region := range data["regions"].(map[string]int) {
		err := redisClient.HSet(ctx, keyRegions, i, region).Err()
		if err != nil {
			log.Panicln(err)
		}
	}
}

func fetchRegionPopsFromRedis(ctx context.Context) map[string]interface{} {
	keyGlobal := fmt.Sprintf("%s:%s", config.CacheNamespacePop, "global")
	keyRegions := fmt.Sprintf("%s:%s", config.CacheNamespacePop, "regions")
	sumGlobal, err := redisClient.Get(ctx, keyGlobal).Int()
	if err != nil {
		log.Panicln(err)
	}
	resultRegions := redisClient.HGetAll(ctx, keyRegions).Val()
	sumRegions := make(map[string]int, len(resultRegions))
	for i, region := range resultRegions {
		sumRegions[i], err = strconv.Atoi(region)
		if err != nil {
			log.Panicln(err)
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
		log.Panicln(err)
	}
	sum := 0
	regions := make(map[string]int)
	for rows.Next() {
		var code string
		var count int
		err = rows.Scan(&code, &count)
		if err != nil {
			log.Panicln(err)
		}
		sum += count
		regions[code] = count
	}
	return map[string]interface{}{
		"global":  sum,
		"regions": regions,
	}
}
