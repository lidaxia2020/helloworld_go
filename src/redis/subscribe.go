package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {

	redisClient()
}

func redisClient() {
	redisdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456",
		DB:       1,
	})

	pubsub := redisdb.Subscribe("channel")
	defer pubsub.Close()
	for msg := range pubsub.Channel() {
		fmt.Println(msg.Payload)
	}
}
