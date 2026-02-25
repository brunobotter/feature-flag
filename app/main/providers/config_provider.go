package providers

import (
	"github.com/brunobotter/feature-flag/infra/logger"
	"github.com/brunobotter/feature-flag/main/config"
	"github.com/brunobotter/feature-flag/main/container"
)

type ConfigServiceProvider struct{}

func NewConfigServiceProvider() *ConfigServiceProvider {
	return &ConfigServiceProvider{}
}

func (p *ConfigServiceProvider) Register(c container.Container) {
	c.Singleton(func() *config.Config {
		cfg := config.Init()
		return cfg
	})

	c.Singleton(func(cfg *config.Config) logger.Logger {
		return logger.NewJammesLogger(cfg.App_Name, cfg.Env, false)
	})
}
