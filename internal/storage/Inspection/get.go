package Inspection

import (
	"Web-App/internal/models"
)

func (s *InspectionStorage) GetInsp(id int) (*models.Inspection, error) {
	var inspection models.Inspection
	query := `SELECT id, automobile_id, card_number, inspection_date, notes FROM inspections WHERE id = $1`
	err := s.DB.QueryRow(query, id).Scan(&inspection.ID, &inspection.AutomobileID, &inspection.CardNumber, &inspection.InspectionDate, &inspection.Notes)
	if err != nil {
		return nil, err
	}
	return &inspection, nil
}
