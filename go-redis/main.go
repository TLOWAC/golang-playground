package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	type Person struct {
		Name  string `json:name`
		Age   int    `json:age`
		Email string `json:email`
	}
	personID := uuid.NewString()
	personKey := fmt.Sprintf("person:%s", personID)

	jsonString, err := json.Marshal(&Person{
		Name:  "홍길동",
		Age:   20,
		Email: "example.com",
	})

	if err != nil {
		log.Fatal("failed to marshal ", err)
	}

	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("redis connection error : ", err)
	}

	log.Println(ping)

	err = client.Set(context.Background(), personKey, jsonString, 0).Err()
	if err != nil {
		log.Fatal("redis can not set the value : ", err)
	}

	person, err := client.Get(context.Background(), personKey).Result()
	if err != nil {
		log.Fatal("redis can not found the key value : ", err)
	}

	log.Printf("result %s", person)
}
