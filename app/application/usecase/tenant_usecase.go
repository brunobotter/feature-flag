package usecase

import (
	"context"

	"github.com/brunobotter/feature-flag/application"
	"github.com/brunobotter/feature-flag/application/command"
	"github.com/brunobotter/feature-flag/application/domain"
	"github.com/brunobotter/feature-flag/application/service"
	"github.com/brunobotter/feature-flag/infra/logger"
)

type TenantUseCase interface {
	Create(ctx context.Context, cmd command.CreateTenant, log logger.Logger) (*domain.TenantDomain, error)
	GetById(ctx context.Context, id string, log logger.Logger) (*domain.TenantDomain, error)
	GetAllTenant(ctx context.Context, cmd command.ListTenant, log logger.Logger) (*domain.TenantPage, error)
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
		return nil, application.Wrap(err)
	}
	tenant, err := u.service.Create(context, cmd)
	if err != nil {
		return nil, application.Wrap(err)
	}
	return tenant, nil
}
func (u *tenantUseCase) GetById(ctx context.Context, id string, log logger.Logger) (*domain.TenantDomain, error) {
	tenant, err := u.service.GetById(ctx, id)
	if err != nil {
		return nil, application.Wrap(err)
	}
	return tenant, nil
}
func (u *tenantUseCase) GetAllTenant(ctx context.Context, cmd command.ListTenant, log logger.Logger) (*domain.TenantPage, error) {
	if err := cmd.Validate(); err != nil {
		return nil, application.Wrap(err)
	}
	tenant, err := u.service.GetAllTenant(ctx, cmd.Page, cmd.Limit)
	if err != nil {
		return nil, application.Wrap(err)
	}
	return tenant, nil
}
