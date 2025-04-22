package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	ApiKey     string
	SecretKey  string
	RedisHost  string
	RedisPort  string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBHost     string
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		ApiKey:     os.Getenv("DADATA_API_KEY"),
		SecretKey:  os.Getenv("DADATA_SECRET_KEY"),
		RedisPort:  os.Getenv("REDIS_PORT"),
		RedisHost:  os.Getenv("REDIS_HOST"),
	}

	if cfg.DBUser == "" || cfg.DBPassword == "" || cfg.DBName == "" || cfg.DBHost == "" || cfg.DBPort == "" || cfg.ApiKey == "" || cfg.SecretKey == "" || cfg.RedisPort == "" || cfg.RedisHost == "" {
		log.Fatal("Не заданы переменные окружения")
	}

	return cfg

}
