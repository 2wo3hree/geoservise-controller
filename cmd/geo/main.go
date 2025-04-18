// @title GeoService API
// @version 1.0
// @description This is a simple geo service using DaData.
// @host localhost:8080
// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
package main

import (
	_ "geoservise-jwt/docs"
	"geoservise-jwt/internal/app"
	"geoservise-jwt/internal/config"
	"geoservise-jwt/internal/server"
	"geoservise-jwt/internal/shutdown"
)

func main() {
	cfg := config.LoadConfig()

	application := app.NewApp(cfg.ApiKey, cfg.SecretKey, cfg.RedisHost, cfg.RedisPort)

	s := server.NewGeoServer(":8080", application.Router)
	shutdown.Gracefully(s)
}
