// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package config

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("./config.toml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("config load failed: ", err)
	}
}
