package main

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("redis connection error : ", err)
	}

	log.Println(ping)

	err = client.Set(context.Background(), "name", "홍길동", 0).Err()
	if err != nil {
		log.Fatal("redis can not set the value : ", err)
	}

	name, err := client.Get(context.Background(), "name").Result()
	if err != nil {
		log.Fatal("redis can not found the key value : ", err)
	}

	log.Printf("result %s", name)
}
