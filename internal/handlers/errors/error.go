package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

//type ReqError struct {
//	Status int    `json:"status"`
//	Text   string `json:"text"`
//}
//
//func ErrorsHandler(c echo.Context, err error, status int, text string) error {
//	if err != nil {
//		return c.JSON(status, errors.ReqError{
//			Status: status,
//			Text:   text,
//		})
//	}
//	return nil
//}

type ReqError struct {
	Status  int               `json:"status"`
	Text    string            `json:"text"`
	Details map[string]string `json:"details,omitempty"`
}

func ErrorsHandler(c echo.Context, err error, status int, text string) error {
	var details map[string]string
	if ve, ok := err.(validator.ValidationErrors); ok {
		details = make(map[string]string)
		for _, fe := range ve {
			field := fe.Field()
			tag := fe.Tag()
			details[field] = "failed on the " + tag + " tag"
		}
	} else {
		details = map[string]string{"error": err.Error()}
	}

	return c.JSON(status, ReqError{
		Status:  status,
		Text:    text,
		Details: details,
	})
}
