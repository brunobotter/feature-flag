package router

import (
	"github.com/brunobotter/feature-flag/api/controllers"
	"github.com/brunobotter/feature-flag/main/config"
	"github.com/brunobotter/feature-flag/main/container"
	"github.com/labstack/echo/v4"
)

func RegisterRouter(e *echo.Echo, cfg *config.Config, c container.Container) {
	var health *controllers.HealthHandler
	c.Resolve(&health)
	e.GET("/health", health.Health)
}
