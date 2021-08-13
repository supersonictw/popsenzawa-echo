package main

import (
	"github.com/gin-gonic/gin"
	"github.com/supersonictw/popcat-echo/internal"
	"github.com/supersonictw/popcat-echo/internal/leaderboard"
	"github.com/supersonictw/popcat-echo/internal/pop"
	"log"
)

func main() {
	go pop.Queue()

	r := gin.Default()
	r.POST("/pop", pop.Response)
	r.POST("/leaderboard", leaderboard.Response)

	err := r.Run(internal.PublishAddress)
	if err != nil {
		log.Fatal(err)
	}
}
