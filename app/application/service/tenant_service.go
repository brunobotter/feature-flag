package service

import (
	"context"

	"github.com/brunobotter/feature-flag/application/command"
	"github.com/brunobotter/feature-flag/application/domain"
	"github.com/brunobotter/feature-flag/application/repo"
)

type TenantService interface {
	Create(ctx context.Context, cmd command.CreateTenant) (*domain.TenantDomain, error)
}

func NewTenantService(repo repo.TenantRepository) TenantService {
	return &tenantService{repo: repo}
}

type tenantService struct {
	repo repo.TenantRepository
}

func (s *tenantService) Create(ctx context.Context, cmd command.CreateTenant) (*domain.TenantDomain, error) {
	tenant, err := s.repo.Create(ctx, cmd)
	if err != nil {

	}
	return tenant, nil
}
