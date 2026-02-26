package providers

import (
	"github.com/brunobotter/feature-flag/application/service"
	"github.com/brunobotter/feature-flag/application/usecase"
	"github.com/brunobotter/feature-flag/infra/logger"
	"github.com/brunobotter/feature-flag/main/container"
)

type UseCaseProvider struct{}

func NewUseCaseProvider() *UseCaseProvider {
	return &UseCaseProvider{}
}
func (p *UseCaseProvider) Register(c container.Container) {
	c.Singleton(func(tenantService service.TenantService, log logger.Logger) usecase.TenantUseCase {
		return usecase.NewTenantUseCase(tenantService, log)
	})
	c.Singleton(func(enviromentService service.EnviromentService, log logger.Logger) usecase.EnviromentUsecase {
		return usecase.NewEnviromentUsecase(enviromentService, log)
	})
}
