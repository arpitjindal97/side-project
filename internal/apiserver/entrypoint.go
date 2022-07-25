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
		Path("/infohash/{id}").Method("GET").
		HandlerFunc(InfohashGet("infohash_get"))

	routeManager.
		Path("/infohash/{id}").Method("PUT").
		HandlerFunc(InfohashPut("infohash_put"))

	routeManager.Path("/metrics").Method("GET").
		Handler(promhttp.Handler())

	router := router.NewRouter(routeManager)
	router.Middlewares.Use(middlewares.Context(0)) // Add Context to support path parameters
	router.Middlewares.Use(middlewares.Logger(1), middlewares.Recover(2))
	return router
}
