package repo

import (
	"context"

	"github.com/brunobotter/feature-flag/application"
	"github.com/brunobotter/feature-flag/application/command"
	"github.com/brunobotter/feature-flag/application/domain"
	apprepo "github.com/brunobotter/feature-flag/application/repo"
	"github.com/brunobotter/feature-flag/infra/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TenantPgRepo struct {
	db  *pgxpool.Pool
	log logger.Logger
}

func NewTenantRepo(db *pgxpool.Pool, log logger.Logger) apprepo.TenantRepository {
	return &TenantPgRepo{
		db:  db,
		log: log,
	}
}

func (r *TenantPgRepo) Create(ctx context.Context, cmd command.CreateTenant) (tenant *domain.TenantDomain, err error) {
	const q = `
    insert into tenants (name)
    values ($1)
    returning id, created_at;
  `
	tenant = &domain.TenantDomain{
		Name: cmd.Name,
	}
	err = r.db.QueryRow(ctx, q,
		tenant.Name,
	).Scan(&tenant.Id, &tenant.CreatedAt)
	if err != nil {
		return nil, application.Wrap(err)
	}

	return tenant, nil
}
