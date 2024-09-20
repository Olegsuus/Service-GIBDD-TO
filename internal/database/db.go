package db

import (
	"Web-App/internal/storage"
	"fmt"
	"log"

	"Web-App/internal/config"
	"database/sql"
	_ "github.com/lib/pq"
)

type DataBase struct {
	DB      *sql.DB
	Storage storage.Storage
}

func (db *DataBase) GetStorage(cfg *config.Config) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DataBase.Host, cfg.DataBase.Port, cfg.DataBase.User, cfg.DataBase.Password, cfg.DataBase.DBName)

	var err error
	db.DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}

	if err = db.DB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
}

func (db *DataBase) Stop() error {
	if db.DB != nil {
		err := db.DB.Close()
		{
			if err != nil {
				log.Fatalf("Failed to closed database: %s", err)
				return err
			}
		}
	}
	return nil
}
