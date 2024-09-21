package handler

import (
	handler_automobile "Web-App/internal/handlers/automobile"
	handler_inspection "Web-App/internal/handlers/inspection"
)

type Handler struct {
	AutomobileHandler *handler_automobile.AutomobileHandlers
	InspectionHandler *handler_inspection.InspectionHandler
}

func NewHandler(autoHandler *handler_automobile.AutomobileHandlers,
	inspectionHandler *handler_inspection.InspectionHandler) *Handler {
	return &Handler{
		AutomobileHandler: autoHandler,
		InspectionHandler: inspectionHandler,
	}
}
