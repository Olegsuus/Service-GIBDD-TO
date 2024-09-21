package handlers

import (
	"github.com/Olegsuus/TZ-WEB-App/internal/errors"
	"github.com/labstack/echo/v4"
)

func ErrorsHandler(c echo.Context, err error, status int, text string) error {
	if err != nil {
		return c.JSON(status, errors.ReqError{
			Status: status,
			Text:   text,
		})
	}
	return nil
}
