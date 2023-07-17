// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package data

import (
	"log"

	"github.com/adjust/rmq/v5"
	"github.com/spf13/viper"
)

const (
	ConnectionName = "rmq"
	QueueName      = ""
)

var (
	configRedisNetwork  = viper.GetString("redis.network")
	configRedisAddress  = viper.GetString("redis.address")
	configRedisDatabase = viper.GetInt("redis.database")
)

var (
	MessageQueue rmq.Queue
	errChan      chan<- error
)

func init() {
	var err error
	connection, err := rmq.OpenConnection(
		ConnectionName,
		configRedisNetwork,
		configRedisAddress,
		configRedisDatabase,
		errChan,
	)
	if err != nil {
		log.Panicln(err)
	}

	MessageQueue, err = connection.OpenQueue(QueueName)
	if err != nil {
		log.Panicln(err)
	}
}
