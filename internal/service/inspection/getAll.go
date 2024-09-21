package service_inspection

import (
	"fmt"
	"github.com/Olegsuus/TZ-WEB-App/internal/models"
	"log/slog"
)

func (s *InspectionService) GetAll() ([]models.Inspection, error) {
	const op = "service_inspection.getAll"

	s.l.With(slog.String("op", op))

	inspectionsStorage, err := s.inP.GetAll()
	if err != nil {
		s.l.Error("Failed to get all inspections", "error", err)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var inspections []models.Inspection
	for _, insp := range inspectionsStorage {
		inspection, err := s.TranslatorToModels(&insp)
		if err != nil {
			s.l.Error("Failed to translate inspection", "error", err)
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		inspections = append(inspections, *inspection)
	}

	s.l.Info("All inspections retrieved successfully")

	return inspections, nil
}
