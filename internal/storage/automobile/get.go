package automobile

import "Web-App/internal/models"

func (s *AutomobileStorage) GetAuto(id int) (*models.Automobile, error) {
	var automobile models.Automobile
	query := "SELECT id, release_data, model, license_plate, registration_date FROM automobiles WHERE id = $1"
	err := s.DB.QueryRow(query, id).Scan(&automobile.ID, &automobile.ReleaseDate, &automobile.Model, &automobile.LicensePlate, &automobile.RegistrationDate)
	if err != nil {
		return nil, err
	}

	return &automobile, nil
}
