package providers

import (
	"github.com/brunobotter/feature-flag/application/repo"
	"github.com/brunobotter/feature-flag/infra/logger"
	infraRepo "github.com/brunobotter/feature-flag/infra/repo"
	"github.com/brunobotter/feature-flag/main/container"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RepositoryProvider struct{}

func NewRepositoryProvider() *RepositoryProvider {
	return &RepositoryProvider{}
}
func (p *RepositoryProvider) Register(c container.Container) {

	c.Singleton(func(db *pgxpool.Pool, log logger.Logger) repo.TenantRepository {
		return infraRepo.NewTenantRepo(db, log)
	})
	c.Singleton(func(db *pgxpool.Pool, log logger.Logger) repo.EnviromentRepository {
		return infraRepo.NewEnviromentRepo(db, log)
	})
}
