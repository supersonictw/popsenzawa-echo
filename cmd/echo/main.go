// PopCat Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/supersonictw/popcat-echo/leaderboard"
	"github.com/supersonictw/popcat-echo/pop"
)

var (
	configPublishAddress string
)

func init() {

}

func init() {
	configPublishAddress = viper.Get("PUBLISH_ADDRESS").(string)
}

func main() {
	fmt.Println("PopCat Echo")
	fmt.Println("===")
	fmt.Println("The server reproduce of https://popcat.click with improvement.")
	fmt.Println("License: MIT LICENSE")
	fmt.Println("Repository: https://github.com/supersonictw/popcat-echo")
	fmt.Println("(c) 2023 SuperSonic. https://github.com/supersonictw")
	fmt.Println()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"application": "popcat-echo",
			"copyright":   "(c)2023 SuperSonic(https://github.com/supersonictw)",
		})
	})
	r.GET("/leaderboard", leaderboard.GetLeaderboard)
	r.POST("/pops", pop.PostPops)

	fmt.Println("Start Echo Server")
	if err := r.Run(configPublishAddress); err != nil {
		log.Fatal(err)
	}
}
