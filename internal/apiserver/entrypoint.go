package apiserver

import (
	_ "embed"
	"encoding/json"
	_ "example.com/m/internal/apiserver/docs"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
	goprom "github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"github.com/xgfone/go-apiserver/entrypoint"
	"github.com/xgfone/go-apiserver/http/router"
	"github.com/xgfone/go-apiserver/http/router/routes/ruler"
	"github.com/xgfone/go-opentelemetry"
	"github.com/xgfone/go-opentelemetry/jaegerexporter"
	"github.com/xgfone/go-opentelemetry/otelhttpx"
	"github.com/xgfone/go-opentelemetry/promexporter"
	"github.com/zitadel/oidc/pkg/client/rp"
	httphelper "github.com/zitadel/oidc/pkg/http"
	"github.com/zitadel/oidc/pkg/oidc"
	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"net/http"
	"os"
	"strings"
	"time"
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

//go:embed docs/swagger.yaml
var swagger []byte

var redirectURI = "/auth/callback"

func getRouter() *router.Router {
	routeManager := ruler.NewRouter()
	routeManager.AddProfileRoutes("/api")

	routeManager.
		Path("/torrents/{id}").Method("GET").
		Handler(otelhttpx.Handler(GetTorrentById("/torrents/{id}"), ""))

	routeManager.
		Path("/torrents/{id}").Method("DELETE").
		Handler(otelhttpx.Handler(DeleteTorrentById("/torrents/{id}"), ""))

	routeManager.
		Path("/torrents").Method("POST").
		Handler(otelhttpx.Handler(PostTorrentById("/torrents"), ""))

	routeManager.
		Path("/search").Method("GET").
		Handler(otelhttpx.Handler(SearchQuery("/search"), ""))

	routeManager.
		Path("/files/{id}").Method("GET").
		Handler(otelhttpx.Handler(GetFilesByInfohash("/files/{id}"), ""))

	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
	routeManager.Path("/docs").Method("GET").
		Handler(middleware.SwaggerUI(opts, nil))

	routeManager.Path("/swagger.yaml").Method("GET").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write(swagger)
		})

	registry := goprom.DefaultRegisterer.(*goprom.Registry)
	routeManager.Path("/metrics").Method("GET").
		Handler(promexporter.Handler(registry))

	routeManager.
		Path("/secureResource").Method("GET").
		Handler(otelhttpx.Handler(SecureResource("/secureResource"), ""))

	/*
		provider, state := authPrepare()
		routeManager.
			Path("/login").Method("GET").
			Handler(otelhttpx.Handler(rp.AuthURLHandler(state, provider), ""))

		routeManager.
			Path(redirectURI).Method("GET").
			Handler(rp.CodeExchangeHandler(rp.UserinfoCallback(marshalUserinfo), provider))

	*/

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

func authPrepare() (rp.RelyingParty, func() string) {
	clientID := "native"
	clientSecret := ""
	keyPath := os.Getenv("KEY_PATH")
	issuer := "http://localhost:9998/"
	scopes := strings.Split("openid profile", " ")
	key := []byte("test1234test1234")

	cookieHandler := httphelper.NewCookieHandler(key, key, httphelper.WithUnsecure())

	options := []rp.Option{
		rp.WithCookieHandler(cookieHandler),
		rp.WithVerifierOpts(rp.WithIssuedAtOffset(5 * time.Second)),
	}
	if clientSecret == "" {
		options = append(options, rp.WithPKCE(cookieHandler))
	}
	if keyPath != "" {
		options = append(options, rp.WithJWTProfile(rp.SignerFromKeyPath(keyPath)))
	}

	provider, err := rp.NewRelyingPartyOIDC(issuer, clientID, clientSecret, "http://localhost:8080"+redirectURI, scopes, options...)
	if err != nil {
		logrus.Fatalf("error creating provider %s", err.Error())
	}

	state := func() string {
		return uuid.New().String()
	}

	return provider, state
}

var marshalUserinfo = func(w http.ResponseWriter, r *http.Request, tokens *oidc.Tokens, state string, rp rp.RelyingParty, info oidc.UserInfo) {
	fmt.Println("State: " + state)
	go getAccessToken(tokens)
	data, err := json.Marshal(info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func getAccessToken(tokens *oidc.Tokens) {
	fmt.Println("IDToken: " + tokens.IDToken)
	fmt.Println("Access Token: " + tokens.AccessToken)
	fmt.Println("Refresh Token: " + tokens.RefreshToken)
	fmt.Println("TokenType: " + tokens.TokenType)
	fmt.Println(tokens.IDTokenClaims)
}
