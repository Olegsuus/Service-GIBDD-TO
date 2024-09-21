package service_inspection

import (
	"fmt"
	"log/slog"
)

func (s *InspectionService) Delete(id int) error {
	const op = "service_inspection.delete"

	s.l.With(slog.String("op", op))

	if err := s.inP.Delete(id); err != nil {
		s.l.Error("Failed to delete inspection", "id", id, "error", err)
		return fmt.Errorf("%s : %w", op, err)
	}

	s.l.Info("Inspection deleted successfully", "id", id)

	return nil
}
