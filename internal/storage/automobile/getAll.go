package storage_automobile

import (
	"fmt"
	storage_models "github.com/Olegsuus/TZ-WEB-App/internal/storage/models"
)

func (s *AutomobileStorage) GetAll() ([]storage_models.Automobile, error) {
	const op = "storage_automobile.getAll"

	var automobiles []storage_models.Automobile
	query := `SELECT id, release_date, model, license_plate, registration_date FROM automobiles`
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer rows.Close()

	for rows.Next() {
		var automobile storage_models.Automobile
		if err = rows.Scan(&automobile.ID, &automobile.ReleaseDate, &automobile.Model,
			&automobile.LicensePlate, &automobile.RegistrationDate); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		automobiles = append(automobiles, automobile)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return automobiles, err
}
