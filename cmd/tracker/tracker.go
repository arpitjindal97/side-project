package main

import (
	"example.com/m/internal/tracker"
	"github.com/go-redis/redis/v9"
	"log"
	"net"
)

func main() {

	tracker.OtherTrackers = []string{
		"udp://tracker.openbittorrent.com:6969/announce",
		"udp://tracker.coppersurfer.tk:6969/announce",
		"udp://9.rarbg.to:2710/announce",
		"udp://tracker.opentrackr.org:1337/announce",
	}

	// Redis Setup
	log.Println("Initializing Redis")
	tracker.Rdb = redis.NewClient(&redis.Options{
		Addr:     "vergon-redis-master:6379",
		Password: "bhXvm2p7Xj", // no password set
		DB:       1,            // use default DB
	})
	defer tracker.Rdb.Close()

	tracker.APIServerURL = "http://apiserver:8080"

	packetConn, err := net.ListenPacket("udp4", ":1337")
	if err != nil {
		log.Fatal(err)
	}
	tracker.UDPTrackerServer(packetConn)

}
