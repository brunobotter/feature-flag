package command

import (
	"github.com/brunobotter/feature-flag/application"
	"github.com/brunobotter/feature-flag/application/validator"
)

type CreateTenant struct {
	Name string
}

func (c *CreateTenant) Validate() error {
	v := validator.NewFieldValidatorControl()
	v.AddFieldValidator("name", c.Name, validator.Required())
	return application.NewValidationApplicationError(v.Error())
}
