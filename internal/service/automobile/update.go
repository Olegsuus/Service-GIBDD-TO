package service_automobile

import (
	"Web-App/internal/models"
	"fmt"
	"log/slog"
)

func (s *AutomobileService) Update(automobile *models.Automobile) error {
	const op = "service_automobile.update"

	s.l.With(slog.String("op", op))

	automobileStorage, err := s.TranslatorToStorage(automobile)
	if err != nil {
		s.l.Error("Failed to translator to storage", err)
		return fmt.Errorf("%s: %w", op, err)
	}

	if err = s.auP.Update(automobileStorage); err != nil {
		s.l.Error("Failed to update automobile", "id", automobile.ID, "error", err)
		return fmt.Errorf("%s: %w", op, err)
	}

	s.l.Info("Automobile updated successfully", "id", automobile.ID)

	return nil
}
