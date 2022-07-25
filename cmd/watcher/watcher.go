package main

import (
	"example.com/m/internal/apiserver"
	"example.com/m/internal/watcher"
	"github.com/go-redis/redis/v9"
	"strconv"
	"time"
)

func main() {
	watcher.JobStartTime = strconv.FormatInt(time.Now().Unix(), 10)

	apiserver.Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	apiserver.RedisKey = "apiserver"
	watcher.StartJob()
}
