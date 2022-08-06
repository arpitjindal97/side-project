package apiserver

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/xgfone/go-apiserver/entrypoint"
	"github.com/xgfone/go-apiserver/http/middlewares"
	"github.com/xgfone/go-apiserver/http/router"
	"github.com/xgfone/go-apiserver/http/router/routes/ruler"
)

// StartHTTPServer is a simple convenient function to start a http server.
func StartHTTPServer(addr string) (err error) {
	handler := getRouter()
	ep, err := entrypoint.NewEntryPoint("apiserver", addr, handler)
	if err == nil {
		ep.Start()
	}
	return
}

func getRouter() *router.Router {
	routeManager := ruler.NewRouter()

	routeManager.
		Path("/torrents/{id}").Method("GET").
		HandlerFunc(GetTorrentById("GetTorrentById"))

	routeManager.
		Path("/torrents").Method("POST").
		HandlerFunc(PostTorrentById("PostTorrentById"))

	routeManager.
		Path("/search/{query}").Method("GET").
		HandlerFunc(SearchQuery("SearchQuery"))

	/*
		routeManager.
			Path("/torrents/search/findByInfoHashEquals").Method("GET").
			HandlerFunc(storegateway.TorrentFindByInfoHashEquals("torrent_findByInfoHashEquals"))
	*/

	routeManager.Path("/metrics").Method("GET").
		Handler(promhttp.Handler())

	router := router.NewRouter(routeManager)
	router.Middlewares.Use(middlewares.Context(0)) // Add Context to support path parameters
	router.Middlewares.Use(middlewares.Logger(1), middlewares.Recover(2))
	return router
}
