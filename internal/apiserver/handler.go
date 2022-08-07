package apiserver

import (
	"encoding/json"
	"errors"
	"example.com/m/internal/pkg"
	"example.com/m/internal/pkg/cassandra"
	"example.com/m/internal/pkg/elasticsearch"
	"fmt"
	"github.com/xgfone/go-apiserver/http/reqresp"
	"io/ioutil"
	"net/http"
)

func PostTorrentById(route string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := reqresp.GetContext(w, r)
		w.Header().Set("Content-Type", "application/json")
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
		_, _ = fmt.Fprintf(w, pkg.JsonMessage("Successfully received submission"))
	}
}

func GetTorrentById(route string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := reqresp.GetContext(w, r)
		id := fmt.Sprintf("%s", c.Data["id"])

		w.Header().Set("Content-Type", "application/json")
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

func SearchQuery(route string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := reqresp.GetContext(w, r)
		query := c.GetQuery("q")
		data, _ := elasticsearch.Search(query)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(data)
	}
}
