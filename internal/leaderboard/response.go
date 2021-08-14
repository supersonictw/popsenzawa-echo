package leaderboard

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/r3labs/sse/v2"
	"time"
)

func Response(c *gin.Context) {
	server := sse.New()
	server.HTTPHandler(c.Writer, c.Request)
	for {
		jsonBytes, err := json.Marshal(fetchRegionPops())
		if err != nil {
			panic(err)
		}
		server.Publish("message", &sse.Event{
			Data: jsonBytes,
		})
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
