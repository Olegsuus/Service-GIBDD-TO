package handler_inspection

import (
	handlers "github.com/Olegsuus/TZ-WEB-App/internal/handlers/errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// Get Обработчик для получения тех осмотра по id
// @Summary Получить тех осмотр по id
// @Description Возвращает тех осмотр
// @Tags тех осмотр
// @Accept  json
// @Produce  json
// @Success 200  "OK"
// @Failure 500  "Ошибка на сервере"
// @Router /inspection/:id [get]
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
