package main

import (
	"Web-App/internal/app"
	"Web-App/internal/config"
	"Web-App/internal/database"
	migration "Web-App/internal/migrations"
	"log"
)

func main() {
	cfg := config.GetConfig()
	db := database.DataBase{}
	db.GetStorage(cfg)
	migration.Migrations(cfg, db.DB)

	App := app.App{Config: cfg, DB: &db}

	srv := &app.Server{}
	App.ServerInterface = srv

	if err := App.Start(); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}

	//todo: task-7 написать обработку запросов через storage функции

	//todo: task-8 написать маршруты в handlers по методам storage

	//todo: task-9 написать роуты

}
