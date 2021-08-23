package pop

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/supersonictw/popcat-echo/internal/config"
	"log"
	"sync"
	"time"
)

func Queue() {
	for {
		stepTimestamp := getPreviousStepTimestamp()
		go doTask(stepTimestamp)
		for {
			if getPreviousStepTimestamp() == stepTimestamp {
				time.Sleep(time.Second)
			} else {
				break
			}
		}
	}
}

func doTask(stepTimestamp int64) {
	ctx := context.Background()
	key := fmt.Sprintf("%s:%d", config.CacheNamespacePop, stepTimestamp)
	length := redisClient.LLen(ctx, key).Val()
	if length == 0 {
		return
	}
	allResource := redisClient.LRange(ctx, key, 0, length).Val()
	regionPops := make(map[string]*Pop)
	addressPops := make(map[string]*Pop)
	for _, value := range allResource {
		pop := new(Pop)
		err := json.Unmarshal([]byte(value), pop)
		if err != nil {
			log.Panicln(err)
		}
		if origin, ok := regionPops[pop.Region]; ok {
			origin.Count += pop.Count
		} else {
			regionPops[pop.Region] = pop
		}
		if origin, ok := addressPops[pop.Address]; ok {
			if config.ForceFixRate && config.RateLimit != 0 && pop.Count > config.RateLimit {
				recycleCount := pop.Count - config.RateLimit
				regionPops[pop.Region].Count -= recycleCount
				AppendRegionCount(ctx, pop.Region, -recycleCount)
				pop.Count = config.RateLimit
			}
			origin.Count += pop.Count
		} else {
			addressPops[pop.Address] = pop
		}
	}
	sg := new(sync.WaitGroup)
	sg.Add(int(length * 2))
	for _, value := range regionPops {
		go updateRegionPop(sg, value)
	}
	for _, value := range addressPops {
		go updateAddressPop(sg, value)
	}
	err := redisClient.Del(ctx, key).Err()
	if err != nil {
		log.Panicln(err)
	}
	sg.Wait()
}

func updateRegionPop(sg *sync.WaitGroup, pop *Pop) {
	stmt, err := mysqlClient.Prepare(
		"INSERT INTO `region`(`code`, `count`) VALUES(?, ?) ON DUPLICATE KEY UPDATE `count` = `count` + ?",
	)
	if err != nil {
		log.Panicln(err)
	}
	_, err = stmt.Exec(pop.Region, pop.Count, pop.Count)
	if err != nil {
		log.Panicln(err)
	}
	sg.Done()
}

func updateAddressPop(sg *sync.WaitGroup, pop *Pop) {
	stmt, err := mysqlClient.Prepare(
		"INSERT INTO `address`(`address`, `count`, `region`) VALUES(?, ?, ?) ON DUPLICATE KEY UPDATE `count` = `count` + ?",
	)
	if err != nil {
		log.Panicln(err)
	}
	_, err = stmt.Exec(pop.Address, pop.Count, pop.Region, pop.Count)
	if err != nil {
		log.Panicln(err)
	}
	sg.Done()
}

func getCurrentStepTimestamp() int64 {
	timestamp := time.Now().Unix()
	return timestamp / config.RefreshInterval * config.RefreshInterval
}

func getPreviousStepTimestamp() int64 {
	return getCurrentStepTimestamp() - config.RefreshInterval*config.RefreshDelay
}
