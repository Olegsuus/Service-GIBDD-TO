package handler_automobile

import (
	handlers "github.com/Olegsuus/TZ-WEB-App/internal/handlers/errors"
	"github.com/Olegsuus/TZ-WEB-App/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// Delete Обработчик для удаления автомобиля
// @Summary Удалить автомобиль
// @Description Удаляет  автомобиль
// @Tags автомобиль
// @Accept  json
// @Produce  json
// @Param id path int true "ID автомобиля"
// @Param automobile body models.Automobile true "Удаление автомобиля"
// @Success 200  "OK"
// @Failure 400  "Неверные данные запроса"
// @Failure 500  "Ошибка на сервере"
// @Router /car/:id [delete]
func (h *AutomobileHandlers) Delete(c echo.Context) error {
	strId := c.Param("id")

	id, err := strconv.Atoi(strId)
	if err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusBadRequest, "Неверный формат id")
	}

	if err = h.Service.Delete(id); err != nil {
		return handlers.ErrorsHandler(c, err, http.StatusInternalServerError, "Ошибка при удалении авто")
	}

	response := models.DeleteRes{Success: true}

	return c.JSON(http.StatusOK, response)

}
