// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package leaderboard

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func GetLeaderboard(c *gin.Context) {
	sendMessage := useSendMessage(c)

	for {
		globalSum, regionSum := fetchPopsOverview()
		regionMap := regionSumToMap(regionSum)

		if err := sendMessage(&Response{
			Global:  globalSum,
			Regions: regionMap,
		}); err == nil {
			time.Sleep(time.Second)
		} else {
			log.Panicln(err)
		}
	}
}
