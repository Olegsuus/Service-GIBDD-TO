package handler_automobile

import (
	handlers "github.com/Olegsuus/TZ-WEB-App/internal/handlers/errors"
	"github.com/Olegsuus/TZ-WEB-App/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Update Обработчик для обновления данных автомобиля
// @Summary Обновить автомобиль
// @Description Обновляет  автомобиль
// @Tags автомобиль
// @Accept  json
// @Produce  json
// @Success 200  "OK"
// @Failure 500  "Ошибка на сервере"
// @Router /car/:id [patch]
func (h *AutomobileHandlers) Update(c echo.Context) error {
	var automobile models.Automobile

	if err := c.Bind(&automobile); err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusBadRequest, "Неверные формат данных для обновления авто")
	}

	if err := h.Service.Update(&automobile); err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusInternalServerError, "Ошибка при обновлении данных авто")
	}

	return c.JSON(http.StatusOK, automobile)
}
