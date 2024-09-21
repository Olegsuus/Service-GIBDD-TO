package handler_inspection

import (
	handlers "github.com/Olegsuus/TZ-WEB-App/internal/handlers/errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetAll Обработчик для получения списка тех осмотров
// @Summary Получить список тех осмотров
// @Description Возвращает список тех осмотров
// @Tags тех осмотр
// @Accept  json
// @Produce  json
// @Success 200  "OK"
// @Failure 500  "Ошибка на сервере"
// @Router /inspections [get]
func (h *InspectionHandler) GetAll(c echo.Context) error {
	inspections, err := h.Service.GetAll()
	if err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusInternalServerError, "Ошибка при получении всех тех. осмотров")
	}

	return c.JSON(http.StatusOK, inspections)
}
