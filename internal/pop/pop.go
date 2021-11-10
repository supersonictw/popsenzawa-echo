// PopCat Echo
// (c) 2021 SuperSonic (https://github.com/supersonictw).

package pop

import (
	"encoding/json"
	"log"
)

type Pop struct {
	Count   int    `json:"count"`
	Address string `json:"address"`
	Region  string `json:"region"`
}

func NewPop(count int, address, region string) *Pop {
	p := new(Pop)
	p.Count = count
	p.Address = address
	p.Region = region
	return p
}

func (p *Pop) JSON() []byte {
	val, err := json.Marshal(p)
	if err != nil {
		log.Panicln(err)
	}
	return val
}
