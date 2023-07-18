// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package data

import (
	"log"
)

var (
	nextPop = make(chan *BrokerNextPop, configMessageQueuePrefetchLimit)
)

type BrokerInitPop struct {
	GlobalSum int64            `json:"global_sum"`
	RegionMap map[string]int64 `json:"region_map"`
}

type BrokerNextPop struct {
	RegionCode  string `json:"region_code"`
	CountAppend int64  `json:"count_append"`
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

func BrokerOnUpdated(callback func(nextPop *BrokerNextPop), done <-chan struct{}) {
	for pop := range nextPop {
		select {
		case <-done:
			return
		default:
			callback(pop)
		}
	}
}
