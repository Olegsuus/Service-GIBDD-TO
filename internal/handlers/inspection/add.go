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
// @Param inspection body models.AddInspectionDTO true "Техосмотр"
// @Success 200   "OK"
// @Failure 400  "Неверные данные запроса"
// @Failure 500  "Ошибка на сервере"
// @Router /inspection [post]
func (h *InspectionHandler) Add(c echo.Context) error {
	var dto models.AddInspectionDTO

	if err := c.Bind(&dto); err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusBadRequest, "Неверный формат данных на добавление тех. осмотра")
	}

	if err := c.Validate(&dto); err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusBadRequest, "Валидация не пройдена")
	}

	inspection := models.Inspection{
		AutomobileID:   dto.AutomobileID,
		CardNumber:     dto.CardNumber,
		InspectionDate: dto.InspectionDate,
		Notes:          dto.Notes,
	}

	if err := h.Service.Add(&inspection); err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusInternalServerError, "Ошибка при добавлении нового тех. осмотра")
	}

	response := models.AddRes{
		ID: inspection.ID,
	}

	return c.JSON(http.StatusCreated, response)
}
