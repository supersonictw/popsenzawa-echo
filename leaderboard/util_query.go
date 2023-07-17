// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package leaderboard

import (
	"log"

	"github.com/supersonictw/popsenzawa-echo/data"
)

func fetchPopsOverview() (int64, []*data.RegionPop) {
	var regionSum []*data.RegionPop
	if tx := data.Database.Find(&regionSum); tx.Error != nil {
		log.Panicln(tx.Error)
	}

	var globalSum int64
	for _, region := range regionSum {
		globalSum += region.Count
	}

	return globalSum, regionSum
}

func regionSumToMap(regionSum []*data.RegionPop) map[string]int64 {
	regionMap := make(map[string]int64, len(regionSum))

	for _, region := range regionSum {
		regionMap[region.RegionCode] = region.Count
	}

	return regionMap
}
