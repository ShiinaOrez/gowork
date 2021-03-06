package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"os"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("Elite_RedisServer"),
		Password: os.Getenv("Elite_RedisPassword"),
		DB:       0,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}
