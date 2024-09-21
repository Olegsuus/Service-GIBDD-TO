package service_automobile

import (
	"github.com/Olegsuus/TZ-WEB-App/internal/models"
	storage_models "github.com/Olegsuus/TZ-WEB-App/internal/storage/models"
)

func (s *AutomobileService) TranslatorToModels(automobileStorage *storage_models.Automobile) (*models.Automobile, error) {
	return &models.Automobile{
		ID:               automobileStorage.ID,
		ReleaseDate:      automobileStorage.ReleaseDate,
		Model:            automobileStorage.Model,
		LicensePlate:     automobileStorage.LicensePlate,
		RegistrationDate: automobileStorage.RegistrationDate,
	}, nil
}

func (s *AutomobileService) TranslatorToStorage(automobileModels *models.Automobile) (*storage_models.Automobile, error) {
	return &storage_models.Automobile{
		ID:               automobileModels.ID,
		ReleaseDate:      automobileModels.ReleaseDate,
		Model:            automobileModels.Model,
		LicensePlate:     automobileModels.LicensePlate,
		RegistrationDate: automobileModels.RegistrationDate,
	}, nil
}
