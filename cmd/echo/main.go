package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/supersonictw/popcat-echo/internal/config"
	"github.com/supersonictw/popcat-echo/internal/leaderboard"
	"github.com/supersonictw/popcat-echo/internal/pop"
	"log"
)

func main() {
	fmt.Println("PopCat Echo")
	fmt.Println("===")
	fmt.Println("The server reproduce of https://popcat.click with improvement.")
	fmt.Println("License: MIT LICENSE")
	fmt.Println("Repository: https://github.com/supersonictw/popcat-echo")
	fmt.Println("(c) 2021 SuperSonic. https://github.com/supersonictw")
	fmt.Println()

	go pop.Queue()

	r := gin.Default()
	r.GET("/leaderboard", leaderboard.Response)
	r.POST("/pop", pop.Response)

	err := r.Run(config.PublishAddress)
	if err != nil {
		log.Fatal(err)
	}
}
