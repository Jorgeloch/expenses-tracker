package utils

import "github.com/go-playground/validator/v10"

func ValidateDate(fl validator.FieldLevel) bool {
	date := fl.Field().Int()
	return date > 0 && date < 32
}

func RegisterDateValidation(validate *validator.Validate) error {
	return validate.RegisterValidation("day", ValidateDate)
}
