package leaderboard

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/r3labs/sse/v2"
	"time"
)

func Response(c *gin.Context) {
	server := sse.New()
	ctx, cancel := context.WithCancel(context.Background())
	server.CreateStream("messages")
	go listen(ctx, server)
	server.HTTPHandler(c.Writer, c.Request)
	cancel()
}

func listen(ctx context.Context, server *sse.Server) {
	for {
		jsonBytes, err := json.Marshal(fetchRegionPops())
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

func fetchRegionPops() map[string]interface{} {
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
