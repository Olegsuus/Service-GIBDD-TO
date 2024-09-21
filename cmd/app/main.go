package main

import (
	_ "github.com/Olegsuus/TZ-WEB-App/docs"
	"github.com/Olegsuus/TZ-WEB-App/internal/app"
	"github.com/Olegsuus/TZ-WEB-App/internal/config"
	"github.com/Olegsuus/TZ-WEB-App/internal/database"
	"github.com/Olegsuus/TZ-WEB-App/internal/migrations"
	"github.com/Olegsuus/TZ-WEB-App/internal/storage"
	"log"
	"log/slog"
	"os"
)

// @title Web-App API
// @version 1.0
// @description API для управления автопарком.

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com

// @host localhost:8001
// @BasePath /
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
