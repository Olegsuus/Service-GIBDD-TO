package Inspection

import "Web-App/internal/models"

func (s *InspectionStorage) UpdateInsp(inspection *models.Inspection) error {
	query := `UPDATE inspections SET automobile_id = $1, card_number = $2, inspection_date = $3, notes = $4 WHERE id = $5`
	_, err := s.DB.Exec(query, inspection.AutomobileID, inspection.CardNumber, inspection.InspectionDate, inspection.Notes, inspection.ID)
	return err
}
