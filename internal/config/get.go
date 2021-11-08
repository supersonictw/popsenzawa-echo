// PopCat Echo
// (c) 2021 SuperSonic (https://github.com/supersonictw).

package config

import (
	"log"
	"os"
)

func get(key EnvKey) string {
	value, exists := os.LookupEnv(string(key))
	if !exists {
		log.Fatalf("Env lookup failed: %s", key)
	}
	return value
}
