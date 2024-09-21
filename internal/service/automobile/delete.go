package service_automobile

import (
	"fmt"
	"log/slog"
)

func (s *AutomobileService) Delete(id int) error {
	const op = "service_automobile.delete"

	s.l.With(slog.String("op", op))

	if err := s.auP.Delete(id); err != nil {
		s.l.Error("Failed to delete automobile", "id", id, "error", err)
		return fmt.Errorf("%s : %w", op, err)
	}

	s.l.Info("Automobile deleted successfully", "id", id)

	return nil
}
