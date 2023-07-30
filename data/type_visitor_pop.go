// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package data

import (
	"context"
	"encoding/json"
)

type VisitorPop struct {
	IPAddress  VisitorIP `json:"ip_address" gorm:"primaryKey"`
	RegionCode string    `json:"region_code" gorm:"primaryKey"`
	Count      int64     `json:"count" gorm:"not null"`
}

func (v *VisitorPop) Publish(ctx context.Context) error {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return err
	}

	if err := uploadQueue.PublishBytes(jsonBytes); err != nil {
		return err
	}

	if err := redisClient.Publish(
		ctx,
		redisKey(redisKeyBroker),
		jsonBytes,
	).Err(); err != nil {
		return err
	}

	return nil
}

func (v *VisitorPop) FromString(dataString string) error {
	dataBytes := []byte(dataString)
	return json.Unmarshal(dataBytes, v)
}
