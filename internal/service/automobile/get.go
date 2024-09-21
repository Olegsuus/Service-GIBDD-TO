package service_automobile

import (
	"fmt"
	"github.com/Olegsuus/TZ-WEB-App/internal/models"
	"log/slog"
)

func (s *AutomobileService) Get(id int) (*models.Automobile, error) {

	const op = "service_automobile.get"

	s.l.With(slog.String("op", op))

	autoStorage, err := s.auP.Get(id)
	if err != nil {
		s.l.Error("Failed to get automobile", "id", id, "error", err)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	automobile, err := s.TranslatorToModels(autoStorage)
	if err != nil {
		s.l.Error("Failed to translator to models:", err)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	s.l.Info("Automobile retrieved successfully", "id", id)

	return automobile, nil
}
