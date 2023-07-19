// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package leaderboard

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/supersonictw/popsenzawa-echo/data"
)

func GetLeaderboard(c *gin.Context) {
	sessionID, sendPing, sendMessage := useServerSideEvent(c)
	broker := data.NewBroker(sessionID)

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()

	broker.OnConnected(func(initPop *data.BrokerInitPop) {
		if _, err := sendMessage(&Message{
			Type: MessageTypeInitPop,
			Pops: initPop,
		}); err != nil {
			log.Panicln(err)
		}
	})

	go broker.OnActive(ctx, func(t time.Time) {
		if _, err := sendPing(t); err != nil {
			log.Panicln(err)
		}
	})

	go broker.OnUpdated(ctx, func(nextPop *data.BrokerNextPop) {
		if _, err := sendMessage(&Message{
			Type: MessageTypeNextPop,
			Pops: nextPop,
		}); err != nil {
			log.Panicln(err)
		}
	})

	<-c.Request.Context().Done()
}
