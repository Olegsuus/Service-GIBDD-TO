package storage_inspection

import (
	storage_models "Web-App/internal/storage/models"
	"fmt"
)

func (s *InspectionStorage) Add(inspection *storage_models.Inspection) error {
	const op = "storage_inspection.add"

	query := `INSERT INTO inspections (automobile_id, card_number, inspection_date, notes)
              VALUES ($1, $2, $3, $4) RETURNING id`

	err := s.DB.QueryRow(
		query, inspection.AutomobileID, inspection.CardNumber,
		inspection.InspectionDate, inspection.Notes).Scan(&inspection.ID)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
