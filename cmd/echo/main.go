// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package main

import (
	_ "github.com/supersonictw/popsenzawa-echo/config"

	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/supersonictw/popsenzawa-echo/leaderboard"
	"github.com/supersonictw/popsenzawa-echo/pop"
)

var (
	configServerAddress = viper.GetString("server.address")
)

func main() {
	fmt.Println("PopSenzawa Echo")
	fmt.Println("===")
	fmt.Println("The server reproduce of https://popcat.click with improvement.")
	fmt.Println("License: MIT LICENSE")
	fmt.Println("Repository: https://github.com/supersonictw/popsenzawa-echo")
	fmt.Println("(c) 2023 SuperSonic. https://github.com/supersonictw")
	fmt.Println()

	r := gin.Default()
	if len(configAllowOrigins) > 0 {
		cors := getCORS()
		r.Use(cors)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"application": "popsenzawa-echo",
			"copyright":   "(c)2023 SuperSonic(https://github.com/supersonictw)",
		})
	})

	r.GET("/leaderboard",
		leaderboard.GetLeaderboard,
	)

	r.POST("/pops",
		pop.MiddlewareParseJwt,
		pop.MiddlewareCheckRecaptcha,
		pop.MiddlewareCheckRateLimit,
		pop.MiddlewareParseCount,
		pop.PostPops,
	)

	log.Println("echo-server startup successfully!")
	if err := r.Run(configServerAddress); err != nil {
		log.Fatal(err)
	}
}
