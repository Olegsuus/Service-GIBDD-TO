package automobile

import "Web-App/internal/models"

func (s *AutomobileStorage) UpdateAuto(automobile *models.Automobile) error {
	query := `UPDATE automobiles SET release_date = $1, model = $2, license_plate = $3, registration_date = $4 WHERE id = $5`
	_, err := s.DB.Exec(query, automobile.ReleaseDate, automobile.Model, automobile.LicensePlate,
		automobile.RegistrationDate, automobile.ID)

	return err
}
