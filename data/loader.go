// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package data

import (
	"time"

	"github.com/spf13/viper"
)

var (
	configMessageQueuePrefetchLimit = viper.GetInt64("message_queue.prefetcb_limit")
	configMessageQueuePollDuration  = viper.GetFloat64("message_queue.poll_duration")
	configMessageQueueWaveDuration  = viper.GetFloat64("message_queue.wave_duration")
)

var (
	uploader = NewUploader()
	ticker   = time.NewTicker(
		time.Duration(configMessageQueueWaveDuration) * time.Second,
	)
)

func init() {
	MessageQueue.StartConsuming(
		configMessageQueuePrefetchLimit,
		time.Duration(configMessageQueuePollDuration)*time.Second,
	)
	MessageQueue.AddConsumer("uploader", uploader)
}

func init() {
	go func() {
		for range ticker.C {
			uploader.Wave()
		}
	}()
}
