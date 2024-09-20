package Inspection

import "Web-App/internal/models"

func (s *InspectionStorage) AddInsp(inspection *models.Inspection) error {
	query := `INSERT INTO inspections (automobile_id, card_number, inspection_date, notes) VALUES ($1, $2, $3, $4) RETURNING id`
	err := s.DB.QueryRow(query, inspection.AutomobileID, inspection.CardNumber, inspection.InspectionDate, inspection.Notes).Scan(&inspection.ID)
	if err != nil {
		return err
	}
	return nil
}
