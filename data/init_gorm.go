// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package data

import (
	"log"

	"github.com/spf13/viper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	configGORMMySQLDSN = viper.GetString("gorm.mysql_dsn")
)

var (
	database *gorm.DB
)

func init() {
	var err error

	mysqlConnection := mysql.Open(configGORMMySQLDSN)
	database, err = gorm.Open(mysqlConnection, &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
}

func init() {
	if err := database.AutoMigrate(
		&RegionPop{},
		&VisitorPop{},
	); err != nil {
		log.Fatalln(err)
	}
}
