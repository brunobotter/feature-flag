package controllers

import (
	"github.com/brunobotter/feature-flag/api/http"
	"github.com/brunobotter/feature-flag/api/requests"
	"github.com/brunobotter/feature-flag/api/viewmodels"
	"github.com/brunobotter/feature-flag/application/command"
	"github.com/brunobotter/feature-flag/application/usecase"
	"github.com/brunobotter/feature-flag/infra/logger"
)

type EnviromentController struct {
	enviromentUsecase usecase.EnviromentUsecase
	log               logger.Logger
}

func NewEnviromentController(enviromentUsecase usecase.EnviromentUsecase, log logger.Logger) *EnviromentController {
	return &EnviromentController{
		enviromentUsecase: enviromentUsecase,
		log:               log,
	}
}

func (c *EnviromentController) GetAllEnviromentByTenantId(request *requests.Envirement) *http.HttpResponse {
	env := request.Param("env")
	tenantId := request.Param("tenantId")

	cmd := command.GetAllEnviroment{
		TenantId: tenantId,
		Env:      env,
	}
	tenant, err := c.enviromentUsecase.GetAllByTenantId(request.Context(), cmd, c.log)
	if err != nil {
		return http.HandleError(request.Context(), err, c.log)
	}
	vm := viewmodels.NewTenantViewModel(tenant)
	return http.Ok(vm)

}

func (c *EnviromentController) GetEnviromentByTenantId(request *requests.Envirement) *http.HttpResponse {
	return nil
}
