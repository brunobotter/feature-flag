package providers

import (
	"github.com/brunobotter/feature-flag/api/controllers"
	"github.com/brunobotter/feature-flag/main/container"
)

type HandlerServiceProvider struct{}

func NewHandlerServiceProvider() *HandlerServiceProvider {
	return &HandlerServiceProvider{}
}
func (p *HandlerServiceProvider) Register(c container.Container) {
	c.Singleton(func() *controllers.HealthHandler {
		return controllers.NewHealthHandler()
	})

}
