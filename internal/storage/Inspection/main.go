package Inspection

import (
	"Web-App/internal/models"
	"database/sql"
)

type InspectionStorage struct {
	DB *sql.DB
}

func NewInspectionStorage(db *sql.DB) *InspectionStorage {
	return &InspectionStorage{DB: db}
}

type StorageInspectionInterface interface {
	GetInsp(id int) (*models.Inspection, error)
	GetAllInsp() ([]models.Inspection, error)
	AddInsp(automobile *models.Inspection) error
	UpdateInsp(automobile *models.Inspection) error
	DeleteInsp(id int) error
}
