package handler_inspection

import (
	"Web-App/internal/models"
)

type InspectionHandler struct {
	Service InspectionHandlerProvider
}

type InspectionHandlerProvider interface {
	Get(id int) (*models.Inspection, error)
	GetAll() ([]models.Inspection, error)
	Add(automobile *models.Inspection) error
	Update(automobile *models.Inspection) error
	Delete(id int) error
}

func NewInspectionHandler(service InspectionHandlerProvider) *InspectionHandler {
	return &InspectionHandler{
		Service: service,
	}
}
