package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	ApiKey    string
	SecretKey string
	RedisHost string
	RedisPort string
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	apiKey := os.Getenv("DADATA_API_KEY")
	secretKey := os.Getenv("DADATA_SECRET_KEY")
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	if apiKey == "" || secretKey == "" || redisHost == "" || redisPort == "" {
		log.Fatal("переменные окружения не заданы")
	}

	return &Config{
		ApiKey:    apiKey,
		SecretKey: secretKey,
		RedisHost: redisHost,
		RedisPort: redisPort,
	}

}
