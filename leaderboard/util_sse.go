// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package leaderboard

import (
	"log"
	"net"
	"strings"
	"time"

	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SendPingHandler func(t time.Time) (string, error)
type SendMessageHandler func(m *Message) (string, error)

func useServerSideEvent(c *gin.Context) (string, SendPingHandler, SendMessageHandler) {
	// Generate metadata
	sessionID := uuid.New().String()

	// Handshake
	err := sse.Event{
		Id:    sessionID,
		Event: "handshake",
		Data:  "This is PopSenzawa Echo, ùwú.",
	}.Render(c.Writer)
	if err != nil {
		log.Panicln(err)
	}

	// Wrap real handlers
	sendPingHandler := func(t time.Time) (string, error) {
		return sendPing(c, sessionID, t)
	}
	sendMessageHandler := func(m *Message) (string, error) {
		return sendMessage(c, sessionID, m)
	}

	// Return handlers
	return sessionID, sendPingHandler, sendMessageHandler
}

func sendPing(c *gin.Context, sessionID string, timestamp time.Time) (string, error) {
	// Generate metadata
	pingID := uuid.New().String()
	eventId := strings.Join([]string{sessionID, pingID}, "/")

	// Send ping
	err := sse.Encode(c.Writer, sse.Event{
		Event: "ping",
		Id:    eventId,
		Data:  timestamp.Format(time.RFC3339),
	})
	if status, ok := err.(*net.OpError); ok &&
		status.Err.Error() == "write: broken pipe" {
		return pingID, nil
	}
	if err != nil {
		return pingID, err
	}

	// Flush response
	c.Writer.Flush()

	return pingID, nil
}

func sendMessage(c *gin.Context, sessionID string, message *Message) (string, error) {
	// Generate metadata
	messageID := uuid.New().String()
	eventId := strings.Join([]string{sessionID, messageID}, "/")

	// Send message
	err := sse.Encode(c.Writer, sse.Event{
		Event: "message",
		Id:    eventId,
		Data:  message,
	})
	if status, ok := err.(*net.OpError); ok &&
		status.Err.Error() == "write: broken pipe" {
		return messageID, nil
	}
	if err != nil {
		return messageID, err
	}

	// Flush response
	c.Writer.Flush()

	return messageID, nil
}
