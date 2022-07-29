package main

import (
	"example.com/m/internal/apiserver"
	"example.com/m/internal/pkg/cassandra"
	"example.com/m/internal/watchman"
	"github.com/go-redis/redis/v9"
	"strconv"
	"time"
)

func main() {
	watchman.JobStartTime = strconv.FormatInt(time.Now().Unix(), 10)

	apiserver.Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	apiserver.RedisKey = "apiserver"

	cassandra.Conn = cassandra.Cluster{
		URL:      []string{"localhost"},
		KeySpace: "awesome",
		Session:  nil,
	}
	cassandra.Init()
	defer cassandra.Conn.Session.Close()

	watchman.StartJob()
}
