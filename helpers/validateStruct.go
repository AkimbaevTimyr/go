package helpers

import (
	"github.com/go-playground/validator/v10"
)

func ValidateStruct(params interface{}) (map[string]any, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	validationError := validate.Struct(params)

	if validationError != nil {
		msg := []map[string]any{}

		for _, err := range validationError.(validator.ValidationErrors) {
			msg = append(msg, map[string]any{
				"message": err.Field() + " " + err.Tag(),
			})
		}

		return map[string]any{
			"message": "validation error",
			"errors":  msg,
		}, validationError
	}

	return nil, nil
}
