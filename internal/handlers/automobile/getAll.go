package handler_automobile

import (
	handlers "Web-App/internal/handlers/errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *AutomobileHandlers) GetAll(c echo.Context) error {

	automobiles, err := h.Service.GetAll()
	if err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusInternalServerError, "Ошибка в получении списка авто")
	}

	return c.JSON(http.StatusOK, automobiles)
}
