package apiserver

import (
	"fmt"
	"github.com/xgfone/go-apiserver/http/reqresp"
	"net/http"
)

func InfohashGet(route string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := reqresp.GetContext(w, r)
		_, _ = fmt.Fprintf(w, "%s: %s", route, c.Data["id"])
	}
}

func InfohashPut(route string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := reqresp.GetContext(w, r)
		_, _ = fmt.Fprintf(w, "%s: %s", route, c.Data["id"])
		infohash := fmt.Sprintf("%s", c.Data["id"])
		success := RedisAdd(infohash)
		if success == 1 {
			// insert into cassandra
		}
	}
}
