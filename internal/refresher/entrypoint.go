package refresher

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
	ep, err := entrypoint.NewEntryPoint("refresher", addr, handler)
	if err == nil {
		ep.Start()
	}
	return
}

func getRouter() *router.Router {
	routeManager := ruler.NewRouter()

	routeManager.Path("/metrics").Method("GET").
		Handler(promhttp.Handler())

	routeManager.
		Path("/torrent/{id}").Method("GET").
		HandlerFunc(UpdatePeers("UpdatePeers"))

	newRouter := router.NewRouter(routeManager)
	newRouter.Middlewares.Use(middlewares.Context(0)) // Add Context to support path parameters
	newRouter.Middlewares.Use(middlewares.Logger(1), middlewares.Recover(2))
	return newRouter
}
