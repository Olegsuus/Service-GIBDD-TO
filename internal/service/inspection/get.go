package service_inspection

import (
	"Web-App/internal/models"
	"fmt"
	"log/slog"
)

func (s *InspectionService) Get(id int) (*models.Inspection, error) {
	const op = "service_inspection.get"

	s.l.With(slog.String("op", op))

	inspectionStorage, err := s.inP.Get(id)
	if err != nil {
		s.l.Error("Failed to get inspection", "error", err)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	inspection, err := s.TranslatorToModels(inspectionStorage)
	if err != nil {
		s.l.Error("Failed to get inspections", "error", err)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	s.l.Info("successful get inspection where id:", id)

	return inspection, nil
}
