package storage_automobile

import (
	"database/sql"
)

type AutomobileStorage struct {
	DB *sql.DB
}

func NewAutomobileStorage(db *sql.DB) *AutomobileStorage {
	return &AutomobileStorage{DB: db}
}
