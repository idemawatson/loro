package config

import (
	"log"
	"os"
)

func GetEnv(key string) string {
	v := os.Getenv(key)
	log.Printf("key: %v, value: %v\n", key, v)
	return v
}
