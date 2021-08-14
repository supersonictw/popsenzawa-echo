package pop

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/supersonictw/popcat-echo/internal/config"
	"sync"
	"time"
)

func Queue() {
	ctx := context.Background()
	for {
		stepTimestamp := getCurrentStepTimestamp() - config.RefreshInterval*config.RefreshDelay
		key := fmt.Sprintf("%s:%d", config.CacheNamespacePop, stepTimestamp)
		length := redisClient.LLen(ctx, key).Val()
		if length == 0 {
			continue
		}
		allResource := redisClient.LRange(ctx, key, 0, length).Val()
		regionPops := make(map[string]*Pop)
		addressPops := make(map[string]*Pop)
		for _, value := range allResource {
			pop := new(Pop)
			err := json.Unmarshal([]byte(value), pop)
			if err != nil {
				panic(err)
			}
			if origin, ok := regionPops[pop.Region]; ok {
				origin.Count += pop.Count
			} else {
				regionPops[pop.Region] = pop
				regionPops[pop.Region].Address = ""
			}
			if origin, ok := addressPops[pop.Address]; ok {
				origin.Count += pop.Count
			} else {
				addressPops[pop.Region] = pop
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
		redisClient.Del(ctx, key)
		sg.Wait()
		if getCurrentStepTimestamp() == stepTimestamp {
			time.Sleep(time.Second)
		}
	}
}

func updateRegionPop(sg *sync.WaitGroup, pop *Pop) {
	stmt, err := mysqlClient.Prepare(
		"INSERT INTO `region`(`code`, `count`) VALUES(?, ?) ON DUPLICATE KEY UPDATE `count` = `count` + ?",
	)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(pop.Region, pop.Count, pop.Count)
	if err != nil {
		panic(err)
	}
	sg.Done()
}

func updateAddressPop(sg *sync.WaitGroup, pop *Pop) {
	stmt, err := mysqlClient.Prepare(
		"INSERT INTO `address`(`address`, `count`, `region`) VALUES(?, ?) ON DUPLICATE KEY UPDATE `count` = `count` + ?",
	)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(pop.Address, pop.Count, pop.Region, pop.Count)
	if err != nil {
		panic(err)
	}
	sg.Done()
}

func getCurrentStepTimestamp() int64 {
	timestamp := time.Now().Unix()
	return timestamp / config.RefreshInterval * config.RefreshInterval
}
