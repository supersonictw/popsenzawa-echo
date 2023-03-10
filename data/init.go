// PopCat Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package data

import (
	"log"
	"os"

	"github.com/oschwald/maxminddb-golang"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	envMySQLDSN  string
	envIPGeoPath string

	Database      *gorm.DB
	IPGeoDatabase *maxminddb.Reader
)

func init() {
	envMySQLDSN = os.Getenv("MYSQL_DSN")
	envIPGeoPath = os.Getenv("IP_GEO_PATH")
}

func init() {
	var err error

	Database, err = gorm.Open(mysql.Open(envMySQLDSN), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	IPGeoDatabase, err = maxminddb.Open(envIPGeoPath)
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
