package router

import (
	"encoding/json"
	"net/http"

	"github.com/rls/gateway-service/pkg/config"
	"github.com/rls/gateway-service/pkg/meta"
	"github.com/rls/gateway-service/pkg/ping"
	"github.com/rls/gateway-service/utils/errors"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rls/gateway-service/middlewares"
	chttp "github.com/rls/gateway-service/svc/http"
)

var router = chi.NewRouter()

type errResponse struct {
	Err *errors.Err `json:"err"`
}

// Route returns the api router
func Route() http.Handler {
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middlewares.ResolveUser)

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; utf8")
		w.WriteHeader(http.StatusNotFound)
		err := errResponse{errors.NewErr(http.StatusNotFound, "Route not found!")}
		json.NewEncoder(w).Encode(err)
	})

	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; utf8")
		w.WriteHeader(http.StatusMethodNotAllowed)
		err := errResponse{errors.NewErr(http.StatusMethodNotAllowed, "Method not allowed!")}
		json.NewEncoder(w).Encode(err)
	})

	registerRoutes()

	return router
}

func registerRoutes() {
	router.Route("/ping/v1", func(r chi.Router) {
		r.Mount("/locations", pingHandler())
	})

	router.Route("/metadata/api/v1", func(r chi.Router) {
		r.Mount("/users", metaHandler())
	})
}

func pingHandler() http.Handler {
	return ping.MakeHandler(ping.NewService(chttp.NewHTTP(config.PingCfg().RequestTimeout)))
}

func metaHandler() http.Handler {
	return meta.MakeHandler(meta.NewService(chttp.NewHTTP(config.MetaCfg().RequestTimeout)))
}
