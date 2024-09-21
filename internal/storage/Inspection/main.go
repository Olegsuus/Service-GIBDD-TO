package storage_inspection

import (
	"database/sql"
)

type InspectionStorage struct {
	DB *sql.DB
}

func NewInspectionStorage(db *sql.DB) *InspectionStorage {
	return &InspectionStorage{DB: db}
}
