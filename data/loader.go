// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package data

import (
	"time"
)

const (
	unackedLimit = 1000
)

var (
	uploader = NewUploader()
	ticker   = time.NewTicker(1 * time.Second)
)

func init() {
	MessageQueue.StartConsuming(unackedLimit, 5*time.Second)
	MessageQueue.AddConsumer("uploader", uploader)
}

func init() {
	go func() {
		for range ticker.C {
			uploader.Wave()
		}
	}()
}
