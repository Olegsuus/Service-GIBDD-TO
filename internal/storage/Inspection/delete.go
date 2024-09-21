package storage_inspection

import (
	"fmt"
)

func (s *InspectionStorage) Delete(id int) error {
	const op = "storage_inspection.delete"

	query := `DELETE FROM inspections WHERE id = $1`
	result, err := s.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete inspection: %w", err, op)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {

		return fmt.Errorf("failed to get rows affected: %w", err, op)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("Ни одного тех. осмотра не найдено по id: %d", id, op)
	}

	return nil
}
