package apiserver

import (
	"context"
	goprom "github.com/prometheus/client_golang/prometheus"
	"github.com/xgfone/go-apiserver/entrypoint"
	"github.com/xgfone/go-apiserver/http/router"
	"github.com/xgfone/go-apiserver/http/router/routes/ruler"
	"github.com/xgfone/go-opentelemetry"
	"github.com/xgfone/go-opentelemetry/jaegerexporter"
	"github.com/xgfone/go-opentelemetry/otelhttpx"
	"github.com/xgfone/go-opentelemetry/promexporter"
	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/metric/instrument/asyncint64"
	"time"
)

var activeReq asyncint64.UpDownCounter
var activeCount int64

// StartHTTPServer is a simple convenient function to start a http server.
func StartHTTPServer(addr string) (err error) {

	activeReq, _ = global.MeterProvider().
		Meter("go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp").
		AsyncInt64().UpDownCounter("http.server.active_requests")
	activeReq.Observe(context.Background(), activeCount)
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
		Handler(otelhttpx.Handler(GetTorrentById("/torrents/{id}"), ""))

	routeManager.
		Path("/torrents").Method("POST").
		Handler(otelhttpx.Handler(PostTorrentById("/torrents"), ""))

	routeManager.
		Path("/search").Method("GET").
		Handler(otelhttpx.Handler(SearchQuery("/search"), ""))

	registry := goprom.DefaultRegisterer.(*goprom.Registry)
	routeManager.Path("/metrics").Method("GET").
		Handler(promexporter.Handler(registry))

	opentelemetry.SetServiceName("apiserver")
	jaegerexporter.Install(nil, nil)
	promexporter.Install(registry)
	otelhttpx.InstallClient()
	runtime.Start(runtime.WithMinimumReadMemStatsInterval(time.Second))

	newDefaultRouter := router.NewDefaultRouter(routeManager)
	//router.Middlewares.Use(middlewares.Context(0)) // Add Context to support path parameters
	//router.Middlewares.Use(middlewares.Logger(1), middlewares.Recover(2))
	return newDefaultRouter
}
