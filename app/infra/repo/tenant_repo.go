package repo

import (
	"context"
	"time"

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
    insert into tenants (name, created_at)
    values ($1, $2)
    returning created_at;
  `
	var createdAt time.Time
	tenant = &domain.TenantDomain{
		Name:      cmd.Name,
		CreatedAt: time.Now(),
	}
	err = r.db.QueryRow(ctx, q,
		tenant.Id,
		tenant.Name,
		tenant.CreatedAt,
	).Scan(&createdAt)
	if err != nil {
		return nil, err
	}
	tenant.CreatedAt = createdAt
	return tenant, nil
}
