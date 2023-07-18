// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package leaderboard

type MessageType string

const (
	MessageTypeInitPop MessageType = "init_pop"
	MessageTypeNextPop MessageType = "next_pop"
)

type Message struct {
	Type MessageType `json:"type"`
	Pops any         `json:"pops"`
}
