package automobile

import (
	"Web-App/internal/models"
	"database/sql"
)

type AutomobileStorage struct {
	DB *sql.DB
}

func NewAutomobileStorage(db *sql.DB) *AutomobileStorage {
	return &AutomobileStorage{DB: db}
}

type StorageAutomobileInterface interface {
	GetAuto(id int) (*models.Automobile, error)
	GetAllAuto() ([]models.Automobile, error)
	AddAuto(automobile models.Automobile) error
	UpdateAuto(automobile *models.Automobile) error
	DeleteAuto(id int) error
}
