package handler_inspection

import (
	handlers "github.com/Olegsuus/TZ-WEB-App/internal/handlers/errors"
	"github.com/Olegsuus/TZ-WEB-App/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Update Обработчик для обновления данных тех осмотра
// @Summary Обновить данные о тех осмотре
// @Description Обновляет  данные тех осмотра
// @Tags тех осмотр
// @Accept  json
// @Produce  json
// @Success 200  "OK"
// @Failure 500  "Ошибка на сервере"
// @Router /inspection/:id [patch]
func (h *InspectionHandler) Update(c echo.Context) error {
	var inspection models.Inspection

	if err := c.Bind(&inspection); err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusBadRequest, "Неверный формат данных на обновление")
	}

	if err := h.Service.Update(&inspection); err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusInternalServerError, "Ошибка при обновлении данных тех. осмотра")
	}

	return c.JSON(http.StatusOK, inspection)
}
