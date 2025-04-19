package router

import (
	"net/http"
	"net/http/pprof"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
)

func PprofRouter(authMiddleware *jwtauth.JWTAuth) http.Handler {
	r := chi.NewRouter()
	r.Use(jwtauth.Verifier(authMiddleware))
	r.Use(jwtauth.Authenticator)

	r.Get("/", http.HandlerFunc(pprof.Index))
	r.Get("/allocs", http.HandlerFunc(pprof.Handler("allocs").ServeHTTP))
	r.Get("/block", http.HandlerFunc(pprof.Handler("block").ServeHTTP))
	r.Get("/cmdline", http.HandlerFunc(pprof.Cmdline))
	r.Get("/goroutine", http.HandlerFunc(pprof.Handler("goroutine").ServeHTTP))
	r.Get("/heap", http.HandlerFunc(pprof.Handler("heap").ServeHTTP))
	r.Get("/mutex", http.HandlerFunc(pprof.Handler("mutex").ServeHTTP))
	r.Get("/profile", http.HandlerFunc(pprof.Profile))
	r.Get("/threadcreate", http.HandlerFunc(pprof.Handler("threadcreate").ServeHTTP))
	r.Get("/trace", http.HandlerFunc(pprof.Trace))
	r.Get("/debug/pprof/goroutine", http.HandlerFunc(pprof.Handler("goroutine").ServeHTTP))

	return r
}
