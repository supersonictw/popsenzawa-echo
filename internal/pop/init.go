package pop

import (
	"database/sql"
	"github.com/go-redis/redis/v8"
	"github.com/supersonictw/popcat-echo/internal/config"
)

var (
	mysqlClient *sql.DB
	redisClient *redis.Client
)

func init() {
	mysqlClient = config.GetMySQL()
	redisClient = config.GetRedis()
}
