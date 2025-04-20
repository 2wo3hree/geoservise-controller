package router

import (
	"geoservise-jwt/internal/auth"
	"geoservise-jwt/internal/handler"
	custommetrics "geoservise-jwt/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/swaggo/http-swagger"
)

func SetupRouter(h *handler.AddressHandler, tokenAuth *jwtauth.JWTAuth) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(custommetrics.MetricsMiddleware)

	auth.InitJWT()

	r.Post("/api/register", auth.RegisterHandler)
	r.Post("/api/login", auth.LoginHandler)

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/api/address/search", h.Search)
		r.Post("/api/address/geocode", h.Geocode)
	})

	// Swagger UI
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	r.Mount("/mycustompath/pprof", PprofRouter(tokenAuth))

	r.Handle("/metrics", promhttp.Handler())

	return r
}
