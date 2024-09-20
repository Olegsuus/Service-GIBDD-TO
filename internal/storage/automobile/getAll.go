package automobile

import (
	"Web-App/internal/models"
)

func (s *AutomobileStorage) GetAllAuto() ([]models.Automobile, error) {
	var automobiles []models.Automobile
	query := `SELECT id, release_date, model, license_plate, registration_date FROM automobiles`
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var automobile models.Automobile
		if err := rows.Scan(&automobile.ID, &automobile.ReleaseDate, &automobile.Model, &automobile.LicensePlate, &automobile.RegistrationDate); err != nil {
			return nil, err
		}

		automobiles = append(automobiles, automobile)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return automobiles, err
}
