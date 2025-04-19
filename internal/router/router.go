package router

import (
	"geoservise-jwt/internal/auth"
	"geoservise-jwt/internal/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/swaggo/http-swagger"
)

func SetupRouter(h *handler.AddressHandler, tokenAuth *jwtauth.JWTAuth) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

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

	return r
}
