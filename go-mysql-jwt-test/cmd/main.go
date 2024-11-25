package main

import (
	"log"
	"module/cmd/api"
)

func main() {
	server := api.NewAPISever(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
