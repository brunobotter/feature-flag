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

type EnviromentPgRepo struct {
	db  *pgxpool.Pool
	log logger.Logger
}

func NewEnviromentRepo(db *pgxpool.Pool, log logger.Logger) apprepo.EnviromentRepository {
	return &EnviromentPgRepo{
		db:  db,
		log: log,
	}
}

func (r *EnviromentPgRepo) GetAllByTenantId(ctx context.Context, cmd command.GetAllEnviroment) ([]*domain.EnviromentDomain, error) {
	const q = `
		select id, tenant_id, name, created_at, updated_at
		from environments
		where tenant_id = $1
		order by created_at desc;
	`

	rows, err := r.db.Query(ctx, q, cmd.TenantId)
	if err != nil {
		if friendlyErr := mapDatabaseError(err); friendlyErr != nil {
			return nil, friendlyErr
		}
		return nil, application.Wrap(err)
	}
	defer rows.Close()

	enviroments := make([]*domain.EnviromentDomain, 0)
	for rows.Next() {
		enviroment := &domain.EnviromentDomain{}
		if err = rows.Scan(
			&enviroment.Id,
			&enviroment.TenantId,
			&enviroment.Name,
			&enviroment.CreatedAt,
			&enviroment.UpdatedAt,
		); err != nil {
			if friendlyErr := mapDatabaseError(err); friendlyErr != nil {
				return nil, friendlyErr
			}
			return nil, application.Wrap(err)
		}
		enviroments = append(enviroments, enviroment)
	}

	if err = rows.Err(); err != nil {
		if friendlyErr := mapDatabaseError(err); friendlyErr != nil {
			return nil, friendlyErr
		}
		return nil, application.Wrap(err)
	}

	return enviroments, nil
}

func (r *EnviromentPgRepo) GetEnviromentByTenantId(ctx context.Context, cmd command.GetAllEnviroment) (*domain.EnviromentDomain, error) {
	const q = `
		select id, tenant_id, name, created_at, updated_at
		from environments
		where tenant_id = $1 and name = $2;
	`

	enviroment := &domain.EnviromentDomain{}
	if err := r.db.QueryRow(ctx, q, cmd.TenantId, cmd.Env).Scan(
		&enviroment.Id,
		&enviroment.TenantId,
		&enviroment.Name,
		&enviroment.CreatedAt,
		&enviroment.UpdatedAt,
	); err != nil {
		if friendlyErr := mapDatabaseError(err); friendlyErr != nil {
			return nil, friendlyErr
		}
		return nil, application.Wrap(err)
	}

	return enviroment, nil
}
