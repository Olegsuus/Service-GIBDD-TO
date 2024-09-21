package storage

import (
	storage_inspection "Web-App/internal/storage/Inspection"
	storage_automobile "Web-App/internal/storage/automobile"
	"database/sql"
)

type Storage struct {
	AutomobileStorage *storage_automobile.AutomobileStorage
	InspectionStorage *storage_inspection.InspectionStorage
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		AutomobileStorage: storage_automobile.NewAutomobileStorage(db),
		InspectionStorage: storage_inspection.NewInspectionStorage(db),
	}
}
