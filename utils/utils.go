package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

func GetUnmarshalTypeErrorMsg(err error) (bool, string) {
	var unmarshalTypeError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalTypeError) {
		message := fmt.Sprintf("Field '%s' has a type mismatch. Expected '%s'.", unmarshalTypeError.Field, unmarshalTypeError.Type)
		return true, message
	}
	return false, ""
}

func GetMissingFieldErrorMsg(err error, errorMessages map[string]string) (bool, string) {
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		for _, fieldError := range validationErrors {
			if msg, exists := errorMessages[fieldError.Field()]; exists {
				return true, msg
			}
		}
	}
	return false, ""
}
