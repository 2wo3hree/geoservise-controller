package router

import (
	handler2 "geoservise-jwt/internal/delivery/handler"
	auth2 "geoservise-jwt/internal/infrastructure/auth"
	custommetrics "geoservise-jwt/internal/infrastructure/middleware"
	"geoservise-jwt/internal/usecase/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

// TokenFromCookie — middleware, который достаёт токен из cookie и добавляет его в Authorization header
func TokenFromCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt")
		if err == nil {
			r.Header.Set("Authorization", "Bearer "+cookie.Value)
		}
		next.ServeHTTP(w, r)
	})
}

func SetupRouter(h *handler2.AddressHandler, userHandler *handler2.UserHandler, userService service.UserService) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(custommetrics.MetricsMiddleware)

	auth2.InitJWT()

	authHandler := auth2.NewAuthHandler(auth2.TokenAuth, userService, h.Responder)

	r.Post("/api/register", userHandler.Create)
	r.Post("/api/login", authHandler.LoginHandler)

	r.Group(func(r chi.Router) {
		r.Use(TokenFromCookie)
		r.Use(jwtauth.Verifier(authHandler.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/api/address/search", h.Search)
		r.Post("/api/address/geocode", h.Geocode)

	})

	// Swagger UI
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	r.Mount("/mycustompath/pprof", PprofRouter(authHandler.TokenAuth))

	r.Handle("/metrics", promhttp.Handler())

	return r
}
