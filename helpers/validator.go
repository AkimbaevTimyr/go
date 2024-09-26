package helpers

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	// Emailregex pattern https://html.spec.whatwg.org/#valid-e-mail-address.
	EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

func ValidateEmail(email string) bool {
	return EmailRX.MatchString(email)
}

func ValidateStruct(params interface{}) (Envelope, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	validationError := validate.Struct(params)

	if validationError != nil {
		msg := []map[string]any{}

		for _, err := range validationError.(validator.ValidationErrors) {
			msg = append(msg, Envelope{"message": err.Field() + " " + err.Tag()})
		}

		return Envelope{"message": "validation error", "errors": msg}, validationError
	}

	return nil, nil
}
