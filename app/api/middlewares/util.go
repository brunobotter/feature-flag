package middlewares

import (
	"github.com/brunobotter/feature-flag/infra/logger"
	"github.com/brunobotter/feature-flag/main/config"
)

func CommonMiddlewares(logger logger.Logger, cfg *config.Config) []MiddlewareFunc {
	return []MiddlewareFunc{
		NewPanicMiddleware(logger),
		NewLoggerMiddleware(logger, cfg),
	}
}
