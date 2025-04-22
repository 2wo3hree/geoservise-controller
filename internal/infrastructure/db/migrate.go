package db

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func RunMigrations(dbURL string) {
	m, err := migrate.New(
		"file:///root/internal/infrastructure/db/migrations",
		dbURL,
	)
	if err != nil {
		log.Fatalf("ошибка инициализации миграций: %v", err)
	}

	if err := m.Up(); err != nil {
		log.Fatalf("ошибка применения миграций: %v", err)
	}

	log.Println("Миграции применены успешно")
}
