// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package data

import (
	"log"
	"net"

	"github.com/oschwald/maxminddb-golang"
	"github.com/spf13/viper"
)

var (
	configIPGeoPath = viper.GetString("ipgeo.database_path")
)

var (
	ipGeoDatabase *maxminddb.Reader
)

func init() {
	var err error
	ipGeoDatabase, err = maxminddb.Open(configIPGeoPath)
	if err != nil {
		log.Fatal(err)
	}
}

func findRegionCodeByIPAddress(ip net.IP) (string, error) {
	var record struct {
		Country struct {
			ISOCode string `maxminddb:"iso_code"`
		} `maxminddb:"country"`
	}

	err := ipGeoDatabase.Lookup(ip, &record)
	if err != nil {
		return "", err
	}

	if record.Country.ISOCode == "" {
		return "Unknown", nil
	}

	return record.Country.ISOCode, nil
}
