// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package data

import (
	"encoding/json"

	"github.com/adjust/rmq/v5"
)

type VisitorPop struct {
	IPAddress  VisitorIP `json:"ip_address" gorm:"primaryKey"`
	RegionCode string    `json:"region_code" gorm:"primaryKey"`
	Count      int64     `json:"count" gorm:"not null"`
}

func (v *VisitorPop) Publish() error {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return err
	}

	return MessageQueue.PublishBytes(jsonBytes)
}

func (v *VisitorPop) FromMessageQueue(delivery rmq.Delivery) error {
	dataString := delivery.Payload()
	dataBytes := []byte(dataString)

	return json.Unmarshal(dataBytes, v)
}
