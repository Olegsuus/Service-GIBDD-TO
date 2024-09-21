package service_inspection

import (
	storage_models "Web-App/internal/storage/models"
	"log/slog"
)

type InspectionService struct {
	l   *slog.Logger
	inP InspectionProvider
}

type InspectionProvider interface {
	Get(id int) (*storage_models.Inspection, error)
	GetAll() ([]storage_models.Inspection, error)
	Add(automobile *storage_models.Inspection) error
	Update(automobile *storage_models.Inspection) error
	Delete(id int) error
}

func NewInspectionService(inP InspectionProvider, l *slog.Logger) *InspectionService {
	return &InspectionService{inP: inP, l: l}
}
