// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package data

import (
	"log"

	"github.com/adjust/rmq/v5"
)

type Broker struct {
	GlobalSum int64
	RegionMap map[string]int64

	next chan bool
}

func NewBroker() *Broker {
	u := new(Broker)

	regionSum := fetchRegionSum()
	u.GlobalSum = regionSumToGlobal(regionSum)
	u.RegionMap = regionSumToMap(regionSum)

	u.next = make(chan bool, 1)

	return u
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

func (b *Broker) Consume(delivery rmq.Delivery) {
	pop := new(VisitorPop)

	if err := pop.FromMessageQueue(delivery); err != nil {
		log.Println(err)
		b.next <- false
		if err := delivery.Reject(); err != nil {
			log.Println(err)
		}
		return
	}

	b.GlobalSum += pop.Count
	b.RegionMap[pop.RegionCode] += pop.Count

	b.next <- true
	if err := delivery.Ack(); err != nil {
		log.Println(err)
	}
}

func BrokerOnConnected(callback func(globalSum int64, regionMap map[string]int64)) {
	callback(broker.GlobalSum, broker.RegionMap)
}

func BrokerOnUpdated(callback func(globalSum int64, regionMap map[string]int64)) {
	for isNext := range broker.next {
		if !isNext {
			continue
		}
		callback(broker.GlobalSum, broker.RegionMap)
	}
}
