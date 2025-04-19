package server

import (
	"context"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"time"
)

type GeoServer struct {
	server *http.Server
}

func NewGeoServer(addr string, r *chi.Mux) *GeoServer {
	return &GeoServer{
		server: &http.Server{
			Addr:         addr,
			Handler:      r,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 30 * time.Second,
		},
	}
}

func (g *GeoServer) Serve() error {
	log.Println("Сервер запущен на порту", g.server.Addr)
	return g.server.ListenAndServe()
}

func (g *GeoServer) Shutdown(ctx context.Context) error {
	log.Println("Остановка сервера...")
	return g.server.Shutdown(ctx)
}
