package storage_automobile

import (
	"database/sql"
	"fmt"
	storage_models "github.com/Olegsuus/TZ-WEB-App/internal/storage/models"
)

func (s *AutomobileStorage) Get(id int) (*storage_models.Automobile, error) {
	const op = "storage_automobile.get"

	query := `SELECT id, release_date, model, license_plate, registration_date FROM automobiles WHERE id = $1`

	row := s.DB.QueryRow(query, id)
	var automobile storage_models.Automobile

	if err := row.Scan(&automobile.ID, &automobile.ReleaseDate, &automobile.Model,
		&automobile.LicensePlate, &automobile.RegistrationDate); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("automobile with id %d not found", id)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &automobile, nil
}
