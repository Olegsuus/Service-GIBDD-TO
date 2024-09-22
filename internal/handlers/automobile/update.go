package handler_automobile

import (
	handlers "github.com/Olegsuus/TZ-WEB-App/internal/handlers/errors"
	"github.com/Olegsuus/TZ-WEB-App/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// Update Обработчик для обновления данных автомобиля
// @Summary Обновить автомобиль
// @Description Обновляет  автомобиль
// @Tags автомобиль
// @Accept  json
// @Produce  json
// @Param id path int true "ID автомобиля"
// @Param automobile body models.Automobile true "Обновленные данные автомобиля"
// @Success 200  "OK"
// @Failure 400  "Неверные данные запроса"
// @Failure 500  "Ошибка на сервере"
// @Router /car/:id [patch]
func (h *AutomobileHandlers) Update(c echo.Context) error {
	strId := c.Param("id")

	id, err := strconv.Atoi(strId)
	if err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusBadRequest, "Неверный формат id")
	}

	var dto models.UpdateAutomobileDTO
	if err := c.Bind(&dto); err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusBadRequest, "Неверные формат данных для обновления авто")
	}

	if err := c.Validate(&dto); err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusBadRequest, "Валидация данных не пройдена")
	}

	automobile := models.Automobile{
		ID:               id,
		ReleaseDate:      dto.ReleaseDate,
		Model:            dto.Model,
		LicensePlate:     dto.LicensePlate,
		RegistrationDate: dto.RegistrationDate,
	}

	if err := h.Service.Update(&automobile); err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusInternalServerError, "Ошибка при обновлении данных авто")
	}

	response := models.UpdateRes{
		Success: true,
	}

	return c.JSON(http.StatusOK, response)
}
