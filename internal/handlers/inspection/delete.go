package handler_inspection

import (
	handlers "Web-App/internal/handlers/errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *InspectionHandler) Delete(c echo.Context) error {
	strId := c.Param("id")

	id, err := strconv.Atoi(strId)
	if err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusBadRequest, "Неверный формат id")
	}

	if err := h.Service.Delete(id); err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusInternalServerError, "Ошибка при удалении тех. осмотра")
	}

	return c.NoContent(http.StatusNoContent)
}
