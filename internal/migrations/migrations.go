package migrations

import (
	"embed"
	"log"

	"Web-App/internal/config"
	"database/sql"
	"github.com/pressly/goose/v3"
)

//go:embed migrations
var embedMigrations embed.FS

func Migrations(cfg *config.Config, db *sql.DB) {

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Failed to set goose dialect: %v", err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	log.Println("Success migrations")
}
