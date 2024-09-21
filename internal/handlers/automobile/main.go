package handler_automobile

import (
	"github.com/Olegsuus/TZ-WEB-App/internal/models"
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
	GetReport() (*models.Report, error)
}

func NewAutomobileHandler(service AutomobileHandlerProvider) *AutomobileHandlers {
	return &AutomobileHandlers{Service: service}
}
