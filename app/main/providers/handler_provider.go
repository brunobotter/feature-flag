package providers

import (
	handler "github.com/brunobotter/feature-flag/api"
	"github.com/brunobotter/feature-flag/main/container"
)

type HandlerServiceProvider struct{}

func NewHandlerServiceProvider() *HandlerServiceProvider {
	return &HandlerServiceProvider{}
}
func (p *HandlerServiceProvider) Register(c container.Container) {
	c.Singleton(func() *handler.HealthHandler {
		return handler.NewHealthHandler()
	})

}
