package handler_inspection

import (
	handlers "Web-App/internal/handlers/errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *InspectionHandler) Get(c echo.Context) error {
	strId := c.Param("id")

	id, err := strconv.Atoi(strId)
	if err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusBadRequest, "Неверный формат id")
	}

	inspection, err := h.Service.Get(id)
	if err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusNotFound, "Такой тех. осмотр не найден")
	}

	return c.JSON(http.StatusOK, inspection)
}
