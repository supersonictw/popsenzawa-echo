// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package leaderboard

import (
	"log"
	"net"

	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
)

func useSendMessage(c *gin.Context) func(response *Response) error {
	// Handshake
	err := sse.Event{
		Event: "handshake",
		Data:  "This is PopSenzawa Echo, ùwú.",
	}.Render(c.Writer)
	if err != nil {
		log.Panicln(err)
	}

	// Wrap real sendMessage
	return func(response *Response) error {
		return sendMessage(c, response)
	}
}

func sendMessage(c *gin.Context, response *Response) error {
	// Send message
	err := sse.Encode(c.Writer, sse.Event{
		Event: "message",
		Data:  response,
	})
	if status, ok := err.(*net.OpError); ok &&
		status.Err.Error() == "write: broken pipe" {
		return nil
	}
	if err != nil {
		return err
	}

	// Flush response
	c.Writer.Flush()

	return nil
}
