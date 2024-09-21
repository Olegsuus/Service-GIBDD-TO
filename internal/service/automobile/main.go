package service_automobile

import (
	storage_models "github.com/Olegsuus/TZ-WEB-App/internal/storage/models"
	"log/slog"
)

type AutomobileService struct {
	l   *slog.Logger
	auP AutomobileProvider
}

type AutomobileProvider interface {
	Get(id int) (*storage_models.Automobile, error)
	GetAll() ([]storage_models.Automobile, error)
	Add(automobile storage_models.Automobile) error
	Update(automobile *storage_models.Automobile) error
	Delete(id int) error
	GetReport() (*storage_models.Report, error)
}

func NewAutomobileService(auP AutomobileProvider, l *slog.Logger) *AutomobileService {
	return &AutomobileService{auP: auP, l: l}
}
