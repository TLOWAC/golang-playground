package main

import (
	"fmt"
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type EnvConfig struct {
	ServerPort string `env:"SERVER_PORT,required"`
	DBHost     string `env:"DB_HOST,required"`
	DBPort     string `env:"DB_PORT,required"`
	DBUser     string `env:"DB_USER,required"`
	DBPassword string `env:"DB_PASSWORD,required"`
}

func NewEnvConfig() *EnvConfig {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Unable to load .env: %e", err)
	}

	config := &EnvConfig{}

	if err := env.Parse(config); err != nil {
		log.Fatalf("Unable to load variables from .env: %e", err)
	}

	return config
}

func main() {
	config := NewEnvConfig()
	fmt.Println("config.DBHost %+v", config.DBHost)
	fmt.Println("config.DBPort %+v", config.DBPort)
}
