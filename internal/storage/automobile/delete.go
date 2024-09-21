package storage_automobile

import (
	"fmt"
)

func (s *AutomobileStorage) Delete(id int) error {
	const op = "storage_automobile.delete"

	query := `DELETE FROM automobiles WHERE id = $1`
	result, err := s.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("Ни один автомобиль не найден с таким id: %d", id)
	}

	return nil
}
