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
		//_, _ = fmt.Fprintf(w, "%s: %s", route, c.Data["id"])
		//infohash := fmt.Sprintf("%s", c.Data["infohash"])

		body, err := ioutil.ReadAll(c.Body)
		if err != nil {
			_, _ = fmt.Fprintf(w, "%s", pkg.JsonError(err))
			return
		}

		var queue cassandra.Queue
		err = json.Unmarshal(body, &queue)
		if err != nil {
			_, _ = fmt.Fprintf(w, "%s", pkg.JsonError(err))
			return
		}
		if queue.InfoHash == "" {
			w.WriteHeader(400)
			_, _ = fmt.Fprintf(w, "%s", pkg.JsonError(errors.New("bad request json")))
			return
		}

		_ = RedisAdd(queue.InfoHash)
		_, _ = fmt.Fprintf(w, "successfully processed request")
	}
}

func InfohashPost(route string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := reqresp.GetContext(w, r)

		_, _ = fmt.Fprintf(w, "%s: %s", route, c.Body)
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

		str, err := json.Marshal(torrent)
		if err != nil {
			w.WriteHeader(502)
			_, _ = fmt.Fprintf(w, "%s", pkg.JsonError(err))
			return
		}
		_, _ = fmt.Fprintf(w, "%s", str)
		//_, _ = fmt.Fprintf(w, "%s: %s", route, infohash
	}
}
