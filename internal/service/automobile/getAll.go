package service_automobile

import (
	"Web-App/internal/models"
	"fmt"
	"log/slog"
)

func (s *AutomobileService) GetAll() ([]models.Automobile, error) {
	const op = "service_automobiles.getAll"

	s.l.With(slog.String("op", op))

	automobileStorage, err := s.auP.GetAll()
	if err != nil {
		s.l.Error("Failed to get all automobiles", "error", err)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var automobiles []models.Automobile
	for _, auto := range automobileStorage {
		automobile, err := s.TranslatorToModels(&auto)
		if err != nil {
			s.l.Error("Failed to translate automobile", "error", err)
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		automobiles = append(automobiles, *automobile)
	}

	s.l.Info("All automobiles retrieved successfully")
	return automobiles, nil
}
