package main

import (
	"example.com/m/internal/apiserver"
	"example.com/m/internal/pkg/cassandra"
	"github.com/go-redis/redis/v9"
	"log"
)

func main() {

	apiserver.Rdb = redis.NewClient(&redis.Options{
		Addr:     "vergon-redis-master:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer apiserver.Rdb.Close()

	cassandra.Conn = cassandra.Cluster{
		URL:      []string{"vergon-cassandra:9042"},
		KeySpace: "awesome",
		Session:  nil,
		Username: "cassandra",
		Password: "vergon",
	}
	cassandra.Init()
	defer cassandra.Conn.Session.Close()

	apiserver.RefresherURL = "http://refresher:8081"

	err := apiserver.StartHTTPServer("0.0.0.0:8080")
	log.Printf("http server shutdown: %s", err)
}
