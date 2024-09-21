package handler_automobile

import (
	"Web-App/internal/models"
)

type AutomobileHandlers struct {
	Service AutomobileHandlerProvider
}

type AutomobileHandlerProvider interface {
	Get(id int) (*models.Automobile, error)
	GetAll() ([]models.Automobile, error)
	Add(automobile *models.Automobile) error
	Update(automobile *models.Automobile) error
	Delete(id int) error
}

func NewAutomobileHandler(service AutomobileHandlerProvider) *AutomobileHandlers {
	return &AutomobileHandlers{Service: service}
}
