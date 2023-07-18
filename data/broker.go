// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package data

import (
	"log"

	"github.com/adjust/rmq/v5"
)

type Broker struct {
	nextPop chan *BrokerNextPop
}

type BrokerInitPop struct {
	GlobalSum int64            `json:"global_sum"`
	RegionMap map[string]int64 `json:"region_map"`
}

type BrokerNextPop struct {
	RegionCode  string `json:"region_code"`
	CountAppend int64  `json:"count_append"`
}

func NewBroker() *Broker {
	u := new(Broker)
	u.nextPop = make(chan *BrokerNextPop, 1)
	return u
}

func (b *Broker) Consume(delivery rmq.Delivery) {
	pop := new(VisitorPop)

	if err := pop.FromMessageQueue(delivery); err != nil {
		log.Println(err)
		if err := delivery.Reject(); err != nil {
			log.Println(err)
		}
		return
	}

	b.nextPop <- &BrokerNextPop{
		RegionCode:  pop.RegionCode,
		CountAppend: pop.Count,
	}

	if err := delivery.Ack(); err != nil {
		log.Println(err)
	}
}

func fetchRegionSum() []*RegionPop {
	var regionSum []*RegionPop

	if tx := Database.Find(&regionSum); tx.Error != nil {
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

func BrokerOnConnected(callback func(initPop *BrokerInitPop)) {
	regionSum := fetchRegionSum()
	globalSum := regionSumToGlobal(regionSum)
	regionMap := regionSumToMap(regionSum)
	callback(&BrokerInitPop{
		GlobalSum: globalSum,
		RegionMap: regionMap,
	})
}

func BrokerOnUpdated(callback func(nextPop *BrokerNextPop)) {
	for nextPop := range broker.nextPop {
		callback(nextPop)
	}
}
