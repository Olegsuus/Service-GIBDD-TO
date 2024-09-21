package service_inspection

import (
	"github.com/Olegsuus/TZ-WEB-App/internal/models"
	storage_models "github.com/Olegsuus/TZ-WEB-App/internal/storage/models"
)

func (s *InspectionService) TranslatorToModels(inspectionStorage *storage_models.Inspection) (*models.Inspection, error) {
	return &models.Inspection{
		ID:             inspectionStorage.ID,
		AutomobileID:   inspectionStorage.AutomobileID,
		CardNumber:     inspectionStorage.CardNumber,
		InspectionDate: inspectionStorage.InspectionDate,
		Notes:          inspectionStorage.Notes,
	}, nil
}

func (s *InspectionService) TranslatorToStorage(inspectionModels *models.Inspection) (*storage_models.Inspection, error) {
	return &storage_models.Inspection{
		ID:             inspectionModels.ID,
		AutomobileID:   inspectionModels.AutomobileID,
		CardNumber:     inspectionModels.CardNumber,
		InspectionDate: inspectionModels.InspectionDate,
		Notes:          inspectionModels.Notes,
	}, nil
}
