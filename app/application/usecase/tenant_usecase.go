package usecase

import (
	"context"

	"github.com/brunobotter/feature-flag/application/command"
	"github.com/brunobotter/feature-flag/application/domain"
	"github.com/brunobotter/feature-flag/application/service"
	"github.com/brunobotter/feature-flag/infra/logger"
)

type TenantUseCase interface {
	Create(ctx context.Context, cmd command.CreateTenant, log logger.Logger) (*domain.TenantDomain, error)
}

func NewTenantUseCase(service service.TenantService, log logger.Logger) TenantUseCase {
	return &tenantUseCase{
		service: service,
		log:     log,
	}
}

type tenantUseCase struct {
	service service.TenantService
	log     logger.Logger
}

func (u *tenantUseCase) Create(context context.Context, cmd command.CreateTenant, log logger.Logger) (*domain.TenantDomain, error) {
	err := cmd.Validate()
	if err != nil {
		return nil, err
	}
	tenant, err := u.service.Create(context, cmd)
	if err != nil {
		return nil, err
	}
	return tenant, nil
}
