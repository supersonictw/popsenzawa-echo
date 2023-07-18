// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package leaderboard

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/supersonictw/popsenzawa-echo/data"
)

func GetLeaderboard(c *gin.Context) {
	sendMessage := useSendMessage(c)

	data.BrokerOnConnected(func(initPop *data.BrokerInitPop) {
		if err := sendMessage(&Message{
			Type: MessageTypeInitPop,
			Pops: initPop,
		}); err != nil {
			log.Panicln(err)
		}
	})

	data.BrokerOnUpdated(func(nextPop *data.BrokerNextPop) {
		if err := sendMessage(&Message{
			Type: MessageTypeNextPop,
			Pops: nextPop,
		}); err != nil {
			log.Panicln(err)
		}
	}, c.Done())
}
