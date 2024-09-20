package Inspection

import "Web-App/internal/models"

func (s *InspectionStorage) GetAllInsp() ([]models.Inspection, error) {
	var inspections []models.Inspection
	query := `SELECT id, automobile_id, card_number, inspection_date, notes FROM inspections`
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var inspection models.Inspection
		if err := rows.Scan(&inspection.ID, &inspection.AutomobileID, &inspection.CardNumber, &inspection.InspectionDate, &inspection.Notes); err != nil {
			return nil, err
		}
		inspections = append(inspections, inspection)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return inspections, nil
}
