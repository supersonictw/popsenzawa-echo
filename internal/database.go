package internal

import (
	"database/sql"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	DB  *sql.DB
	RDB *redis.Client
)

func init() {
	DB, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}
	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
