package service

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidateApplication(input interface{}, validate *validator.Validate) error {
	err := validate.Struct(input)
	if err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			var customErrors []string
			for _, fieldErr := range validationErrs {
				switch fieldErr.Tag() {
				case "required":
					customErrors = append(customErrors, fmt.Sprintf("%s is required", fieldErr.Field()))
				case "email":
					customErrors = append(customErrors, fmt.Sprintf("%s must be a valid email", fieldErr.Field()))
				case "e164":
					customErrors = append(customErrors, fmt.Sprintf("%s must be a valid E.164 phone number", fieldErr.Field()))
				default:
					customErrors = append(customErrors, fmt.Sprintf("validation error on field %s: %s", fieldErr.Field(), fieldErr.Tag()))
				}
			}
			return fmt.Errorf(fmt.Sprintf("Validation errors: %v", customErrors))
		}
		return err
	}
	return nil
}
