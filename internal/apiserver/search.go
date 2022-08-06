package apiserver

import (
	"example.com/m/internal/pkg/elasticsearch"
	"fmt"
	"net/http"

	"github.com/xgfone/go-apiserver/http/reqresp"
)

func SearchQuery(route string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := reqresp.GetContext(w, r)
		query := fmt.Sprintf("%s", c.Data["query"])
		elasticsearch.Search(query)

	}
}
