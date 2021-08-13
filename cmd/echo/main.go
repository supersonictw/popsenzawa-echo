package main

import (
	"github.com/gin-gonic/gin"
	"github.com/supersonictw/popcat-echo/internal/pop"
	"log"
)

func main() {
	go pop.Queue()

	r := gin.Default()
	r.POST("/pop", pop.Response)

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
