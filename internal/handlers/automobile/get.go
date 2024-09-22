package handler_automobile

import (
	"fmt"
	handlers "github.com/Olegsuus/TZ-WEB-App/internal/handlers/errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// Get Обработчик для получения автомобиля
// @Summary Получить автомобиль по id
// @Description Возвращает  автомобиль
// @Param id path int true "ID автомобиля"
// @Param automobile body models.Automobile true "Получение автомобиля по id"
// @Tags автомобиль
// @Accept  json
// @Produce  json
// @Success 200  "OK"
// @Failure 400 "Неверные данные запроса"
// @Failure 500  "Ошибка на сервере"
// @Router /car/:id [get]
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
	fmt.Println(automobile)

	return c.JSON(http.StatusOK, automobile)
}
