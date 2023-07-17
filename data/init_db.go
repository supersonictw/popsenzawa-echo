// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package data

import (
	"log"
	"net"

	"github.com/oschwald/maxminddb-golang"
	"github.com/spf13/viper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	configMySQLDSN  = viper.GetString("mysql.dsn")
	configIPGeoPath = viper.GetString("geolite2.database_path")
)

var (
	Database      *gorm.DB
	IPGeoDatabase *maxminddb.Reader
)

func init() {
	var err error

	mysqlConnection := mysql.Open(configMySQLDSN)
	Database, err = gorm.Open(mysqlConnection, &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	IPGeoDatabase, err = maxminddb.Open(configIPGeoPath)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	if err := Database.AutoMigrate(
		&RegionPop{},
		&VisitorPop{},
	); err != nil {
		log.Fatalln(err)
	}
}

func FindRegionCodeByIPAddress(ip net.IP) (string, error) {
	var record struct {
		Country struct {
			ISOCode string `maxminddb:"iso_code"`
		} `maxminddb:"country"`
	}

	err := IPGeoDatabase.Lookup(ip, &record)
	if err != nil {
		return "", err
	}

	if record.Country.ISOCode == "" {
		return "Unknown", nil
	}

	return record.Country.ISOCode, nil
}
