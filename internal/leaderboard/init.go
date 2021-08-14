package leaderboard

import (
	"database/sql"
	"github.com/supersonictw/popcat-echo/internal/config"
)

var (
	mysqlClient *sql.DB
)

func init() {
	mysqlClient = config.GetMySQL()
}
