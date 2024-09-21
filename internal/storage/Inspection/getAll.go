package storage_inspection

import (
	"fmt"
	storage_models "github.com/Olegsuus/TZ-WEB-App/internal/storage/models"
)

func (s *InspectionStorage) GetAll() ([]storage_models.Inspection, error) {
	const op = "storage_inspection.getAll"

	var inspections []storage_models.Inspection
	query := `SELECT id, automobile_id, card_number, inspection_date, notes FROM inspections`
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var inspection storage_models.Inspection
		if err := rows.Scan(&inspection.ID, &inspection.AutomobileID, &inspection.CardNumber, &inspection.InspectionDate, &inspection.Notes); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		inspections = append(inspections, inspection)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return inspections, nil
}
