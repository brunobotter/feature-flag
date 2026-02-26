package service

import (
	"context"

	"github.com/brunobotter/feature-flag/application"
	"github.com/brunobotter/feature-flag/application/command"
	"github.com/brunobotter/feature-flag/application/domain"
	"github.com/brunobotter/feature-flag/application/repo"
	"github.com/brunobotter/feature-flag/infra/logger"
)

type EnviromentService interface {
	GetAllByTenantId(ctx context.Context, cmd command.GetAllEnviroment, log logger.Logger) (*domain.TenantDomain, error)
}

func NewEnviromentService(repo repo.EnviromentRepository, log logger.Logger) EnviromentService {
	return &enviromentService{
		repo: repo,
		log:  log,
	}
}

type enviromentService struct {
	repo repo.EnviromentRepository
	log  logger.Logger
}

func (s *enviromentService) GetAllByTenantId(ctx context.Context, cmd command.GetAllEnviroment, log logger.Logger) (*domain.TenantDomain, error) {
	tenant, err := s.repo.GetAllByTenantId(ctx, cmd)
	if err != nil {
		return nil, application.Wrap(err)
	}
	return tenant, nil
}
