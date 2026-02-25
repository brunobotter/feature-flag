package server

import (
	"github.com/brunobotter/feature-flag/api/controllers"
	"github.com/brunobotter/feature-flag/main/server/router"
)

func (s *Server) setupApiRouter(healthController *controllers.HealthHandler, tenantController *controllers.TenantController) {
	var routs router.Router
	s.container.NamedResolve(&routs, "Routes")

	s.echo.GET("/health", healthController.Health)

	routs.Group("/tenant", func(group router.RouteGroup) {
		group.POST("", tenantController.CreateTenant)
	})
}
