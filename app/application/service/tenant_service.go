package service

import (
	"context"

	"github.com/brunobotter/feature-flag/application"
	"github.com/brunobotter/feature-flag/application/command"
	"github.com/brunobotter/feature-flag/application/domain"
	"github.com/brunobotter/feature-flag/application/repo"
	"github.com/brunobotter/feature-flag/infra/logger"
)

type TenantService interface {
	Create(ctx context.Context, cmd command.CreateTenant) (*domain.TenantDomain, error)
}

func NewTenantService(repo repo.TenantRepository, log logger.Logger) TenantService {
	return &tenantService{
		repo: repo,
		log:  log,
	}
}

type tenantService struct {
	repo repo.TenantRepository
	log  logger.Logger
}

func (s *tenantService) Create(ctx context.Context, cmd command.CreateTenant) (*domain.TenantDomain, error) {
	tenant, err := s.repo.Create(ctx, cmd)
	if err != nil {
		return nil, application.Wrap(err)
	}
	return tenant, nil
}
