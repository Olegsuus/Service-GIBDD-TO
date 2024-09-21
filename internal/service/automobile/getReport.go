package service_automobile

import (
	"fmt"
	"github.com/Olegsuus/TZ-WEB-App/internal/models"
	"log/slog"
)

func (s *AutomobileService) GetReport() (*models.Report, error) {
	const op = "service_automobile.getReport"

	s.l.With(slog.String("op", op))

	storageReport, err := s.auP.GetReport()
	if err != nil {
		s.l.Error("Service: Failed to get report", "error", err)
		return nil, fmt.Errorf("%s : %w", op, err)
	}

	report := models.Report{
		TotalCars:       storageReport.TotalCars,
		CarsOlder3Years: storageReport.CarsOlder3Years,
		CarsNewer3Years: storageReport.CarsNewer3Years,
	}

	s.l.Info("Successful get report")

	return &report, nil
}
