package usecase

import (
	"context"

	"github.com/brunobotter/feature-flag/application"
	"github.com/brunobotter/feature-flag/application/command"
	"github.com/brunobotter/feature-flag/application/domain"
	"github.com/brunobotter/feature-flag/application/service"
	"github.com/brunobotter/feature-flag/infra/logger"
)

type EnviromentUsecase interface {
	GetAllByTenantId(ctx context.Context, cmd command.GetAllEnviroment, log logger.Logger) (*domain.TenantDomain, error)
}

func NewEnviromentUsecase(service service.EnviromentService, log logger.Logger) EnviromentUsecase {
	return &enviromentUsecase{
		service: service,
		log:     log,
	}
}

type enviromentUsecase struct {
	service service.EnviromentService
	log     logger.Logger
}

func (u *enviromentUsecase) GetAllByTenantId(ctx context.Context, cmd command.GetAllEnviroment, log logger.Logger) (*domain.TenantDomain, error) {
	err := cmd.Validate()
	if err != nil {
		return nil, application.Wrap(err)
	}
	tenant, err := u.service.GetAllByTenantId(ctx, cmd, log)
	if err != nil {
		return nil, application.Wrap(err)
	}
	return tenant, nil
}
