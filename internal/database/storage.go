package database

import (
	"fmt"
	"log"

	"Web-App/internal/config"
	"database/sql"
	_ "github.com/lib/pq"
)

type DataBase struct {
	DB *sql.DB
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

//func (db *DataBase) GetAutomobileById(id int) (*models.Automobile, error) {
//	var automobile models.Automobile
//	query := "SELECT id, make, model, year, license_plate FROM automobiles WHERE id = 1$"
//	row := db.DB.QueryRow(query, id)
//
//	err := row.Scan(&automobile.ID, &automobile.Make, &automobile.Model, &automobile.Year, &automobile.LicensePlate)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			return nil, err
//		}
//		log.Printf("Failed to scan row: %v\n", err)
//		return nil, err
//	}
//
//	return &automobile, nil
//}
//
//func (db *DataBase) GetInspectionsById(id int) (*models.Inspection, error) {
//	var inspection models.Inspection
//	query := "SELECT id, automobile_id, inspection_date, notes FROM inspections WHERE id = 1$"
//	row := db.DB.QueryRow(query, id)
//
//	err := row.Scan(&inspection.ID, &inspection.AutomobileID, &inspection.InspectionDate, &inspection.Notes)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			return nil, err
//		}
//		log.Printf("Failed to scan row: %v\n", err)
//		return nil, err
//	}
//
//	return &inspection, nil
//}
//
//func (db *DataBase) GetAllAutomobiles() ([]models.Automobile, error) {
//	var automobiles []models.Automobile
//
//	query := "SELECT id, make, model, year, license_plate FROM automobiles"
//	rows, err := db.DB.Query(query)
//	if err != nil {
//		log.Fatalf("Failed to execute query: %v", err)
//	}
//
//	defer rows.Close()
//
//	for rows.Next() {
//		var automobile models.Automobile
//		err := rows.Scan(&automobile.ID, &automobile.Make, &automobile.Model, &automobile.Year, &automobile.LicensePlate)
//		if err != nil {
//			log.Fatalf("Failed to scan rows: %v", err)
//		}
//
//		automobiles = append(automobiles, automobile)
//	}
//
//	if err := rows.Err(); err != nil {
//		log.Fatalf("Failed iteration rows: %v\n", err)
//		return nil, err
//	}
//
//	return automobiles, nil
//}
