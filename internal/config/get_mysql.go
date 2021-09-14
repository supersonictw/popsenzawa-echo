// PopCat Echo
// (c) 2021 SuperSonic (https://github.com/supersonictw).

package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func GetMySQL() *sql.DB {
	mysql, err := sql.Open("mysql", MysqlDSN)
	if err != nil {
		log.Panicln(err)
	}
	mysql.SetConnMaxLifetime(time.Minute * 3)
	mysql.SetMaxOpenConns(10)
	mysql.SetMaxIdleConns(10)
	return mysql
}
