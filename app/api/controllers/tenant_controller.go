package controllers

import (
	"strconv"

	"github.com/brunobotter/feature-flag/api/http"
	"github.com/brunobotter/feature-flag/api/requests"
	"github.com/brunobotter/feature-flag/api/viewmodels"
	"github.com/brunobotter/feature-flag/application/command"
	"github.com/brunobotter/feature-flag/application/usecase"
	"github.com/brunobotter/feature-flag/infra/logger"
	"github.com/brunobotter/feature-flag/util/shared"
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

func (c *TenantController) CreateTenant(request *requests.CreateTenantRequest) *http.HttpResponse {

	cmd := command.CreateTenant{
		Name: request.Name,
	}

	tenant, err := c.tenantUseCase.Create(request.Context(), cmd, c.log)
	if err != nil {
		return http.HandleError(request.Context(), err, c.log)
	}

	vm := viewmodels.NewTenantViewModel(tenant)
	return http.Created(vm)
}

func (c *TenantController) GetByIdTenant(request *requests.TenantRequest) *http.HttpResponse {
	id := request.Param("id")

	tenant, err := c.tenantUseCase.GetById(request.Context(), id, c.log)
	if err != nil {
		return http.HandleError(request.Context(), err, c.log)
	}
	vm := viewmodels.NewTenantViewModel(tenant)
	return http.Ok(vm)
}

func (c *TenantController) GetAllTenant(request *requests.TenantRequest) *http.HttpResponse {
	cmd := buildListTenantCommand(request)

	tenants, err := c.tenantUseCase.GetAllTenant(request.Context(), cmd, c.log)
	if err != nil {
		return http.HandleError(request.Context(), err, c.log)
	}
	vm := viewmodels.NewTenantPageViewModel(tenants)
	return http.Ok(vm)
}

func buildListTenantCommand(request *requests.TenantRequest) command.ListTenant {
	return command.ListTenant{
		Page:  parsePaginationParam(request.QueryParam("page"), shared.DefaultPage),
		Limit: parsePaginationParam(request.QueryParam("limit"), shared.DefaultLimit),
	}
}

func parsePaginationParam(value string, fallback int) int {
	if value == "" {
		return fallback
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return parsed
}
