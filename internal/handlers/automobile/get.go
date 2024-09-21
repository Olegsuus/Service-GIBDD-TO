package handler_automobile

import (
	handlers "Web-App/internal/handlers/errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *AutomobileHandlers) Get(c echo.Context) error {
	strId := c.Param("id")

	id, err := strconv.Atoi(strId)
	if err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusBadRequest, "Неверный формат id")
	}

	automobile, err := h.Service.Get(id)
	if err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusNotFound, "Автомобиль не найден")
	}

	return c.JSON(http.StatusOK, automobile)
}
