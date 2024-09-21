package main

import (
	"Web-App/internal/app"
	"Web-App/internal/config"
	"Web-App/internal/database"
	"Web-App/internal/migrations"
	"Web-App/internal/storage"
	"log"
	"log/slog"
	"os"
)

func main() {
	cfg := config.GetConfig()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db := &database.DataBase{}
	db.GetStorage(cfg)
	defer func() {
		if err := db.Stop(); err != nil {
			logger.Error("Failed to close database", "error", err)
		}
	}()

	migrations.Migrations(cfg, db.DB)

	store := storage.NewStorage(db.DB)

	App := app.NewApp(cfg, store, logger)

	if err := App.Start(store); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
