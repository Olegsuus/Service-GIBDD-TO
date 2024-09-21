package service_inspection

import (
	"Web-App/internal/models"
	"fmt"
	"log/slog"
)

func (s *InspectionService) Update(inspection *models.Inspection) error {
	const op = "service_inspection.update"

	s.l.With(slog.String("op", op))

	inspectionStorage, err := s.TranslatorToStorage(inspection)
	if err != nil {
		s.l.Error("Failed to translator to storage", err)
		return fmt.Errorf("%s: %w", op, err)
	}

	if err = s.inP.Update(inspectionStorage); err != nil {
		s.l.Error("Failed to update inspection", "id", inspection.ID, "error", err)
		return fmt.Errorf("%s: %w", op, err)
	}

	s.l.Info("Inspection updated successfully", "id", inspection.ID)

	return nil
}
