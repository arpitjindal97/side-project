package main

import (
	"example.com/m/internal/pkg/cassandra"
	"example.com/m/internal/pkg/elasticsearch"
	"example.com/m/internal/refresher"
	"log"
)

func main() {

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

	err := refresher.StartHTTPServer("0.0.0.0:8081")
	log.Printf("indexer server shutdown: %s", err)
}
