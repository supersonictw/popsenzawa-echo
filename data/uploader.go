// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package data

import (
	"errors"
	"log"

	"github.com/adjust/rmq/v5"
	"gorm.io/gorm"
)

type Uploader struct {
	VisitorPopSum map[string]*VisitorPop
}

func NewUploader() *Uploader {
	u := new(Uploader)
	u.reset()
	return u
}

func (u *Uploader) Consume(delivery rmq.Delivery) {
	pop := new(VisitorPop)

	if err := pop.FromMessageQueue(delivery); err != nil {
		log.Println(err)
		if err := delivery.Reject(); err != nil {
			log.Println(err)
		}
		return
	}

	ipAddressString := pop.IPAddress.NetIP().String()
	if u.VisitorPopSum[ipAddressString] == nil {
		u.VisitorPopSum[ipAddressString] = pop
	} else {
		u.VisitorPopSum[ipAddressString].Count += pop.Count
	}

	if err := delivery.Ack(); err != nil {
		log.Println(err)
	}
}

func (u *Uploader) Wave() {
	go u.perform()
	u.reset()
}

func (u Uploader) perform() {
	for _, pop := range u.VisitorPopSum {
		upload(pop)
	}
}

func (u *Uploader) reset() {
	u.VisitorPopSum = make(map[string]*VisitorPop)
}

func upload(newPop *VisitorPop) {
	go uploadVisitorPop(
		newPop.IPAddress,
		newPop.RegionCode,
		newPop.Count,
	)
	go uploadRegionPop(
		newPop.RegionCode,
		newPop.Count,
	)
}

func uploadVisitorPop(ipAddress VisitorIP, regionCode string, appendCount int64) {
	pop := new(VisitorPop)
	if tx := Database.First(
		pop,
		"ip_address = ? AND region_code = ?",
		ipAddress,
		regionCode,
	); errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		pop = &VisitorPop{
			IPAddress:  ipAddress,
			RegionCode: regionCode,
			Count:      appendCount,
		}
	} else if tx.Error != nil {
		log.Panicln(tx.Error)
	} else {
		pop.Count += appendCount
	}

	if tx := Database.Save(pop); tx.Error != nil {
		log.Println(tx.Error)
	}
}

func uploadRegionPop(regionCode string, appendCount int64) {
	pop := new(RegionPop)
	if tx := Database.First(
		pop,
		"region_code = ?",
		regionCode,
	); errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		pop = &RegionPop{
			RegionCode: regionCode,
			Count:      appendCount,
		}
	} else if tx.Error != nil {
		log.Panicln(tx.Error)
	} else {
		pop.Count += appendCount
	}

	if tx := Database.Save(pop); tx.Error != nil {
		log.Println(tx.Error)
	}
}