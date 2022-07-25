package main

import (
	"example.com/m/internal/apiserver"
	"github.com/go-redis/redis/v9"
	"log"
)

func main() {

	apiserver.Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	apiserver.RedisKey = "apiserver"

	err := apiserver.StartHTTPServer("0.0.0.0:8081")
	log.Printf("http server shutdown: %s", err)
}
