package storage_automobile

import (
	"fmt"
	storage_models "github.com/Olegsuus/TZ-WEB-App/internal/storage/models"
	"time"
)

func (s *AutomobileStorage) GetReport() (*storage_models.Report, error) {
	const op = "storage_automobile.getReport"

	now := time.Now()
	threeYearsAgo := now.AddDate(-3, 0, 0)

	var total, older3, newer3 int

	err := s.DB.QueryRow(`SELECT COUNT(*) FROM automobiles`).Scan(&total)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = s.DB.QueryRow(`SELECT COUNT(*) FROM automobiles WHERE release_date < $1`, threeYearsAgo).Scan(&older3)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = s.DB.QueryRow(`SELECT COUNT(*) FROM automobiles WHERE release_date >= $1`, threeYearsAgo).Scan(&newer3)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	report := &storage_models.Report{
		TotalCars:       total,
		CarsOlder3Years: older3,
		CarsNewer3Years: newer3,
	}

	return report, nil
}
