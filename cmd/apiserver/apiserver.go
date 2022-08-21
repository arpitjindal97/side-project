package main

import (
	"example.com/m/internal/apiserver"
	"example.com/m/internal/pkg/cassandra"
	"example.com/m/internal/pkg/elasticsearch"
	"github.com/go-redis/redis/v9"
	"log"
)

func main() {
	// Redis Setup
	log.Println("Initializing Redis")
	apiserver.Rdb = redis.NewClient(&redis.Options{
		Addr:     "vergon-redis-master:6379",
		Password: "bhXvm2p7Xj", // no password set
		DB:       0,            // use default DB
	})
	defer apiserver.Rdb.Close()

	// Cassandra Setup
	cassandra.Init(
		[]string{"vergon-cassandra:9042"},
		"cassandra",
		"vergon",
		"awesome",
	)
	defer cassandra.Session.Close()

	// ElasticSearch Setup
	elasticsearch.Init(
		[]string{"http://vergon-elasticsearch:9200", "http://localhost:9200"},
		"elastic",
		"password",
		"torrents",
	)

	apiserver.RefresherURL = "http://refresher:8081"

	err := apiserver.StartHTTPServer("0.0.0.0:8080")
	log.Printf("http server shutdown: %s", err)
}
