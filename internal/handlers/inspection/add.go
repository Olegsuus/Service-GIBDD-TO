package handler_inspection

import (
	handlers "Web-App/internal/handlers/errors"
	"Web-App/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *InspectionHandler) Add(c echo.Context) error {
	var inspection models.Inspection

	if err := c.Bind(&inspection); err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusBadRequest, "Неверный формат данных на добавление тех. осмотра")
	}

	if err := h.Service.Add(&inspection); err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusInternalServerError, "Ошибка при добавлении нового тех. осмотра")
	}

	return c.JSON(http.StatusCreated, inspection)
}
