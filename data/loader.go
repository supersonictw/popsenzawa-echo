// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package data

import (
	"time"

	"github.com/spf13/viper"
)

var (
	configLoaderPrefetchLimit = viper.GetInt64("loader.prefetch_limit")
	configLoaderPollDuration  = viper.GetFloat64("loader.poll_duration")
	configLoaderWaveDuration  = viper.GetFloat64("loader.wave_duration")
)

var (
	uploader = NewUploader()
	ticker   = time.NewTicker(
		time.Duration(configLoaderWaveDuration) * time.Second,
	)
)

func init() {
	MessageQueue.StartConsuming(
		configLoaderPrefetchLimit,
		time.Duration(configLoaderPollDuration)*time.Second,
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
