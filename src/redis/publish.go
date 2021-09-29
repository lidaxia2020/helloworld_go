package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"math/rand"
	"time"
)

func main() {
	redisConnect()
}

func redisConnect() {
	redisdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use default Addr
		Password: "123456",         // no password set
		DB:       1,                // use default DB
	})
	for i := 0; i < 20; i++ {
		time.Sleep(time.Second)
		rand.Seed(time.Now().Unix())
		num := rand.Intn(1000)
		n, _ := redisdb.Publish("channel", num).Result()
		fmt.Println(n)
	}
}
