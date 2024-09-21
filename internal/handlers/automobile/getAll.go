package handler_automobile

import (
	handlers "github.com/Olegsuus/TZ-WEB-App/internal/handlers/errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetAll Обработчик для получения списка автомобилей
// @Summary Получить список автомобилей
// @Description Возвращает  список автомобилей
// @Tags автомобиль
// @Accept  json
// @Produce  json
// @Success 200  "OK"
// @Failure 500  "Ошибка на сервере"
// @Router /cars [get]
func (h *AutomobileHandlers) GetAll(c echo.Context) error {

	automobiles, err := h.Service.GetAll()
	if err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusInternalServerError, "Ошибка в получении списка авто")
	}

	return c.JSON(http.StatusOK, automobiles)
}
