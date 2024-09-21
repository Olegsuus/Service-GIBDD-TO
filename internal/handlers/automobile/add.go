package handler_automobile

import (
	handlers "Web-App/internal/handlers/errors"
	"Web-App/internal/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

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
