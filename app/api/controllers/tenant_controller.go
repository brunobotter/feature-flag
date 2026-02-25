package controllers

import (
	"github.com/brunobotter/feature-flag/api/http"
	"github.com/brunobotter/feature-flag/api/requests"
	"github.com/brunobotter/feature-flag/application/command"
	"github.com/brunobotter/feature-flag/application/usecase"
	"github.com/brunobotter/feature-flag/infra/logger"
)

type TenantController struct {
	tenantUseCase usecase.TenantUseCase
	log           logger.Logger
}

func NewTenantController(tenantUseCase usecase.TenantUseCase, logger logger.Logger) *TenantController {
	return &TenantController{
		tenantUseCase: tenantUseCase,
		log:           logger,
	}
}

func (c *TenantController) CreateTenant(request requests.TenantRequest) *http.HttpResponse {

	cmd := command.CreateTenant{
		Name: request.Name,
	}

	tenant, err := c.tenantUseCase.Create(request.Request.Context(), cmd, c.log)
	if err != nil {
		return http.BadRequest(err.Error())
	}

	return http.Created(tenant)
}
