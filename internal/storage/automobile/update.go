package storage_automobile

import (
	"fmt"
	storage_models "github.com/Olegsuus/TZ-WEB-App/internal/storage/models"
)

func (s *AutomobileStorage) Update(automobile *storage_models.Automobile) error {
	const op = "storage_automobile.update"

	query := `UPDATE automobiles SET release_date = $1, model = $2, license_plate = $3, registration_date = $4 WHERE id = $5`
	result, err := s.DB.Exec(query, automobile.ReleaseDate, automobile.Model, automobile.LicensePlate,
		automobile.RegistrationDate, automobile.ID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("Автомобиль не найден по этому id: %d", automobile.ID)
	}

	return nil
}
