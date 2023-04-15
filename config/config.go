package config

import "os"

var RedisEndpoint string = "localhost:6379"

func init() {
	if os.Getenv("REDIS_ENDPOINT") != "" {
		RedisEndpoint = os.Getenv("REDIS_ENDPOINT")
	}
}
