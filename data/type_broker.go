// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package data

import (
	"context"
	"log"
	"time"

	"github.com/spf13/viper"
)

var (
	configBrokerPingDuration = viper.GetFloat64("broker.ping_duration")
)

type Broker struct {
	SessionID string
}

type BrokerInitPop struct {
	GlobalSum int64            `json:"global_sum"`
	RegionMap map[string]int64 `json:"region_map"`
}

type BrokerNextPop struct {
	RegionCode  string `json:"region_code"`
	CountAppend int64  `json:"count_append"`
}

func NewBroker(sessionID string) *Broker {
	b := new(Broker)
	b.SessionID = sessionID
	return b
}

func fetchRegionSum() []*RegionPop {
	var regionSum []*RegionPop

	if tx := database.Find(&regionSum); tx.Error != nil {
		log.Panicln(tx.Error)
	}

	return regionSum
}

func regionSumToGlobal(regionSum []*RegionPop) int64 {
	var globalSum int64

	for _, region := range regionSum {
		globalSum += region.Count
	}

	return globalSum
}

func regionSumToMap(regionSum []*RegionPop) map[string]int64 {
	regionMap := make(map[string]int64, len(regionSum))

	for _, region := range regionSum {
		regionMap[region.RegionCode] = region.Count
	}

	return regionMap
}

func (b *Broker) OnConnected(callback func(initPop *BrokerInitPop)) {
	regionSum := fetchRegionSum()
	globalSum := regionSumToGlobal(regionSum)
	regionMap := regionSumToMap(regionSum)

	callback(&BrokerInitPop{
		GlobalSum: globalSum,
		RegionMap: regionMap,
	})
}

func (b *Broker) OnActive(ctx context.Context, callback func(timestamp time.Time)) {
	brokerTickerDuration := time.Duration(configBrokerPingDuration) * time.Second
	brokerTicker := time.NewTicker(brokerTickerDuration)

	defer func() {
		brokerTicker.Stop()
	}()

	for time := range brokerTicker.C {
		select {
		case <-ctx.Done():
			return
		default:
			callback(time)
		}
	}
}

func (b *Broker) OnUpdated(ctx context.Context, callback func(nextPop *BrokerNextPop)) {
	pubSub := redisClient.Subscribe(ctx, redisKey(redisKeyBroker))

	for {
		select {
		case message := <-pubSub.Channel():
			pubPop := new(VisitorPop)
			if err := pubPop.FromString(message.Payload); err != nil {
				log.Panicln(err)
			}
			callback(&BrokerNextPop{
				RegionCode:  pubPop.RegionCode,
				CountAppend: pubPop.Count,
			})
		case <-ctx.Done():
			return
		}
	}
}
