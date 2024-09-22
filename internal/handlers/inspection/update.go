package handler_inspection

import (
	handlers "github.com/Olegsuus/TZ-WEB-App/internal/handlers/errors"
	"github.com/Olegsuus/TZ-WEB-App/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// Update Обработчик для обновления данных тех осмотра
// @Summary Обновить данные о тех осмотре
// @Description Обновляет  данные тех осмотра
// @Tags тех осмотр
// @Accept  json
// @Produce  json
// @Param id path int true "ID техосмотра"
// @Param inspection body models.UpdateInspectionDTO true "Обновленные данные техосмотра"
// @Success 200  "Успешно"
// @Failure 400  "Неверные данные запроса"
// @Failure 500  "Ошибка на сервере"
// @Router /inspection/:id [patch]
func (h *InspectionHandler) Update(c echo.Context) error {
	strId := c.Param("id")

	id, err := strconv.Atoi(strId)
	if err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusBadRequest, "Неверный формат id")
	}

	var dto models.UpdateInspectionDTO

	if err := c.Bind(&dto); err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusBadRequest, "Неверный формат данных на обновление")
	}

	inspection := models.Inspection{
		ID:             id,
		AutomobileID:   dto.AutomobileID,
		CardNumber:     dto.CardNumber,
		InspectionDate: dto.InspectionDate,
		Notes:          dto.Notes,
	}

	if err = h.Service.Update(&inspection); err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusInternalServerError, "Ошибка при обновлении данных тех. осмотра")
	}

	response := models.UpdateRes{
		Success: true,
	}

	return c.JSON(http.StatusOK, response)
}
