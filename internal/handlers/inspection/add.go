package handler_inspection

import (
	handlers "github.com/Olegsuus/TZ-WEB-App/internal/handlers/errors"
	"github.com/Olegsuus/TZ-WEB-App/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Add Обработчик для добавления тех осмотра
// @Summary Добавить тех осмотр
// @Description Добавляет тех осмотр
// @Tags тех осмотр
// @Accept  json
// @Produce  json
// @Success 200  "OK"
// @Failure 500  "Ошибка на сервере"
// @Router /inspection [post]
func (h *InspectionHandler) Add(c echo.Context) error {
	var inspection models.Inspection

	if err := c.Bind(&inspection); err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusBadRequest, "Неверный формат данных на добавление тех. осмотра")
	}

	if err := h.Service.Add(&inspection); err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusInternalServerError, "Ошибка при добавлении нового тех. осмотра")
	}

	return c.JSON(http.StatusCreated, inspection)
}
