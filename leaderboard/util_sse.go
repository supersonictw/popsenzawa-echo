// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package leaderboard

import (
	"log"
	"net"
	"strings"

	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func useSendMessage(c *gin.Context) func(message *Message) error {
	// Handshake
	sessionID := uuid.New().String()
	eventId := strings.Join([]string{
		c.ClientIP(),
		sessionID,
	}, "/")
	err := sse.Event{
		Id:    eventId,
		Event: "handshake",
		Data:  "This is PopSenzawa Echo, ùwú.",
	}.Render(c.Writer)
	if err != nil {
		log.Panicln(err)
	}

	// Wrap real sendMessage
	return func(message *Message) error {
		return sendMessage(c, sessionID, message)
	}
}

func sendMessage(c *gin.Context, sessionID string, message *Message) error {
	// Send message
	messageID := uuid.New().String()
	eventId := strings.Join([]string{
		sessionID,
		messageID,
	}, "/")
	err := sse.Encode(c.Writer, sse.Event{
		Id:    eventId,
		Event: "message",
		Data:  message,
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
