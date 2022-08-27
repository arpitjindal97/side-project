package apiserver

import (
	"example.com/m/internal/pkg/cassandra"
	"net/http"
)

var RefresherURL string

func addTorrent(infohash string) {
	if RedisGet(infohash) == nil {
		// already in cache
		return
	}
	RedisSet(infohash)
	_, err := cassandra.FindTorrentByInfohash(infohash)
	if err == nil {
		// already indexed
		// send to refresher
		go func() {
			resp, _ := http.Get(RefresherURL + "/torrent/" + infohash)
			_ = resp.Body.Close()
		}()
	} else {
		// check if it's already in queue
		_, err = cassandra.FindQueueByInfohash(infohash)
		if err != nil {
			// was not present in queue
			_ = cassandra.InsertQueueByInfohash(infohash)
		}
	}
}
