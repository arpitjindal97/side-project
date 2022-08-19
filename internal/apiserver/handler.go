package apiserver

import (
	"encoding/json"
	"errors"
	"example.com/m/internal/pkg/cassandra"
	"example.com/m/internal/pkg/elasticsearch"
	"example.com/m/internal/pkg/utils"
	"fmt"
	"github.com/xgfone/go-apiserver/http/reqresp"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func PostTorrentById(route string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		labeler, _ := otelhttp.LabelerFromContext(r.Context())
		labeler.Add(semconv.HTTPRouteKey.String(route))
		c := reqresp.GetContext(w, r)
		w.Header().Set("Content-Type", "application/json")
		body, err := ioutil.ReadAll(c.Body)
		if err != nil {
			w.WriteHeader(502)
			labeler.Add(semconv.HTTPStatusCodeKey.Int(502))
			_, _ = fmt.Fprintf(w, "%s", utils.JsonError(err))
			return
		}

		defer func() {
			if nil != any(recover()) {
				w.WriteHeader(500)
				labeler.Add(semconv.HTTPStatusCodeKey.Int(500))
				_, _ = fmt.Fprintf(w, "%s", utils.JsonError(errors.New("internal server error")))
			}
		}()

		var queue cassandra.Queue
		err = json.Unmarshal(body, &queue)
		if err != nil || queue.InfoHash == "" {
			w.WriteHeader(400)
			labeler.Add(semconv.HTTPStatusCodeKey.Int(400))
			_, _ = fmt.Fprintf(w, "%s", utils.JsonError(errors.New("bad request json")))
			return
		}

		w.WriteHeader(200)
		go addTorrent(strings.ToLower(queue.InfoHash))
		_, _ = fmt.Fprintf(w, utils.JsonMessage("Successfully received submission"))
		labeler.Add(semconv.HTTPStatusCodeKey.Int(200))
	}
}

func GetTorrentById(route string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		labeler, _ := otelhttp.LabelerFromContext(r.Context())
		labeler.Add(semconv.HTTPRouteKey.String(route))
		c := reqresp.GetContext(w, r)
		id := strings.ToLower(fmt.Sprintf("%s", c.Data["id"]))
		defer func() {
			if nil != any(recover()) {
				w.WriteHeader(500)
				labeler.Add(semconv.HTTPStatusCodeKey.Int(500))
				_, _ = fmt.Fprintf(w, "%s", utils.JsonError(errors.New("internal server error")))
			}
		}()

		w.Header().Set("Content-Type", "application/json")
		torrent, err := cassandra.FindTorrentByInfohash(id)
		if err != nil {
			w.WriteHeader(404)
			labeler.Add(semconv.HTTPStatusCodeKey.Int(404))
			_, _ = fmt.Fprintf(w, "%s", utils.JsonError(err))
			return
		}

		// https://pkg.go.dev/encoding/json#Marshal
		// < > & will be escaped in magnet url
		str, err := json.MarshalIndent(torrent, "", "    ")
		if err != nil {
			w.WriteHeader(502)
			labeler.Add(semconv.HTTPStatusCodeKey.Int(502))
			_, _ = fmt.Fprintf(w, "%s", utils.JsonError(err))
			return
		}
		_, _ = fmt.Fprintf(w, "%s", str)
		//_, _ = fmt.Fprintf(w, "%s: %s", route, infohash
		labeler.Add(semconv.HTTPStatusCodeKey.Int(200))
	}
}

func SearchQuery(route string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		labeler, _ := otelhttp.LabelerFromContext(r.Context())
		labeler.Add(semconv.HTTPRouteKey.String(route))
		c := reqresp.GetContext(w, r)
		defer func() {
			if nil != any(recover()) {
				w.WriteHeader(500)
				labeler.Add(semconv.HTTPStatusCodeKey.Int(500))
				_, _ = fmt.Fprintf(w, "%s", utils.JsonError(errors.New("internal server error")))
			}
		}()
		query := c.GetQuery("q")
		sort := c.GetQuery("sort")
		size, err := strconv.Atoi(c.GetQuery("size"))
		if err != nil {
			size = 10
		}
		from, _ := strconv.Atoi(c.GetQuery("from"))
		if err != nil {
			from = 0
		}
		data, _ := elasticsearch.Search(query, sort, size, from)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		_, _ = w.Write(data)
		labeler.Add(semconv.HTTPStatusCodeKey.Int(200))
	}
}

func GetFilesByInfohash(route string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		labeler, _ := otelhttp.LabelerFromContext(r.Context())
		labeler.Add(semconv.HTTPRouteKey.String(route))
		c := reqresp.GetContext(w, r)
		defer func() {
			if nil != any(recover()) {
				w.WriteHeader(500)
				labeler.Add(semconv.HTTPStatusCodeKey.Int(500))
				_, _ = fmt.Fprintf(w, "%s", utils.JsonError(errors.New("internal server error")))
			}
		}()
		id := strings.ToLower(fmt.Sprintf("%s", c.Data["id"]))
		files, err := cassandra.FindFilesByInfohash(id)
		if err != nil {
			w.WriteHeader(404)
			labeler.Add(semconv.HTTPStatusCodeKey.Int(404))
			_, _ = fmt.Fprintf(w, "%s", utils.JsonError(err))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		data, _ := json.MarshalIndent(files, "", "    ")
		_, _ = w.Write(data)
		labeler.Add(semconv.HTTPStatusCodeKey.Int(200))
	}
}

func DeleteTorrentById(route string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		labeler, _ := otelhttp.LabelerFromContext(r.Context())
		labeler.Add(semconv.HTTPRouteKey.String(route))
		c := reqresp.GetContext(w, r)
		defer func() {
			if nil != any(recover()) {
				w.WriteHeader(500)
				labeler.Add(semconv.HTTPStatusCodeKey.Int(500))
				_, _ = fmt.Fprintf(w, "%s", utils.JsonError(errors.New("internal server error")))
			}
		}()
		id := strings.ToLower(fmt.Sprintf("%s", c.Data["id"]))
		_ = cassandra.DeleteTorrentById(id)
		_ = elasticsearch.Delete(id)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		_, _ = fmt.Fprintf(w, "%s", utils.JsonMessage("Successfully deleted "+id))
		labeler.Add(semconv.HTTPStatusCodeKey.Int(200))
	}
}
