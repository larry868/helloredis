package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Testing Golang Redis")

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env variables: %s", err)
	}

	redis_host := os.Getenv("REDIS_HOST")
	redis_port := os.Getenv("REDIS_PORT")
	redis_addr := redis_host + ":" + redis_port
	redis_password := os.Getenv("REDIS_PASSWORD")

	client := redis.NewClient(&redis.Options{
		Addr:     redis_addr,
		Password: redis_password,
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

}
