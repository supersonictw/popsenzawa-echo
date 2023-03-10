// PopCat Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package data

import (
	"log"

	"github.com/oschwald/maxminddb-golang"
	"github.com/spf13/viper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	configMySQLDSN  string
	configIPGeoPath string

	Database      *gorm.DB
	IPGeoDatabase *maxminddb.Reader
)

func init() {
	configMySQLDSN = viper.Get("MYSQL_DSN").(string)
	configIPGeoPath = viper.Get("IP_GEO_PATH").(string)
}

func init() {
	var err error

	Database, err = gorm.Open(mysql.Open(configMySQLDSN), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	IPGeoDatabase, err = maxminddb.Open(configIPGeoPath)
	if err != nil {
		log.Fatal(err)
	}
}

func AutoMigrate() error {
	if err := Database.AutoMigrate(
		&RegionPop{},
		&VisitorPop{},
	); err != nil {
		return err
	}
	return nil
}
