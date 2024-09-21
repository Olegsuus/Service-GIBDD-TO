package handler_automobile

import (
	"fmt"
	handlers "github.com/Olegsuus/TZ-WEB-App/internal/handlers/errors"
	"github.com/Olegsuus/TZ-WEB-App/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Add Обработчик для добавления автомобиля
// @Summary Добавить новый автомобиль
// @Description Добавляет новый автомобиль
// @Param automobile body models.Automobile true "Автомобиль"
// @Tags автомобиль
// @Accept  json
// @Produce  json
// @Success 200  "OK"
// @Failure 500  "Ошибка на сервере"
// @Router /car [get]
func (h *AutomobileHandlers) Add(c echo.Context) error {
	var automobile models.Automobile

	if err := c.Bind(&automobile); err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusBadRequest, "Введены неверный формат значений")
	}
	fmt.Println("Received automobile data", "data", automobile)

	if err := h.Service.Add(&automobile); err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusInternalServerError, "Ошибка при добавлении нового авто")
	}

	fmt.Println("Automobile added successfully", "automobile", automobile)

	return c.JSON(http.StatusCreated, automobile)
}
