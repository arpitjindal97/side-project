package main

import (
	"example.com/m/internal/pkg/cassandra"
	"example.com/m/internal/refresher"
	"log"
)

func main() {

	cassandra.Conn = cassandra.Cluster{
		URL:      []string{"vergon-cassandra-0"},
		KeySpace: "awesome",
		Session:  nil,
		Username: "cassandra",
		Password: "vergon",
	}
	cassandra.Init()
	defer cassandra.Conn.Session.Close()

	err := refresher.StartHTTPServer("0.0.0.0:8081")
	log.Printf("indexer server shutdown: %s", err)
}
