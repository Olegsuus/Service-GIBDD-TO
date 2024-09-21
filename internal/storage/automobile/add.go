package storage_automobile

import (
	storage_models "Web-App/internal/storage/models"
	"fmt"
)

func (s *AutomobileStorage) Add(automobile storage_models.Automobile) error {
	const op = "storage_automobile.add"

	query := `INSERT INTO automobiles (release_date, model, license_plate, registration_date) 
			  VALUES ($1, $2, $3, $4) RETURNING id`

	err := s.DB.QueryRow(
		query, automobile.ReleaseDate, automobile.Model, automobile.LicensePlate,
		automobile.RegistrationDate).Scan(&automobile.ID)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
