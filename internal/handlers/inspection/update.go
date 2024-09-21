package handler_inspection

import (
	handlers "Web-App/internal/handlers/errors"
	"Web-App/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *InspectionHandler) Update(c echo.Context) error {
	var inspection models.Inspection

	if err := c.Bind(&inspection); err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusBadRequest, "Неверный формат данных на обновление")
	}

	if err := h.Service.Update(&inspection); err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusInternalServerError, "Ошибка при обновлении данных тех. осмотра")
	}

	return c.JSON(http.StatusOK, inspection)
}
