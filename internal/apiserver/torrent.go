package apiserver

import (
	"encoding/json"
	"errors"
	"example.com/m/internal/pkg"
	"example.com/m/internal/pkg/cassandra"
	"fmt"
	"github.com/xgfone/go-apiserver/http/reqresp"
	"io/ioutil"
	"net/http"
)

func PostTorrentById(route string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := reqresp.GetContext(w, r)
		body, err := ioutil.ReadAll(c.Body)
		if err != nil {
			w.WriteHeader(502)
			_, _ = fmt.Fprintf(w, "%s", pkg.JsonError(err))
			return
		}

		var queue cassandra.Queue
		err = json.Unmarshal(body, &queue)
		if err != nil || queue.InfoHash == "" {
			w.WriteHeader(400)
			_, _ = fmt.Fprintf(w, "%s", pkg.JsonError(errors.New("bad request json")))
			return
		}

		w.WriteHeader(200)
		go addTorrent(queue.InfoHash)
		_, _ = fmt.Fprintf(w, "successfully processed request")
	}
}

func GetTorrentById(route string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := reqresp.GetContext(w, r)
		id := fmt.Sprintf("%s", c.Data["id"])

		torrent, err := cassandra.FindTorrentByInfohash(id)
		if err != nil {
			w.WriteHeader(404)
			_, _ = fmt.Fprintf(w, "%s", pkg.JsonError(err))
			return
		}

		// https://pkg.go.dev/encoding/json#Marshal
		// < > & will be escaped in magnet url
		str, err := json.MarshalIndent(torrent, "", "    ")
		if err != nil {
			w.WriteHeader(502)
			_, _ = fmt.Fprintf(w, "%s", pkg.JsonError(err))
			return
		}
		_, _ = fmt.Fprintf(w, "%s", str)
		//_, _ = fmt.Fprintf(w, "%s: %s", route, infohash
	}
}

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
		fmt.Println("Sending to refresher")
		go func() {
			_, _ = http.Get(RefresherURL + "/torrent/" + infohash)
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
