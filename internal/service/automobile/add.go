package service_automobile

import (
	"fmt"
	"github.com/Olegsuus/TZ-WEB-App/internal/models"
	"log/slog"
)

func (s *AutomobileService) Add(automobile *models.Automobile) error {

	const op = "service_automobile.Add"

	s.l.With(slog.String("op", op))

	automobileStorage, err := s.TranslatorToStorage(automobile)
	if err != nil {
		s.l.Error("Failed to translator to storage:", err)
		return fmt.Errorf("%s: %w", op, err)
	}

	if err = s.auP.Add(*automobileStorage); err != nil {
		s.l.Error("Failed to add automobile", "error", err)
		return fmt.Errorf("%s: %w", op, err)
	}

	automobile.ID = automobileStorage.ID

	s.l.Info("Automobile added successfully", "id", automobile.ID)

	return nil
}
