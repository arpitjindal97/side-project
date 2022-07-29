package cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/xgfone/go-apiserver/http/reqresp"
	"io/ioutil"
	"net/http"
	"time"
)

func FindTorrentByInfohash(id string) (Torrent, error) {
	var torrent Torrent
	err := Conn.Session.Query(find_torrent_by_infohash, id).Consistency(gocql.One).Scan(&torrent)
	return torrent, err
}

func InsertQueueByInfohash(infohash string) error {
	return Conn.Session.Query(insert_queue_by_infohash, infohash, time.Now(), 0).Exec()
}

func FindQueueByInfohash(id string) (queue Queue, err error) {
	err = Conn.Session.Query(find_queue_by_infohash, id).Consistency(gocql.One).Scan(&queue)
	return
}

func TorrentPost(route string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := reqresp.GetContext(w, r)

		//str := fmt.Sprintf("%s", c.Body)

		_, err := ioutil.ReadAll(c.Body)
		if err != nil {
			_, _ = fmt.Fprintf(w, "%s", err)
			return
		}

		/*
			var torrent Torrent
			err := Conn.Session.Query(find_torrent_by_id, id).Consistency(gocql.One).Scan(&torrent)

			err := json.Unmarshal(body, &torrent)
			if err != nil {
				_, _ = fmt.Fprintf(w, "%s", err)
				return
			}
			_, _ = fmt.Fprintf(w, "%s", str)
		*/
		//_, _ = fmt.Fprintf(w, "%s: %s", route, infohash
	}
}
