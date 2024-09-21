package handler_automobile

import (
	handlers "github.com/Olegsuus/TZ-WEB-App/internal/handlers/errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetReport Обработчик для получения отчёта по автомобилям
// @Summary Получить отчёт по автомобилям
// @Description Возвращает агрегированные данные по автомобилям
// @Tags автомобиль
// @Accept  json
// @Produce  json
// @Success 200  "OK"
// @Failure 500  "Ошибка на сервере"
// @Router /report [get]
func (h *AutomobileHandlers) GetReport(c echo.Context) error {

	report, err := h.Service.GetReport()
	if err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusInternalServerError, "Ошибка при получении отчета")
	}

	return c.JSON(http.StatusOK, report)

}
