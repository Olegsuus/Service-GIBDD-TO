package handler_automobile

import (
	handlers "github.com/Olegsuus/TZ-WEB-App/internal/handlers/errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// Delete Обработчик для удаления автомобиля
// @Summary Удалить автомобиль
// @Description Удаляет  автомобиль
// @Tags автомобиль
// @Accept  json
// @Produce  json
// @Success 200  "OK"
// @Failure 500  "Ошибка на сервере"
// @Router /car/:id [delete]
func (h *AutomobileHandlers) Delete(c echo.Context) error {
	strId := c.Param("id")

	id, err := strconv.Atoi(strId)
	if err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusBadRequest, "Неверный формат id")
	}

	if err = h.Service.Delete(id); err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusInternalServerError, "Ошибка при удалении авто")
	}

	return c.NoContent(http.StatusNoContent)

}
