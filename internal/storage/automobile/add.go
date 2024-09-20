package automobile

import "Web-App/internal/models"

func (s *AutomobileStorage) AddAuto(automobile models.Automobile) error {
	query := "INSERT INTO automobiles (release_date, model, license_plate, registration_date) VALUES ($1, $2, $3,$4) RETURNING id"
	err := s.DB.QueryRow(query, automobile.ReleaseDate, automobile.Model, automobile.LicensePlate,
		automobile.RegistrationDate).Scan(&automobile.ID)

	if err != nil {
		return err
	}

	return nil
}
