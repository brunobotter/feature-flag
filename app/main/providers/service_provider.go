package providers

import (
	"github.com/brunobotter/feature-flag/application/repo"
	"github.com/brunobotter/feature-flag/application/service"
	"github.com/brunobotter/feature-flag/infra/logger"
	"github.com/brunobotter/feature-flag/main/container"
)

type ServiceProvider struct{}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}
func (p *ServiceProvider) Register(c container.Container) {

	c.Singleton(func(tenantRepo repo.TenantRepository, log logger.Logger) service.TenantService {
		return service.NewTenantService(tenantRepo, log)
	})
}
