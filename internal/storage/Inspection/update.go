package storage_inspection

import (
	storage_models "Web-App/internal/storage/models"
	"fmt"
)

func (s *InspectionStorage) Update(inspection *storage_models.Inspection) error {
	const op = "storage_inspection.update"

	query := `UPDATE inspections SET automobile_id = $1, card_number = $2, inspection_date = $3, notes = $4 WHERE id = $5`

	result, err := s.DB.Exec(query, inspection.AutomobileID, inspection.CardNumber,
		inspection.InspectionDate, inspection.Notes, inspection.ID)
	if err != nil {
		return fmt.Errorf("failed to update inspection: %w", err, op)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err, op)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("Ни одного тех. осмотра не найдено по id: %d", inspection.ID)
	}

	return nil
}
