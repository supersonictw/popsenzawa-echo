// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package data

import (
	"log"
	"time"

	"github.com/adjust/rmq/v5"
	"github.com/spf13/viper"
)

var (
	configUploadPrefetchLimit = viper.GetInt64("upload.prefetch_limit")
	configUploadPollDuration  = viper.GetFloat64("upload.poll_duration")
	configUploadWaveDuration  = viper.GetFloat64("upload.wave_duration")
)

var (
	uploader    *Uploader
	uploadQueue rmq.Queue
)

func init() {
	var err error

	uploadQueue, err = rmqConnection.OpenQueue(rmqQueueNameUpload)
	if err != nil {
		log.Panicln(err)
	}

	prefetchLimit := configUploadPrefetchLimit
	pollDuration := time.Duration(configUploadPollDuration) * time.Second
	uploadQueue.StartConsuming(prefetchLimit, pollDuration)

	uploader = NewUploader()
	uploadQueue.AddConsumer("uploader", uploader)
}

func init() {
	uploadTickerDuration := time.Duration(configUploadWaveDuration) * time.Second
	uploadTicker := time.NewTicker(uploadTickerDuration)
	go func() {
		for range uploadTicker.C {
			uploader.Wave()
		}
	}()
}
