package command

import (
	"github.com/brunobotter/feature-flag/application"
	"github.com/brunobotter/feature-flag/application/validator"
)

type GetAllEnviroment struct {
	TenantId string
	Env      string
}

func (c *GetAllEnviroment) Validate() error {
	v := validator.NewFieldValidatorControl()
	v.AddFieldValidator("tenantId", c.TenantId, validator.Required())
	v.AddFieldValidator("env", c.Env, validator.Required())
	return application.NewValidationApplicationError(application.ValidationDomain, v.Error())
}
