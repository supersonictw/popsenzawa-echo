package pop

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
)

func Queue() {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := context.Background()

	for {
		timestamp := time.Now().Unix()
		stepTimestamp := timestamp / 100 * 100

		val, err := rdb.Get(ctx, "key").Result()
		if err != nil {
			panic(err)
		}
		
	}
}
