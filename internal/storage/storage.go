package storage

import (
	"database/sql"
	storage_inspection "github.com/Olegsuus/TZ-WEB-App/internal/storage/Inspection"
	storage_automobile "github.com/Olegsuus/TZ-WEB-App/internal/storage/automobile"
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
