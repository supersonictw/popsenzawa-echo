// PopCat Echo
// (c) 2022uperSonic (https://github.com/supersonictw).

package leaderboard

import (
	"log"
	"net"
	"time"

	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Global  int
	Regions map[string]int
}

func GetLeaderboard(c *gin.Context) {
	err := sse.Event{
		Event: "handshake",
		Data:  "This is PopCat Echo, ùwú.",
	}.Render(c.Writer)
	if err != nil {
		log.Panicln(err)
	}
	for {
		err = sse.Encode(c.Writer, sse.Event{
			Event: "message",
			Data:  &Response{},
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
