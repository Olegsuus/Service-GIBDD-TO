package customvalidator

import (
	"github.com/Olegsuus/TZ-WEB-App/internal/models"
	"github.com/go-playground/validator/v10"
	"regexp"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func New() *CustomValidator {
	v := validator.New()

	v.RegisterValidation("license_plate", func(fl validator.FieldLevel) bool {
		licensePlate := fl.Field().String()
		matched, _ := regexp.MatchString("^[A-Za-z0-9-]+$", licensePlate)
		return matched
	})

	v.RegisterValidation("customdate", func(fl validator.FieldLevel) bool {
		cd, ok := fl.Field().Interface().(models.CustomDate)
		if !ok {
			return false
		}
		return !cd.Time.IsZero()
	})

	return &CustomValidator{Validator: v}
}
