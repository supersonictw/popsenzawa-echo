// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package data

import (
	"log"

	"github.com/adjust/rmq/v5"
)

const (
	rmqConnectionName = "popsenzawa_rmq"
)

const (
	rmqQueueNameUpload = "upload"
)

var (
	rmqConnection rmq.Connection
	errChan       chan<- error
)

func init() {
	var err error
	rmqConnection, err = rmq.OpenConnectionWithRedisClient(
		rmqConnectionName,
		redisClient,
		errChan,
	)
	if err != nil {
		log.Panicln(err)
	}
}
