package handler_inspection

import (
	handlers "github.com/Olegsuus/TZ-WEB-App/internal/handlers/errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// Get Обработчик для удаления тех осмотра по id
// @Summary Удалить тех осмотр по id
// @Description Удаляет тех осмотр
// @Tags тех осмотр
// @Accept  json
// @Produce  json
// @Success 200  "OK"
// @Failure 500  "Ошибка на сервере"
// @Router /inspection/:id [delete]
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
