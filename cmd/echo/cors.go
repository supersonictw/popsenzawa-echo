// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	allowOrigins = viper.GetStringSlice("cors.allow_origins")
)

func getCORS() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = allowOrigins
	return cors.New(corsConfig)
}
