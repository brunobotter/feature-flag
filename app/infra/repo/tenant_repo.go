package repo

import (
	"context"
	"time"

	"github.com/brunobotter/feature-flag/application/command"
	"github.com/brunobotter/feature-flag/application/domain"
	apprepo "github.com/brunobotter/feature-flag/application/repo"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TenantPgRepo struct {
	db *pgxpool.Pool
}

func NewTenantRepo(db *pgxpool.Pool) apprepo.TenantRepository {
	return &TenantPgRepo{db: db}
}

func (r *TenantPgRepo) Create(ctx context.Context, cmd command.CreateTenant) (tenant *domain.TenantDomain, err error) {
	const q = `
    insert into tenants (id, name, created_at,)
    values ($1, $2, $3)
    returning created_at;
  `
	var createdAt time.Time
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
