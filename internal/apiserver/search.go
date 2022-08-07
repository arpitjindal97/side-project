package apiserver

import (
	"example.com/m/internal/pkg/elasticsearch"
	"github.com/xgfone/go-apiserver/http/reqresp"
	"net/http"
)

func SearchQuery(route string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := reqresp.GetContext(w, r)
		query := c.GetQuery("q")
		data, _ := elasticsearch.Search(query)
		_, _ = w.Write(data)
	}
}
