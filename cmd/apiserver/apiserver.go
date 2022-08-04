package main

import (
	"example.com/m/internal/apiserver"
	"example.com/m/internal/pkg/cassandra"
	"github.com/go-redis/redis/v9"
	"log"
)

func main() {

	apiserver.Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer apiserver.Rdb.Close()

	cassandra.Conn = cassandra.Cluster{
		URL:      []string{"localhost:9042"},
		KeySpace: "awesome",
		Session:  nil,
	}
	cassandra.Init()
	defer cassandra.Conn.Session.Close()

	err := apiserver.StartHTTPServer("0.0.0.0:8080")
	log.Printf("http server shutdown: %s", err)
}
