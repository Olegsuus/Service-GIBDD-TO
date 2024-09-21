package handler_inspection

import (
	handlers "Web-App/internal/handlers/errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *InspectionHandler) GetAll(c echo.Context) error {
	inspections, err := h.Service.GetAll()
	if err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusInternalServerError, "Ошибка при получении всех тех. осмотров")
	}

	return c.JSON(http.StatusOK, inspections)
}
