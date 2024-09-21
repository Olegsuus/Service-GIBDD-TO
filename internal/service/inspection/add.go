package service_inspection

import (
	"Web-App/internal/models"
	"fmt"
	"log/slog"
)

func (s *InspectionService) Add(inspection *models.Inspection) error {

	const op = "service_inspection.Add"

	s.l.With(slog.String("op", op))

	inspectionStorage, err := s.TranslatorToStorage(inspection)
	if err != nil {
		s.l.Error("Failed to add new inspection", err)
		return fmt.Errorf("%s: %w", op, err)
	}

	if err = s.inP.Add(inspectionStorage); err != nil {
		s.l.Error("Failed to add inspection", "error", err)
		return fmt.Errorf("%s: %w", op, err)
	}

	inspection.ID = inspectionStorage.ID

	s.l.Info("Inspection added successfully", "id", inspection.ID)

	return nil
}
