package service_automobile

import (
	"github.com/Olegsuus/TZ-WEB-App/internal/models"
	storage_models "github.com/Olegsuus/TZ-WEB-App/internal/storage/models"
)

func (s *AutomobileService) TranslatorToModels(automobileStorage *storage_models.Automobile) (*models.Automobile, error) {
	return &models.Automobile{
		ID:               automobileStorage.ID,
		ReleaseDate:      models.CustomDate{Time: automobileStorage.ReleaseDate},
		Model:            automobileStorage.Model,
		LicensePlate:     automobileStorage.LicensePlate,
		RegistrationDate: models.CustomDate{Time: automobileStorage.RegistrationDate},
	}, nil
}

func (s *AutomobileService) TranslatorToStorage(automobileModels *models.Automobile) (*storage_models.Automobile, error) {
	return &storage_models.Automobile{
		ID:               automobileModels.ID,
		ReleaseDate:      automobileModels.ReleaseDate.Time,
		Model:            automobileModels.Model,
		LicensePlate:     automobileModels.LicensePlate,
		RegistrationDate: automobileModels.RegistrationDate.Time,
	}, nil
}
