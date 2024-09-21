package handler_automobile

import (
	handlers "Web-App/internal/handlers/errors"
	"Web-App/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *AutomobileHandlers) Update(c echo.Context) error {
	var automobile models.Automobile

	if err := c.Bind(&automobile); err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusBadRequest, "Неверные формат данных для обновления авто")
	}

	if err := h.Service.Update(&automobile); err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusInternalServerError, "Ошибка при обновлении данных авто")
	}

	return c.JSON(http.StatusOK, automobile)
}
