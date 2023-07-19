// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package leaderboard

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/supersonictw/popsenzawa-echo/data"
)

func GetLeaderboard(c *gin.Context) {
	_, sendPing, sendMessage := useServerSideEvent(c)

	data.BrokerOnConnected(func(initPop *data.BrokerInitPop) {
		if _, err := sendMessage(&Message{
			Type: MessageTypeInitPop,
			Pops: initPop,
		}); err != nil {
			log.Panicln(err)
		}
	})

	go data.BrokerOnActive(func(t time.Time) {
		if _, err := sendPing(t); err != nil {
			log.Panicln(err)
		}
	}, c.Done())

	go data.BrokerOnUpdated(func(nextPop *data.BrokerNextPop) {
		if _, err := sendMessage(&Message{
			Type: MessageTypeNextPop,
			Pops: nextPop,
		}); err != nil {
			log.Panicln(err)
		}
	}, c.Done())

	<-c.Done()
}
