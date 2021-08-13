package config

import (
	"log"
	"os"
)

func Get(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Env lookup failed: %s", key)
	}
	return value
}
