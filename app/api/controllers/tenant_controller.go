package controllers

import (
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

func (c *TenantController) CreateTenant(request requests.TenantRequest) error {

	cmd := command.CreateTenant{
		Name: request.Name,
	}

	err := c.tenantUseCase.Create(request.Request.Context(), cmd, c.log)
	if err != nil {
		return nil
	}
	return nil
}
