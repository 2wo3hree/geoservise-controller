package app

import (
	"fmt"
	"geoservise-jwt/internal/config"
	handler2 "geoservise-jwt/internal/delivery/handler"
	"geoservise-jwt/internal/delivery/router"
	"geoservise-jwt/internal/infrastructure/cache"
	db2 "geoservise-jwt/internal/infrastructure/db"
	"geoservise-jwt/internal/infrastructure/metrics"
	"geoservise-jwt/internal/infrastructure/repository/postgres"
	"geoservise-jwt/internal/usecase/responder"
	service2 "geoservise-jwt/internal/usecase/service"
	"github.com/ekomobile/dadata/v2"
	"github.com/ekomobile/dadata/v2/client"
	"github.com/go-chi/chi/v5"
	"time"
)

type App struct {
	Router *chi.Mux
}

func NewApp(cfg *config.Config) *App {
	metrics.Init()
	// int db
	pool := db2.NewPostgres(cfg)

	// run Migrations
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db2.RunMigrations(dbURL)

	// init Dadata client
	creds := client.Credentials{
		ApiKeyValue:    cfg.ApiKey,
		SecretKeyValue: cfg.SecretKey,
	}
	api := dadata.NewSuggestApi(
		client.WithCredentialProvider(&creds),
	)

	rdb := cache.NewRedisClient(cfg.RedisHost, cfg.RedisPort)

	// init repo
	userRepo := postgres.NewUserRepo(pool)

	// init service
	s := service2.NewService(api)
	userService := service2.NewUserService(userRepo)

	cachedService := service2.NewCachedGeoService(s, rdb, time.Hour)

	// init responder
	resp := responder.NewJSONResponder()

	// init handlers
	addressHandler := handler2.NewAddressHandler(cachedService, resp)
	userHandler := handler2.NewUserHandler(userService, resp)

	// init router
	r := router.SetupRouter(addressHandler, userHandler, userService)

	return &App{
		Router: r,
	}

}
