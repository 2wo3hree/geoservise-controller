package app

import (
	"geoservise-jwt/internal/auth"
	"geoservise-jwt/internal/cache"
	"geoservise-jwt/internal/handler"
	"geoservise-jwt/internal/metrics"
	"geoservise-jwt/internal/responder"
	"geoservise-jwt/internal/router"
	"geoservise-jwt/internal/service"
	"github.com/ekomobile/dadata/v2"
	"github.com/ekomobile/dadata/v2/client"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	"time"
)

type App struct {
	Router   *chi.Mux
	JWTAuth  *jwtauth.JWTAuth
	Handlers *handler.AddressHandler
}

func NewApp(apiKey, secretKey, redisHost, redisPort string) *App {

	metrics.Init()

	// init Dadata client
	creds := client.Credentials{
		ApiKeyValue:    apiKey,
		SecretKeyValue: secretKey,
	}
	api := dadata.NewSuggestApi(
		client.WithCredentialProvider(&creds),
	)

	rdb := cache.NewRedisClient(redisHost, redisPort)

	// init service
	s := service.NewService(api)

	cachedService := service.NewCachedGeoService(s, rdb, time.Hour)

	// init responder
	resp := responder.NewJSONResponder()

	// init handlers
	h := handler.NewAddressHandler(cachedService, resp)

	// init jwt
	auth.InitJWT()

	// init router
	r := router.SetupRouter(h, auth.TokenAuth)

	return &App{
		Router:   r,
		JWTAuth:  auth.TokenAuth,
		Handlers: h,
	}

}
