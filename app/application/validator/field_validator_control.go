package validator

import (
	"errors"

	"github.com/brunobotter/feature-flag/application"
)

type InvalidFieldError struct {
	Field string
	Error string
}

type fieldValidatorControl struct {
	invalidFields []InvalidFieldError
}

func NewFieldValidatorControl() *fieldValidatorControl {
	return &fieldValidatorControl{}
}

func (i *fieldValidatorControl) AddFieldValidator(fieldName string, fieldValue any, validators ...Validator) {
	for _, validator := range validators {
		if err := validator.Validate(fieldValue); err != nil {
			i.AddInvalidField(fieldName, err.Error())
		}
	}
}

func (i *fieldValidatorControl) AddInvalidField(name, message string) {
	i.invalidFields = append(i.invalidFields, InvalidFieldError{Field: name, Error: message})
}

func (i *fieldValidatorControl) Error() error {
	if len(i.invalidFields) == 0 {
		return nil
	}
	message := ""
	for _, field := range i.invalidFields {
		message += field.Field + ": " + field.Error + ", "
	}
	return application.NewValidationApplicationError(errors.New(message[:len(message)-2]))
}
