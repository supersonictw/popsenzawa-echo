package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/supersonictw/popcat-echo/internal/config"
	"github.com/supersonictw/popcat-echo/internal/leaderboard"
	"github.com/supersonictw/popcat-echo/internal/pop"
	"log"
	"net/http"
)

const (
	EnvCORSSupport      config.EnvKey = "CORS_SUPPORT"
	EnvFrontendHostname config.EnvKey = "FRONTEND_HOSTNAME"
	EnvFrontendSSL      config.EnvKey = "FRONTEND_SSL"
)

func main() {
	fmt.Println("PopCat Echo")
	fmt.Println("===")
	fmt.Println("The server reproduce of https://popcat.click with improvement.")
	fmt.Println("License: MIT LICENSE")
	fmt.Println("Repository: https://github.com/supersonictw/popcat-echo")
	fmt.Println("(c) 2021 SuperSonic. https://github.com/supersonictw")
	fmt.Println()

	leaderboard.PrepareCache()
	go pop.Queue()

	r := gin.Default()

	if config.Get(EnvCORSSupport) == "yes" {
		var frontendURI string
		if hostname := config.Get(EnvFrontendHostname); config.Get(EnvFrontendSSL) == "yes" {
			frontendURI = fmt.Sprintf("https://%s", hostname)
		} else {
			frontendURI = fmt.Sprintf("http://%s", hostname)
		}
		corsConfig := cors.DefaultConfig()
		corsConfig.AllowOrigins = []string{frontendURI}
		r.Use(cors.New(corsConfig))
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"application": "popcat-echo",
			"copyright":   "(c)2021 SuperSonic(https://github.com/supersonictw)",
		})
	})
	r.GET("/leaderboard", leaderboard.Response)
	r.POST("/pop", pop.Response)

	fmt.Println("Start")
	err := r.Run(config.PublishAddress)
	if err != nil {
		log.Fatal(err)
	}
}
