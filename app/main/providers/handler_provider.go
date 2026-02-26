package providers

import (
	"github.com/brunobotter/feature-flag/api/controllers"
	"github.com/brunobotter/feature-flag/application/usecase"
	"github.com/brunobotter/feature-flag/infra/logger"
	"github.com/brunobotter/feature-flag/main/container"
)

type ControllerProvider struct{}

func NewControllereProvider() *ControllerProvider {
	return &ControllerProvider{}
}
func (p *ControllerProvider) Register(c container.Container) {
	c.Singleton(func() *controllers.HealthHandler {
		return controllers.NewHealthHandler()
	})
	c.Singleton(func(tenantUseCase usecase.TenantUseCase, log logger.Logger) *controllers.TenantController {
		return controllers.NewTenantController(tenantUseCase, log)
	})
	c.Singleton(func(enviromentUsecase usecase.EnviromentUsecase, log logger.Logger) *controllers.EnviromentController {
		return controllers.NewEnviromentController(enviromentUsecase, log)
	})
}
