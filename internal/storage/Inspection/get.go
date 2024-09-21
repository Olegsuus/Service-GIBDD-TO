package storage_inspection

import (
	storage_models "Web-App/internal/storage/models"
	"database/sql"
	"fmt"
)

func (s *InspectionStorage) Get(id int) (*storage_models.Inspection, error) {
	const op = "storage_inspection.get"

	query := `SELECT id, automobile_id, card_number, inspection_date, notes FROM inspections WHERE id = $1`

	row := s.DB.QueryRow(query, id)
	var inspection storage_models.Inspection
	if err := row.Scan(&inspection.ID, &inspection.AutomobileID, &inspection.CardNumber,
		&inspection.InspectionDate, &inspection.Notes); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		return nil, err
	}

	return &inspection, nil
}
