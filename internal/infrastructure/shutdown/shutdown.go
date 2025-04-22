package shutdown

import (
	"context"
	"geoservise-jwt/internal/infrastructure/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Gracefully(server *server.GeoServer) {
	go func() {
		if err := server.Serve(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}

	log.Println("Server stopped gracefully")
}
