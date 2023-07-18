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
	sendResponse := func(globalSum int64, regionMap map[string]int64) {
		if err := sendMessage(&Response{
			Global:  globalSum,
			Regions: regionMap,
		}); err != nil {
			log.Panicln(err)
		}
	}

	data.BrokerOnConnected(sendResponse)
	data.BrokerOnUpdated(sendResponse)
}
