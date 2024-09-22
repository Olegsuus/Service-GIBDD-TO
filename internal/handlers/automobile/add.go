package handler_automobile

import (
	handlers "github.com/Olegsuus/TZ-WEB-App/internal/handlers/errors"
	"github.com/Olegsuus/TZ-WEB-App/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Add Обработчик для добавления автомобиля
// @Summary Добавить новый автомобиль
// @Description Добавляет новый автомобиль
// @Param automobile body models.AddAutomobileDTO true "Автомобиль"
// @Tags автомобиль
// @Accept  json
// @Produce  json
// @Success 200  "OK"
// @Failure 400 "Неверные данные запроса"
// @Failure 500  "Ошибка на сервере"
// @Router /car [post]
func (h *AutomobileHandlers) Add(c echo.Context) error {
	var dto models.AddAutomobileDTO

	if err := c.Bind(&dto); err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusBadRequest, "Введены неверный формат значений")
	}

	if err := c.Validate(&dto); err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusBadRequest, "Валидация данных не пройдена")
	}

	automobile := models.Automobile{
		ReleaseDate:      dto.ReleaseDate,
		Model:            dto.Model,
		LicensePlate:     dto.LicensePlate,
		RegistrationDate: dto.RegistrationDate,
	}

	if err := h.Service.Add(&automobile); err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusInternalServerError, "Ошибка при добавлении нового авто")
	}

	response := models.AddRes{
		ID: automobile.ID,
	}

	return c.JSON(http.StatusCreated, response)
}
