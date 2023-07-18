// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	configAllowOrigins = viper.GetStringSlice("cors.allow_origins")
)

func getCORS() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = configAllowOrigins
	return cors.New(corsConfig)
}
