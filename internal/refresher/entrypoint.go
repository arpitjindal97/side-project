package refresher

import (
	goprom "github.com/prometheus/client_golang/prometheus"
	"github.com/xgfone/go-apiserver/entrypoint"
	"github.com/xgfone/go-apiserver/http/router"
	"github.com/xgfone/go-apiserver/http/router/routes/ruler"
	"github.com/xgfone/go-opentelemetry"
	"github.com/xgfone/go-opentelemetry/jaegerexporter"
	"github.com/xgfone/go-opentelemetry/otelhttpx"
	"github.com/xgfone/go-opentelemetry/promexporter"
	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"time"
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

	registry := goprom.DefaultRegisterer.(*goprom.Registry)
	routeManager.Path("/metrics").Method("GET").
		Handler(promexporter.Handler(registry))

	routeManager.
		Path("/torrent/{id}").Method("GET").
		Handler(otelhttpx.Handler(UpdatePeers("/torrent/{id}"), ""))

	opentelemetry.SetServiceName("refresher")
	jaegerexporter.Install(nil, nil)
	promexporter.Install(registry)
	otelhttpx.InstallClient()
	runtime.Start(runtime.WithMinimumReadMemStatsInterval(time.Second))

	newDefaultRouter := router.NewDefaultRouter(routeManager)
	return newDefaultRouter
}
