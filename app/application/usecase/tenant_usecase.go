package usecase

import (
	"context"

	"github.com/brunobotter/feature-flag/application/command"
	"github.com/brunobotter/feature-flag/application/service"
	"github.com/brunobotter/feature-flag/infra/logger"
)

type TenantUseCase struct {
	service service.TenantService
}

func NewTenantUseCase() *TenantUseCase {
	return &TenantUseCase{}
}

func (u *TenantUseCase) Create(context context.Context, cmd command.CreateTenant, log logger.Logger) error {
	err := cmd.Validate()
	if err != nil {
		return nil
	}

	return nil
}
