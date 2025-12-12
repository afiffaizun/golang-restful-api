package helper

import (
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Tambahkan type alias ini
type ValidationErrors []ValidationError

func ValidateStruct(s interface{}) ValidationErrors {
	var errors ValidationErrors
	validate := validator.New()
	err := validate.Struct(s)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ValidationError
			element.Field = err.Field()
			element.Message = getErrorMessage(err)
			errors = append(errors, element)
		}
	}

	return errors
}

func getErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return err.Field() + " is required"
	case "min":
		return err.Field() + " must be at least " + err.Param() + " characters"
	case "max":
		return err.Field() + " must be at most " + err.Param() + " characters"
	default:
		return err.Field() + " is invalid"
	}
}